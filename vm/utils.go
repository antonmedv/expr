package vm

import (
	"reflect"
	"sort"
	"time"
)

type (
	Function     = func(params ...any) (any, error)
	SafeFunction = func(params ...any) (any, uint, error)
)

var (
	errorType = reflect.TypeOf((*error)(nil)).Elem()
)

type Scope struct {
	Array reflect.Value
	Index int
	Len   int
	Count int
	Acc   any
	// Map iteration state. Keys holds the sorted keys of the iterated map;
	// it is non-nil if and only if the scope iterates over a map. Marked
	// collects the keys of elements kept by collecting builtins.
	Keys   []string
	Marked []string
	Map    map[string]any // Fast path for map[string]any
	// Fast paths
	Ints    []int
	Floats  []float64
	Strings []string
	Anys    []any
}

// Item returns the current element from the scope using fast paths when available.
func (s *Scope) Item() any {
	if s.Ints != nil {
		return s.Ints[s.Index]
	}
	if s.Floats != nil {
		return s.Floats[s.Index]
	}
	if s.Strings != nil {
		return s.Strings[s.Index]
	}
	if s.Anys != nil {
		return s.Anys[s.Index]
	}
	if s.Keys != nil {
		key := s.Keys[s.Index]
		if s.Map != nil {
			return s.Map[key]
		}
		k := reflect.ValueOf(key)
		if kt := s.Array.Type().Key(); kt.Kind() != reflect.Interface {
			k = k.Convert(kt)
		}
		return s.Array.MapIndex(k).Interface()
	}
	return s.Array.Index(s.Index).Interface()
}

func sortedKeys(m map[string]any) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

type groupBy = map[any][]any

type Span struct {
	Name       string  `json:"name"`
	Expression string  `json:"expression"`
	Duration   int64   `json:"duration"`
	Children   []*Span `json:"children"`
	start      time.Time
}

func GetSpan(program *Program) *Span {
	return program.span
}
