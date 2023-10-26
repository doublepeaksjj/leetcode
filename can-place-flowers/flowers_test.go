package canplaceflowers

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCanPlaceFlowers(t *testing.T) {
	result := canPlaceFlowers([]int{1, 0, 0, 0, 1}, 1)
	require.Equal(t, true, result)

	result = canPlaceFlowers([]int{1, 0, 0, 0, 1}, 2)
	require.Equal(t, false, result)
}

func BenchmarkCanPlaceFlowers(b *testing.B) {
	flowerbed := [20000]int{}
	for i := 0; i < len(flowerbed); i++ {
		if i%4 == 0 {
			flowerbed[i] = 1
		}
	}
	for n := 0; n < b.N; n++ {
		result := canPlaceFlowers(flowerbed[:], 5000)
		if !result {
			panic("expected to be able to place all flowers")
		}
	}
}
