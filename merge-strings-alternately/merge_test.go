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

func BenchmarkMerge(b *testing.B) {
	for n := 0; n < b.N; n++ {
		mergeAlternately("abcdefghijklmnopq", "123456789012345678901234567890")
	}
}
