package greatestcandies

func kidsWithCandies(candies []int, extraCandies int) []bool {
	maxVal := -1
	n := len(candies)
	for i := 0; i < n; i++ {
		if candies[i] > maxVal {
			maxVal = candies[i]
		}
	}

	result := [100]bool{}
	for i := 0; i < n; i++ {
		result[i] = candies[i]+extraCandies >= maxVal
	}
	return result[:n]
}
