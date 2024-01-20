package b75

/*
198. House Robber

You are a professional robber planning to rob houses along a street.

Each house has a certain amount of money stashed,
 the only constraint stopping you from robbing each of
 them is that adjacent houses have security systems connected and
 it will automatically contact the police if
 two adjacent houses were broken into on the same night.

Given an integer array nums representing the amount of money of each house,
return the maximum amount of money you can rob tonight without alerting the police.



Example 1:

Input: nums = [1,2,3,1]
Output: 4
Explanation: Rob house 1 (money = 1) and then rob house 3 (money = 3).
Total amount you can rob = 1 + 3 = 4.
Example 2:

Input: nums = [2,7,9,3,1]
Output: 12
Explanation: Rob house 1 (money = 2), rob house 3 (money = 9) and rob house 5 (money = 1).
Total amount you can rob = 2 + 9 + 1 = 12.


Constraints:

1 <= nums.length <= 100
0 <= nums[i] <= 400
*/

func houseRobber40(nums []int) int {
	r1, r2 := 0, 0

	for _, num := range nums {
		tmp := max(r1+num, r2)
		r1 = r2
		r2 = tmp
	}

	return r2
}

func houseRobber39(nums []int) int {
	r1, r2 := 0, 0

	for _, num := range nums {
		tmp := max(r1+num, r2)
		r1 = r2
		r2 = tmp
	}

	return r2
}

func houseRobber38(nums []int) int {
	r1, r2 := 0, 0

	for _, num := range nums {
		tmp := max(r1+num, r2)
		r1 = r2
		r2 = tmp
	}

	return r2
}

func houseRobber37(nums []int) int {
	r1, r2 := 0, 0

	for _, num := range nums {
		tmp := max(r1+num, r2)
		r1 = r2
		r2 = tmp
	}

	return r2
}

func houseRobber36(nums []int) int {
	r1, r2 := 0, 0

	for _, num := range nums {
		tmp := max(r1+num, r2)
		r1 = r2
		r2 = tmp
	}

	return r2
}

func houseRobber35(nums []int) int {
	r1, r2 := 0, 0

	for _, num := range nums {
		tmp := max(r1+num, r2)
		r1 = r2
		r2 = tmp
	}

	return r2
}

func houseRobber34(nums []int) int {
	r1, r2 := 0, 0

	for _, num := range nums {
		tmp := max(r2, r1+num)
		r1 = r2
		r2 = tmp
	}

	return r2
}

func houseRobber33(nums []int) int {
	r1, r2 := 0, 0

	for _, num := range nums {
		tmp := max(r1+num, r2)
		r1 = r2
		r2 = tmp
	}

	return r2
}

func houseRobber32(nums []int) int {
	r1, r2 := 0, 0

	for _, num := range nums {
		tmp := max(r1+num, r2)
		r1 = r2
		r2 = tmp
	}

	return r2
}

func houseRobber31(nums []int) int {
	r1, r2 := 0, 0

	for i := 0; i < len(nums); i++ {
		tmp := max(r1+nums[i], r2)
		r1 = r2
		r2 = tmp
	}

	return r2
}

func houseRobber30(nums []int) int {
	r1, r2 := 0, 0

	for _, num := range nums {
		tmp := max(r1+num, r2)
		r1 = r2
		r2 = tmp
	}

	return r2
}

func houseRobber29(nums []int) int {
	r1, r2 := 0, 0

	for _, num := range nums {
		tmp := max(r1+num, r2)
		r1 = r2
		r2 = tmp
	}

	return r2
}

func houseRobber28(nums []int) int {
	r1, r2 := 0, 0

	for _, num := range nums {
		tmp := max(r1+num, r2)
		r1 = r2
		r2 = tmp
	}

	return r2
}

func houseRobber27(nums []int) int {
	r1, r2 := 0, 0

	for _, num := range nums {
		tmp := max(r1+num, r2)
		r1 = r2
		r2 = tmp
	}

	return r2
}

func houseRobber26(nums []int) int {
	r1, r2 := 0, 0

	for _, num := range nums {
		tmp := max(r1+num, r2)
		r1 = r2
		r2 = tmp
	}

	return r2
}

func houseRobber25(nums []int) int {
	r1, r2 := 0, 0

	for _, num := range nums {
		tmp := max(r1+num, r2)
		r1 = r2
		r2 = tmp
	}

	return r2
}

func houseRobber24(nums []int) int {
	r1, r2 := 0, 0

	for _, num := range nums {
		tmp := max(r2, r1+num)
		r1 = r2
		r2 = tmp
	}

	return r2
}

func houseRobber23(nums []int) int {
	r1, r2 := 0, 0

	for _, num := range nums {
		tmp := max(r1+num, r2)
		r1 = r2
		r2 = tmp
	}

	return r2
}

func houseRobber22(nums []int) int {
	r1, r2 := 0, 0

	for _, num := range nums {
		tmp := max(r1+num, r2)
		r1 = r2
		r2 = tmp
	}

	return r2
}

func houseRobber21(nums []int) int {
	r1, r2 := 0, 0

	for _, num := range nums {
		tmp := max(r1+num, r2)
		r1 = r2
		r2 = tmp
	}

	return r2
}

func houseRobber20(nums []int) int {
	r1, r2 := 0, 0

	for _, num := range nums {
		tmp := max(r1+num, r2)
		r1 = r2
		r2 = tmp
	}

	return r2
}

func houseRobber19(nums []int) int {
	r1, r2 := 0, 0

	for _, num := range nums {
		tmp := max(r1+num, r2)
		r1 = r2
		r2 = tmp
	}

	return r2
}

func houseRobber18(nums []int) int {
	r1, r2 := 0, 0

	for _, num := range nums {
		tmp := max(r1+num, r2)
		r1 = r2
		r2 = tmp
	}

	return r2
}

func houseRobber17(nums []int) int {
	r1, r2 := 0, 0

	for _, num := range nums {
		tmp := max(r1+num, r2)
		r1 = r2
		r2 = tmp
	}

	return r2
}

func houseRobber16(nums []int) int {
	r1, r2 := 0, 0

	for _, num := range nums {
		tmp := max(r1+num, r2)
		r1 = r2
		r2 = tmp
	}
	return r2
}

func houseRobber15(nums []int) int {
	r1, r2 := 0, 0

	for _, num := range nums {
		tmp := max(r1+num, r2)
		r1 = r2
		r2 = tmp
	}

	return r2
}

func houseRobberv14(nums []int) int {
	r1, r2 := 0, 0

	for _, num := range nums {
		tmp := max(r1+num, r2)
		r1 = r2
		r2 = tmp
	}

	return r2
}

func houseRobberv13(nums []int) int {
	r1, r2 := 0, 0

	for _, num := range nums {
		tmp := max(r1+num, r2)
		r1 = r2
		r2 = tmp
	}

	return r2
}

func houseRobberv12(nums []int) int {
	r1, r2 := 0, 0

	for _, num := range nums {
		tmp := max(r1+num, r2)
		r1 = r2
		r2 = tmp
	}

	return r2
}

func houseRobberv11(nums []int) int {
	r1, r2 := 0, 0

	for _, num := range nums {
		tmp := max(r1+num, r2)
		r1 = r2
		r2 = tmp
	}

	return r2
}

func houseRobberv10(nums []int) int {
	r1, r2 := 0, 0

	for _, num := range nums {
		tmp := max(r1+num, r2)
		r1 = r2
		r2 = tmp
	}

	return r2
}

func houseRobberv9(nums []int) int {
	r1, r2 := 0, 0

	for _, num := range nums {
		tmp := max(r1+num, r2)
		r1 = r2
		r2 = tmp
	}

	return r2
}

func houseRobberv8(nums []int) int {
	r1, r2 := 0, 0

	for _, num := range nums {
		tmp := max(r1+num, r2)
		r1 = r2
		r2 = tmp
	}

	return r2
}

func houseRobberv7(nums []int) int {
	// 0, 0, nums[0], nums[1],...

	r1, r2 := 0, 0

	for _, num := range nums {
		tmp := max(r1+num, r2)
		r1 = r2
		r2 = tmp
	}

	return r2
}

func houseRobberv6(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	if n == 1 {
		return nums[0]
	}

	if n == 2 {
		return max(nums[0], nums[1])
	}

	m := make([]int, n)
	m[0] = nums[0]
	m[1] = max(nums[0], nums[1])
	for i := 2; i < n; i++ {
		m[i] = max(m[i-1], m[i-2]+nums[i])
	}

	return m[n-1]
}

func houseRobberv5(nums []int) int {

	r1, r2 := 0, 0

	for _, num := range nums {

		tmp := max(r1+num, r2)
		r1 = r2
		r2 = tmp

		//r1, r2 = r2, max(r1+num, r2)
	}

	return r2

}

func Rob(nums []int) int {
	n := len(nums)
	m := make([]int, n+1)

	for i := 1; i <= n; i++ {
		m[i] = nums[i-1]
		for j := i + 1; j < n; j += 2 {
			m[i] += nums[j]
		}
		m[i] = max(m[i-1], m[i])
	}
	return m[n]
}

func RobBetterDP(nums []int) int {

	n := len(nums)
	m := make([]int, n)

	m[0] = nums[0]
	m[1] = max(m[0], nums[1])

	for i := 2; i < n; i++ {
		m[i] = max(m[i-2]+nums[i], m[i-1])
	}

	return m[n-1]

}

func RobBetterDPv2(nums []int) int {

	r1, r2 := 0, 0

	// [ r1, r2, nums[0], nums[1],...nums[n] ]
	for _, num := range nums {
		tmp := max(r1+num, r2)
		r1 = r2
		r2 = tmp
	}

	return r2

}

/*
// Bad implementation here, please read them and keep learning what mistake I did.
func houseRobber(houses []int) int {

	if houses == nil {
		return 0
	}

	if len(houses) == 1 {
		return houses[0]
	}

	for i := len(houses) - 1; i > 1; i-- {
		houses[i-2] += houses[i]
	}

	return max(houses[0], houses[1])
}


func houseRobberv4(nums []int) int {
	for i := 2; i < len(nums); i++ {
		nums[i] += nums[i-2]
	}

	return max(nums[len(nums)-1], nums[len(nums)-2])
}

func houseRobberv3(nums []int) int {
	nums = append(nums, 0)

	for i := len(nums) - 1; i >= 2; i-- {
		nums[i-2] += nums[i]
	}

	return max(nums[0], nums[1])

}

func houseRobberv2(nums []int) int {

	for i := len(nums) - 1; i >= 2; i-- {
		nums[i-2] += nums[i]
	}

	return max(nums[0], nums[1])

}

func houseRobberv1(nums []int) int {
	for i := len(nums) - 1; i >= 2; i-- {
		nums[i-2] += nums[i]
	}
	return max(nums[0], nums[1])
}

*/
