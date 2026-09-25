package issue990

import (
	"testing"

	"github.com/expr-lang/expr"
	"github.com/expr-lang/expr/internal/testify/require"
)

type Outer struct {
	Inner map[string]struct{} `expr:"bar"`
}

type Env struct {
	Outer Outer `expr:"foo"`
}

// TestIssue990 tests that the `in` operator on a struct respects the `expr`
// struct tag the same way plain field access does: once a field is renamed
// via `expr:"..."`, only the tagged name identifies it, and the original Go
// field name is invisible.
func TestIssue990(t *testing.T) {
	env := Env{Outer: Outer{Inner: map[string]struct{}{}}}

	out, err := expr.Eval(`"bar" in foo`, env)
	require.NoError(t, err)
	require.Equal(t, true, out)

	out, err = expr.Eval(`"Inner" in foo`, env)
	require.NoError(t, err)
	require.Equal(t, false, out)
}
