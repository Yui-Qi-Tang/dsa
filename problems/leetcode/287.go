package b75

// linked list cycle
// floyd's algorithm

/*
287. Find the Duplicate Number
Medium
18.7K
2.7K
Companies
Given an array of integers nums containing n + 1 integers where each integer is in the range [1, n] inclusive.

There is only one repeated number in nums, return this repeated number.

You must solve the problem without modifying the array nums and uses only constant extra space.



Example 1:

Input: nums = [1,3,4,2,2]
Output: 2
Example 2:

Input: nums = [3,1,3,4,2]
Output: 3


Constraints:

1 <= n <= 105
nums.length == n + 1
1 <= nums[i] <= n
All the integers in nums appear only once except for precisely one integer which appears two or more times.


Follow up:

How can we prove that at least one duplicate number must exist in nums?
Can you solve the problem in linear runtime complexity?
*/

func findDuplicatev28(nums []int) int {
	s, f := 0, 0

	for {
		s = nums[s]
		f = nums[nums[f]]
		if s == f {
			break
		}
	}

	s2 := 0
	for {
		s = nums[s]
		s2 = nums[s2]
		if s == s2 {
			return s2
		}
	}
}

func findDuplicatev27(nums []int) int {
	s, f := 0, 0

	for {
		s = nums[s]
		f = nums[nums[f]]
		if s == f {
			break
		}
	}

	s2 := 0
	for {
		s = nums[s]
		s2 = nums[s2]
		if s == s2 {
			return s2
		}
	}
}

func findDuplicatev26(nums []int) int {
	s, f := 0, 0

	for {
		s = nums[s]
		f = nums[nums[f]]
		if s == f {
			break
		}
	}

	s2 := 0
	for {
		s = nums[s]
		s2 = nums[s2]
		if s == s2 {
			return s2
		}
	}
}

func findDuplicatev25(nums []int) int {
	s, f := 0, 0

	for {
		s = nums[s]
		f = nums[nums[f]]
		if s == f {
			break
		}
	}

	s2 := 0

	for {
		s = nums[s]
		s2 = nums[s2]
		if s == s2 {
			return s2
		}
	}
}

func findDuplicatev24(nums []int) int {
	s, f := 0, 0

	for {
		s = nums[s]
		f = nums[nums[f]]
		if s == f {
			break
		}
	}

	s2 := 0

	for {
		s = nums[s]
		s2 = nums[s2]

		if s == s2 {
			return s2
		}
	}
}

func findDuplicatev23(nums []int) int {
	s, f := 0, 0

	for {
		s = nums[s]
		f = nums[nums[f]]
		if s == f {
			break
		}
	}

	s2 := 0
	for {
		s = nums[s]
		s2 = nums[s2]
		if s == s2 {
			return s2
		}
	}
}

func findDuplicatev22(nums []int) int {
	s, f := 0, 0

	for {
		s = nums[s]
		f = nums[nums[f]]
		if s == f {
			break
		}
	}

	s2 := 0
	for {
		s = nums[s]
		s2 = nums[s2]
		if s == s2 {
			return s2
		}
	}
}

func findDuplicatev21(nums []int) int {
	s, f := 0, 0

	for {
		s = nums[s]
		f = nums[nums[f]]
		if s == f {
			break
		}
	}

	s2 := 0
	for {
		s = nums[s]
		s2 = nums[s2]
		if s == s2 {
			return s2
		}
	}
}

func findDuplicatev20(nums []int) int {

	s, f := 0, 0

	for {
		s = nums[s]
		f = nums[nums[f]]
		if s == f {
			break
		}
	}

	s2 := 0
	for {
		s = nums[s]
		s2 = nums[s2]
		if s == s2 {
			return s2
		}
	}
}

func findDuplicatev19(nums []int) int {
	s, f := 0, 0

	for {
		s = nums[s]
		f = nums[nums[f]]
		if s == f {
			break
		}
	}

	s2 := 0

	for {
		s = nums[s]
		s2 = nums[s2]
		if s == s2 {
			break
		}
	}

	return s2
}

func findDuplicatev18(nums []int) int {
	s, f := 0, 0

	for {
		s = nums[s]
		f = nums[nums[f]]
		if s == f {
			break
		}
	}

	s2 := 0

	for {
		s = nums[s]
		s2 = nums[s2]
		if s == s2 {
			break
		}
	}

	return s2
}

func findDuplicatev17(nums []int) int {
	s, f := 0, 0

	for {
		s = nums[s]
		f = nums[nums[f]]
		if s == f {
			break
		}
	}

	s2 := 0

	for {
		s = nums[s]
		s2 = nums[s2]
		if s == s2 {
			break
		}
	}

	return s2
}

func findDuplicatev16(nums []int) int {
	s, f := 0, 0

	for {
		s = nums[s]
		f = nums[nums[f]]
		if s == f {
			break
		}
	}

	s2 := 0

	for {
		s = nums[s]
		s2 = nums[s2]
		if s == s2 {
			break
		}
	}

	return s2
}

func findDuplicatev15(nums []int) int {
	s, f := 0, 0

	for {
		s = nums[s]
		f = nums[nums[f]]
		if s == f {
			break
		}
	}

	s2 := 0

	for {
		s = nums[s]
		s2 = nums[s2]
		if s == s2 {
			break
		}
	}

	return s2
}

func findDuplicatev14(nums []int) int {
	s, f := 0, 0

	for {
		s = nums[s]
		f = nums[nums[f]]
		if s == f {
			break
		}
	}

	s2 := 0
	for {
		s = nums[s]
		s2 = nums[s2]
		if s == s2 {
			break
		}
	}

	return s2
}

func findDuplicatev13(nums []int) int {
	s, f := 0, 0

	for {
		s = nums[s]
		f = nums[nums[f]]
		if s == f {
			break
		}
	}

	s2 := 0
	for {
		s = nums[s]
		s2 = nums[s2]
		if s == s2 {
			break
		}
	}

	return s2
}

func findDuplicatev12(nums []int) int {
	s, f := 0, 0

	for {
		s = nums[s]
		f = nums[nums[f]]
		if s == f {
			break
		}
	}

	s2 := 0

	for {
		s = nums[s]
		s2 = nums[s2]
		if s == s2 {
			break
		}
	}

	return s2
}

func findDuplicatev11(nums []int) int {
	s, f := 0, 0
	for {
		s = nums[s]
		f = nums[nums[f]]
		if s == f {
			break
		}
	}

	s2 := 0
	for {
		s = nums[s]
		s2 = nums[s2]
		if s == s2 {
			break
		}
	}

	return s2
}

func findDuplicatev10(nums []int) int {
	s, f := 0, 0
	for {
		s = nums[s]
		f = nums[nums[f]]
		if s == f {
			break
		}
	}

	s2 := 0

	for {
		s = nums[s]
		s2 = nums[s2]
		if s == s2 {
			break
		}
	}

	return s2
}

func findDuplicatev9(nums []int) int {
	s, f := 0, 0

	for {
		s = nums[s]
		f = nums[nums[f]]
		if s == f {
			break
		}
	}

	s2 := 0

	for {
		s = nums[s]
		s2 = nums[s2]
		if s == s2 {
			break
		}
	}

	return s2
}

func findDuplicatev8(nums []int) int {
	s, f := 0, 0
	for {
		s = nums[s]
		f = nums[nums[f]]
		if s == f {
			break
		}
	}

	s2 := 0
	for {
		s = nums[s]
		s2 = nums[s2]
		if s == s2 {
			break
		}
	}

	return s2
}

func findDuplicatev7(nums []int) int {
	slow, fast := 0, 0

	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			break
		}
	}

	slow2 := 0

	for {
		slow = nums[slow]
		slow2 = nums[slow2]
		if slow2 == slow {
			break
		}
	}

	return slow2
}

func findDuplicatev6(nums []int) int {
	slow, fast := 0, 0

	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			break
		}
	}

	slow2 := 0
	for {
		slow = nums[slow]
		slow2 = nums[slow2]
		if slow == slow2 {
			break
		}
	}

	return slow2
}

func findDuplicatev5(nums []int) int {
	slow, fast := 0, 0

	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			break
		}
	}
	slow2 := 0

	for {
		slow = nums[slow]
		slow2 = nums[slow2]
		if slow == slow2 {
			break
		}
	}

	return slow2
}

func findDuplicatev4(nums []int) int {
	slow, fast := 0, 0

	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			break
		}
	}

	slow2 := 0

	for {
		slow = nums[slow]
		slow2 = nums[slow2]
		if slow == slow2 {
			break
		}
	}

	return slow2
}

func findDuplicatev3(nums []int) int {
	slow, fast := 0, 0
	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			break
		}
	}

	slow2 := 0
	for {
		slow = nums[slow]
		slow2 = nums[slow2]
		if slow2 == slow {
			break
		}
	}

	return slow2
}

func findDuplicatev2(nums []int) int {
	slow, fast := 0, 0

	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			break
		}
	}

	slow2 := 0

	for {
		slow = nums[slow]
		slow2 = nums[slow2]
		if slow2 == slow {
			break
		}
	}

	return slow2
}

func findDuplicatev1(nums []int) int {
	slow, fast := 0, 0

	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			break
		}
	}

	slow2 := 0

	for {
		slow = nums[slow]
		slow2 = nums[slow2]
		if slow == slow2 {
			break
		}
	}

	return slow2
}
