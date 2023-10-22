package b75

/*
There is an integer array nums sorted in ascending order (with distinct values).

Prior to being passed to your function, nums is possibly rotated at an unknown pivot index k (1 <= k < nums.length) such that the resulting array is [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]] (0-indexed). For example, [0,1,2,4,5,6,7] might be rotated at pivot index 3 and become [4,5,6,7,0,1,2].

Given the array nums after the possible rotation and an integer target, return the index of target if it is in nums, or -1 if it is not in nums.

You must write an algorithm with O(log n) runtime complexity.



Example 1:

Input: nums = [4,5,6,7,0,1,2], target = 0
Output: 4
Example 2:

Input: nums = [4,5,6,7,0,1,2], target = 3
Output: -1
Example 3:

Input: nums = [1], target = 0
Output: -1


Constraints:

1 <= nums.length <= 5000
-104 <= nums[i] <= 104
All values of nums are unique.
nums is an ascending array that is possibly rotated.
-104 <= target <= 104
*/

func SearchInRotatedArrayv20(nums []int, target int) int {
	l, r := 0, len(nums)-1

	for l != r {
		m := (l + r) / 2

		if target >= nums[l] && target < nums[m] {
			r = m - 1
		} else {
			l = m + 1
		}

		if nums[l] == target {
			return l
		}
	}

	return -1
}

func SearchInRotatedArrayv19(nums []int, target int) int {
	l, r := 0, len(nums)-1

	for l != r {
		m := (l + r) / 2

		if target >= nums[l] && target < nums[m] {
			r = m - 1
		} else {
			l = m + 1
		}

		if nums[l] == target {
			return l
		}
	}

	return -1
}

func SearchInRotatedArrayv18(nums []int, target int) int {
	l, r := 0, len(nums)-1

	for l != r {
		m := (l + r) / 2

		if target >= nums[l] && target < nums[m] {
			r = m - 1
		} else {
			l = m + 1
		}

		if nums[l] == target {
			return l
		}
	}

	return -1
}

func SearchInRotatedArrayv17(nums []int, target int) int {
	l, r := 0, len(nums)-1

	for l != r {

		m := (l + r) / 2
		if target >= nums[l] && target < nums[m] {
			r = m - 1
		} else {
			l = m + 1
		}

		if nums[l] == target {
			return l
		}

	}

	return -1
}

func SearchInRotatedArrayv16(nums []int, target int) int {
	l, r := 0, len(nums)-1

	for l != r {
		m := (l + r) / 2

		if target >= nums[l] && target < nums[m] {
			r = m - 1
		} else {
			l = m + 1
		}

		if nums[l] == target {
			return l
		}
	}

	return -1
}

func SearchInRotatedArrayv15(nums []int, target int) int {
	l, r := 0, len(nums)-1

	for l != r {
		m := (l + r) / 2

		if target >= nums[l] && target < nums[m] {
			r = m - 1
		} else {
			l = m + 1
		}

		if nums[l] == target {
			return l
		}
	}

	return -1
}

func SearchInRotatedArrayv14(nums []int, target int) int {
	l, r := 0, len(nums)-1

	for l != r {
		m := (l + r) / 2

		if target >= nums[l] && target < nums[m] {
			r = m - 1
		} else {
			l = m + 1
		}

		if nums[l] == target {
			return l
		}
	}

	return -1
}

func SearchInRotatedArrayv13(nums []int, target int) int {
	l, r := 0, len(nums)-1

	for l != r {
		m := (l + r) / 2

		if target >= nums[l] && target < nums[m] {
			r = m - 1
		} else {
			l = m + 1
		}

		if nums[l] == target {
			return l
		}
	}

	return -1
}

func SearchInRotatedArrayv12(nums []int, target int) int {
	l, r := 0, len(nums)-1

	for l != r {
		m := (l + r) / 2

		if target >= nums[l] && target < nums[m] {
			r = m - 1
		} else {
			l = m + 1
		}

		if nums[l] == target {
			return l
		}
	}

	return -1
}

func SearchInRotatedArrayv11(nums []int, target int) int {
	l, r := 0, len(nums)-1

	for l != r {
		m := (l + r) / 2

		if target >= nums[l] && target < nums[m] {
			r = m - 1
		} else {
			l = m + 1
		}

		if nums[l] == target {
			return l
		}

	}

	return -1
}

func SearchInRotatedArrayv10(nums []int, target int) int {
	l, r := 0, len(nums)-1

	for l < r {
		m := (l + r) / 2
		if target >= nums[l] && target < nums[m] {
			r = m - 1
		} else {
			l = m + 1
		}

		if nums[l] == target {
			return l
		}
	}

	return -1
}

func SearchInRotatedArrayv9(nums []int, target int) int {
	l, r := 0, len(nums)-1

	for l != r {
		m := (l + r) / 2

		if target >= nums[l] && target < nums[m] {
			r = m - 1
		} else {
			l = m + 1
		}

		if nums[l] == target {
			return l
		}
	}

	return -1
}

func SearchInRotatedArrayv8(nums []int, target int) int {
	l, r := 0, len(nums)-1

	for l != r {
		m := (l + r) / 2

		if target >= nums[m] && target < nums[r] {
			r = m - 1
		} else {
			l = m + 1
		}

		if nums[l] == target {
			return l
		}

	}

	return -1
}

func SearchInRotatedArrayv7(nums []int, target int) int {
	l, r := 0, len(nums)-1

	for l != r {
		m := (l + r) / 2

		if target >= nums[m] && target < nums[r] {
			r = m - 1
		} else {
			l = m + 1
		}

		if nums[l] == target {
			return l
		}
	}

	return -1
}

func SearchInRotatedArrayv6(nums []int, target int) int {
	l, r := 0, len(nums)-1

	for l != r {
		m := (l + r) / 2

		if target >= nums[m] && target < nums[r] {
			r = m - 1
		} else {
			l = m + 1
		}

		if nums[l] == target {
			return l
		}
	}

	return -1
}

func SearchInRotatedArrayv5(nums []int, target int) int {
	l, r := 0, len(nums)-1

	for l != r {
		m := (l + r) / 2

		if target >= nums[m] && target < nums[r] {
			r = m - 1
		} else {
			l = m + 1
		}

		if nums[l] == target {
			return l
		}
	}

	return -1
}

func SearchInRotatedArrayv4(nums []int, target int) int {
	l, r := 0, len(nums)-1

	for l != r {
		m := (l + r) / 2

		if nums[m] <= target && target < nums[r] {
			r = m - 1
		} else {
			l = m + 1
		}

		if nums[l] == target {
			return l
		}

	}

	return -1
}

func SearchInRotatedArrayv3(nums []int, target int) int {
	l, r := 0, len(nums)-1

	for l != r {
		m := (l + r) / 2

		if nums[m] <= target && nums[r] > target {
			r = m - 1
		} else {
			l = m + 1
		}
		if nums[l] == target {
			return l
		}
	}

	return -1
}

func SearchInRotatedArrayv2(nums []int, target int) int {
	l, r := 0, len(nums)-1

	for l != r {
		m := (l + r) / 2
		if target > nums[m] {
			if nums[m] > nums[r] {
				r = m - 1
			} else {
				l = m + 1
			}
		} else {
			if nums[m] > nums[r] {
				l = m + 1
			} else {
				r = m - 1
			}
		}
		if nums[l] == target {
			return l
		}
	}

	return -1
}

// return index of target in the nums, -1 if target does not in nums
func SearchInRotatedArrayv1(nums []int, target int) int {

	if len(nums) == 0 {
		return -1
	}

	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		}
		return -1
	}

	n := len(nums)

	l := 0
	r := n - 1

	for l != r {

		m := (l + r) / 2

		// v in l, .... , m
		if nums[l] <= target && nums[m] > target {
			r = m - 1
		} else { // v in m+1, ... r
			l = m + 1
		}

		if nums[l] == target {
			return l
		}

	}

	return -1
}
