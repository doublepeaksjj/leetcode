package canplaceflowers

func canPlaceFlowers(flowerbed []int, n int) bool {
	placed := 0
	lenFb := len(flowerbed)
	for i := 0; i < lenFb && placed < n; i++ {
		canPlace := (i == 0 || flowerbed[i-1] == 0) && (i == lenFb-1 || flowerbed[i+1] == 0) && flowerbed[i] == 0
		if canPlace {
			placed++
			i++
		}
	}
	return placed == n
}
