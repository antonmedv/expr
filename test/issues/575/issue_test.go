// Package issue575 tests that the collection builtins (all, any, none, one,
// count, sum, filter, map, find, findLast, reduce, min, max, mean, median)
// work on maps with string keys, not just arrays.
//
// See https://github.com/expr-lang/expr/issues/575.
package issue575

import (
	"reflect"
	"testing"

	"github.com/expr-lang/expr"
	"github.com/expr-lang/expr/internal/testify/assert"
	"github.com/expr-lang/expr/internal/testify/require"
)

// asStringMap normalizes any map result into a map[string]any so assertions
// do not depend on the concrete map type a program returns.
func asStringMap(t *testing.T, out any) map[string]any {
	t.Helper()
	rv := reflect.ValueOf(out)
	require.Equal(t, reflect.Map, rv.Kind(), "expected a map result, got %T", out)
	m := make(map[string]any, rv.Len())
	iter := rv.MapRange()
	for iter.Next() {
		k := iter.Key()
		if k.Kind() == reflect.Interface {
			k = k.Elem()
		}
		require.Equal(t, reflect.String, k.Kind(), "expected string map keys, got %v", k.Kind())
		m[k.String()] = iter.Value().Interface()
	}
	return m
}

func runTyped(t *testing.T, code string, env any) any {
	t.Helper()
	program, err := expr.Compile(code, expr.Env(env))
	require.NoError(t, err)
	out, err := expr.Run(program, env)
	require.NoError(t, err)
	return out
}

func runUntyped(t *testing.T, code string, env any) any {
	t.Helper()
	program, err := expr.Compile(code)
	require.NoError(t, err)
	out, err := expr.Run(program, env)
	require.NoError(t, err)
	return out
}

func testEnv() map[string]any {
	return map[string]any{
		"prices": map[string]int{"banana": 3, "apple": 1, "cherry": 5},
		"ratios": map[string]float64{"x": 1.5, "y": 2.5, "z": 3.5},
		"flags":  map[string]bool{"a": true, "b": false, "c": true},
		"tags":   map[string]string{"b": "beta", "a": "alpha", "c": "gamma"},
		"nested": map[string]any{"evens": []int{2, 4}, "odds": []int{1, 3, 5}},
		"items":  []int{10, 20, 30},
	}
}

func TestIssue575_filter_and_map_over_maps(t *testing.T) {
	env := testEnv()

	out := runTyped(t, `filter(prices, # > 2)`, env)
	assert.Equal(t, map[string]any{"banana": 3, "cherry": 5}, asStringMap(t, out))

	out = runTyped(t, `filter(prices, # > 100)`, env)
	assert.Len(t, asStringMap(t, out), 0)

	out = runTyped(t, `map(prices, # * 10)`, env)
	assert.Equal(t, map[string]any{"apple": 10, "banana": 30, "cherry": 50}, asStringMap(t, out))

	out = runTyped(t, `map(tags, upper(#))`, env)
	assert.Equal(t, map[string]any{"a": "ALPHA", "b": "BETA", "c": "GAMMA"}, asStringMap(t, out))

	// map() over a map keeps the original keys.
	out = runTyped(t, `map(prices, #index)`, env)
	assert.Equal(t, map[string]any{"apple": "apple", "banana": "banana", "cherry": "cherry"}, asStringMap(t, out))
}

func TestIssue575_all_any_none_one_over_maps(t *testing.T) {
	env := testEnv()
	assert.Equal(t, true, runTyped(t, `all(prices, # > 0)`, env))
	assert.Equal(t, false, runTyped(t, `all(prices, # > 1)`, env))
	assert.Equal(t, true, runTyped(t, `any(prices, # > 4)`, env))
	assert.Equal(t, false, runTyped(t, `any(prices, # > 100)`, env))
	assert.Equal(t, true, runTyped(t, `none(prices, # > 100)`, env))
	assert.Equal(t, false, runTyped(t, `none(prices, # > 4)`, env))
	assert.Equal(t, true, runTyped(t, `one(prices, # == 3)`, env))
	assert.Equal(t, false, runTyped(t, `one(prices, # > 1)`, env))
}

func TestIssue575_count_and_sum_over_maps(t *testing.T) {
	env := testEnv()
	assert.Equal(t, 2, runTyped(t, `count(prices, # > 2)`, env))
	assert.Equal(t, 0, runTyped(t, `count(prices, # > 100)`, env))
	// Without a predicate, count() counts truthy values.
	assert.Equal(t, 2, runTyped(t, `count(flags)`, env))

	assert.Equal(t, 9, runTyped(t, `sum(prices)`, env))
	assert.Equal(t, 7.5, runTyped(t, `sum(ratios)`, env))
	assert.Equal(t, 18, runTyped(t, `sum(prices, # * 2)`, env))
}

func TestIssue575_find_and_findLast_use_key_order(t *testing.T) {
	env := testEnv()
	// Ascending key order: apple=1, banana=3, cherry=5.
	assert.Equal(t, 1, runTyped(t, `find(prices, # > 0)`, env))
	assert.Equal(t, 3, runTyped(t, `find(prices, # > 1)`, env))
	assert.Equal(t, 5, runTyped(t, `findLast(prices, # > 0)`, env))
	assert.Equal(t, 3, runTyped(t, `findLast(prices, # < 5)`, env))
	assert.Nil(t, runTyped(t, `find(prices, # > 100)`, env))
	assert.Nil(t, runTyped(t, `findLast(prices, # > 100)`, env))
}

func TestIssue575_reduce_folds_values_in_key_order(t *testing.T) {
	env := testEnv()
	assert.Equal(t, 9, runTyped(t, `reduce(prices, #acc + #, 0)`, env))
	assert.Equal(t, "alpha|beta|gamma|", runTyped(t, `reduce(tags, #acc + # + "|", "")`, env))
	// #index is available in reduce too, and holds the key.
	assert.Equal(t, "abc", runTyped(t, `reduce(tags, #acc + #index, "")`, env))
}

func TestIssue575_index_is_the_map_key(t *testing.T) {
	env := testEnv()
	out := runTyped(t, `filter(prices, #index startsWith "b")`, env)
	assert.Equal(t, map[string]any{"banana": 3}, asStringMap(t, out))
	assert.Equal(t, 1, runTyped(t, `count(prices, #index == "cherry")`, env))
	assert.Equal(t, 5, runTyped(t, `find(prices, #index == "cherry")`, env))
	assert.Equal(t, true, runTyped(t, `any(tags, #index == "a")`, env))
}

func TestIssue575_index_still_holds_position_for_arrays(t *testing.T) {
	env := testEnv()
	// #index becomes available in all these predicates for arrays too,
	// but keeps holding the element's position, not a key.
	out := runTyped(t, `filter(items, #index > 0)`, env)
	assert.Equal(t, []any{20, 30}, out)
	assert.Equal(t, 20, runTyped(t, `find(items, #index == 1)`, env))
	assert.Equal(t, 1, runTyped(t, `count(items, #index == 2)`, env))
	assert.Equal(t, true, runTyped(t, `all(items, #index < 3)`, env))
}

func TestIssue575_empty_and_nil_maps_behave_like_empty_arrays(t *testing.T) {
	var nilMap map[string]int
	env := map[string]any{"empty": map[string]int{}, "vacant": nilMap}
	for _, name := range []string{"empty", "vacant"} {
		assert.Equal(t, true, runTyped(t, `all(`+name+`, # > 0)`, env), name)
		assert.Equal(t, false, runTyped(t, `any(`+name+`, # > 0)`, env), name)
		assert.Equal(t, true, runTyped(t, `none(`+name+`, # > 0)`, env), name)
		assert.Equal(t, false, runTyped(t, `one(`+name+`, # > 0)`, env), name)
		assert.Equal(t, 0, runTyped(t, `count(`+name+`, # > 0)`, env), name)
		assert.Equal(t, 0, runTyped(t, `sum(`+name+`)`, env), name)
		assert.Nil(t, runTyped(t, `find(`+name+`, # > 0)`, env), name)
		assert.Len(t, asStringMap(t, runTyped(t, `filter(`+name+`, # > 0)`, env)), 0, name)
		assert.Len(t, asStringMap(t, runTyped(t, `map(`+name+`, #)`, env)), 0, name)
	}
}

func TestIssue575_nested_map_and_array_predicates(t *testing.T) {
	env := testEnv()
	// Outer predicate iterates a map, inner one iterates the value array.
	out := runTyped(t, `filter(nested, all(#, # % 2 == 0))`, env)
	assert.Equal(t, map[string]any{"evens": []int{2, 4}}, asStringMap(t, out))
}

func TestIssue575_filter_and_map_compose(t *testing.T) {
	env := testEnv()
	assert.Equal(t, 2, runTyped(t, `len(filter(prices, # > 2))`, env))

	out := runTyped(t, `map(filter(prices, # > 2), # * 100)`, env)
	assert.Equal(t, map[string]any{"banana": 300, "cherry": 500}, asStringMap(t, out))

	assert.Equal(t, 8, runTyped(t, `sum(filter(prices, # > 2))`, env))

	out = runTyped(t, `filter(map(prices, # * 2), # > 5)`, env)
	assert.Equal(t, map[string]any{"banana": 6, "cherry": 10}, asStringMap(t, out))
}

func TestIssue575_map_and_filter_results_feed_map_operations(t *testing.T) {
	env := testEnv()
	assert.Equal(t, 30, runTyped(t, `map(prices, # * 10)["banana"]`, env))
	assert.Equal(t, 5, runTyped(t, `filter(prices, # > 2)["cherry"]`, env))
	assert.Equal(t, true, runTyped(t, `"banana" in filter(prices, # > 2)`, env))
	assert.Equal(t, false, runTyped(t, `"apple" in filter(prices, # > 2)`, env))
	keys := runTyped(t, `keys(filter(prices, # > 2))`, env)
	assert.ElementsMatch(t, []any{"banana", "cherry"}, keys)
}

type issue575Event struct {
	Author string
	Rating float64
}

func TestIssue575_groupBy_pipeline(t *testing.T) {
	// This is the motivating example from the issue: groupBy() produces a
	// map, and previously nothing downstream could consume it.
	env := map[string]any{
		"events": []issue575Event{
			{"ann", 7.0},
			{"bob", 9.1},
			{"bob", 8.5},
			{"cat", 9.9},
		},
	}
	out := runTyped(t, `events | groupBy(.Author) | filter(len(#) > 1)`, env)
	m := asStringMap(t, out)
	require.Len(t, m, 1)
	require.Contains(t, m, "bob")

	out = runTyped(t, `events | groupBy(.Author) | map(len(#))`, env)
	assert.Equal(t, map[string]any{"ann": 1, "bob": 2, "cat": 1}, asStringMap(t, out))

	assert.Equal(t, 2, runTyped(t, `events | groupBy(.Author) | count(all(#, .Rating > 8.0))`, env))
}

func TestIssue575_min_max_mean_median_aggregate_map_values(t *testing.T) {
	env := testEnv()
	assert.Equal(t, 1, runTyped(t, `min(prices)`, env))
	assert.Equal(t, 5, runTyped(t, `max(prices)`, env))
	assert.Equal(t, 2.5, runTyped(t, `mean(ratios)`, env))
	assert.Equal(t, 2.5, runTyped(t, `median(ratios)`, env))
	// A map argument mixes freely with array and scalar arguments.
	assert.Equal(t, 1, runTyped(t, `min(items, prices)`, env))
	assert.Equal(t, 30, runTyped(t, `max(items, prices)`, env))
}

func TestIssue575_works_without_a_typed_environment(t *testing.T) {
	env := testEnv()
	out := runUntyped(t, `filter(prices, # > 2)`, env)
	assert.Equal(t, map[string]any{"banana": 3, "cherry": 5}, asStringMap(t, out))

	out = runUntyped(t, `map(prices, # * 10)`, env)
	assert.Equal(t, map[string]any{"apple": 10, "banana": 30, "cherry": 50}, asStringMap(t, out))

	out = runUntyped(t, `filter(prices, #index startsWith "b")`, env)
	assert.Equal(t, map[string]any{"banana": 3}, asStringMap(t, out))

	assert.Equal(t, 9, runUntyped(t, `sum(prices)`, env))
	assert.Equal(t, 1, runUntyped(t, `find(prices, # > 0)`, env))
	assert.Equal(t, "abc", runUntyped(t, `reduce(tags, #acc + #index, "")`, env))
	assert.Equal(t, 1, runUntyped(t, `min(prices)`, env))
	assert.Equal(t, 2.5, runUntyped(t, `mean(ratios)`, env))

	out = runUntyped(t, `groupBy(events, .Author) | filter(len(#) > 1)`, map[string]any{
		"events": []issue575Event{{"ann", 7.0}, {"bob", 9.1}, {"bob", 8.5}},
	})
	m := asStringMap(t, out)
	require.Len(t, m, 1)
	require.Contains(t, m, "bob")
}

func TestIssue575_non_string_keyed_maps_are_rejected(t *testing.T) {
	env := map[string]any{
		"m":  map[string]int{"a": 1},
		"mi": map[int]string{1: "x"},
	}
	// String-keyed maps work...
	assert.Equal(t, 1, runTyped(t, `sum(m)`, env))
	assert.Equal(t, 1, runTyped(t, `min(m)`, env))

	// ...while non-string-keyed maps are rejected at compile time with a
	// typed environment...
	for _, code := range []string{`sum(mi)`, `filter(mi, true)`, `map(mi, #)`, `min(mi)`, `all(mi, true)`} {
		_, err := expr.Compile(code, expr.Env(env))
		assert.Error(t, err, code)
	}

	// ...and at runtime without one.
	for _, code := range []string{`sum(mi)`, `filter(mi, true)`, `min(mi)`} {
		program, err := expr.Compile(code)
		require.NoError(t, err, code)
		_, err = expr.Run(program, env)
		assert.Error(t, err, code)
	}
}

func TestIssue575_positional_and_reordering_builtins_still_reject_maps(t *testing.T) {
	env := map[string]any{"m": map[string]int{"a": 1, "b": 2}}
	// findIndex, findLastIndex, groupBy and sortBy don't translate to map
	// entries, since positions and reordering don't make sense for a map.
	for _, code := range []string{
		`findIndex(m, # > 0)`,
		`findLastIndex(m, # > 0)`,
		`groupBy(m, #)`,
		`sortBy(m, #)`,
	} {
		_, err := expr.Compile(code, expr.Env(env))
		assert.Error(t, err, code)
	}
	// ...while the supported builtins accept maps just fine.
	assert.Equal(t, 2, runTyped(t, `findLast(m, # > 0)`, env))
}

func TestIssue575_array_behavior_is_unchanged(t *testing.T) {
	env := testEnv()
	assert.Equal(t, []any{20, 30}, runTyped(t, `filter(items, # > 10)`, env))
	assert.Equal(t, []any{20, 40, 60}, runTyped(t, `map(items, # * 2)`, env))
	assert.Equal(t, []any{0, 1, 2}, runTyped(t, `map(items, #index)`, env))
	assert.Equal(t, 60, runTyped(t, `sum(items)`, env))
	assert.Equal(t, 30, runTyped(t, `findLast(items, # > 10)`, env))
	assert.Equal(t, 10, runTyped(t, `min(items)`, env))
	assert.Equal(t, 20.0, runTyped(t, `mean(items)`, env))
	assert.Equal(t, true, runTyped(t, `one(items, # == 20)`, env))
	assert.Equal(t, 60, runTyped(t, `reduce(items, #acc + #, 0)`, env))
}
