package b75

/*
Given an array nums containing n distinct numbers in the range [0, n], return the only number in the range that is missing from the array.



Example 1:

Input: nums = [3,0,1]
Output: 2
Explanation: n = 3 since there are 3 numbers, so all numbers are in the range [0,3]. 2 is the missing number in the range since it does not appear in nums.
Example 2:

Input: nums = [0,1]
Output: 2
Explanation: n = 2 since there are 2 numbers, so all numbers are in the range [0,2]. 2 is the missing number in the range since it does not appear in nums.
Example 3:

Input: nums = [9,6,4,2,3,5,7,0,1]
Output: 8
Explanation: n = 9 since there are 9 numbers, so all numbers are in the range [0,9]. 8 is the missing number in the range since it does not appear in nums.


Constraints:

n == nums.length
1 <= n <= 104
0 <= nums[i] <= n
All the numbers of nums are unique.


Follow up: Could you implement a solution using only O(1) extra space complexity and O(n) runtime complexity?
*/

func missingNumberv17(nums []int) int {
	total := ((1 + len(nums)) * len(nums)) / 2

	for _, num := range nums {
		total -= num
	}

	return total
}

func missingNumberv16(nums []int) int {
	total := ((1 + len(nums)) * len(nums)) / 2

	for _, num := range nums {
		total -= num
	}

	return total
}

func missingNumberv15(nums []int) int {
	total := ((1 + len(nums)) * len(nums)) / 2

	for _, num := range nums {
		total -= num
	}

	return total
}

func missingNumberv14(nums []int) int {
	total := ((1 + len(nums)) * len(nums)) / 2

	for _, num := range nums {
		total -= num
	}

	return total
}

func missingNumberv13(nums []int) int {
	total := ((1 + len(nums)) * len(nums)) / 2

	for _, num := range nums {
		total -= num
	}
	return total
}

func missingNumberv12(nums []int) int {
	total := ((1 + len(nums)) * len(nums)) / 2

	for _, num := range nums {
		total -= num
	}

	return total
}

func missingNumberv11(nums []int) int {
	total := ((1 + len(nums)) * len(nums)) / 2

	for _, num := range nums {
		total -= num
	}

	return total
}

func missingNumberv10(nums []int) int {
	total := ((1 + len(nums)) * len(nums)) / 2

	for _, num := range nums {
		total -= num
	}

	return total
}

func missingNumberv9(nums []int) int {
	total := ((1 + len(nums)) * len(nums)) / 2

	for _, num := range nums {
		total -= num
	}

	return total
}

func missingNumberv8(nums []int) int {
	total := ((1 + len(nums)) * len(nums)) / 2
	for _, num := range nums {
		total -= num
	}
	return total
}

func missingNumberv7(nums []int) int {
	total := ((1 + len(nums)) * len(nums)) / 2

	for _, num := range nums {
		total -= num
	}
	return total
}

func missingNumberv6(nums []int) int {
	n := len(nums)

	total := ((1 + n) * n) / 2

	for _, n := range nums {
		total -= n
	}

	return total
}

func missingNumberv5(nums []int) int {
	n := len(nums)

	total := ((1 + n) * n) / 2

	for _, num := range nums {
		total -= num
	}

	return total
}

func missingNumberv4(nums []int) int {
	n := len(nums)
	total := ((1 + n) * n) / 2

	for _, n := range nums {
		total -= n
	}

	return total
}

func missingNumberv3(nums []int) int {
	h := len(nums)
	total := ((1 + h) * h) / 2

	for _, n := range nums {
		total -= n
	}

	return total
}

func missingNumberv2(nums []int) int {
	n := len(nums)

	total := ((1 + n) * n) / 2

	for _, num := range nums {
		total -= num
	}
	return total
}
func missingNumberv1(nums []int) int {

	var max int = ((1 + len(nums)) * len(nums)) / 2

	for _, v := range nums {
		max -= v
	}

	return max
}
