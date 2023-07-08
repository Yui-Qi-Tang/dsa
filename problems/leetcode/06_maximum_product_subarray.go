package b75

/*
Given an integer array nums, find a contiguous non-empty subarray within the array that has the largest product,
and return the product.

The test cases are generated so that the answer will fit in a 32-bit integer.

A subarray is a contiguous subsequence of the array.

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

func MaximumProductSubarrayv6(nums []int) int {
	n := len(nums)

	if n == 0 {
		return 0
	}

	maxs, mins := make([]int, n), make([]int, n)
	maxs[0], mins[0] = nums[0], nums[0]

	for i := 1; i < n; i++ {

		maxs[i] = max(maxs[i-1], nums[i]*nums[i-1])
		mins[i] = min(mins[i-1], nums[i]*nums[i-1])

		if maxs[i] != 0 {
			maxs[i] = max(maxs[i], nums[i]*mins[i-1])
		}

	}

	return maxs[n-1]
}

func MaximumProductSubarrayv5(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	maxs, mins := make([]int, n), make([]int, n)
	maxs[0], mins[0] = nums[0], nums[0]

	for i := 1; i < n; i++ {
		maxs[i] = max(maxs[i-1], nums[i]*nums[i-1])
		mins[i] = min(mins[i-1], nums[i]*nums[i-1])

		if maxs[i] != 0 {
			maxs[i] = max(maxs[i], nums[i]*mins[i-1])
		}
	}

	return maxs[n-1]
}

func MaximumProductSubarrayv4(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	mins, maxs := make([]int, n), make([]int, n)
	mins[0], maxs[0] = nums[0], nums[0]

	for i := 1; i < n; i++ {
		maxs[i] = max(maxs[i-1], nums[i]*nums[i-1])
		mins[i] = min(mins[i-1], nums[i]*nums[i-1])

		if maxs[i] != 0 {
			maxs[i] = max(maxs[i], nums[i]*mins[i-1])
		}
	}

	return maxs[n-1]
}

func MaximumProductSubarrayv3(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	n := len(nums)
	maxs, mins := make([]int, n), make([]int, n)
	maxs[0], mins[0] = nums[0], nums[0]
	for i := 1; i < n; i++ {
		maxs[i] = max(maxs[i-1], nums[i]*nums[i-1])
		mins[i] = min(mins[i-1], nums[i]*nums[i-1])
		if maxs[i] != 0 {
			maxs[i] = max(maxs[i], nums[i]*mins[i-1])
		}
	}

	return maxs[n-1]
}

func MaximumProductSubarrayv2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	mins := make([]int, len(nums))
	maxs := make([]int, len(nums))
	maxs[0], mins[0] = nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		maxs[i] = max(maxs[i-1], nums[i]*nums[i-1])
		mins[i] = min(mins[i-1], nums[i]*nums[i-1])

		if maxs[i] != 0 {
			maxs[i] = max(nums[i]*mins[i-1], maxs[i])
		}
	}

	return maxs[len(nums)-1]
}

func MaximumProductSubarrayv1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0] // nums[0] * 1
	}

	n := len(nums)
	maxs := make([]int, n)
	mins := make([]int, n)
	maxs[0] = nums[0]
	mins[0] = nums[0]
	for i := 1; i < n; i++ {
		maxs[i] = max(nums[i]*nums[i-1], maxs[i-1])
		mins[i] = min(nums[i]*nums[i-1], mins[i-1])

		if maxs[i] != 0 {
			maxs[i] = max(nums[i]*mins[i-1], maxs[i])
		}
	}

	return maxs[n-1]
}

func MaximumProductSubarray(nums []int) int {
	n := len(nums)

	if n == 0 {
		return 0
	}

	if n == 1 {
		return nums[0]
	}

	maxs := make([]int, n)
	mins := make([]int, n)
	maxs[0] = nums[0] // game 0: nums[0] * 1 wins the game
	mins[0] = nums[0]

	for i := 1; i < n; i++ {
		maxs[i] = max(nums[i]*nums[i-1], maxs[i-1])
		mins[i] = min(nums[i]*nums[i-1], mins[i-1])

		if maxs[i] != 0 {
			maxs[i] = max(maxs[i], nums[i]*mins[i-1])
		}

	}
	return maxs[n-1]
}
