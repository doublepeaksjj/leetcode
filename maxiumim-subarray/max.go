package maximumsubarray

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	curSum := nums[0]
	maxSum := nums[0]
	for i := 1; i < len(nums); i++ {
		num := nums[i]
		if curSum < 0 {
			curSum = num
		} else {
			curSum += num
		}
		if curSum > maxSum {
			maxSum = curSum
		}
	}
	return maxSum
}
