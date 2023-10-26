package greatestcandies

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestKidsWithCandies(t *testing.T) {
	result := kidsWithCandies([]int{2, 3, 5, 1, 3}, 3)
	require.Equal(t, []bool{true, true, true, false, true}, result)

	result = kidsWithCandies([]int{4, 2, 1, 1, 2}, 1)
	require.Equal(t, []bool{true, false, false, false, false}, result)

	result = kidsWithCandies([]int{12, 1, 12}, 10)
	require.Equal(t, []bool{true, false, true}, result)
}

func BenchmarkCandies(b *testing.B) {
	candies := [100]int{}
	for i := 0; i < len(candies); i++ {
		candies[i] = i
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		kidsWithCandies(candies[:], 50)
	}
}
