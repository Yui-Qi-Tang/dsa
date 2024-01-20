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

func ContainDuplicatev39(nums []int) bool {
	dup := make(map[int]bool, len(nums))
	for _, num := range nums {
		if dup[num] {
			return true
		}
		dup[num] = true
	}
	return false
}

func ContainDuplicatev38(nums []int) bool {
	dup := make(map[int]bool)
	for _, num := range nums {
		if dup[num] {
			return true
		}
		dup[num] = true
	}
	return false
}

func ContainDuplicatev37(nums []int) bool {
	dup := make(map[int]bool)

	for _, num := range nums {
		if dup[num] {
			return true
		}
		dup[num] = true
	}

	return false
}

func ContainDuplicatev36(nums []int) bool {
	dup := make(map[int]bool, len(nums))
	for _, num := range nums {
		if dup[num] {
			return true
		}
		dup[num] = true
	}

	return false
}

func ContainDuplicatev35(nums []int) bool {
	dup := make(map[int]bool, len(nums))

	for _, num := range nums {
		if dup[num] {
			return true
		}
		dup[num] = true
	}

	return false
}

func ContainDuplicatev34(nums []int) bool {
	dup := make(map[int]bool, len(nums))

	for _, num := range nums {
		if dup[num] {
			return true
		}

		dup[num] = true
	}

	return false
}

func ContainDuplicatev33(nums []int) bool {
	dup := make(map[int]bool)

	for _, num := range nums {
		if dup[num] {
			return true
		}

		dup[num] = true
	}

	return false
}

func ContainDuplicatev32(nums []int) bool {

	dup := make(map[int]bool)

	for _, num := range nums {
		if dup[num] {
			return true
		}

		dup[num] = true
	}

	return false
}

func ContainDuplicatev31(nums []int) bool {
	exist := make(map[int]bool, len(nums))

	for _, num := range nums {
		if exist[num] {
			return true
		}
		exist[num] = true
	}

	return false
}

func ContainDuplicatev30(nums []int) bool {
	dup := make(map[int]bool, len(nums))

	for _, num := range nums {
		if dup[num] {
			return true
		}

		dup[num] = true
	}

	return false
}

func ContainDuplicatev29(nums []int) bool {
	dup := make(map[int]bool, len(nums))

	for _, num := range nums {
		if dup[num] {
			return true
		}

		dup[num] = true
	}

	return false
}

func ContainDuplicatev28(nums []int) bool {
	dup := make(map[int]bool, len(nums))

	for _, num := range nums {
		if dup[num] {
			return true
		}

		dup[num] = true
	}

	return false
}

func ContainDuplicatev27(nums []int) bool {
	dup := make(map[int]bool, len(nums))

	for _, num := range nums {
		if dup[num] {
			return true
		}
		dup[num] = true
	}

	return false
}

func ContainDuplicatev26(nums []int) bool {
	dup := make(map[int]bool, len(nums))

	for _, num := range nums {
		if dup[num] {
			return true
		}

		dup[num] = true
	}

	return false
}

func ContainDuplicatev25(nums []int) bool {
	dup := make(map[int]bool, len(nums))

	for _, num := range nums {
		if dup[num] {
			return true
		}

		dup[num] = true
	}

	return false
}

func ContainDuplicatev24(nums []int) bool {
	dup := make(map[int]bool, len(nums))

	for _, num := range nums {
		if dup[num] {
			return true
		}
		dup[num] = true
	}

	return false
}

func ContainDuplicatev23(nums []int) bool {
	dup := make(map[int]bool, len(nums))

	for _, num := range nums {
		if dup[num] {
			return true
		}

		dup[num] = true
	}

	return false
}

func ContainDuplicatev22(nums []int) bool {
	dup := make(map[int]bool, len(nums))

	for _, num := range nums {
		if dup[num] {
			return true
		}

		dup[num] = true
	}

	return false
}

func ContainDuplicatev21(nums []int) bool {
	dup := make(map[int]bool, len(nums))

	for _, num := range nums {
		if dup[num] {
			return true
		}

		dup[num] = true
	}

	return false
}

func ContainDuplicatev20(nums []int) bool {
	dup := make(map[int]bool, len(nums))

	for _, num := range nums {
		if dup[num] {
			return true
		}

		dup[num] = true
	}

	return false
}

func ContainDuplicatev19(nums []int) bool {
	dup := make(map[int]bool, len(nums))

	for _, num := range nums {
		if dup[num] {
			return true
		}

		dup[num] = true
	}

	return false
}

func ContainDuplicatev18(nums []int) bool {
	dup := make(map[int]bool, len(nums))

	for _, num := range nums {
		if dup[num] {
			return true
		}

		dup[num] = true
	}

	return false
}

func ContainDuplicatev17(nums []int) bool {
	dup := make(map[int]bool, len(nums))

	for _, num := range nums {
		if dup[num] {
			return true
		}

		dup[num] = true
	}

	return false
}

func ContainDuplicatev16(nums []int) bool {
	dup := make(map[int]bool, len(nums))

	for _, num := range nums {
		if dup[num] {
			return true
		}

		dup[num] = true
	}

	return false
}

func ContainDuplicatev15(nums []int) bool {
	dup := make(map[int]bool, len(nums))

	for _, num := range nums {
		if dup[num] {
			return true
		}

		dup[num] = true
	}

	return false
}

func ContainDuplicatev14(nums []int) bool {
	dup := make(map[int]bool, len(nums))

	for _, num := range nums {
		if dup[num] {
			return true
		}

		dup[num] = true
	}

	return false
}

func ContainDuplicatev13(nums []int) bool {
	dup := make(map[int]bool, len(nums))

	for _, num := range nums {
		if dup[num] {
			return true
		}
		dup[num] = true
	}

	return false
}

func ContainDuplicatev12(nums []int) bool {
	exist := make(map[int]bool, len(nums))
	for _, num := range nums {
		if exist[num] {
			return true
		}
		exist[num] = true
	}
	return false
}

func ContainDuplicatev11(nums []int) bool {
	dup := make(map[int]bool, len(nums))

	for _, num := range nums {
		if dup[num] {
			return true
		}

		dup[num] = true
	}

	return false
}

func ContainDuplicatev10(nums []int) bool {
	dup := make(map[int]bool, len(nums))
	for _, num := range nums {
		if dup[num] {
			return true
		}

		dup[num] = true
	}
	return false
}

func ContainDuplicatev9(nums []int) bool {
	exist := make(map[int]bool, len(nums))

	for _, n := range nums {
		if exist[n] {
			return true
		}
		exist[n] = true
	}

	return false
}

func ContainDuplicatev8(nums []int) bool {
	exist := make(map[int]bool, len(nums))

	for _, num := range nums {
		if exist[num] {
			return true
		}
		exist[num] = true
	}

	return false
}

func ContainDuplicatev7(nums []int) bool {
	dup := make(map[int]bool, len(nums))

	for i := range nums {
		if dup[nums[i]] {
			return true
		}
		dup[nums[i]] = true
	}

	return false
}

func ContainDuplicatev6(nums []int) bool {
	exist := make(map[int]bool, len(nums))

	for _, num := range nums {
		if exist[num] {
			return true
		}
		exist[num] = true
	}

	return false
}

func ContainDuplicatev5(nums []int) bool {
	exist := make(map[int]bool, len(nums))

	for _, num := range nums {
		if exist[num] {
			return true
		}

		exist[num] = true
	}

	return false
}

func ContainDuplicatev4(nums []int) bool {
	exist := make(map[int]bool, len(nums))

	for _, num := range nums {
		if exist[num] {
			return true
		}

		exist[num] = true
	}

	return false
}

func ContainDuplicatev3(nums []int) bool {
	exist := make(map[int]bool, len(nums))

	for _, v := range nums {
		if exist[v] {
			return true
		}
		exist[v] = true
	}

	return false
}

func ContainDuplicatev2(nums []int) bool {

	exist := make(map[int]bool, len(nums))

	for _, num := range nums {
		if exist[num] {
			return true
		}
		exist[num] = true
	}

	return false
}

func ContainDuplicatev1(nums []int) bool {
	exist := make(map[int]bool, len(nums))

	for _, num := range nums {
		if exist[num] {
			return true
		}
		exist[num] = true
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
