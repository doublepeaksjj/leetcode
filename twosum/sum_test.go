package twosum

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTwoSum(t *testing.T) {
	ans := twoSum([]int{2, 7, 11, 15}, 9)
	require.Equal(t, []int{0, 1}, ans)

	ans = twoSum([]int{3, 2, 4}, 6)
	require.Equal(t, []int{1, 2}, ans)

	ans = twoSum([]int{3, 3}, 6)
	require.Equal(t, []int{0, 1}, ans)
}
