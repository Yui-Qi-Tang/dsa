package b75

import (
	"math"
)

/*
Given an integer array nums,
return the length of the longest strictly increasing subsequence.

A subsequence is a sequence that can be derived from an array by deleting some or no elements without changing the order of the remaining elements.
For example, [3,6,2,7] is a subsequence of the array [0,3,1,6,2,2,7].

Example 1:

Input: nums = [10,9,2,5,3,7,101,18]
Output: 4
Explanation: The longest increasing subsequence is [2,3,7,101], therefore the length is 4.
Example 2:

Input: nums = [0,1,0,3,2,3]
Output: 4
Example 3:

Input: nums = [7,7,7,7,7,7,7]
Output: 1

Constraints:

1 <= nums.length <= 2500
-104 <= nums[i] <= 104
*/

func LengthOfLISv34(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = 1
	result := 1
	for i := 1; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], 1+dp[j])
			}
		}
		result = max(result, dp[i])
	}
	return result
}

func LengthOfLISv33(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = 1
	result := 1

	for i := 1; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], 1+dp[j])
			}
		}
		result = max(result, dp[i])
	}

	return result
}

func LengthOfLISv32(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = 1
	result := 1
	for i := 1; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], 1+dp[j])
			}
		}
		result = max(dp[i], result)
	}

	return result
}

func LengthOfLISv31(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = 1
	result := 1
	for i := 1; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], 1+dp[j])
			}
		}
		result = max(result, dp[i])
	}

	return result
}

func LengthOfLISv30(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = 1

	result := 1
	for i := 1; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], 1+dp[j])
			}
		}
		result = max(result, dp[i])
	}

	return result
}

func LengthOfLISv29(nums []int) int {
	result := 1
	dp := make([]int, len(nums))
	dp[0] = 1
	for i := 1; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], 1+dp[j])
			}
		}
		result = max(dp[i], result)
	}
	return result
}

func LengthOfLISv28(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = 1
	result := 1
	for i := 1; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], 1+dp[j])
			}
		}
		result = max(result, dp[i])
	}
	return result
}

func LengthOfLISv27(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = 1
	result := 1
	for i := 1; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], 1+dp[j])
			}
		}
		result = max(result, dp[i])
	}

	return result
}

func LengthOfLISv26(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = 1
	result := 1
	for i := 1; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], 1+dp[j])
			}
		}
		result = max(result, dp[i])
	}
	return result
}

func LengthOfLISv25(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = 1
	result := 1
	for i := 1; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], 1+dp[j])
			}
		}
		result = max(result, dp[i])
	}

	return result
}

func LengthOfLISv24(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = 1
	result := 1

	for i := 1; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(1+dp[j], dp[i])
			}
		}
		result = max(result, dp[i])
	}

	return result
}

func LengthOfLISv23(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = 1
	result := 1
	for i := 1; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], 1+dp[j])
			}
		}
		result = max(result, dp[i])
	}
	return result
}

func LengthOfLISv22(nums []int) int {
	dp := make([]int, len(nums))
	result := 1
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], 1+dp[j])
			}
		}
		result = max(result, dp[i])
	}

	return result
}

func LengthOfLISv21(nums []int) int {
	dp := make([]int, len(nums))

	result := 1
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], 1+dp[j])
			}
		}
		result = max(result, dp[i])
	}

	return result
}

func LengthOfLISv20(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return n
	}

	result := 1
	dp := make([]int, n)

	for i := range nums {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], 1+dp[j])
			}
		}
		result = max(result, dp[i])
	}

	return result
}

func LengthOfLISv19(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	var result int = 1
	dp := make([]int, n)

	for i := 0; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], 1+dp[j])
			}
		}
		result = max(result, dp[i])
	}

	return result
}

func LengthOfLISv18(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return n
	}

	var longest int = 1
	dp := make([]int, n)

	for i := 0; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], 1+dp[j])
			}
		}

		longest = max(longest, dp[i])
	}

	return longest
}

func LengthOfLISv17(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return n
	}

	dp := make([]int, n)

	var longest int = 0

	for i := 0; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], 1+dp[j])
			}
		}
		longest = max(longest, dp[i])
	}

	return longest
}

func LengthOfLISv16(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return 1
	}

	var longest int = 1
	dp := make([]int, n)

	for i := 0; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], 1+dp[j])
			}
		}
		longest = max(dp[i], longest)
	}

	return longest
}

func LengthOfLISv15(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return n
	}

	dp := make([]int, n)
	var longest int = 0

	for i := 0; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], 1+dp[j])
			}
		}
		longest = max(longest, dp[i])
	}

	return longest
}

func LengthOfLISv14(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return n
	}

	dp := make([]int, n)
	var longest int = 0
	for i := 0; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], 1+dp[j])
			}
		}

		longest = max(dp[i], longest)
	}

	return longest
}

func LengthOfLISv13(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return n
	}

	dp := make([]int, n)
	var longest int = 0

	for i := 0; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], 1+dp[j])
			}
		}
		longest = max(dp[i], longest)
	}

	return longest
}

func LengthOfLISv12(nums []int) int {
	n := len(nums)

	if n <= 1 {
		return n
	}

	dp := make([]int, n)
	var longest int = 0
	for i := 0; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], 1+dp[j])
			}
		}
		longest = max(longest, dp[i])
	}
	return longest

}

func LengthOfLISv11(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return n
	}

	dp := make([]int, n)

	var longest int = 0

	for i := 0; i < n; i++ {
		dp[i] = 1
		for j := 0; j < n; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], 1+dp[j])
			}
		}
		longest = max(dp[i], longest)
	}

	return longest
}

func LengthOfLISv10(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	dp := make([]int, n)

	var longest int = 0

	for i := 0; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], 1+dp[j])
			}
		}
		longest = max(longest, dp[i])
	}
	return longest
}

func LengthOfLISv9(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	dp := make([]int, n)

	var longest int = 0
	for i := 0; i < n; i++ {
		dp[i] = 1

		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], 1+dp[j])
			}
		}

		longest = max(longest, dp[i])
	}

	return longest
}

func LengthOfLISv8(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	dp := make([]int, n)
	var longest int = 0
	for i := 0; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], 1+dp[j])
			}
		}
		longest = max(dp[i], longest)
	}

	return longest
}

func LengthOfLISv7(nums []int) int {
	n := len(nums)
	if n == 0 || n == 1 {
		return n
	}

	dp := make([]int, n)

	var longest int = 0

	for i := 0; i < n; i++ {
		dp[i] = 1

		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], 1+dp[j])
			}
		}

		longest = max(longest, dp[i])
	}

	return longest
}

func LengthOfLISv6(nums []int) int {
	n := len(nums)

	if n == 0 {
		return 0
	}

	dp := make([]int, n)

	var longest int = 0
	for i := 0; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], 1+dp[j])
			}
		}

		longest = max(longest, dp[i])
	}

	return longest
}

func LengthOfLISv5(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	dp := make([]int, n)

	var longest int = 0
	for i := 0; i < n; i++ {
		dp[i] = 1

		// 0<j<i
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], 1+dp[j])
			}
		}

		longest = max(longest, dp[i])

	}

	return longest
}

func LengthOfLISv4(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	dp := make([]int, n)

	result := 0
	for i := 0; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], 1+dp[j])
			}
		}
		result = max(result, dp[i])

	}

	return result
}

func LengthOfLISv3(nums []int) int {
	n := len(nums)

	if n == 0 {
		return 0
	}

	var longest int = 0
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1

		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		longest = max(longest, dp[i])
	}

	return longest
}

func LengthOfLISv2(nums []int) int {

	n := len(nums)
	if n == 0 {
		return 0
	}

	dp := make([]int, n)
	dp[0] = 1

	result := 0
	for i := 0; i < n; i++ {
		dp[i] = 1

		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		result = max(result, dp[i])

	}
	return result
}

func LengthOfLIS(nums []int) int {

	n := len(nums)

	if n <= 1 {
		return n
	}

	m := make([]int, n)
	q := math.MinInt

	for i := 0; i < n; i++ {
		m[i] = 1

		for j := 0; j < i; j++ {
			if nums[i] > nums[j] && m[i] < m[j]+1 {
				m[i] = m[j] + 1
			}

			q = max(q, m[i])
		}
	}

	// fmt.Println(m)

	return q
}
