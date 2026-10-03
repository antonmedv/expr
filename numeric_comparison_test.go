package expr_test

import (
	"math"
	"math/big"
	"testing"

	"github.com/expr-lang/expr"
	"github.com/expr-lang/expr/types"
)

func TestNumericComparisonUint64(t *testing.T) {
	env := map[string]any{"n": uint64(math.MaxUint64), "zero": uint64(0)}
	program, err := expr.Compile("n > zero", expr.Env(env))
	if err != nil {
		t.Fatal(err)
	}
	got, err := expr.Run(program, env)
	if err != nil {
		t.Fatal(err)
	}
	if got != true {
		t.Fatalf("MaxUint64 > uint64(0) = %v, want true", got)
	}
}

func FuzzNumericComparison(f *testing.F) {
	for _, seed := range []struct {
		i int64
		u uint64
		v float64
	}{
		{-1, math.MaxUint64, -1},
		{9007199254740993, 9007199254740993, 9007199254740992},
		{math.MaxInt64, math.MaxUint64, 0x1p64},
		{math.MinInt64, 0, -0x1p63},
		{0, 0, math.SmallestNonzeroFloat64},
		{0, 0, -math.SmallestNonzeroFloat64},
		{1, 1, 1.5},
		{-1, 1, -1.5},
	} {
		f.Add(seed.i, seed.u, math.Float64bits(seed.v))
	}
	program, err := expr.Compile("[a == b, a != b, a < b, a <= b, a > b, a >= b]", expr.Env(types.Map{"a": types.Any, "b": types.Any}))
	if err != nil {
		f.Fatal(err)
	}
	f.Fuzz(func(t *testing.T, i int64, u, floatBits uint64) {
		v := math.Float64frombits(floatBits)
		if math.IsNaN(v) || math.IsInf(v, 0) {
			return // IEEE unordered/infinite cases are covered by the table test.
		}
		integer := new(big.Rat).SetInt64(i)
		unsigned := new(big.Rat).SetInt(new(big.Int).SetUint64(u))
		floating := new(big.Rat).SetFloat64(v)
		for _, pair := range []struct {
			a, b  any
			order int
		}{
			{i, u, integer.Cmp(unsigned)},
			{i, v, integer.Cmp(floating)},
			{u, v, unsigned.Cmp(floating)},
		} {
			for reverse := 0; reverse < 2; reverse++ {
				a, b, order := pair.a, pair.b, pair.order
				if reverse == 1 {
					a, b, order = b, a, -order
				}
				got, err := expr.Run(program, map[string]any{"a": a, "b": b})
				if err != nil {
					t.Fatal(err)
				}
				want := []any{order == 0, order != 0, order < 0, order <= 0, order > 0, order >= 0}
				for index, value := range got.([]any) {
					if value != want[index] {
						t.Errorf("%T(%v) versus %T(%v): comparison %d = %v, want %v", a, a, b, b, index, value, want[index])
					}
				}
			}
		}
	})
}

func TestNumericComparisonSignedUnsigned(t *testing.T) {
	for _, test := range []struct {
		name  string
		a, b  any
		order int
	}{
		{"negative and unsigned max", int64(-1), uint64(math.MaxUint64), -1},
		{"signed and unsigned max", int64(math.MaxInt64), uint64(math.MaxUint64), -1},
		{"min signed and unsigned zero", int64(math.MinInt64), uint64(0), -1},
		{"equal zero", int8(0), uint8(0), 0},
		{"equal positive", int64(42), uint32(42), 0},
		{"larger signed positive", int32(43), uint64(42), 1},
	} {
		t.Run(test.name, func(t *testing.T) {
			checkNumericComparisons(t, test.a, test.b, test.order, true)
			checkNumericComparisons(t, test.b, test.a, -test.order, true)
		})
	}
}

func TestNumericComparisonIntegerFloat(t *testing.T) {
	for _, test := range []struct {
		name  string
		a, b  any
		order int
	}{
		{"signed beyond exact float range", int64(9007199254740993), float64(9007199254740992), 1},
		{"unsigned beyond exact float range", uint64(9007199254740993), float64(9007199254740992), 1},
		{"signed max rounded up", int64(math.MaxInt64), float64(0x1p63), -1},
		{"unsigned max rounded up", uint64(math.MaxUint64), float64(0x1p64), -1},
		{"signed min", int64(math.MinInt64), float64(-0x1p63), 0},
		{"signed min rounded down", int64(math.MinInt64), math.Nextafter(-0x1p63, math.Inf(-1)), 1},
		{"positive fraction", int64(1), float64(1.5), -1},
		{"negative fraction", int64(-1), float64(-1.5), 1},
		{"unsigned and negative fraction", uint64(0), float64(-0.5), 1},
		{"negative zero", int64(0), math.Copysign(0, -1), 0},
		{"positive infinity", uint64(math.MaxUint64), math.Inf(1), -1},
		{"negative infinity", int64(math.MinInt64), math.Inf(-1), 1},
		{"float32 integer", int64(16777217), float32(16777216), 1},
	} {
		t.Run(test.name, func(t *testing.T) {
			checkNumericComparisons(t, test.a, test.b, test.order, true)
			checkNumericComparisons(t, test.b, test.a, -test.order, true)
		})
	}
	for _, value := range []any{int64(0), uint64(math.MaxUint64)} {
		checkNumericComparisons(t, value, math.NaN(), 0, false)
		checkNumericComparisons(t, math.NaN(), value, 0, false)
	}
}

func checkNumericComparisons(t *testing.T, a, b any, order int, ordered bool) {
	t.Helper()
	env := map[string]any{"a": a, "b": b}
	for _, mode := range []struct {
		name  string
		shape any
	}{
		{"typed", env},
		{"dynamic", types.Map{"a": types.Any, "b": types.Any}},
	} {
		for _, operator := range []struct {
			name string
			want bool
		}{
			{"==", ordered && order == 0}, {"!=", !ordered || order != 0},
			{"<", ordered && order < 0}, {"<=", ordered && order <= 0},
			{">", ordered && order > 0}, {">=", ordered && order >= 0},
		} {
			program, err := expr.Compile("a "+operator.name+" b", expr.Env(mode.shape))
			if err != nil {
				t.Fatal(err)
			}
			got, err := expr.Run(program, env)
			if err != nil {
				t.Fatal(err)
			}
			if got != operator.want {
				t.Errorf("%s: %T(%v) %s %T(%v) = %v, want %v", mode.name, a, a, operator.name, b, b, got, operator.want)
			}
		}
	}
}
