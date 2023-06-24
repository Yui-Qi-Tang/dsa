package b75

import (
	"fmt"
)

/*
238. Product of Array Except Self

Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i].

The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.

You must write an algorithm that runs in O(n) time and without using the division operation.



Example 1:

    Input: nums = [1,2,3,4]
    Output: [24,12,8,6]

Example 2:

    Input: nums = [-1,1,0,-3,3]
    Output: [0,0,9,0,0]

Constraints:

2 <= nums.length <= 105
-30 <= nums[i] <= 30
The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.


Follow up: Can you solve the problem in O(1) extra space complexity? (The output array does not count as extra space for space complexity analysis.)
*/

func ProductOfArrayExceptSelf40(nums []int) []int {
	result := make([]int, len(nums))
	prefix := 1

	for i, num := range nums {
		result[i] = prefix
		prefix *= num
	}

	suffix := 1
	for i := len(result) - 1; i >= 0; i-- {
		result[i] *= suffix
		suffix *= nums[i]
	}

	return result
}

func ProductOfArrayExceptSelf39(nums []int) []int {
	result := make([]int, len(nums))

	prefix := 1
	for i, num := range nums {
		result[i] = prefix
		prefix *= num
	}

	suffix := 1
	for i := len(result) - 1; i >= 0; i-- {
		result[i] *= suffix
		suffix *= nums[i]
	}

	return result
}

func ProductOfArrayExceptSelf38(nums []int) []int {
	result := make([]int, len(nums))

	prefix := 1
	for i := range nums {
		result[i] = prefix
		prefix *= nums[i]
	}

	suffix := 1
	for i := len(nums) - 1; i >= 0; i-- {
		result[i] *= suffix
		suffix *= nums[i]
	}

	return result
}

func ProductOfArrayExceptSelf37(nums []int) []int {
	result := make([]int, len(nums))

	prefix := 1
	for i, num := range nums {
		result[i] = prefix
		prefix *= num
	}

	suffix := 1
	for i := len(result) - 1; i >= 0; i-- {
		result[i] *= suffix
		suffix *= nums[i]
	}

	return result
}

func ProductOfArrayExceptSelf36(nums []int) []int {
	result := make([]int, len(nums))
	prefix := 1
	for i, num := range nums {
		result[i] = prefix
		prefix *= num
	}

	suffix := 1
	for i := len(result) - 1; i >= 0; i-- {
		result[i] *= suffix
		suffix *= nums[i]
	}

	return result
}

func ProductOfArrayExceptSelf35(nums []int) []int {
	result := make([]int, len(nums))

	prefix := 1
	for i, num := range nums {
		result[i] = prefix
		prefix *= num
	}

	suffix := 1
	for i := len(nums) - 1; i >= 0; i-- {
		result[i] *= suffix
		suffix *= nums[i]
	}

	return result
}

func ProductOfArrayExceptSelf34(nums []int) []int {
	result := make([]int, len(nums))
	prefix := 1
	for i, num := range nums {
		result[i] = prefix
		prefix *= num
	}

	suffix := 1
	for i := len(result) - 1; i >= 0; i-- {
		result[i] *= suffix
		suffix *= nums[i]
	}

	return result
}

func ProductOfArrayExceptSelf33(nums []int) []int {
	result := make([]int, len(nums))

	prefix := 1
	for i := range result {
		result[i] = prefix
		prefix *= nums[i]
	}

	suffix := 1
	for i := len(nums) - 1; i >= 0; i-- {
		result[i] *= suffix
		suffix *= nums[i]
	}

	return result
}

func ProductOfArrayExceptSelf32(nums []int) []int {
	result := make([]int, len(nums))

	prefix := 1
	for i, num := range nums {
		result[i] = prefix
		prefix *= num
	}

	suffix := 1
	for i := len(nums) - 1; i >= 0; i-- {
		result[i] *= suffix
		suffix *= nums[i]
	}

	return result
}

func ProductOfArrayExceptSelf31(nums []int) []int {
	result := make([]int, len(nums))

	prefix := 1
	for i, num := range nums {
		result[i] = prefix
		prefix *= num
	}

	suffix := 1
	for i := len(nums) - 1; i >= 0; i-- {
		result[i] *= suffix
		suffix *= nums[i]
	}

	return result
}

func ProductOfArrayExceptSelf30(nums []int) []int {
	result := make([]int, len(nums))

	prefix := 1
	for i := range nums {
		result[i] = prefix
		prefix *= nums[i]
	}

	suffix := 1
	for i := len(nums) - 1; i >= 0; i-- {
		result[i] *= suffix
		suffix *= nums[i]
	}

	return result
}

func ProductOfArrayExceptSelf29(nums []int) []int {
	result := make([]int, len(nums))
	prefix := 1

	for i, num := range nums {
		result[i] = prefix
		prefix *= num
	}

	suffix := 1
	for i := len(nums) - 1; i >= 0; i-- {
		result[i] *= suffix
		suffix *= nums[i]
	}

	return result
}

func ProductOfArrayExceptSelf28(nums []int) []int {
	result := make([]int, len(nums))

	prefix := 1
	for i, num := range nums {
		result[i] = prefix
		prefix *= num
	}

	suffix := 1
	for i := len(nums) - 1; i >= 0; i-- {
		result[i] *= suffix
		suffix *= nums[i]
	}

	return result
}

func ProductOfArrayExceptSelf27(nums []int) []int {
	result := make([]int, len(nums))

	prefix := 1
	for i := 0; i < len(nums); i++ {
		result[i] = prefix
		prefix *= nums[i]
	}

	suffix := 1
	for i := len(nums) - 1; i >= 0; i-- {
		result[i] *= suffix
		suffix *= nums[i]
	}
	return result
}

func ProductOfArrayExceptSelf26(nums []int) []int {
	result := make([]int, len(nums))

	prefix := 1
	for i := 0; i < len(nums); i++ {
		result[i] = prefix
		prefix *= nums[i]
	}

	suffix := 1
	for i := len(nums) - 1; i >= 0; i-- {
		result[i] *= suffix
		suffix *= nums[i]
	}

	return result
}

func ProductOfArrayExceptSelf25(nums []int) []int {
	result := make([]int, len(nums))

	prefix := 1
	for i := 0; i < len(nums); i++ {
		result[i] = prefix
		prefix *= nums[i]
	}

	sufiix := 1
	for i := len(nums) - 1; i >= 0; i-- {
		result[i] *= sufiix
		sufiix *= nums[i]
	}

	return result
}

func ProductOfArrayExceptSelf24(nums []int) []int {
	result := make([]int, len(nums))

	prefix := 1
	for i := 0; i < len(nums); i++ {
		result[i] = prefix
		prefix *= nums[i]
	}

	suffix := 1
	for i := len(nums) - 1; i >= 0; i-- {
		result[i] *= suffix
		suffix *= nums[i]
	}

	return result
}

func ProductOfArrayExceptSelf23(nums []int) []int {
	result := make([]int, len(nums))
	prefix := 1
	for i := 0; i < len(nums); i++ {
		result[i] = prefix
		prefix *= nums[i]
	}

	suffix := 1
	for i := len(nums) - 1; i >= 0; i-- {
		result[i] *= suffix
		suffix *= nums[i]
	}
	return result
}

func ProductOfArrayExceptSelf22(nums []int) []int {
	result := make([]int, len(nums))

	prefix := 1
	for i := 0; i < len(nums); i++ {
		result[i] = prefix
		prefix *= nums[i]
	}

	suffix := 1
	for i := len(nums) - 1; i >= 0; i-- {
		result[i] *= suffix
		suffix *= nums[i]
	}

	return result
}

func ProductOfArrayExceptSelf21(nums []int) []int {
	result := make([]int, len(nums))

	prefix := 1
	for i, num := range nums {
		result[i] = prefix
		prefix *= num
	}

	suffix := 1
	for i := len(nums) - 1; i >= 0; i-- {
		result[i] *= suffix
		suffix *= nums[i]
	}

	return result
}

func ProductOfArrayExceptSelf20(nums []int) []int {
	result := make([]int, len(nums))
	prefix := 1
	for i := range result {
		result[i] = prefix
		prefix *= nums[i]
	}

	suffix := 1
	for i := len(nums) - 1; i >= 0; i-- {
		result[i] *= suffix
		suffix *= nums[i]
	}

	return result
}

func ProductOfArrayExceptSelf19(nums []int) []int {
	result := make([]int, len(nums))

	prefix := 1
	for i := range result {
		result[i] = prefix
		prefix *= nums[i]
	}

	suffix := 1
	for i := len(nums) - 1; i >= 0; i-- {
		result[i] *= suffix
		suffix *= nums[i]
	}

	return result
}

func ProductOfArrayExceptSelf18(nums []int) []int {
	result := make([]int, len(nums))

	prefix := 1
	for i := 0; i < len(nums); i++ {
		result[i] = prefix
		prefix *= nums[i]
	}

	suffix := 1
	for i := len(nums) - 1; i >= 0; i-- {
		result[i] *= suffix
		suffix *= nums[i]
	}

	return result
}

func ProductOfArrayExceptSelf17(nums []int) []int {
	result := make([]int, len(nums))

	prefix := 1
	for i := 0; i < len(nums); i++ {
		result[i] = prefix
		prefix *= nums[i]
	}

	suffix := 1
	for i := len(nums) - 1; i >= 0; i-- {
		result[i] *= suffix
		suffix *= nums[i]
	}

	return result
}

func ProductOfArrayExceptSelf16(nums []int) []int {

	result := make([]int, len(nums))

	prefix := 1
	for i := 0; i < len(nums); i++ {
		result[i] = prefix
		prefix *= nums[i]
	}

	suffix := 1
	for i := len(nums) - 1; i >= 0; i-- {
		result[i] *= suffix
		suffix *= nums[i]
	}

	return result
}

func ProductOfArrayExceptSelf15(nums []int) []int {
	n := len(nums)
	result := make([]int, n)

	for i := range result {
		result[i] = 1
	}

	for i := 1; i < n; i++ {
		result[i] = nums[i-1] * result[i-1]
	}

	suffix := 1
	for i := n - 1; i >= 0; i-- {
		result[i] *= suffix
		suffix *= nums[i]
	}

	return result
}

func ProductOfArrayExceptSelf14(nums []int) []int {
	n := len(nums)
	result := make([]int, n)
	for i := range result {
		result[i] = 1
	}

	for i := 1; i < n; i++ {
		result[i] = result[i-1] * nums[i-1]
	}

	suffix := 1
	for i := n - 1; i >= 0; i-- {
		result[i] *= suffix
		suffix *= nums[i]
	}

	return result
}

func ProductOfArrayExceptSelf13(nums []int) []int {
	result := make([]int, len(nums))

	for i := range result {
		result[i] = 1
	}

	for i := 1; i < len(nums); i++ {
		result[i] = nums[i-1] * result[i-1]
	}

	suffix := 1

	for i := len(nums) - 1; i >= 0; i-- {
		result[i] *= suffix
		suffix *= nums[i]
	}

	return result
}

func ProductOfArrayExceptSelf12(nums []int) []int {
	result := make([]int, len(nums))

	for i := range result {
		result[i] = 1
	}

	for i := 1; i < len(nums); i++ {
		result[i] *= nums[i-1] * result[i-1]
	}

	suffix := 1
	for i := len(nums) - 1; i >= 0; i-- {
		result[i] *= suffix
		suffix *= nums[i]
	}

	return result
}

func ProductOfArrayExceptSelf11(nums []int) []int {
	n := len(nums)

	ans := make([]int, n)
	for i := range ans {
		ans[i] = 1
	}

	for i := 1; i < n; i++ {
		ans[i] *= nums[i-1] * ans[i-1]
	}

	fmt.Println(ans)

	suffix := 1
	for i := n - 2; i >= 0; i-- {
		ans[i] *= nums[i+1] * suffix
		suffix *= nums[i+1]
	}

	return ans
}

func ProductOfArrayExceptSelf10(nums []int) []int {

	output := make([]int, len(nums))

	for i := range output {
		output[i] = 1
	}

	for i := 1; i < len(nums); i++ {
		output[i] = output[i-1] * nums[i-1]
	}

	suffix := 1
	for i := len(nums) - 1; i >= 0; i-- {
		output[i] *= suffix
		suffix *= nums[i]
	}

	return output
}

func ProductOfArrayExceptSelf9(nums []int) []int {
	n := len(nums)
	lp := make([]int, n)
	rp := make([]int, n)

	lp[0] = 1
	for i := 1; i < n; i++ {
		lp[i] = nums[i-1] * lp[i-1]
	}

	rp[n-1] = 1
	for i := n - 2; i >= 0; i-- {
		rp[i] = nums[i+1] * rp[i+1]
	}

	for i := range lp {
		lp[i] *= rp[i]
	}

	return lp
}

func ProductOfArrayExceptSelf8(nums []int) []int {
	n := len(nums)
	lp := make([]int, n)
	rp := make([]int, n)

	lp[0] = 1
	for i := 1; i < n; i++ {
		lp[i] = nums[i-1] * lp[i-1]
	}

	rp[n-1] = 1
	for i := n - 2; i >= 0; i-- {
		rp[i] = nums[i+1] * rp[i+1]
	}

	for i := 0; i < n; i++ {
		lp[i] = lp[i] * rp[i]
	}

	return lp
}

func ProductOfArrayExceptSelf7(nums []int) []int {
	n := len(nums)
	lp := make([]int, n)
	rp := make([]int, n)

	lp[0] = 1
	for i := 1; i < n; i++ {
		lp[i] = nums[i-1] * lp[i-1]
	}

	rp[n-1] = 1
	for i := n - 2; i >= 0; i-- {
		rp[i] = nums[i+1] * rp[i+1]
	}

	for i := 0; i < n; i++ {
		lp[i] *= rp[i]
	}

	return lp
}

func ProductOfArrayExceptSelf6(nums []int) []int {

	n := len(nums)
	lprod := make([]int, n)
	rprod := make([]int, n)

	rprod[n-1] = 1
	for i := n - 2; i >= 0; i-- {
		rprod[i] = nums[i+1] * rprod[i+1]
	}

	lprod[0] = 1
	for i := 1; i < n; i++ {
		lprod[i] = nums[i-1] * lprod[i-1]
	}

	for i := 0; i < n; i++ {
		lprod[i] *= rprod[i]
	}

	return lprod
}

func ProductOfArrayExceptSelf5(nums []int) []int {
	n := len(nums)
	lprod := make([]int, n)
	rprod := make([]int, n)

	lprod[0] = 1
	for i := 1; i < n; i++ {
		lprod[i] = nums[i-1] * lprod[i-1]
	}

	rprod[n-1] = 1
	for i := n - 2; i >= 0; i-- {
		rprod[i] = nums[i+1] * rprod[i+1]
	}

	for i := 0; i < n; i++ {
		lprod[i] *= rprod[i]
	}

	return lprod
}

func ProductOfArrayExceptSelf4(nums []int) []int {
	n := len(nums)

	lr := make([]int, n)
	rl := make([]int, n)

	lr[0] = 1
	for i := 1; i < n; i++ {
		lr[i] = nums[i-1] * lr[i-1]
	}

	rl[n-1] = 1
	for i := n - 2; i >= 0; i-- {
		rl[i] = nums[i+1] * rl[i+1]
	}

	for i := 0; i < n; i++ {
		lr[i] *= rl[i]
	}

	return lr
}

func ProductOfArrayExceptSelf3(nums []int) []int {
	n := len(nums)
	lr := make([]int, n)
	rl := make([]int, n)

	lr[0] = 1
	for i := 1; i < n; i++ {
		lr[i] = lr[i-1] * nums[i-1]
	}

	rl[n-1] = 1
	for i := n - 2; i >= 0; i-- {
		rl[i] = nums[i+1] * rl[i+1]
	}

	for i := range nums {
		lr[i] = lr[i] * rl[i]
	}

	return lr
}

func ProductOfArrayExceptSelf(nums []int) []int {

	// product of nums[i] = product of [0,...,i-1] * product of [i+1,...n] where n is the length of nums

	n := len(nums)
	result := make([]int, n)

	// compute [n,....0] part
	product := 1 // base
	for i := n - 1; i >= 0; i-- {
		result[i] = product
		product *= nums[i]
	}

	// compute [0,...i-1]
	product = 1
	for i := 0; i < n; i++ {
		result[i] *= product
		product *= nums[i]
	}

	return result
}

func ProductOfArrayExceptSelf2(nums []int) []int {

	n := len(nums)

	r := make([]int, n)

	p := 1
	for i := n - 1; i >= 0; i-- {
		r[i] = p
		p *= nums[i]
	}

	p = 1
	for i := 0; i < n; i++ {
		r[i] *= p
		p *= nums[i]
	}

	return r
}
