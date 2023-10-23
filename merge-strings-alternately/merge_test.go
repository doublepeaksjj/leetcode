package leetmerge

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMerge(t *testing.T) {
	result := mergeAlternately("abc", "qpr")
	require.Equal(t, "aqbpcr", result)

	result = mergeAlternately("ab", "qprs")
	require.Equal(t, "aqbprs", result)

	result = mergeAlternately("abcd", "qp")
	require.Equal(t, "aqbpcd", result)
}
