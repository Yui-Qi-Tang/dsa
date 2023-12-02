package b75

/*
152. Maximum Product Subarray

Given an integer array nums, find a subarray that has the largest product, and return the product.

The test cases are generated so that the answer will fit in a 32-bit integer.

Example 1:

Input: nums = [2,3,-2,4]
Output: 6
Explanation: [2,3] has the largest product 6.
Example 2:

Input: nums = [-2,0,-1]
Output: 0
Explanation: The result cannot be 2, because [-2,-1] is not a subarray.


Constraints:

1 <= nums.length <= 2 * 104
-10 <= nums[i] <= 10
The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.
*/

func maxProductv17(nums []int) int {
	result := nums[0]
	currMax := nums[0]
	currMin := nums[0]

	for i := 1; i < len(nums); i++ {
		num := nums[i]
		if num < 0 {
			currMax, currMin = currMin, currMax
		}

		currMax = max(num, num*currMax)
		currMin = min(num, currMin*num)
		result = max(result, currMax)
	}

	return result
}

func maxProductv16(nums []int) int {
	curMax, curMin, result := nums[0], nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		num := nums[i]
		if num < 0 {
			curMax, curMin = curMin, curMax
		}
		curMax = max(num, num*curMax)
		curMin = min(num, num*curMin)
		result = max(curMax, result)
	}

	return result
}

func maxProductv15(nums []int) int {
	curMax, curMin, result := nums[0], nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		num := nums[i]
		if num < 0 {
			curMax, curMin = curMin, curMax
		}

		curMax = max(num, num*curMax)
		curMin = min(num, num*curMin)
		result = max(result, curMax)
	}

	return result
}

func maxProductv14(nums []int) int {
	curMax, curMin, result := nums[0], nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		num := nums[i]
		if num < 0 {
			curMax, curMin = curMin, curMax
		}

		curMax = max(num, num*curMax)
		curMin = min(num, num*curMin)
		result = max(result, curMax)
	}

	return result
}

func maxProductv13(nums []int) int {
	currMax, currMin, result := nums[0], nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		num := nums[i]
		if num < 0 {
			currMax, currMin = currMin, currMax
		}

		currMax = max(num, num*currMax)
		currMin = min(num, num*currMin)
		result = max(currMax, result)
	}

	return result
}

func maxProductv12(nums []int) int {
	currMax, currMin, result := nums[0], nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		num := nums[i]
		if num < 0 {
			currMax, currMin = currMin, currMax
		}

		currMax = max(num, num*currMax)
		currMin = min(num, num*currMin)
		result = max(result, currMax)
	}

	return result
}

func maxProductv11(nums []int) int {
	currMax, currMin, result := nums[0], nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		num := nums[i]
		if num < 0 {
			currMax, currMin = currMin, currMax
		}

		currMax = max(num, num*currMax)
		currMin = min(num, num*currMin)
		result = max(result, currMax)
	}

	return result
}

func maxProductv10(nums []int) int {
	currMax, currMin, result := nums[0], nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		num := nums[i]
		if num < 0 {
			currMax, currMin = currMin, currMax
		}
		currMax = max(num, num*currMax)
		currMin = min(num, num*currMin)
		result = max(result, currMax)
	}

	return result
}

func maxProductv9(nums []int) int {
	curMax, curMin, result := nums[0], nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		num := nums[i]
		if num < 0 {
			curMax, curMin = curMin, curMax
		}

		curMax = max(num, num*curMax)
		curMin = min(num, num*curMin)
		result = max(result, curMax)
	}

	return result
}

func maxProductv8(nums []int) int {
	curMax, curMin, result := nums[0], nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		num := nums[i]
		if num < 0 {
			curMax, curMin = curMin, curMax
		}

		curMax = max(num, num*curMax)
		curMin = min(num, num*curMin)
		result = max(result, curMax)
	}

	return result
}

func maxProductv7(nums []int) int {
	curMax, curMin, result := nums[0], nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		num := nums[i]
		if num < 0 {
			curMax, curMin = curMin, curMax
		}
		curMax = max(num, num*curMax)
		curMin = min(num, num*curMin)
		result = max(result, curMax)
	}

	return result
}

func maxProductv6(nums []int) int {
	curMax, curMin, result := nums[0], nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		num := nums[i]
		if num < 0 {
			curMax, curMin = curMin, curMax
		}

		curMax = max(num, num*curMax)
		curMin = min(num, num*curMin)
		result = max(result, curMax)
	}

	return result
}

func maxProductv5(nums []int) int {
	curMax, curMin, result := nums[0], nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		num := nums[i]
		if num < 0 {
			curMax, curMin = curMin, curMax
		}

		curMax = max(curMax*num, num)
		curMin = min(curMin*num, num)
		result = max(curMax, result)
	}

	return result
}

func maxProductv4(nums []int) int {
	curMax, curMin, result := nums[0], nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		num := nums[i]
		if num < 0 {
			curMax, curMin = curMin, curMax
		}

		curMax = max(num, num*curMax)
		curMin = min(num, num*curMin)
		result = max(result, curMax)
	}

	return result
}

func maxProductv3(nums []int) int {
	currMax, currMin, result := nums[0], nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		num := nums[i]
		if num < 0 {
			currMax, currMin = currMin, currMax
		}

		currMax = max(num, num*currMax)
		currMin = min(num, num*currMin)
		result = max(result, currMax)
	}

	return result
}

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
