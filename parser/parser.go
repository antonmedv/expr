package parser

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"

	. "github.com/antonmedv/expr/ast"
	"github.com/antonmedv/expr/file"
	. "github.com/antonmedv/expr/parser/lexer"
)

type associativity int

const (
	left associativity = iota + 1
	right
)

type operator struct {
	precedence    int
	associativity associativity
}

type builtin struct {
	arity int
}

var unaryOperators = map[string]operator{
	"not": {50, left},
	"!":   {50, left},
	"-":   {500, left},
	"+":   {500, left},
}

var binaryOperators = map[string]operator{
	"or":         {10, left},
	"||":         {10, left},
	"and":        {15, left},
	"&&":         {15, left},
	"==":         {20, left},
	"!=":         {20, left},
	"<":          {20, left},
	">":          {20, left},
	">=":         {20, left},
	"<=":         {20, left},
	"not in":     {20, left},
	"in":         {20, left},
	"matches":    {20, left},
	"contains":   {20, left},
	"startsWith": {20, left},
	"endsWith":   {20, left},
	"..":         {25, left},
	"+":          {30, left},
	"-":          {30, left},
	"*":          {60, left},
	"/":          {60, left},
	"%":          {60, left},
	"**":         {70, right},
}

var builtins = map[string]builtin{
	"len":    {1},
	"all":    {2},
	"none":   {2},
	"any":    {2},
	"one":    {2},
	"filter": {2},
	"map":    {2},
}

type parser struct {
	tokens  []Token
	current Token
	pos     int
	err     *file.Error
}

type Tree struct {
	Node   Node
	Source *file.Source
}

func Parse(input string) (*Tree, error) {
	source := file.NewSource(input)

	tokens, err := Lex(source)
	if err != nil {
		return nil, err
	}

	p := &parser{
		tokens:  tokens,
		current: tokens[0],
	}

	node := p.parseExpression(0)

	if !p.current.Is(EOF) {
		p.error("unexpected token %v", p.current)
	}

	if p.err != nil {
		return nil, fmt.Errorf("%v", p.err.Format(source))
	}

	return &Tree{
		Node:   node,
		Source: source,
	}, nil
}

func (p *parser) error(format string, args ...interface{}) {
	if p.err == nil { // show first error
		p.err = &file.Error{
			Location: p.current.Location,
			Message:  fmt.Sprintf(format, args...),
		}
	}
}

func (p *parser) next() {
	p.pos++
	if p.pos >= len(p.tokens) {
		p.error("unexpected end of expression")
		return
	}
	p.current = p.tokens[p.pos]
}

func (p *parser) expect(kind Kind, values ...string) {
	if p.current.Is(kind, values...) {
		p.next()
		return
	}
	p.error("unexpected token %v", p.current)
}

// parse functions

func (p *parser) parseExpression(precedence int) Node {
	nodeLeft := p.parsePrimary()

	token := p.current
	for token.Is(Operator) {
		if op, ok := binaryOperators[token.Value]; ok {
			if op.precedence >= precedence {
				p.next()

				var nodeRight Node
				if op.associativity == left {
					nodeRight = p.parseExpression(op.precedence + 1)
				} else {
					nodeRight = p.parseExpression(op.precedence)
				}

				if token.Is(Operator, "matches") {
					var r *regexp.Regexp
					var err error

					if s, ok := nodeRight.(*StringNode); ok {
						r, err = regexp.Compile(s.Value)
						if err != nil {
							p.error("%v", err)
						}
					}
					nodeLeft = &MatchesNode{
						Base:   Loc(token.Location),
						Regexp: r,
						Left:   nodeLeft,
						Right:  nodeRight,
					}
				} else {
					nodeLeft = &BinaryNode{
						Base:     Loc(token.Location),
						Operator: token.Value,
						Left:     nodeLeft,
						Right:    nodeRight,
					}
				}
				token = p.current
				continue
			}
		}
		break
	}

	if precedence == 0 {
		nodeLeft = p.parseConditionalExpression(nodeLeft)
	}

	return nodeLeft
}

func (p *parser) parsePrimary() Node {
	token := p.current

	if token.Is(Operator) {
		if op, ok := unaryOperators[token.Value]; ok {
			p.next()
			expr := p.parseExpression(op.precedence)
			return p.parsePostfixExpression(&UnaryNode{
				Base:     Loc(token.Location),
				Operator: token.Value,
				Node:     expr,
			})
		}
	}

	if token.Is(Bracket, "(") {
		p.next()
		expr := p.parseExpression(0)
		p.expect(Bracket, ")") // "an opened parenthesis is not properly closed"
		return p.parsePostfixExpression(expr)
	}

	if token.Is(Operator, "#") || token.Is(Operator, ".") {
		if token.Is(Operator, "#") {
			p.next()
		}
		node := &PointerNode{Base: Loc(token.Location)}
		return p.parsePostfixExpression(node)
	}

	return p.parsePrimaryExpression()
}

func (p *parser) parseConditionalExpression(node Node) Node {
	var expr1, expr2 Node
	for p.current.Is(Operator, "?") {
		p.next()

		if !p.current.Is(Operator, ":") {
			expr1 = p.parseExpression(0)
			p.expect(Operator, ":")
			expr2 = p.parseExpression(0)
		} else {
			p.next()
			expr1 = node
			expr2 = p.parseExpression(0)
		}

		node = &ConditionalNode{
			Cond: node,
			Exp1: expr1,
			Exp2: expr2,
		}
	}
	return node
}

func (p *parser) parsePrimaryExpression() Node {
	var node Node
	token := p.current

	switch token.Kind {

	case Identifier:
		p.next()
		switch token.Value {
		case "true":
			return &BoolNode{Base: Loc(token.Location), Value: true}
		case "false":
			return &BoolNode{Base: Loc(token.Location), Value: false}
		case "nil":
			return &NilNode{Base: Loc(token.Location)}
		default:
			node = p.parseIdentifierExpression(token)
		}

	case Number:
		p.next()
		value := strings.Replace(token.Value, "_", "", -1)
		if strings.Contains(value, ".") {
			number, err := strconv.ParseFloat(value, 64)
			if err != nil {
				p.error("invalid float literal: %v", err)
			}
			return &FloatNode{Base: Loc(token.Location), Value: number}
		} else if strings.Contains(value, "x") {
			number, err := strconv.ParseInt(value, 0, 64)
			if err != nil {
				p.error("invalid hex literal: %v", err)
			}
			return &IntegerNode{Base: Loc(token.Location), Value: int(number)}
		} else {
			number, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				p.error("invalid integer literal: %v", err)
			}
			return &IntegerNode{Base: Loc(token.Location), Value: int(number)}
		}

	case String:
		p.next()
		return &StringNode{Base: Loc(token.Location), Value: token.Value}

	default:
		if token.Is(Bracket, "[") {
			node = p.parseArrayExpression(token)
		} else if token.Is(Bracket, "{") {
			node = p.parseMapExpression(token)
		} else {
			p.error("unexpected token %v", token)
		}
	}

	return p.parsePostfixExpression(node)
}

func (p *parser) parseIdentifierExpression(token Token) Node {
	var node Node
	if p.current.Is(Bracket, "(") {
		var arguments []Node

		if b, ok := builtins[token.Value]; ok {
			p.expect(Bracket, "(")
			// TODO: Add builtins signatures.
			if b.arity == 1 {
				arguments = make([]Node, 1)
				arguments[0] = p.parseExpression(0)
			} else if b.arity == 2 {
				arguments = make([]Node, 2)
				arguments[0] = p.parseExpression(0)
				p.expect(Operator, ",")
				arguments[1] = p.parseClosure()
			}
			p.expect(Bracket, ")")

			node = &BuiltinNode{
				Base:      Loc(token.Location),
				Name:      token.Value,
				Arguments: arguments,
			}
		} else {
			arguments = p.parseArguments()
			node = &FunctionNode{
				Base:      Loc(token.Location),
				Name:      token.Value,
				Arguments: arguments,
			}
		}
	} else {
		node = &IdentifierNode{Base: Loc(token.Location), Value: token.Value}
	}
	return node
}

func (p *parser) parseClosure() Node {
	token := p.current
	p.expect(Bracket, "{")

	node := p.parseExpression(0)

	p.expect(Bracket, "}")
	return &ClosureNode{
		Base: Loc(token.Location),
		Node: node,
	}
}

func (p *parser) parseArrayExpression(token Token) Node {
	nodes := p.parseList("[", "]")
	return &ArrayNode{Base: Loc(token.Location), Nodes: nodes}
}

func (p *parser) parseMapExpression(token Token) Node {
	p.expect(Bracket, "{")

	nodes := make([]Node, 0)
	for !p.current.Is(Bracket, "}") {
		if len(nodes) > 0 {
			p.expect(Operator, ",")
		}

		var key Node
		// a map key can be:
		//  * a number
		//  * a string
		//  * a identifier, which is equivalent to a string
		//  * an expression, which must be enclosed in parentheses -- (1 + 2)
		if p.current.Is(Number) || p.current.Is(String) || p.current.Is(Identifier) {
			key = &StringNode{Base: Loc(token.Location), Value: p.current.Value}
			p.next()
		} else if p.current.Is(Bracket, "(") {
			key = p.parseExpression(0)
		} else {
			p.error("a map key must be a quoted string, a number, a identifier, or an expression enclosed in parentheses (unexpected token %v)", p.current)
		}

		p.expect(Operator, ":")

		node := p.parseExpression(0)
		nodes = append(nodes, &PairNode{Base: Loc(token.Location), Key: key, Value: node})
	}

	p.expect(Bracket, "}")

	return &MapNode{Base: Loc(token.Location), Pairs: nodes}
}

func (p *parser) parsePostfixExpression(node Node) Node {
	token := p.current
	for token.Is(Operator) || token.Is(Bracket) {
		if token.Value == "." {
			p.next()

			token = p.current
			p.next()

			if token.Kind != Identifier &&
				// Operators like "not" and "matches" are valid method or property names,
				//
				// In other words, besides name token kind, operator kind could also be parsed as a property or method.
				// This is because operators are processed by the lexer prior to names. So "not" in "foo.not()"
				// or "matches" in "foo.matches" will be recognized as an operator first. But in fact, "not"
				// and "matches" in such expressions shall be parsed as method or property names.
				//
				// And this ONLY works if the operator consists of valid characters for a property or method name.
				//
				// Other types, such as text kind and number kind, can't be parsed as property nor method names.
				//
				// As a result, if token is NOT an operator OR token.value is NOT a valid property or method name,
				// an error shall be returned.
				(token.Kind != Operator || !isValidIdentifier(token.Value)) {
				p.error("expected name")
			}

			if p.current.Is(Bracket, "(") {
				arguments := p.parseArguments()
				node = &MethodNode{
					Base:      Loc(token.Location),
					Node:      node,
					Method:    token.Value,
					Arguments: arguments,
				}
			} else {
				node = &PropertyNode{
					Base:     Loc(token.Location),
					Node:     node,
					Property: token.Value,
				}
			}

		} else if token.Value == "[" {
			p.next()
			var from, to Node

			if p.current.Is(Operator, ":") { // slice without from [:1]
				p.next()

				if !p.current.Is(Bracket, "]") { // slice without from and to [:]
					to = p.parseExpression(0)
				}

				node = &SliceNode{
					Base: Loc(p.current.Location),
					Node: node,
					To:   to,
				}
				p.expect(Bracket, "]")

			} else {

				from = p.parseExpression(0)

				if p.current.Is(Operator, ":") {
					p.next()

					if !p.current.Is(Bracket, "]") { // slice without to [1:]
						to = p.parseExpression(0)
					}

					node = &SliceNode{
						Base: Loc(p.current.Location),
						Node: node,
						From: from,
						To:   to,
					}
					p.expect(Bracket, "]")

				} else {
					// Slice operator [:] was not found, it should by just index node.

					node = &IndexNode{
						Base:  Loc(token.Location),
						Node:  node,
						Index: from,
					}
					p.expect(Bracket, "]")
				}
			}

		} else {
			break
		}

		token = p.current
	}
	return node
}

func isValidIdentifier(str string) bool {
	if len(str) == 0 {
		return false
	}
	h, w := utf8.DecodeRuneInString(str)
	if !isAlphabetic(h) {
		return false
	}
	for _, r := range str[w:] {
		if !isAlphaNumeric(r) {
			return false
		}
	}
	return true
}

func isAlphaNumeric(r rune) bool {
	return isAlphabetic(r) || unicode.IsDigit(r)
}

func isAlphabetic(r rune) bool {
	return r == '_' || unicode.IsLetter(r)
}

func (p *parser) parseArguments() []Node {
	return p.parseList("(", ")")
}

func (p *parser) parseList(start, end string) []Node {
	p.expect(Bracket, start)

	nodes := make([]Node, 0)
	for !p.current.Is(Bracket, end) {
		if len(nodes) > 0 {
			p.expect(Operator, ",")
		}
		node := p.parseExpression(0)
		nodes = append(nodes, node)
	}

	p.expect(Bracket, end)
	return nodes
}
