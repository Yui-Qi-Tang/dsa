package b75

/*
Given an integer array nums, return true if any value appears at least twice in the array, and return false if every element is distinct.



Example 1:

Input: nums = [1,2,3,1]
Output: true
Example 2:

Input: nums = [1,2,3,4]
Output: false
Example 3:

Input: nums = [1,1,1,3,3,4,3,2,4,2]
Output: true

*/

func ContainDuplicatev1(nums []int) bool {
	freq := make([]int, len(nums))

	for _, n := range nums {
		freq[n-1]++
		if freq[n-1] > 1 {
			return true
		}
	}
	return false
}

func ContainDuplicate(nums []int) bool {
	m := make(map[int]bool, 0)

	for _, num := range nums {
		if _, ok := m[num]; ok {
			return true
		}
		m[num] = true
	}
	return false
}
