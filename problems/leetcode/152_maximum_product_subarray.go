package b75

func maxProductv2(nums []int) int {
	currMax, currMin, result := nums[0], nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		num := nums[i]
		if num < 0 {
			currMax, currMin = currMin, currMax
		}

		currMax = max(num, currMax*num)
		currMin = min(num, currMin*num)
		result = max(currMax, result)
	}

	return result
}

func maxProductv1(nums []int) int {
	currMax, currMin, result := nums[0], nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		num := nums[i]
		if num < 0 {
			currMax, currMin = currMin, currMax
		}
		currMax = max(num, currMax*num)
		currMin = min(num, currMin*num)
		result = max(result, currMax)
	}

	return result
}
