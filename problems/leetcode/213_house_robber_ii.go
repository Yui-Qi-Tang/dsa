package b75

import (
	"fmt"
)

/*
213. House Robber II


You are a professional robber planning to rob houses along a street.
Each house has a certain amount of money stashed.
All houses at this place are arranged in a circle.
 That means the first house is the neighbor of the last one. Meanwhile,
 adjacent houses have a security system connected,
 and it will automatically contact the police if
 two adjacent houses were broken into on the same night.

Given an integer array nums representing the amount of money of each house,
return the maximum amount of money you can rob tonight without alerting the police.



Example 1:

Input: nums = [2,3,2]
Output: 3
Explanation: You cannot rob house 1 (money = 2) and then rob house 3 (money = 2), because they are adjacent houses.
Example 2:

Input: nums = [1,2,3,1]
Output: 4
Explanation: Rob house 1 (money = 1) and then rob house 3 (money = 3).
Total amount you can rob = 1 + 3 = 4.
Example 3:

Input: nums = [1,2,3]
Output: 3


Constraints:

1 <= nums.length <= 100
0 <= nums[i] <= 1000
*/

func houseRobberIIv22(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	rob := func(hs []int) int {
		r1, r2 := 0, 0
		for _, h := range hs {
			tmp := max(r1+h, r2)
			r1 = r2
			r2 = tmp
		}
		return r2
	}

	return max(rob(nums[1:]), rob(nums[:len(nums)-1]))
}

func houseRobberIIv21(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	rob := func(hs []int) int {
		r1, r2 := 0, 0
		for _, h := range hs {
			tmp := max(r1+h, r2)
			r1 = r2
			r2 = tmp
		}
		return r2
	}

	return max(rob(nums[1:]), rob(nums[:len(nums)-1]))
}

func houseRobberIIv20(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	rob := func(hs []int) int {
		r1, r2 := 0, 0
		for _, h := range hs {
			tmp := max(r1+h, r2)
			r1 = r2
			r2 = tmp
		}
		return r2
	}
	return max(rob(nums[1:]), rob(nums[:len(nums)-1]))
}

func houseRobberIIv19(nums []int) int {

	if len(nums) == 1 {
		return nums[0]
	}

	rob := func(hs []int) int {
		r1, r2 := 0, 0

		for _, h := range hs {
			tmp := max(r1+h, r2)
			r1 = r2
			r2 = tmp
		}

		return r2
	}

	return max(rob(nums[1:]), rob(nums[:len(nums)-1]))
}

func houseRobberIIv18(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	rob := func(hs []int) int {
		r1, r2 := 0, 0

		for _, h := range hs {
			tmp := max(r1+h, r2)
			r1 = r2
			r2 = tmp
		}

		return r2
	}

	return max(rob(nums[1:]), rob(nums[:len(nums)-1]))
}

func houseRobberIIv17(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	rob := func(houses []int) int {
		r1, r2 := 0, 0
		for _, h := range houses {
			tmp := max(r2, r1+h)
			r1 = r2
			r2 = tmp
		}
		return r2
	}

	return max(rob(nums[1:]), rob(nums[:len(nums)-1]))
}

func houseRobberIIv16(nums []int) int {

	if len(nums) == 1 {
		return nums[0]
	}

	rob := func(house []int) int {
		r1, r2 := 0, 0

		for _, h := range house {
			tmp := max(r1+h, r2)
			r1 = r2
			r2 = tmp
		}

		return r2
	}
	return max(rob(nums[1:]), rob(nums[:len(nums)-1]))
}

func houseRobberIIv15(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	rob := func(ns []int) int {
		r1, r2 := 0, 0

		for i := range ns {
			tmp := max(r1+ns[i], r2)
			r1 = r2
			r2 = tmp
		}

		return r2
	}

	return max(rob(nums[1:]), rob(nums[:len(nums)-1]))
}

func houseRobberIIv14(nums []int) int {

	if nums == nil {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}

	rob := func(in []int) int {
		r1, r2 := 0, 0

		for _, v := range in {
			tmp := max(r1+v, r2)
			r1 = r2
			r2 = tmp
		}
		return r2
	}

	return max(rob(nums[1:]), rob(nums[:len(nums)-1]))

}

func houseRobberIIv13(nums []int) int {
	if nums == nil {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}

	rob := func(ns []int) int {
		r1, r2 := 0, 0

		for _, n := range ns {
			tmp := max(r1+n, r2)
			r1 = r2
			r2 = tmp
		}

		return r2
	}

	return max(rob(nums[1:]), rob(nums[:len(nums)-1]))

}

func houseRobberIIv12(nums []int) int {
	if nums == nil {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}

	rob := func(in []int) int {
		r1, r2 := 0, 0
		for _, v := range in {
			tmp := max(r1+v, r2)
			r1 = r2
			r2 = tmp
		}
		return r2
	}

	return max(rob(nums[1:]), rob(nums[:len(nums)-1]))
}

func houseRobberIIv11(nums []int) int {
	if nums == nil {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}

	rob := func(in []int) int {
		r1, r2 := 0, 0

		for _, v := range in {
			tmp := max(r1+v, r2)
			r1 = r2
			r2 = tmp
		}

		return r2
	}

	return max(rob(nums[1:]), rob(nums[:len(nums)-1]))
}

func houseRobberIIv10(nums []int) int {

	if nums == nil {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}

	rob := func(in []int) int {
		f1, f2 := 0, 0

		for _, v := range in {
			tmp := max(f1+v, f2)
			f1 = f2
			f2 = tmp
		}

		return f2
	}

	return max(rob(nums[1:]), rob(nums[:len(nums)-1]))
}

func houseRobberIIv9(nums []int) int {
	if nums == nil {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}

	rob := func(n []int) int {
		r1, r2 := 0, 0

		for _, v := range n {
			tmp := max(r1+v, r2)
			r1 = r2
			r2 = tmp
		}

		return r2
	}

	return max(rob(nums[1:]), rob(nums[:len(nums)-1]))

}

func houseRobberIIv8(nums []int) int {
	if nums == nil {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}

	rob := func(nums []int) int {
		r1, r2 := 0, 0

		for _, num := range nums {
			tmp := max(r1+num, r2)
			r1 = r2
			r2 = tmp
		}
		return r2
	}

	return max(
		rob(nums[1:]), rob(nums[:len(nums)-1]),
	)
}

func houseRobberIIv7(nums []int) int {
	if nums == nil {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}

	rob := func(nums []int) int {
		r1, r2 := 0, 0

		for _, num := range nums {
			tmp := max(r1+num, r2)
			r1 = r2
			r2 = tmp
		}

		return r2
	}

	return max(
		rob(nums[1:]),           // start from index 1 -> ignore the first element
		rob(nums[:len(nums)-1]), // ignore the last element
	)
}

func houseRobberIIv6(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	rob := func(in []int) int {

		n := len(in)

		if n == 1 {
			return in[0]
		}

		if n == 2 {
			return max(in[0], in[1])
		}

		m := make([]int, n)

		m[0], m[1] = in[0], max(in[0], in[1])

		for i := 2; i < n; i++ {
			m[i] = max(m[i-1], in[i]+m[i-2])
		}

		fmt.Println(m)
		return m[n-1]
	}

	return max(rob(nums[1:]), rob(nums[:len(nums)-1]))
}

func houseRobberIIv5(nums []int) int {

	if nums == nil {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	rob := func(n []int) int {
		r1, r2 := 0, 0

		for _, v := range n {
			tmp := max(r1+v, r2)
			r1 = r2
			r2 = tmp
		}
		return r2
	}

	return max(rob(nums[1:]), rob(nums[:len(nums)-1]))
}

// RobIIBetter is from leetcode discuss: https://leetcode.com/problems/house-robber-ii/discuss/538040/Go
func RobIIBetter(nums []int) int {
	n := len(nums)

	if n == 0 {
		return 0
	}

	if n == 1 {
		return nums[0]
	}

	/*
		given a slice nums
		    case 1: start from 0 without latest element => nums[:n-1] = nums[0,.....n-1]
		    case 2: start from 1 to latest element => nums[1:] = nums[1,.....n]
		    where n is length of nums
	*/
	return max(robIIHelper(nums[:n-1]), robIIHelper(nums[1:]))
}

func robIIHelper(nums []int) int {

	n := len(nums)

	if n == 1 {
		return nums[0]
	}

	if n == 2 {
		return max(nums[0], nums[1])
	}

	m := make([]int, n)
	m[0], m[1] = nums[0], max(nums[0], nums[1])

	for i := 2; i < n; i++ {
		m[i] = max(m[i-1], nums[i]+m[i-2])
	}

	return m[n-1]
}

/*

// Bad implementation below, just keep them and learning what mistake I did

func houseRobberIIv4(nums []int) int {
	rob := func(nums []int) int {
		if len(nums) == 0 {
			return 0
		}

		for i := 2; i < len(nums); i++ {
			nums[i] += nums[i-2]
		}
		return max(nums[len(nums)-1], nums[len(nums)-2])
	}

	return max(rob(nums[:len(nums)-1]), rob(nums[1:]))
}

func houseRobberIIv3(nums []int) int {
	rob := func(nums []int) int {
		nums = append(nums, 0)
		for i := len(nums) - 1; i >= 2; i-- {
			nums[i-2] += nums[i]
		}
		return max(nums[0], nums[1])
	}

	return max(rob(nums[1:]), rob(nums[:len(nums)-1]))
}

func houseRobberIIv2(nums []int) int {
	helper := func(nums []int) int {
		for i := len(nums) - 1; i >= 2; i-- {
			nums[i-2] += nums[i]
		}
		return max(nums[0], nums[1])
	}

	return max(helper(nums[1:]), helper(nums[:len(nums)-1]))
}

// RobII my method, but this is not run on leetcode 2022/9/4, this should be tested on leetcode
// This is wrong implementation, just keep it
func RobII(nums []int) int {

	n := len(nums)
	m := make([]int, n)
	m[0] = nums[0]
	m[1] = max(nums[0], nums[1])

	for i := 2; i < n; i++ {

		if i%(n-1) == 0 {
			m[i] = max(m[i-1], +nums[i])
		} else {
			m[i] = max(m[i-1], nums[i]+m[i-2])
		}

	}
	return m[n-1]
}
*/
