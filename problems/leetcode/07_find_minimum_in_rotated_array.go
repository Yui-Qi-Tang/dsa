package b75

/*
Suppose an array of length n sorted in ascending order is rotated between 1 and n times. For example, the array nums = [0,1,2,4,5,6,7] might become:

[4,5,6,7,0,1,2] if it was rotated 4 times.
[0,1,2,4,5,6,7] if it was rotated 7 times.
Notice that rotating an array [a[0], a[1], a[2], ..., a[n-1]] 1 time results in the array [a[n-1], a[0], a[1], a[2], ..., a[n-2]].

Given the sorted rotated array nums of unique elements, return the minimum element of this array.

You must write an algorithm that runs in O(log n) time.

Example 1:

Input: nums = [3,4,5,1,2]
Output: 1
Explanation: The original array was [1,2,3,4,5] rotated 3 times.
Example 2:

Input: nums = [4,5,6,7,0,1,2]
Output: 0
Explanation: The original array was [0,1,2,4,5,6,7] and it was rotated 4 times.
Example 3:

Input: nums = [11,13,15,17]
Output: 11
Explanation: The original array was [11,13,15,17] and it was rotated 4 times.

Constraints:

n == nums.length
1 <= n <= 5000
-5000 <= nums[i] <= 5000
All the integers of nums are unique.
nums is sorted and rotated between 1 and n times.
*/

func findMinRotatedSortedArrayv27(nums []int) int {
	l, r := 0, len(nums)-1

	for l < r {
		m := (l + r) / 2

		if nums[m] > nums[r] {
			l = m + 1
		} else {
			r = m
		}
	}

	return nums[l]
}

func findMinRotatedSortedArrayv26(nums []int) int {
	l, r := 0, len(nums)-1

	for l < r {
		m := (l + r) / 2
		if nums[m] > nums[r] {
			l = m + 1
		} else {
			r = m
		}
	}

	return nums[l]
}

func findMinRotatedSortedArrayv25(nums []int) int {
	l, r := 0, len(nums)-1

	for l < r {
		m := (l + r) / 2

		if nums[m] > nums[r] {
			l = m + 1
		} else {
			r = m
		}

	}

	return nums[l]
}

func findMinRotatedSortedArrayv24(nums []int) int {

	l, r := 0, len(nums)-1
	for l < r {
		m := (l + r) / 2

		if nums[m] > nums[r] {
			l = m + 1
		} else {
			r = m
		}
	}

	return nums[l]
}

func findMinRotatedSortedArrayv23(nums []int) int {
	l, r := 0, len(nums)-1

	for l < r {
		m := (l + r) / 2
		if nums[m] > nums[r] {
			l = m + 1
		} else {
			r = m
		}
	}

	return nums[l]
}

func findMinRotatedSortedArrayv22(nums []int) int {
	l, r := 0, len(nums)-1

	for l < r {
		m := (l + r) / 2

		if nums[m] > nums[r] {
			l = m + 1
		} else {
			r = m
		}
	}

	return nums[l]
}

func findMinRotatedSortedArrayv21(nums []int) int {
	l, r := 0, len(nums)-1

	for l < r {
		m := (l + r) / 2
		if nums[m] > nums[r] {
			l = m + 1
		} else {
			r = m
		}
	}

	return nums[l]
}

func findMinRotatedSortedArrayv20(nums []int) int {
	l, r := 0, len(nums)-1

	for l < r {
		m := (l + r) / 2

		if nums[m] > nums[r] {
			l = m + 1
		} else {
			r = m
		}
	}

	return nums[l]
}

func findMinRotatedSortedArrayv19(nums []int) int {
	l, r := 0, len(nums)-1

	for l < r {
		m := (l + r) / 2

		if nums[m] > nums[r] {
			l = m + 1
		} else {
			r = m
		}
	}

	return nums[l]
}

func findMinRotatedSortedArrayv18(nums []int) int {
	l, r := 0, len(nums)-1

	for l < r {
		m := (l + r) / 2

		if nums[m] > nums[r] {
			l = m + 1
		} else {
			r = m
		}
	}

	return nums[l]
}

func findMinRotatedSortedArrayv17(nums []int) int {
	l, r := 0, len(nums)-1

	for l < r {
		m := (l + r) / 2
		if nums[m] > nums[r] {
			l = m + 1
		} else {
			r = m
		}
	}

	return nums[l]
}

func findMinRotatedSortedArrayv16(nums []int) int {
	l, r := 0, len(nums)-1

	for l < r {
		m := (l + r) / 2

		if nums[m] > nums[r] {
			l = m + 1
		} else {
			r = m
		}
	}

	return nums[l]
}

func findMinRotatedSortedArrayv15(nums []int) int {
	l, r := 0, len(nums)-1

	for l < r {
		m := (r + l) / 2

		if nums[m] > nums[r] {
			l = m + 1
		} else {
			r = m
		}
	}

	return nums[l]
}

func findMinRotatedSortedArrayv14(nums []int) int {
	l, r := 0, len(nums)-1

	for l < r {
		m := (l + r) / 2

		if nums[m] > nums[r] {
			l = m + 1
		} else {
			r = m
		}
	}

	return nums[l]
}

func findMinRotatedSortedArrayv13(nums []int) int {
	l, r := 0, len(nums)-1

	for l < r {
		m := (r + l) / 2
		if nums[m] > nums[r] {
			l = m + 1
		} else {
			r = m
		}

	}

	return nums[l]
}

func findMinRotatedSortedArrayv12(nums []int) int {
	l, r := 0, len(nums)-1

	for l < r {
		m := (l + r) / 2
		if nums[m] > nums[r] {
			l = m + 1
		} else {
			r = m
		}
	}

	return nums[l]
}

func findMinRotatedSortedArrayv11(nums []int) int {
	l, r := 0, len(nums)-1

	for l < r {
		m := (l + r) / 2

		if nums[m] > nums[r] {
			l = m + 1
		} else {
			r = m
		}
	}

	return nums[l]
}

func findMinRotatedSortedArrayv10(nums []int) int {
	l, r := 0, len(nums)-1

	for l < r {
		m := (l + r) / 2
		if nums[m] > nums[r] {
			l = m + 1
		} else {
			r = m
		}
	}

	return nums[l]
}

func findMinRotatedSortedArrayv9(nums []int) int {
	l, r := 0, len(nums)-1

	for l < r {
		m := (l + r) / 2
		if nums[m] > nums[r] {
			l = m + 1
		} else {
			r = m
		}
	}

	return nums[l]
}

func findMinRotatedSortedArrayv8(nums []int) int {

	l, r := 0, len(nums)-1

	for l < r {
		m := (l + r) / 2
		if nums[m] > nums[r] {
			l = m + 1
		} else {
			r = m
		}
	}

	return nums[l]
}

func findMinRotatedSortedArrayv7(nums []int) int {
	l, r := 0, len(nums)-1

	for l < r {
		m := (l + r) / 2
		if nums[m] > nums[r] {
			l = m + 1
		} else {
			r = m
		}
	}

	return nums[l]
}

func findMinRotatedSortedArrayv6(nums []int) int {
	l, r := 0, len(nums)-1

	for l < r {
		m := (l + r) / 2
		if nums[m] > nums[r] {
			l = m + 1
		} else {
			r = m
		}
	}

	return nums[l]
}

func findMinRotatedSortedArrayv5(nums []int) int {
	l, r := 0, len(nums)-1

	for l < r {
		m := (l + r) / 2
		if nums[m] > nums[r] {
			l = m + 1
		} else {
			r = m
		}
	}

	return nums[l]
}

func findMinRotatedSortedArrayv4(nums []int) int {
	l, r := 0, len(nums)-1

	for l < r {
		m := (l + r) / 2

		if nums[m] > nums[r] {
			l = m + 1
		} else {
			r = m
		}
	}

	return nums[l]
}

func findMinRotatedSortedArrayv3(nums []int) int {
	l, r := 0, len(nums)-1

	for l < r {
		m := (l + r) / 2

		if nums[m] > nums[r] {
			l = m + 1
		} else {
			r = m
		}
	}

	return nums[l]
}

func findMinRotatedSortedArrayv2(nums []int) int {
	l, r := 0, len(nums)-1

	for l < r {
		m := (l + r) / 2
		if nums[m] > nums[r] {
			l = m + 1
		} else {
			r = m
		}
	}

	return nums[l]
}

func findMinRotatedSortedArrayv1(nums []int) int {
	l, r := 0, len(nums)-1

	for l < r {
		m := (l + r) / 2
		if nums[m] > nums[r] {
			l = m + 1
		} else {
			r = m
		}
	}
	return nums[l]
}
