package b75

/*
494. Target Sum
Medium

You are given an integer array nums and an integer target.

You want to build an expression out of nums by adding one of the symbols '+' and '-' before each integer in nums and then concatenate all the integers.

For example, if nums = [2, 1], you can add a '+' before 2 and a '-' before 1 and concatenate them to build the expression "+2-1".
Return the number of different expressions that you can build, which evaluates to target.



Example 1:

Input: nums = [1,1,1,1,1], target = 3
Output: 5
Explanation: There are 5 ways to assign symbols to make the sum of nums be target 3.
-1 + 1 + 1 + 1 + 1 = 3
+1 - 1 + 1 + 1 + 1 = 3
+1 + 1 - 1 + 1 + 1 = 3
+1 + 1 + 1 - 1 + 1 = 3
+1 + 1 + 1 + 1 - 1 = 3
Example 2:

Input: nums = [1], target = 1
Output: 1


Constraints:

1 <= nums.length <= 20
0 <= nums[i] <= 1000
0 <= sum(nums[i]) <= 1000
-1000 <= target <= 1000
*/

func findTargetSumWaysv37(nums []int, target int) int {
	dp := make(map[int]map[int]int)
	var backtrack func(i, sum int) int
	backtrack = func(i, sum int) int {
		if i == len(nums) {
			if sum == target {
				return 1
			}
			return 0
		}

		if ans, exist := dp[i][sum]; exist {
			return ans
		}

		dp[i] = make(map[int]int)
		dp[i][sum] = backtrack(i+1, sum+nums[i]) + backtrack(i+1, sum-nums[i])
		return dp[i][sum]
	}

	return backtrack(0, 0)
}

func findTargetSumWaysv36(nums []int, target int) int {
	dp := make(map[int]map[int]int)
	var backtrace func(i, sum int) int
	backtrace = func(i, sum int) int {
		if i == len(nums) {
			if sum == target {
				return 1
			}
			return 0
		}

		if ans, exist := dp[i][sum]; exist {
			return ans
		}

		dp[i] = make(map[int]int)
		dp[i][sum] = backtrace(i+1, sum+nums[i]) + backtrace(i+1, sum-nums[i])
		return dp[i][sum]
	}
	return backtrace(0, 0)
}

func findTargetSumWaysv35(nums []int, target int) int {
	dp := make(map[int]map[int]int)
	var backtrace func(i, sum int) int
	backtrace = func(i, sum int) int {
		if i == len(nums) {
			if sum == target {
				return 1
			}
			return 0
		}

		if ans, exist := dp[i][sum]; exist {
			return ans
		}

		dp[i] = make(map[int]int)
		dp[i][sum] = backtrace(i+1, sum+nums[i]) + backtrace(i+1, sum-nums[i])
		return dp[i][sum]
	}

	return backtrace(0, 0)
}

func findTargetSumWaysv34(nums []int, target int) int {
	dp := make(map[int]map[int]int)
	var backtrace func(i, sum int) int
	backtrace = func(i, sum int) int {
		if i == len(nums) {
			if sum == target {
				return 1
			}
			return 0
		}

		if ans, exist := dp[i][sum]; exist {
			return ans
		}

		dp[i] = make(map[int]int)
		dp[i][sum] = backtrace(i+1, sum+nums[i]) + backtrace(i+1, sum-nums[i])
		return dp[i][sum]
	}
	return backtrace(0, 0)
}

func findTargetSumWaysv33(nums []int, target int) int {
	dp := make(map[int]map[int]int)
	var backtrace func(i, sum int) int
	backtrace = func(i, sum int) int {
		if i == len(nums) {
			if sum == target {
				return 1
			}
			return 0
		}

		if ans, exist := dp[i][sum]; exist {
			return ans
		}

		dp[i] = make(map[int]int)
		dp[i][sum] = backtrace(i+1, sum+nums[i]) + backtrace(i+1, sum-nums[i])
		return dp[i][sum]
	}

	return backtrace(0, 0)
}

func findTargetSumWaysv32(nums []int, target int) int {
	dp := make(map[int]map[int]int)
	var backtrace func(i, sum int) int
	backtrace = func(i, sum int) int {
		if i == len(nums) {
			if sum == target {
				return 1
			}
			return 0
		}

		if ans, exist := dp[i][sum]; exist {
			return ans
		}
		dp[i] = make(map[int]int)
		dp[i][sum] = backtrace(i+1, sum+nums[i]) + backtrace(i+1, sum-nums[i])
		return dp[i][sum]
	}
	return backtrace(0, 0)
}

func findTargetSumWaysv31(nums []int, target int) int {
	dp := make(map[int]map[int]int)
	var backtrace func(i, sum int) int
	backtrace = func(i, sum int) int {
		if i == len(nums) {
			if sum == target {
				return 1
			}
			return 0
		}

		if ans, exist := dp[i][sum]; exist {
			return ans
		}

		dp[i] = make(map[int]int)
		dp[i][sum] = backtrace(i+1, sum+nums[i]) + backtrace(i+1, sum-nums[i])
		return dp[i][sum]
	}
	return backtrace(0, 0)
}

func findTargetSumWaysv30(nums []int, target int) int {
	dp := make(map[int]map[int]int) // idx: sum -> ways
	var backtrace func(i, sum int) int
	backtrace = func(i, sum int) int {
		if i == len(nums) {
			if sum == target {
				return 1
			}
			return 0
		}

		if ans, exist := dp[i][sum]; exist {
			return ans
		}

		dp[i] = make(map[int]int)
		dp[i][sum] = backtrace(i+1, sum+nums[i]) + backtrace(i+1, sum-nums[i])
		return dp[i][sum]
	}
	return backtrace(0, 0)
}

func findTargetSumWaysv29(nums []int, target int) int {
	dp := make(map[int]map[int]int)
	var backtrace func(i, sum int) int
	backtrace = func(i, sum int) int {
		if i == len(nums) {
			if sum == target {
				return 1
			}
			return 0
		}

		if ans, exist := dp[i][sum]; exist {
			return ans
		}

		dp[i] = make(map[int]int)
		dp[i][sum] = backtrace(i+1, sum+nums[i]) + backtrace(i+1, sum-nums[i])
		return dp[i][sum]
	}
	return backtrace(0, 0)
}

func findTargetSumWaysv28(nums []int, target int) int {
	dp := make(map[int]map[int]int) // idx: sum -> ways
	var backtrace func(i, sum int) int
	backtrace = func(i, sum int) int {
		if i == len(nums) {
			if sum == target {
				return 1
			}
			return 0
		}

		if ans, exist := dp[i][sum]; exist {
			return ans
		}

		dp[i] = make(map[int]int)
		dp[i][sum] = backtrace(i+1, sum+nums[i]) + backtrace(i+1, sum-nums[i])
		return dp[i][sum]
	}
	return backtrace(0, 0)
}

func findTargetSumWaysv27(nums []int, target int) int {
	dp := make(map[int]map[int]int)
	var backtrace func(i, sum int) int
	backtrace = func(i, sum int) int {
		if i == len(nums) {
			if sum == target {
				return 1
			}
			return 0
		}

		if ans, exist := dp[i][sum]; exist {
			return ans
		}

		dp[i] = make(map[int]int)
		dp[i][sum] = backtrace(i+1, sum+nums[i]) + backtrace(i+1, sum-nums[i])
		return dp[i][sum]
	}
	return backtrace(0, 0)
}

func findTargetSumWaysv26(nums []int, target int) int {
	dp := make(map[int]map[int]int)
	var backtrace func(i, sum int) int
	backtrace = func(i, sum int) int {
		if i == len(nums) {
			if sum == target {
				return 1
			}
			return 0
		}

		if ans, exist := dp[i][sum]; exist {
			return ans
		}

		dp[i] = make(map[int]int)
		dp[i][sum] = backtrace(i+1, sum+nums[i]) + backtrace(i+1, sum-nums[i])
		return dp[i][sum]
	}
	return backtrace(0, 0)
}

func findTargetSumWaysv25(nums []int, target int) int {
	dp := make(map[int]map[int]int) // idx: sum -> ways
	var backtrace func(i, sum int) int
	backtrace = func(i, sum int) int {
		if i == len(nums) {
			if sum == target {
				return 1
			}
			return 0
		}

		if v, exist := dp[i][sum]; exist {
			return v
		}
		dp[i] = make(map[int]int)
		dp[i][sum] = backtrace(i+1, sum+nums[i]) + backtrace(i+1, sum-nums[i])
		return dp[i][sum]
	}
	return backtrace(0, 0)
}

func findTargetSumWaysv24(nums []int, target int) int {
	var backtrack func(i, sum int) int
	dp := make(map[int]map[int]int)
	backtrack = func(i, sum int) int {
		if i == len(nums) {
			if sum == target {
				return 1
			}
			return 0
		}

		if v, exist := dp[i][sum]; exist {
			return v
		}
		dp[i] = make(map[int]int)
		dp[i][sum] = backtrack(i+1, sum+nums[i]) + backtrack(i+1, sum-nums[i])
		return dp[i][sum]
	}
	return backtrack(0, 0)
}

func findTargetSumWaysv23(nums []int, target int) int {
	dp := make(map[int]map[int]int) // idx: sum -> ways
	var backtrack func(i, sum int) int
	backtrack = func(i, sum int) int {
		if i == len(nums) {
			if sum == target {
				return 1
			}
			return 0
		}

		if v, exist := dp[i][sum]; exist {
			return v
		}

		dp[i] = make(map[int]int)
		dp[i][sum] = backtrack(i+1, sum+nums[i]) + backtrack(i+1, sum-nums[i])
		return dp[i][sum]
	}

	return backtrack(0, 0)
}

func findTargetSumWaysv22(nums []int, target int) int {
	var dfs func(i, sum int) int
	dp := make(map[int]map[int]int) // idx: sum -> ways
	dfs = func(i, sum int) int {
		if i == len(nums) {
			if sum == target {
				return 1
			}
			return 0
		}

		if v, exist := dp[i][sum]; exist {
			return v
		}

		dp[i] = make(map[int]int)
		dp[i][sum] = dfs(i+1, sum+nums[i]) + dfs(i+1, sum-nums[i])
		return dp[i][sum]
	}
	return dfs(0, 0)
}

func findTargetSumWaysv21(nums []int, target int) int {
	var dfs func(i, sum int) int
	dp := make(map[int]map[int]int)
	dfs = func(i, sum int) int {
		if i == len(nums) {
			if sum == target {
				return 1
			} else {
				return 0
			}
		}

		if v, exist := dp[i][sum]; exist {
			return v
		}
		dp[i] = make(map[int]int)
		dp[i][sum] = dfs(i+1, sum+nums[i]) + dfs(i+1, sum-nums[i])
		return dp[i][sum]
	}
	return dfs(0, 0)
}

func findTargetSumWaysv20(nums []int, target int) int {
	var dfs func(i, sum int) int
	dp := make(map[int]map[int]int)
	dfs = func(i, sum int) int {
		if i == len(nums) {
			if sum == target {
				return 1
			}
			return 0

		}

		if v, exist := dp[i][sum]; exist {
			return v
		}
		dp[i] = make(map[int]int)
		dp[i][sum] = dfs(i+1, sum+nums[i]) + dfs(i+1, sum-nums[i])
		return dp[i][sum]
	}

	return dfs(0, 0)
}

func findTargetSumWaysv19(nums []int, target int) int {

	dp := make(map[int]map[int]int) // idx: sum -> ways
	var dfs func(i, sum int) int
	dfs = func(i, sum int) int {
		if i == len(nums) {
			if sum == target {
				return 1
			} else {
				return 0
			}
		}
		if v, ok := dp[i][sum]; ok {
			return v
		}
		dp[i] = make(map[int]int)
		dp[i][sum] = dfs(i+1, sum+nums[i]) + dfs(i+1, sum-nums[i])
		return dp[i][sum]
	}

	return dfs(0, 0)
}

func findTargetSumWaysv18(nums []int, target int) int {
	dp := make(map[int]map[int]int) // idx: sum -> ways

	var dfs func(i int, sum int) int

	dfs = func(i, sum int) int {
		if i == len(nums) {
			if sum == target {
				return 1
			} else {
				return 0
			}
		}
		if v, ok := dp[i][sum]; ok {
			return v
		}
		dp[i] = make(map[int]int)
		dp[i][sum] = dfs(i+1, sum+nums[i]) + dfs(i+1, sum-nums[i])
		return dp[i][sum]
	}

	return dfs(0, 0)
}

func findTargetSumWaysv17(nums []int, target int) int {
	dp := make(map[int]map[int]int) // num idx: sum -> ways

	var dfs func(i, sum int) int

	dfs = func(i, sum int) int {

		if i == len(nums) {
			if sum == target {
				return 1
			}
			return 0
		}

		if v, exist := dp[i][sum]; exist {
			return v
		}

		dp[i] = make(map[int]int)
		dp[i][sum] = dfs(i+1, sum+nums[i]) + dfs(i+1, sum-nums[i])
		return dp[i][sum]
	}

	return dfs(0, 0)
}

func findTargetSumWaysv16(nums []int, target int) int {
	dp := make(map[int]map[int]int)

	var dfs func(i, sum int) int

	dfs = func(i, sum int) int {
		if i == len(nums) {
			if sum == target {
				return 1
			}
			return 0
		}

		if v, exist := dp[i][sum]; exist {
			return v
		}

		dp[i] = make(map[int]int)
		dp[i][sum] = dfs(i+1, sum+nums[i]) + dfs(i+1, sum-nums[i])
		return dp[i][sum]
	}

	return dfs(0, 0)
}

func findTargetSumWaysv15(nums []int, target int) int {
	dp := make(map[int]map[int]int) // num idx: sum -> ways

	var dfs func(i, sum int) int
	dfs = func(i, sum int) int {
		if i == len(nums) {
			if sum == target {
				return 1
			}
			return 0
		}

		if v, exist := dp[i][sum]; exist {
			return v
		}
		dp[i] = make(map[int]int)
		dp[i][sum] = dfs(i+1, sum+nums[i]) + dfs(i+1, sum-nums[i])
		return dp[i][sum]
	}

	return dfs(0, 0)
}

func findTargetSumWaysv14(nums []int, target int) int {
	dp := make(map[int]map[int]int) // num idx: sum -> ways

	var dfs func(i, sum int) int
	dfs = func(i, sum int) int {

		if i == len(nums) {
			if sum == target {
				return 1
			}
			return 0
		}

		if v, exist := dp[i][sum]; exist {
			return v
		}

		dp[i] = make(map[int]int)
		dp[i][sum] = dfs(i+1, sum+nums[i]) + dfs(i+1, sum-nums[i])
		return dp[i][sum]

	}

	return dfs(0, 0)
}

func findTargetSumWaysv13(nums []int, target int) int {

	dp := make(map[int]map[int]int) // num index: sum -> ways

	var dfs func(i, sum int) int
	dfs = func(i, sum int) int {
		if i == len(nums) {
			if sum == target {
				return 1
			}
			return 0
		}

		if v, exist := dp[i][sum]; exist {
			return v
		}

		dp[i] = make(map[int]int)

		dp[i][sum] = dfs(i+1, sum+nums[i]) + dfs(i+1, sum-nums[i])
		return dp[i][sum]
	}

	return dfs(0, 0)
}

func findTargetSumWaysv12(nums []int, target int) int {
	dp := make(map[int]map[int]int) // num index: sum -> ways

	var dfs func(i, sum int) int

	dfs = func(i, sum int) int {

		if i == len(nums) {
			if sum == target {
				return 1
			}
			return 0
		}

		if v, exist := dp[i][sum]; exist {
			return v
		}

		dp[i] = make(map[int]int)
		dp[i][sum] = dfs(i+1, sum+nums[i]) + dfs(i+1, sum-nums[i])
		return dp[i][sum]

	}

	return dfs(0, 0)
}

func findTargetSumWaysv11(nums []int, target int) int {

	dp := make(map[int]map[int]int) // num index: sum -> ways

	var dfs func(i, sum int) int

	dfs = func(i, sum int) int {

		if i == len(nums) {
			if sum == target {
				return 1
			}
			return 0
		}

		if v, exist := dp[i][sum]; exist {
			return v
		}

		dp[i] = make(map[int]int)
		dp[i][sum] = dfs(i+1, sum+nums[i]) + dfs(i+1, sum-nums[i])
		return dp[i][sum]
	}

	return dfs(0, 0)
}

func findTargetSumWaysv10(nums []int, target int) int {
	dp := make(map[int]map[int]int) // num index: sum -> ways

	var dfs func(i, sum int) int
	dfs = func(i, sum int) int {
		if i == len(nums) {
			if sum == target {
				return 1
			}
			return 0
		}

		if v, exist := dp[i][sum]; exist {
			return v
		}

		dp[i] = make(map[int]int)
		dp[i][sum] = dfs(i+1, sum+nums[i]) + dfs(i+1, sum-nums[i])
		return dp[i][sum]
	}

	return dfs(0, 0)
}

// dfs way
func findTargetSumWaysv9(nums []int, target int) int {
	dp := make(map[int]map[int]int) // num index: sum -> ways

	var dfs func(i, sum int) int

	dfs = func(i, sum int) int {

		if i == len(nums) {
			if sum == target {
				return 1
			}
			return 0
		}

		if v, exist := dp[i][sum]; exist {
			return v
		}

		dp[i] = make(map[int]int)
		dp[i][sum] = dfs(i+1, sum+nums[i]) + dfs(i+1, sum-nums[i])
		return dp[i][sum]
	}

	return dfs(0, 0)
}

// dfs way
func findTargetSumWaysv8(nums []int, target int) int {
	dp := make(map[int]map[int]int) // nums index: sum -> ways

	var backtrack func(i, sum int) int

	backtrack = func(i, sum int) int {

		if i == len(nums) {
			if sum == target {
				return 1
			} else {
				return 0
			}
		}

		if v, exist := dp[i][sum]; exist {
			return v
		}

		dp[i] = make(map[int]int)
		dp[i][sum] = backtrack(i+1, sum+nums[i]) + backtrack(i+1, sum-nums[i])
		return dp[i][sum]
	}
	return backtrack(0, 0)
}

func findTargetSumWaysv7(nums []int, target int) int {
	dp := make(map[int]map[int]int) // num index: target -> ways
	var backtrack func(i, sum int) int
	backtrack = func(i, sum int) int {

		if i == len(nums) {
			if sum == target {
				return 1
			}
			return 0
		}

		if v, exist := dp[i][sum]; exist {
			return v
		}

		dp[i] = make(map[int]int)
		dp[i][sum] = backtrack(i+1, sum+nums[i]) + backtrack(i+1, sum-nums[i])
		return dp[i][sum]

	}

	return backtrack(0, 0)
}

func findTargetSumWaysv6(nums []int, target int) int {
	dp := make(map[int]map[int]int) // num_idx: target -> ways

	var backtrack func(i, sum int) int

	backtrack = func(i, sum int) int {

		if i == len(nums) {
			if sum == target {
				return 1
			} else {
				return 0
			}
		}

		if v, exist := dp[i][sum]; exist {
			return v
		}

		dp[i] = make(map[int]int)

		dp[i][sum] = backtrack(i+1, sum+nums[i]) + backtrack(i+1, sum-nums[i])

		return dp[i][sum]

	}

	return backtrack(0, 0)
}

func findTargetSumWaysv5(nums []int, target int) int {

	dp := make(map[int]map[int]int) // index: target -> ways

	var backtrack func(i, sum int) int
	backtrack = func(i, sum int) int {
		if i == len(nums) {
			if sum == target {
				return 1
			}
			return 0
		}

		if v, exist := dp[i][sum]; exist {
			return v
		}

		dp[i] = make(map[int]int)
		dp[i][sum] = backtrack(i+1, sum+nums[i]) + backtrack(i+1, sum-nums[i])
		return dp[i][sum]

	}

	return backtrack(0, 0)

}
func findTargetSumWaysv4(nums []int, target int) int {
	dp := make(map[int]map[int]int) // num_index:sum -> ways

	var dfs func(i, sum int) int
	dfs = func(i, sum int) int {

		if i == len(nums) {
			if sum == target {
				return 1
			}
			return 0
		}

		if v, exist := dp[i][sum]; exist {
			return v
		}

		dp[i] = make(map[int]int)
		dp[i][sum] = dfs(i+1, sum+nums[i]) + dfs(i+1, sum-nums[i])
		return dp[i][sum]

	}

	return dfs(0, 0)

}

func findTargetSumWaysv3(nums []int, target int) int {
	dp := make(map[int]map[int]int) // (i, total) -> ways

	var backtrack func(i, total int) int

	backtrack = func(i, total int) int {
		if i == len(nums) {
			if total == target {
				return 1
			}
			return 0
		}
		if v, exists := dp[i][total]; exists {
			return v
		}

		dp[i] = make(map[int]int)
		// dfs way=> root value = left done + right done
		dp[i][total] = backtrack(i+1, total+nums[i]) + backtrack(i+1, total-nums[i])
		return dp[i][total]
	}
	return backtrack(0, 0)
}

func findTargetSumWaysv2(nums []int, target int) int {

	dp := make(map[int]map[int]int) // (index, total) -> ways

	var backtrack func(i, total int) int
	backtrack = func(i, total int) int {

		// check the indexing of nums
		if i == len(nums) {
			if total == target {
				return 1
			}
			return 0
		}

		if v, exist := dp[i][total]; exist {
			return v
		}

		dp[i] = make(map[int]int)
		// +1, -1
		dp[i][total] = backtrack(i+1, total+nums[i]) + backtrack(i+1, total-nums[i])

		return dp[i][total]
	}

	return backtrack(0, 0)

}

func findTargetSumWaysv1(nums []int, target int) int {
	dp := make(map[int]map[int]int) // (index, total) -> ways

	var backtrack func(i, total int) int

	backtrack = func(i, total int) int {
		if i == len(nums) {
			if total == target {
				return 1
			} else {
				return 0
			}
		}

		if v, exist := dp[i][total]; exist {
			return v
		}

		dp[i] = make(map[int]int)
		dp[i][total] = backtrack(i+1, total+nums[i]) +
			backtrack(i+1, total-nums[i])
		return dp[i][total]
	}

	return backtrack(0, 0)
}
