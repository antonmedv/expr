package expr

import (
	"reflect"

	"github.com/antonmedv/expr/checker"
	"github.com/antonmedv/expr/compiler"
	"github.com/antonmedv/expr/internal/conf"
	"github.com/antonmedv/expr/optimizer"
	"github.com/antonmedv/expr/parser"
	"github.com/antonmedv/expr/vm"
)

// Eval parses, compiles and runs given input.
func Eval(input string, env interface{}) (interface{}, error) {
	tree, err := parser.Parse(input)
	if err != nil {
		return nil, err
	}

	program, err := compiler.Compile(tree, nil)
	if err != nil {
		return nil, err
	}

	output, err := vm.Run(program, env)
	if err != nil {
		return nil, err
	}

	return output, nil
}

// Env specifies expected input of env for type checks.
// If struct is passed, all fields will be treated as variables,
// as well as all fields of embedded structs and struct itself.
// If map is passed, all items will be treated as variables.
// Methods defined on this type will be available as functions.
func Env(i interface{}) conf.Option {
	return func(c *conf.Config) {
		if _, ok := i.(map[string]interface{}); ok {
			c.MapEnv = true
		}
		c.Types = conf.CreateTypesTable(i)
	}
}

// Operator allows to override binary operator with function.
func Operator(operator string, fn ...string) conf.Option {
	return func(c *conf.Config) {
		c.Operators[operator] = append(c.Operators[operator], fn...)
	}
}

// AsBool tells the compiler to expect boolean result.
func AsBool() conf.Option {
	return func(c *conf.Config) {
		c.Expect = reflect.Bool
	}
}

// AsInt64 tells the compiler to expect int64 result.
func AsInt64() conf.Option {
	return func(c *conf.Config) {
		c.Expect = reflect.Int64
	}
}

// AsFloat64 tells the compiler to expect float64 result.
func AsFloat64() conf.Option {
	return func(c *conf.Config) {
		c.Expect = reflect.Float64
	}
}

// Optimize turns optimizations on or off.
func Optimize(b bool) conf.Option {
	return func(c *conf.Config) {
		c.Optimize = b
	}
}

// Compile parses and compiles given input expression to bytecode program.
func Compile(input string, ops ...conf.Option) (*vm.Program, error) {
	config := &conf.Config{
		Operators: make(map[string][]string),
		Optimize:  true,
	}

	for _, op := range ops {
		op(config)
	}

	if err := config.Check(); err != nil {
		return nil, err
	}

	tree, err := parser.Parse(input)
	if err != nil {
		return nil, err
	}

	if config.Types != nil {
		_, err = checker.Check(tree, config)
		if err != nil {
			return nil, err
		}
		checker.PatchOperators(tree, config)
	}

	if config.Optimize {
		optimizer.Optimize(&tree.Node)
	}

	program, err := compiler.Compile(tree, config)
	if err != nil {
		return nil, err
	}

	return program, nil
}

// Run evaluates given bytecode program.
func Run(program *vm.Program, env interface{}) (interface{}, error) {
	return vm.Run(program, env)
}

// Error is implemented by errors returned by the package.
type Error interface {
	error
	// The error message.
	Message() string
	// The 1-based line and 0-based column where the error occurred.
	// Line is 0 when no location information is available.
	Location() (line, column int)
	// The expression source.
	// An empty string is returned when the expression source is unavailable.
	Source() string
}
