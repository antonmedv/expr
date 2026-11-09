package issues

import (
	"testing"

	"github.com/expr-lang/expr"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIssue830(t *testing.T) {
	program, err := expr.Compile("varNotExist", expr.AllowUndefinedVariables(), expr.AsBool())
	require.NoError(t, err)

	output, err := expr.Run(program, map[string]interface{}{})
	require.NoError(t, err)

	// The user expects output to be false (bool), but gets nil.
	assert.Equal(t, false, output)
}
