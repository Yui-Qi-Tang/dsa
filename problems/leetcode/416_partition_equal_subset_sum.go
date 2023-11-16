package b75

/*
416. Partition Equal Subset Sum
Given an integer array nums,
return true if you can partition the array into two subsets
such that the sum of the elements in both subsets is equal or
false otherwise.



Example 1:

Input: nums = [1,5,11,5]
Output: true
Explanation: The array can be partitioned as [1, 5, 5] and [11].
Example 2:

Input: nums = [1,2,3,5]
Output: false
Explanation: The array cannot be partitioned into equal sum subsets.


Constraints:

1 <= nums.length <= 200
1 <= nums[i] <= 100
*/

func canPartitionv12(nums []int) bool {
	if len(nums) < 2 {
		return false
	}

	sum := 0
	for i := range nums {
		sum += nums[i]
	}

	if sum%2 != 0 {
		return false
	}

	target := sum / 2
	dp := make(map[int]bool)
	dp[0] = true

	for i := len(nums) - 1; i >= 0; i-- {
		nextDP := make(map[int]bool)
		for k := range dp {
			if k+nums[i] == target {
				return true
			} else if k+nums[i] < target {
				nextDP[k+nums[i]] = true
			}
			nextDP[k] = true
		}
		dp = nextDP
	}

	return dp[target]
}

func canPartitionv11(nums []int) bool {
	n := len(nums)
	if n < 2 {
		return false
	}

	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%2 != 0 {
		return false
	}

	target := sum / 2
	dp := make(map[int]bool)
	dp[0] = true

	for i := n - 1; i >= 0; i-- {
		nextDP := make(map[int]bool)

		for k := range dp {
			if nums[i]+k == target {
				return true
			} else if nums[i]+k < target {
				nextDP[nums[i]+k] = true
			}
			nextDP[k] = true
		}

		dp = nextDP
	}

	return dp[target]
}

func canPartitionv10(nums []int) bool {
	n := len(nums)
	if n < 2 {
		return false
	}

	sum := 0
	for _, num := range nums {
		sum += num
	}

	if sum%2 != 0 {
		return false
	}

	target := sum / 2
	dp := make(map[int]bool)
	dp[0] = true

	for i := n - 1; i >= 0; i-- {
		nextDP := make(map[int]bool)
		for k := range dp {
			if k+nums[i] == target {
				return true
			} else if k+nums[i] < target {
				nextDP[k+nums[i]] = true
			}
			nextDP[k] = true
		}
		dp = nextDP
	}

	return dp[target]
}

func canPartitionv9(nums []int) bool {
	n := len(nums)
	if n < 2 {
		return false
	}

	sum := 0
	for _, num := range nums {
		sum += num
	}

	if sum%2 != 0 {
		return false
	}

	target := sum / 2
	dp := make(map[int]bool)
	dp[0] = true

	for i := n - 1; i >= 0; i-- {
		nextDP := make(map[int]bool)

		for k := range dp {
			if k+nums[i] == target {
				return true
			}
			if k+nums[i] < target {
				nextDP[k+nums[i]] = true
			}
			nextDP[k] = true
		}

		dp = nextDP
	}

	return dp[target]
}

func canPartitionv8(nums []int) bool {
	n := len(nums)
	if n < 2 {
		return false
	}

	sum := 0
	for _, num := range nums {
		sum += num
	}

	if sum%2 != 0 {
		return false
	}

	target := sum / 2
	dp := make(map[int]bool)
	dp[0] = true

	for i := n - 1; i >= 0; i-- {
		nextDP := make(map[int]bool)

		for k := range dp {
			if k+nums[i] == target {
				return true
			} else if k+nums[i] < target {
				nextDP[k+nums[i]] = true
			}
			nextDP[k] = true

		}

		dp = nextDP
	}

	return dp[target]
}

func canPartitionv7(nums []int) bool {
	n := len(nums)

	if n <= 1 {
		return false
	}

	sum := 0
	for _, num := range nums {
		sum += num
	}

	if sum%2 != 0 {
		return false
	}

	target := sum / 2
	dp := make(map[int]bool)
	dp[0] = true

	for i := n - 1; i >= 0; i-- {
		nextDP := make(map[int]bool)
		for k := range dp {
			nextDP[k] = true
			if k+nums[i] == target {
				return true
			} else if k+nums[i] < target {
				nextDP[k+nums[i]] = true
			}
		}
		dp = nextDP
	}

	return dp[target]

}

func canPartitionv6(nums []int) bool {
	n := len(nums)
	if n <= 1 {
		return false
	}

	sum := 0
	for _, num := range nums {
		sum += num
	}

	if sum%2 != 0 {
		return false
	}

	target := sum / 2

	dp := make(map[int]bool)
	dp[0] = true

	for i := n - 1; i >= 0; i-- {
		nextDP := make(map[int]bool)
		for k := range dp {
			if k+nums[i] == target {
				return true
			} else if k+nums[i] < target {
				nextDP[k+nums[i]] = true
			}
			nextDP[k] = true
		}
		dp = nextDP
	}

	return dp[target]
}

func canPartitionv5(nums []int) bool {
	n := len(nums)
	if n <= 1 {
		return false
	}

	sum := 0
	for _, num := range nums {
		sum += num
	}

	if sum%2 != 0 {
		return false
	}

	target := sum / 2

	dp := make(map[int]bool)
	dp[0] = true
	for i := n - 1; i >= 0; i-- {
		nextDP := make(map[int]bool)

		for k := range dp {
			nextDP[k] = true
			if k+nums[i] <= target {
				nextDP[k+nums[i]] = true
			}
		}

		dp = nextDP
	}

	return dp[target]
}

func canPartitionv4(nums []int) bool {
	n := len(nums)
	if n <= 1 {
		return false
	}

	sum := 0
	for _, num := range nums {
		sum += num
	}

	if sum%2 != 0 {
		return false
	}

	target := sum / 2
	dp := make(map[int]bool) // sum: exist
	dp[0] = true

	for i := n - 1; i >= 0; i-- {
		nextDP := make(map[int]bool)
		for k := range dp {
			nextDP[k] = true
			if k+nums[i] <= target {
				nextDP[k+nums[i]] = true
			}
		}
		dp = nextDP
	}

	return dp[target]

}

func canPartitionv3(nums []int) bool {
	n := len(nums)
	if n <= 1 {
		return false
	}

	sum := 0
	for _, num := range nums {
		sum += num
	}

	if sum%2 != 0 {
		return false
	}
	target := sum / 2

	dp := make(map[int]bool) // target: exist
	dp[0] = true

	for i := n - 1; i >= 0; i-- {
		nextDP := make(map[int]bool)
		for k := range dp {
			nextDP[k] = true
			if k+nums[i] <= target {
				nextDP[k+nums[i]] = true
			}
		}
		dp = nextDP
	}

	return dp[target]
}

func canPartitionv2(nums []int) bool {
	n := len(nums)

	if n <= 2 {
		return false
	}

	sum := 0
	for _, num := range nums {
		sum += num
	}

	if sum%2 != 0 {
		return false
	}
	target := sum / 2

	dp := make(map[int]bool)
	dp[0] = true
	for i := n - 1; i >= 0; i-- {
		nextDP := make(map[int]bool)
		for k := range dp {
			if nums[i] <= target {
				nextDP[k] = true
				nextDP[k+nums[i]] = true
			}
		}
		dp = nextDP
	}

	return dp[target]
}

func canPartitionv1(nums []int) bool {

	if len(nums) == 0 {
		return false
	}

	sum := 0
	for _, num := range nums {
		sum += num
	}

	if sum%2 != 0 {
		return false
	}

	target := sum / 2

	dp := make(map[int]bool)
	dp[0] = true
	for i := len(nums) - 1; i >= 0; i-- {
		newDP := make(map[int]bool)
		for k := range dp {
			if nums[i] <= target {
				newDP[k] = true
			}

			if nums[i]+k <= target {
				newDP[nums[i]+k] = true
			}
		}
		dp = newDP
	}

	return dp[target]
}
