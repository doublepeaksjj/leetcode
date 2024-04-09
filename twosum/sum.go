package twosum

func twoSum(nums []int, target int) []int {

	meets := make(map[int]int)

	for i := range nums {
		if idx, ok := meets[nums[i]]; ok {
			return []int{idx, i}
		}
		meets[target-nums[i]] = i
	}

	for i := range nums {
		if idx, ok := meets[nums[i]]; ok && idx != i {
			return []int{i, idx}
		}
	}
	return nil
}
