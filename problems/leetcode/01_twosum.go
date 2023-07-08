package b75

/*
1. Two Sum

Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.


Example 1:

Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
Example 2:

Input: nums = [3,2,4], target = 6
Output: [1,2]
Example 3:

Input: nums = [3,3], target = 6
Output: [0,1]


Constraints:

2 <= nums.length <= 104
-109 <= nums[i] <= 109
-109 <= target <= 109
Only one valid answer exists.
*/

func TwoSumv17(nums []int, target int) []int {
	diff := make(map[int]int, len(nums))

	for i, num := range nums {
		if v, exist := diff[num]; exist {
			return []int{v, i}
		}
		diff[target-num] = i
	}

	return []int{}
}

func TwoSumv16(nums []int, target int) []int {
	diff := make(map[int]int, len(nums))

	for i, num := range nums {
		if v, ok := diff[num]; ok {
			return []int{v, i}
		}
		diff[target-num] = i
	}

	return []int{}
}

func TwoSumv15(nums []int, target int) []int {
	diff := make(map[int]int, len(nums))

	for i, num := range nums {
		if d, ok := diff[num]; ok {
			return []int{d, i}
		}

		diff[target-num] = i
	}

	return []int{}
}

func TwoSumv14(nums []int, target int) []int {
	diff := make(map[int]int, len(nums))
	for i, num := range nums {
		if v, ok := diff[num]; ok {
			return []int{v, i}
		}
		diff[target-num] = i
	}
	return []int{}
}

func TwoSumv13(nums []int, target int) []int {
	diff := make(map[int]int)

	for i, num := range nums {
		if v, ok := diff[num]; ok {
			return []int{v, i}
		}
		diff[target-num] = i
	}

	return []int{}
}

func TwoSumv12(nums []int, target int) []int {
	exist := make(map[int]int, len(nums)) // diff: idx

	for i, num := range nums {
		if ii, ok := exist[num]; ok {
			return []int{ii, i}
		}

		exist[target-num] = i
	}

	return nil
}

func TwoSumv11(nums []int, target int) []int {
	m := make(map[int]int)

	for i, num := range nums {
		if v, exist := m[num]; exist {
			return []int{v, i}
		}
		m[target-num] = i
	}

	return nil
}

func TwoSumv10(nums []int, target int) []int {
	m := make(map[int]int)

	for i, num := range nums {
		if v, exist := m[num]; exist {
			return []int{v, i}
		}

		m[target-num] = i
	}

	return nil
}

func TwoSumv9(nums []int, target int) []int {
	m := make(map[int]int)

	for i, num := range nums {
		if v, exist := m[num]; exist {
			return []int{v, i}
		}

		m[target-num] = i
	}

	return nil
}

func TwoSumv8(nums []int, target int) []int {
	m := make(map[int]int)

	for i, num := range nums {
		if v, exist := m[num]; exist {
			return []int{v, i}
		}
		m[target-num] = i
	}
	return nil
}

func TwoSumv7(nums []int, target int) []int {
	m := make(map[int]int)

	for i, num := range nums {
		if v, exist := m[num]; exist {
			return []int{v, i}
		}
		m[target-num] = i
	}

	return nil
}

func TwoSumv6(nums []int, target int) []int {
	m := make(map[int]int)

	for i, num := range nums {
		if v, exist := m[num]; exist {
			return []int{v, i}
		}
		m[target-num] = i
	}

	return nil
}

func TwoSumv5(nums []int, target int) []int {
	m := make(map[int]int)

	for i, num := range nums {
		if v, exist := m[num]; exist {
			return []int{v, i}
		}
		m[target-num] = i
	}

	return nil
}

func TwoSumv4(nums []int, target int) []int {
	m := make(map[int]int)

	for i, num := range nums {
		if idx, exist := m[num]; exist {
			return []int{idx, i}
		}

		m[target-num] = i
	}

	return nil
}

func TwoSumv3(nums []int, target int) []int {
	m := make(map[int]int)

	for i, num := range nums {
		if v, exist := m[num]; exist {
			return []int{v, i}
		}
		m[target-num] = i
	}

	return nil
}

func TwoSumv2(nums []int, target int) []int {
	m := make(map[int]int) // diff, index

	for i, num := range nums {

		if v, exist := m[num]; exist {
			return []int{v, i}
		}
		m[target-num] = i

	}

	return nil
}

func TwoSumv1(nums []int, target int) []int {
	m := make(map[int]int)

	for i, num := range nums {

		if idx, exist := m[target-num]; exist {
			return []int{idx, i}
		} else {
			m[num] = i
		}

	}
	return nil
}

func TwoSum(nums []int, target int) []int {
	m := make(map[int]int) // v:index
	for i, num := range nums {
		if idx, ok := m[target-num]; ok {
			return []int{idx, i}
		}
		m[num] = i
	}
	return nil
}
