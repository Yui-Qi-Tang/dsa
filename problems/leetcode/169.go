package b75

/*
169. Majority Element
Easy
14.7K
446
Companies
Given an array nums of size n, return the majority element.

The majority element is the element that appears more than ⌊n / 2⌋ times. You may assume that the majority element always exists in the array.



Example 1:

Input: nums = [3,2,3]
Output: 3
Example 2:

Input: nums = [2,2,1,1,1,2,2]
Output: 2


Constraints:

n == nums.length
1 <= n <= 5 * 104
-109 <= nums[i] <= 109

normal way:
1. create a hash map to store value:freq through the loop
2. find the value with max freq

Follow-up: Could you solve the problem in linear time and in O(1) space?
^^^HINT: Boyer–Moore majority vote algorithm

Accepted
1.7M
Submissions
2.7M
Acceptance Rate
63.9%

*/

func majorityElementv7(nums []int) int {
	candidate := 0
	cnt := 0

	for _, num := range nums {
		if cnt == 0 {
			candidate = num
		}

		if candidate == num {
			cnt++
		} else {
			cnt--
		}
	}

	return candidate
}

func majorityElementv6(nums []int) int {
	candidate := 0
	cnt := 0

	for _, num := range nums {
		if cnt == 0 {
			candidate = num
		}

		if candidate == num {
			cnt++
		} else {
			cnt--
		}
	}

	return candidate
}

func majorityElementv5(nums []int) int {
	candidate := 0
	cnt := 0

	for _, num := range nums {
		if cnt == 0 {
			candidate = num
		}

		if candidate == num {
			cnt++
		} else {
			cnt--
		}
	}

	return candidate
}

func majorityElementv4(nums []int) int {
	candidate := 0
	cnt := 0

	for _, num := range nums {
		if cnt == 0 {
			candidate = num
		}

		if candidate == num {
			cnt++
		} else {
			cnt--
		}
	}

	return candidate
}

func majorityElementv3(nums []int) int {
	candidate := 0
	cnt := 0

	for _, num := range nums {
		if cnt == 0 {
			candidate = num
		}

		if num == candidate {
			cnt++
		} else {
			cnt--
		}
	}

	return candidate
}

func majorityElementv2(nums []int) int {
	candidate := 0
	cnt := 0

	for _, num := range nums {
		if cnt == 0 {
			candidate = num
		}

		if candidate == num {
			cnt++
		} else {
			cnt--
		}
	}

	return candidate
}

func majorityElementv1(nums []int) int {
	candidate := 0
	cnt := 0

	for _, num := range nums {
		if cnt == 0 {
			candidate = num
		}

		if candidate == num {
			cnt++
		} else {
			cnt--
		}
	}

	return candidate
}

func majorityElement(nums []int) int {
	candidate := -1
	count := 0

	for _, num := range nums {
		if count == 0 {
			candidate = num
		}

		if num == candidate {
			count++
		} else {
			count--
		}
	}

	return candidate
}