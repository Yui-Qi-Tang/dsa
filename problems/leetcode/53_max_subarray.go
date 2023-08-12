package b75

import "math"

/*
53. Maximum Subarray
Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.

A subarray is a contiguous part of an array.

Example 1:

Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
Output: 6
Explanation: [4,-1,2,1] has the largest sum = 6
Example 2:

Input: nums = [1]
Output: 1
Example 3:

Input: nums = [5,4,-1,7,8]
Output: 23

use Kadane's algorithm. Kadane's algorithm is a dynamic programming approach to find the maximum subarray sum in an array
*/

func MaximumSubarrayv39(nums []int) int {
	g, l := nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		l = max(nums[i], l+nums[i])
		g = max(g, l)
	}

	return g
}

func MaximumSubarrayv38(nums []int) int {
	g, l := nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		l = max(nums[i], l+nums[i])
		g = max(g, l)
	}

	return g
}

func MaximumSubarrayv37(nums []int) int {
	g, l := nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		l = max(nums[i], l+nums[i])
		g = max(g, l)
	}

	return g
}

func MaximumSubarrayv36(nums []int) int {
	g, l := nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		l = max(nums[i], l+nums[i])
		g = max(g, l)
	}

	return g
}

func MaximumSubarrayv35(nums []int) int {
	g, l := nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		l = max(nums[i], l+nums[i])
		g = max(g, l)
	}

	return g
}

func MaximumSubarrayv34(nums []int) int {
	g, l := nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		l = max(nums[i], nums[i]+l)
		g = max(g, l)
	}

	return g
}

func MaximumSubarrayv33(nums []int) int {
	g, l := nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		l = max(nums[i], l+nums[i])
		g = max(g, l)
	}

	return g
}

func MaximumSubarrayv32(nums []int) int {
	g, l := nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		l = max(nums[i], l+nums[i])
		g = max(g, l)
	}

	return g
}

func MaximumSubarrayv31(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	global, local := nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		local = max(nums[i], nums[i]+local)
		global = max(global, local)
	}

	return global
}

func MaximumSubarrayv30(nums []int) int {
	g, l := nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		l = max(nums[i], nums[i]+l)
		g = max(g, l)
	}

	return g
}

func MaximumSubarrayv29(nums []int) int {
	g, l := nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		l = max(nums[i], l+nums[i])
		g = max(g, l)
	}

	return g
}

func MaximumSubarrayv28(nums []int) int {
	g, l := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		l = max(nums[i], nums[i]+l)
		g = max(g, l)
	}

	return g
}

func MaximumSubarrayv27(nums []int) int {
	g, l := nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		l = max(nums[i], l+nums[i])
		g = max(g, l)
	}

	return g
}

func MaximumSubarrayv26(nums []int) int {
	g, l := nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		l = max(nums[i], nums[i]+l)
		g = max(l, g)
	}

	return g
}

func MaximumSubarrayv25(nums []int) int {
	g, l := nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		l = max(nums[i], nums[i]+l)
		g = max(g, l)
	}

	return g
}

func MaximumSubarrayv24(nums []int) int {
	g, l := nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		l = max(nums[i], l+nums[i])
		g = max(l, g)
	}

	return g
}

func MaximumSubarrayv23(nums []int) int {
	g, l := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		l = max(nums[i], l+nums[i])
		g = max(g, l)
	}

	return g
}

func MaximumSubarrayv22(nums []int) int {
	g, l := nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		l = max(nums[i], l+nums[i])
		g = max(g, l)
	}

	return g
}

func MaximumSubarrayv21(nums []int) int {
	g, l := nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		l = max(nums[i], nums[i]+l)
		g = max(g, l)
	}

	return g
}

func MaximumSubarrayv20(nums []int) int {
	g, l := nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		l = max(nums[i], nums[i]+l)
		g = max(g, l)
	}

	return g
}

func MaximumSubarrayv19(nums []int) int {
	g, l := nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		l = max(nums[i], nums[i]+l)
		g = max(g, l)
	}
	return g
}

func MaximumSubarrayv18(nums []int) int {
	g, l := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		l = max(nums[i], nums[i]+l)
		g = max(g, l)
	}
	return g
}

func MaximumSubarrayv17(nums []int) int {
	g, l := nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		l = max(nums[i], nums[i]+l)
		g = max(g, l)
	}

	return g
}

func MaximumSubarrayv16(nums []int) int {
	g, l := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		l = max(nums[i], nums[i]+l)
		g = max(g, l)
	}

	return g
}

func MaximumSubarrayv15(nums []int) int {
	g, l := nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		l = max(nums[i], l+nums[i])
		g = max(l, g)
	}

	return g
}

func MaximumSubarrayv14(nums []int) int {
	g, l := nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		l = max(nums[i], nums[i]+l)
		g = max(g, l)
	}

	return g
}

func MaximumSubarrayv13(nums []int) int {
	g, l := nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		l = max(nums[i], nums[i]+l)
		g = max(g, l)
	}

	return g
}

func MaximumSubarrayv12(nums []int) int {
	g, l := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		l = max(nums[i], nums[i]+l)
		g = max(g, l)
	}

	return g
}

func MaximumSubarrayv11(nums []int) int {
	g, l := nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		l = max(nums[i], nums[i]+l)
		g = max(g, l)
	}

	return g
}

func MaximumSubarrayv10(nums []int) int {

	g, l := nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		l = max(nums[i], l+nums[i])
		g = max(g, l)
	}

	return g
}

func MaximumSubarrayv9(nums []int) int {
	g, l := nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		l = max(nums[i], nums[i]+l)
		g = max(g, l)
	}

	return g
}

func MaximumSubarrayv8(nums []int) int {
	globalOpt, localOpt := nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		localOpt = max(nums[i], localOpt+nums[i])
		globalOpt = max(localOpt, globalOpt)
	}

	return globalOpt
}

func MaximumSubarrayv7(nums []int) int {
	gm, lm := nums[0], nums[0] // gm: global max; lm: localmax

	for i := 1; i < len(nums); i++ {

		lm = max(nums[i], nums[i]+lm)
		gm = max(lm, gm)

	}
	return gm
}

func MaximumSubarrayv6(nums []int) int {
	globalMax, localMax := nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		localMax = max(nums[i], nums[i]+localMax)
		globalMax = max(localMax, globalMax)
	}

	return globalMax
}

func MaximumSubarrayv5(nums []int) int {
	localMax, globalMax := nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		localMax = max(nums[i], nums[i]+localMax)
		globalMax = max(globalMax, localMax)
	}

	return globalMax
}

func MaximumSubarrayv4(nums []int) int {
	globalMax, localMax := nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		// according to substring, so we only care about the max of each subarray
		// once the subarray+nums[i] is less than nums[i] which means we shouldn't hold the result(nums[i]+localMAX),
		// we must start with nums[i] to find the new subarray, and use globalMax to store the global max for answer
		//check if  nums[i] or localMax+nums[i] is the max one
		localMax = max(nums[i], localMax+nums[i])
		globalMax = max(localMax, globalMax)
	}

	return globalMax
}

func MaximumSubarrayv3(nums []int) int {

	localMax, globalMax := math.MinInt32, math.MinInt32

	for _, num := range nums {
		localMax = max(num, localMax+num)
		globalMax = max(localMax, globalMax)
	}

	return globalMax
}

func MaximumSubarrayv2(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}

	globalMax, localMax := nums[0], nums[0]

	for i := 1; i < n; i++ {
		localMax = max(nums[i], nums[i]+localMax)
		globalMax = max(localMax, globalMax)

	}

	return globalMax
}

func MaximumSubarrayv1(nums []int) int {

	n := len(nums)
	if n == 0 {
		return 0
	}

	result, curMax := nums[0], nums[0]
	for i := 1; i < n; i++ {
		curMax = max(nums[i], nums[i]+curMax)
		result = max(curMax, result)
	}

	return result
}
