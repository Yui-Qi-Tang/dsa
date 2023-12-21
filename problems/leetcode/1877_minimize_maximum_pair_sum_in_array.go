package b75

import "sort"

/*
1877. Minimize Maximum Pair Sum in Array
Medium
The pair sum of a pair (a,b) is equal to a + b. The maximum pair sum is the largest pair sum in a list of pairs.

For example, if we have pairs (1,5), (2,3), and (4,4), the maximum pair sum would be max(1+5, 2+3, 4+4) = max(6, 5, 8) = 8.
Given an array nums of even length n, pair up the elements of nums into n / 2 pairs such that:

Each element of nums is in exactly one pair, and
The maximum pair sum is minimized.
Return the minimized maximum pair sum after optimally pairing up the elements.

Example 1:

Input: nums = [3,5,2,3]
Output: 7
Explanation: The elements can be paired up into pairs (3,3) and (5,2).
The maximum pair sum is max(3+3, 5+2) = max(6, 7) = 7.
Example 2:

Input: nums = [3,5,4,2,4,6]
Output: 8
Explanation: The elements can be paired up into pairs (3,5), (4,4), and (6,2).
The maximum pair sum is max(3+5, 4+4, 6+2) = max(8, 8, 8) = 8.

Constraints:

n == nums.length
2 <= n <= 105
n is even.
1 <= nums[i] <= 105
*/

func minPairSumv28(nums []int) int {
	sort.Ints(nums)
	l, r, result := 0, len(nums)-1, 0

	for l < r {
		result = max(result, nums[l]+nums[r])
		l++
		r--
	}

	return result
}

func minPairSumv27(nums []int) int {
	sort.Ints(nums)
	l, r := 0, len(nums)-1
	result := 0

	for l < r {
		result = max(result, nums[l]+nums[r])
		l++
		r--
	}

	return result
}

func minPairSumv26(nums []int) int {
	sort.Ints(nums)
	l, r := 0, len(nums)-1
	result := 0

	for l < r {
		result = max(result, nums[l]+nums[r])
		l++
		r--
	}

	return result
}

func minPairSumv25(nums []int) int {
	sort.Ints(nums)
	l, r := 0, len(nums)-1
	result := 0

	for l < r {
		result = max(result, nums[l]+nums[r])
		l++
		r--
	}

	return result
}

func minPairSumv24(nums []int) int {
	sort.Ints(nums)
	l, r := 0, len(nums)-1
	result := 0

	for l < r {
		result = max(result, nums[l]+nums[r])
		l++
		r--
	}

	return result
}

func minPairSumv23(nums []int) int {
	sort.Ints(nums)
	l, r := 0, len(nums)-1
	result := 0
	for l < r {
		result = max(result, nums[l]+nums[r])
		l++
		r--
	}

	return result
}

func minPairSumv22(nums []int) int {
	sort.Ints(nums)
	l, r := 0, len(nums)-1
	result := 0
	for l < r {
		result = max(result, nums[l]+nums[r])
		l++
		r--
	}

	return result
}

func minPairSumv21(nums []int) int {
	sort.Ints(nums)
	l, r := 0, len(nums)-1

	result := 0
	for l < r {
		result = max(result, nums[l]+nums[r])
		l++
		r--
	}

	return result
}

func minPairSumv20(nums []int) int {
	sort.Ints(nums)
	l, r := 0, len(nums)-1

	result := 0
	for l < r {
		result = max(result, nums[l]+nums[r])
		l++
		r--
	}

	return result
}

func minPairSumv19(nums []int) int {
	sort.Ints(nums)
	l, r := 0, len(nums)-1

	result := 0
	for l < r {
		result = max(result, nums[l]+nums[r])
		l++
		r--
	}

	return result
}

func minPairSumv18(nums []int) int {
	sort.Ints(nums)
	l, r := 0, len(nums)-1
	result := 0

	for l < r {
		result = max(result, nums[l]+nums[r])
		l++
		r--
	}

	return result
}

func minPairSumv17(nums []int) int {
	sort.Ints(nums)
	l, r := 0, len(nums)-1
	result := 0
	for l < r {
		result = max(result, nums[l]+nums[r])
		l++
		r--
	}

	return result
}

func minPairSumv16(nums []int) int {
	sort.Ints(nums)
	l, r := 0, len(nums)-1
	result := 0

	for l < r {
		result = max(result, nums[l]+nums[r])
		l++
		r--
	}

	return result
}

func minPairSumv15(nums []int) int {
	sort.Ints(nums)
	l, r := 0, len(nums)-1
	result := 0
	for l < r {
		result = max(result, nums[l]+nums[r])
		l++
		r--
	}
	return result
}

func minPairSumv14(nums []int) int {
	sort.Ints(nums)
	l, r := 0, len(nums)-1

	result := 0
	for l < r {
		result = max(result, nums[l]+nums[r])
		l++
		r--
	}
	return result
}

func minPairSumv13(nums []int) int {
	sort.Ints(nums)
	l, r := 0, len(nums)-1
	result := 0

	for l < r {
		result = max(result, nums[l]+nums[r])
		l++
		r--
	}
	return result
}

func minPairSumv12(nums []int) int {
	sort.Ints(nums)
	l, r := 0, len(nums)-1
	result := 0
	for l < r {
		result = max(result, nums[l]+nums[r])
		l++
		r--
	}
	return result
}

func minPairSumv11(nums []int) int {
	sort.Ints(nums)
	l, r := 0, len(nums)-1

	result := 0

	for l < r {
		result = max(result, nums[l]+nums[r])
		r--
		l++
	}

	return result
}

func minPairSumv10(nums []int) int {
	sort.Ints(nums)

	l, r := 0, len(nums)-1
	result := 0

	for l < r {
		result = max(result, nums[l]+nums[r])
		l++
		r--
	}

	return result
}

func minPairSumv9(nums []int) int {
	sort.Ints(nums)
	l, r := 0, len(nums)-1
	result := 0
	for l < r {
		result = max(result, nums[l]+nums[r])
		l++
		r--
	}

	return result
}

func minPairSumv8(nums []int) int {
	sort.Ints(nums)
	result := 0
	l, r := 0, len(nums)-1

	for l < r {
		result = max(result, nums[l]+nums[r])
		l++
		r--
	}

	return result
}

func minPairSumv7(nums []int) int {
	sort.Ints(nums)

	l, r := 0, len(nums)-1
	result := 0
	for l < r {
		result = max(result, nums[l]+nums[r])
		l++
		r--
	}
	return result
}

func minPairSumv6(nums []int) int {
	sort.Ints(nums)
	l, r := 0, len(nums)-1

	result := 0
	for l < r {
		result = max(result, nums[l]+nums[r])
		l++
		r--
	}

	return result
}

func minPairSumv5(nums []int) int {
	sort.Ints(nums)
	l, r := 0, len(nums)-1
	result := 0

	for l < r {
		result = max(result, nums[l]+nums[r])
		l++
		r--
	}
	return result
}

func minPairSumv4(nums []int) int {
	sort.Ints(nums)
	l, r := 0, len(nums)-1
	result := 0

	for l < r {
		result = max(result, nums[l]+nums[r])
		l++
		r--
	}
	return result
}

func minPairSumv3(nums []int) int {
	sort.Ints(nums)
	l, r := 0, len(nums)-1
	result := 0

	for l < r {
		result = max(result, nums[l]+nums[r])
		l++
		r--
	}

	return result
}

func minPairSumv2(nums []int) int {
	sort.Ints(nums)
	result := 0
	l, r := 0, len(nums)-1

	for l < r {
		result = max(result, nums[l]+nums[r])
		l++
		r--
	}

	return result
}

func minPairSumv1(nums []int) int {
	sort.Ints(nums)
	l, r := 0, len(nums)-1
	result := 0
	for l < r {
		result = max(result, nums[l]+nums[r])
		l++
		r--
	}

	return result
}
