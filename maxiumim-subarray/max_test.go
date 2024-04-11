package maximumsubarray

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMaxSubArray(t *testing.T) {
	require.Equal(t, 6, maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	require.Equal(t, 1, maxSubArray([]int{1}))
	require.Equal(t, 23, maxSubArray([]int{5, 4, -1, 7, 8}))
	require.Equal(t, 0, maxSubArray([]int{}))
	require.Equal(t, -1, maxSubArray([]int{-1, -2, -3}))
	require.Equal(t, -1, maxSubArray([]int{-3, -2, -1}))
}
