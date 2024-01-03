package b75

/*
39. Combination Sum
Given an array of distinct integers candidates and a target integer target,
return a list of all unique combinations of candidates where the chosen numbers sum to target.

You may return the combinations in any order.

The same number may be chosen from candidates an unlimited number of times.

Two combinations are unique if the frequencyof at least one of the chosen numbers is different.

The test cases are generated such that the number of unique combinations that sum up to target is less than 150 combinations for the given input.

- Example 1:

Input: candidates = [2,3,6,7], target = 7
Output: [[2,2,3],[7]]
Explanation:
2 and 3 are candidates, and 2 + 2 + 3 = 7. Note that 2 can be used multiple times.
7 is a candidate, and 7 = 7.
These are the only two combinations.

- Example 2:

Input: candidates = [2,3,5], target = 8
Output: [[2,2,2,2],[2,3,3],[3,5]]
Example 3:

Input: candidates = [2], target = 1
Output: []

- Constraints:

1 <= candidates.length <= 30
2 <= candidates[i] <= 40
All elements of candidates are distinct.
1 <= target <= 40

HINT: we should avoid the duplicated premutation
*/

func combinationSumv38(candicates []int, target int) [][]int {
	result := make([][]int, 0)
	var dfs func(i, sum int, curr []int)
	dfs = func(i, sum int, curr []int) {
		if i >= len(candicates) || sum > target {
			return
		}

		if sum == target {
			cp := make([]int, len(curr))
			copy(cp, curr)
			result = append(result, cp)
			return
		}

		dfs(i+1, sum, curr)
		curr = append(curr, candicates[i])
		dfs(i, sum+candicates[i], curr)
	}
	dfs(0, 0, []int{})
	return result
}

func combinationSumv37(candicates []int, target int) [][]int {
	result := make([][]int, 0)
	var dfs func(i, sum int, curr []int)
	dfs = func(i, sum int, curr []int) {
		if i >= len(candicates) || sum > target {
			return
		}
		if sum == target {
			cp := make([]int, len(curr))
			copy(cp, curr)
			result = append(result, cp)
			return
		}

		dfs(i+1, sum, curr)
		curr = append(curr, candicates[i])
		dfs(i, sum+candicates[i], curr)
	}
	dfs(0, 0, []int{})
	return result
}

func combinationSumv36(candicates []int, target int) [][]int {
	result := make([][]int, 0)
	var dfs func(i, sum int, curr []int)
	dfs = func(i, sum int, curr []int) {
		if i >= len(candicates) || sum > target {
			return
		}

		if sum == target {
			cp := make([]int, len(curr))
			copy(cp, curr)
			result = append(result, cp)
			return
		}

		dfs(i+1, sum, curr)
		curr = append(curr, candicates[i])
		dfs(i, sum+candicates[i], curr)
	}
	dfs(0, 0, []int{})
	return result
}

func combinationSumv35(candicates []int, target int) [][]int {
	result := make([][]int, 0)
	var dfs func(i, sum int, curr []int)
	dfs = func(i, sum int, curr []int) {
		if i >= len(candicates) || sum > target {
			return
		}

		if sum == target {
			cp := make([]int, len(curr))
			copy(cp, curr)
			result = append(result, curr)
			return
		}

		dfs(i+1, sum, curr)
		curr = append(curr, candicates[i])
		dfs(i, sum+candicates[i], curr)
	}
	dfs(0, 0, []int{})
	return result
}

func combinationSumv34(candicates []int, target int) [][]int {
	result := make([][]int, 0)

	var dfs func(i, sum int, curr []int)
	dfs = func(i, sum int, curr []int) {
		if i >= len(candicates) || sum > target {
			return
		}

		if sum == target {
			cp := make([]int, len(curr))
			copy(cp, curr)
			result = append(result, cp)
			return
		}

		dfs(i+1, sum, curr)
		curr = append(curr, candicates[i])
		dfs(i, sum+candicates[i], curr)
	}
	dfs(0, 0, []int{})
	return result
}

func combinationSumv33(candicates []int, target int) [][]int {
	result := make([][]int, 0)

	var dfs func(i, sum int, curr []int)
	dfs = func(i, sum int, curr []int) {
		if i >= len(candicates) || sum > target {
			return
		}

		if sum == target {
			cp := make([]int, len(curr))
			copy(cp, curr)
			result = append(result, cp)
			return
		}

		dfs(i+1, sum, curr)
		curr = append(curr, candicates[i])
		dfs(i, sum+candicates[i], curr)
	}
	dfs(0, 0, []int{})
	return result
}

func combinationSumv32(candidates []int, target int) [][]int {
	result := make([][]int, 0)
	var dfs func(i, sum int, curr []int)
	dfs = func(i, sum int, curr []int) {
		if i >= len(candidates) || sum > target {
			return
		}

		if sum == target {
			cp := make([]int, len(curr))
			copy(cp, curr)
			result = append(result, cp)
			return
		}

		dfs(i+1, sum, curr)
		curr = append(curr, candidates[i])
		dfs(i, sum+candidates[i], curr)
	}
	dfs(0, 0, []int{})
	return result
}

func combinationSumv31(candidates []int, target int) [][]int {
	result := make([][]int, 0)
	var dfs func(i, sum int, curr []int)
	dfs = func(i, sum int, curr []int) {
		if i >= len(candidates) || sum > target {
			return
		}

		if sum == target {
			cp := make([]int, len(curr))
			copy(cp, curr)
			result = append(result, cp)
			return
		}

		dfs(i+1, sum, curr)
		curr = append(curr, candidates[i])
		dfs(i, sum+candidates[i], curr)
	}
	dfs(0, 0, []int{})
	return result
}

func combinationSumv30(candidates []int, target int) [][]int {
	result := make([][]int, 0)
	var dfs func(i, sum int, curr []int)
	dfs = func(i, sum int, curr []int) {
		if i >= len(candidates) || sum > target {
			return
		}

		if sum == target {
			cp := make([]int, len(curr))
			copy(cp, curr)
			result = append(result, cp)
			return
		}

		dfs(i+1, sum, curr)
		curr = append(curr, candidates[i])
		dfs(i, sum+candidates[i], curr)
	}
	dfs(0, 0, []int{})
	return result
}

func combinationSumv29(candidates []int, target int) [][]int {
	result := make([][]int, 0)
	var dfs func(i, sum int, curr []int)

	dfs = func(i, sum int, curr []int) {
		if i >= len(candidates) || sum > target {
			return
		}

		if sum == target {
			cp := make([]int, len(curr))
			copy(cp, curr)
			result = append(result, cp)
			return
		}

		dfs(i+1, sum, curr)
		curr = append(curr, candidates[i])
		dfs(i, sum+candidates[i], curr)
	}
	dfs(0, 0, []int{})
	return result
}

func combinationSumv28(candidates []int, target int) [][]int {
	result := make([][]int, 0)

	var dfs func(i, sum int, curr []int)
	dfs = func(i, sum int, curr []int) {
		if i >= len(candidates) || sum > target {
			return
		}

		if sum == target {
			cp := make([]int, len(curr))
			copy(cp, curr)
			result = append(result, cp)
			return
		}
		dfs(i+1, sum, curr)
		curr = append(curr, candidates[i])
		dfs(i, sum+candidates[i], curr)
	}
	dfs(0, 0, []int{})
	return result
}

func combinationSumv27(candidates []int, target int) [][]int {
	result := make([][]int, 0)

	var dfs func(i, sum int, curr []int)
	dfs = func(i, sum int, curr []int) {
		if i >= len(candidates) || sum > target {
			return
		}

		if sum == target {
			cp := make([]int, len(curr))
			copy(cp, curr)
			result = append(result, cp)
			return
		}

		dfs(i+1, sum, curr)
		curr = append(curr, candidates[i])
		dfs(i, sum+candidates[i], curr)
	}
	dfs(0, 0, []int{})
	return result
}

func combinationSumv26(candidates []int, target int) [][]int {
	result := make([][]int, 0)
	var dfs func(i, sum int, curr []int)

	dfs = func(i, sum int, curr []int) {
		if i >= len(candidates) || sum > target {
			return
		}

		if sum == target {
			cp := make([]int, len(curr))
			copy(cp, curr)
			result = append(result, cp)
			return
		}

		dfs(i+1, sum, curr)
		curr = append(curr, candidates[i])
		dfs(i, sum+candidates[i], curr)
	}
	dfs(0, 0, []int{})
	return result
}

func combinationSumv25(candidates []int, target int) [][]int {
	result := make([][]int, 0)

	var dfs func(i, sum int, curr []int)
	dfs = func(i, sum int, curr []int) {
		if i >= len(candidates) || sum > target {
			return
		}

		if sum == target {
			cp := make([]int, len(curr))
			copy(cp, curr)
			result = append(result, cp)
			return
		}

		dfs(i+1, sum, curr)
		curr = append(curr, candidates[i])
		dfs(i, sum+candidates[i], curr)
	}
	dfs(0, 0, []int{})
	return result
}

func combinationSumv24(candidates []int, target int) [][]int {
	result := make([][]int, 0)

	var dfs func(i, sum int, curr []int)

	dfs = func(i, sum int, curr []int) {
		if i >= len(candidates) || sum > target {
			return
		}

		if sum == target {
			cp := make([]int, len(curr))
			copy(cp, curr)
			result = append(result, curr)
			return
		}

		dfs(i+1, sum, curr)
		curr = append(curr, candidates[i])
		dfs(i, sum+candidates[i], curr)
	}
	dfs(0, 0, []int{})
	return result
}

func combinationSumv23(candidates []int, target int) [][]int {
	result := make([][]int, 0)

	var dfs func(i, sum int, curr []int)
	dfs = func(i, sum int, curr []int) {
		if i >= len(candidates) || sum > target {
			return
		}

		if sum == target {
			cp := make([]int, len(curr))
			copy(cp, curr)
			result = append(result, cp)
			return
		}

		dfs(i+1, sum, curr)
		curr = append(curr, candidates[i])
		dfs(i, sum+candidates[i], curr)
	}
	dfs(0, 0, []int{})
	return result
}

func combinationSumv22(candidates []int, target int) [][]int {
	result := make([][]int, 0)

	var dfs func(i, sum int, curr []int)
	dfs = func(i, sum int, curr []int) {
		if i >= len(candidates) || sum > target {
			return
		}

		if sum == target {
			cp := make([]int, len(curr))
			copy(cp, curr)
			result = append(result, cp)
			return
		}

		dfs(i+1, sum, curr)
		curr = append(curr, candidates[i])
		dfs(i, sum+candidates[i], curr)
	}
	dfs(0, 0, []int{})
	return result
}

func combinationSumv21(candidates []int, target int) [][]int {
	result := make([][]int, 0)

	var dfs func(i, sum int, curr []int)

	dfs = func(i, sum int, curr []int) {
		if i >= len(candidates) || sum > target {
			return
		}

		if sum == target {
			cp := make([]int, len(curr))
			copy(cp, curr)
			result = append(result, cp)
			return
		}

		dfs(i+1, sum, curr)
		curr = append(curr, candidates[i])
		dfs(i, sum+candidates[i], curr)
	}
	dfs(0, 0, []int{})
	return result
}

func combinationSumv20(candidates []int, target int) [][]int {
	result := make([][]int, 0)

	var dfs func(i, sum int, curr []int)
	dfs = func(i, sum int, curr []int) {
		if i >= len(candidates) || sum > target {
			return
		}

		if sum == target {
			cp := make([]int, len(curr))
			copy(cp, curr)
			result = append(result, cp)
			return
		}

		dfs(i+1, sum, curr)
		curr = append(curr, candidates[i])
		dfs(i, sum+candidates[i], curr)
	}
	dfs(0, 0, []int{})
	return result
}

func combinationSumv19(candidates []int, target int) [][]int {
	result := make([][]int, 0)

	var dfs func(i, sum int, curr []int)
	dfs = func(i, sum int, curr []int) {
		if i >= len(candidates) || sum > target {
			return
		}

		if sum == target {
			cp := make([]int, len(curr))
			copy(cp, curr)
			result = append(result, cp)
			return
		}

		dfs(i+1, sum, curr)
		curr = append(curr, candidates[i])
		dfs(i, sum+candidates[i], curr)
	}

	dfs(0, 0, []int{})

	return result
}

func combinationSumv18(candidates []int, target int) [][]int {
	result := make([][]int, 0)

	var dfs func(i, sum int, curr []int)
	dfs = func(i, sum int, curr []int) {
		if i >= len(candidates) || sum > target {
			return
		}

		if sum == target {
			cp := make([]int, len(curr))
			copy(cp, curr)
			result = append(result, cp)
			return
		}

		dfs(i+1, sum, curr)
		curr = append(curr, candidates[i])
		dfs(i, sum+candidates[i], curr)
	}
	dfs(0, 0, []int{})
	return result
}

func combinationSumv17(candidates []int, target int) [][]int {
	result := make([][]int, 0)

	var dfs func(i, sum int, curr []int)

	dfs = func(i, sum int, curr []int) {
		if i >= len(candidates) || sum > target {
			return
		}

		if sum == target {
			cp := make([]int, len(curr))
			copy(cp, curr)
			result = append(result, cp)
			return
		}

		dfs(i+1, sum, curr)

		curr = append(curr, candidates[i])
		dfs(i, sum+candidates[i], curr)
	}

	dfs(0, 0, []int{})
	return result
}

func combinationSumv16(candidates []int, target int) [][]int {
	result := make([][]int, 0)

	var dfs func(i, sum int, curr []int)
	dfs = func(i, sum int, curr []int) {
		if i >= len(candidates) || sum > target {
			return
		}

		if sum == target {
			cp := make([]int, len(curr))
			copy(cp, curr)
			result = append(result, cp)
			return
		}

		dfs(i+1, sum, curr)

		curr = append(curr, candidates[i])
		dfs(i, candidates[i]+sum, curr)
	}

	dfs(0, 0, []int{})

	return result
}

func combinationSumv15(candidates []int, target int) [][]int {
	result := make([][]int, 0)

	var dfs func(i, sum int, curr []int)
	dfs = func(i, sum int, curr []int) {
		if i >= len(candidates) || sum > target {
			return
		}

		if sum == target {
			cp := make([]int, len(curr))
			copy(cp, curr)
			result = append(result, cp)
			return
		}

		dfs(i+1, sum, curr)

		curr = append(curr, candidates[i])
		dfs(i, sum+candidates[i], curr)

	}

	dfs(0, 0, []int{})
	return result
}

func combinationSumv14(candidates []int, target int) [][]int {
	result := make([][]int, 0)

	var dfs func(i, sum int, curr []int)

	dfs = func(i, sum int, curr []int) {
		if i >= len(candidates) || sum > target {
			return
		}

		if sum == target {
			cp := make([]int, len(curr))
			copy(cp, curr)
			result = append(result, cp)
			return
		}

		dfs(i+1, sum, curr)

		curr = append(curr, candidates[i])
		dfs(i, sum+candidates[i], curr)
	}

	dfs(0, 0, []int{})
	return result
}

func combinationSumv13(candidates []int, target int) [][]int {
	result := make([][]int, 0)

	var dfs func(i, sum int, curr []int)
	dfs = func(i, sum int, curr []int) {
		if i >= len(candidates) || sum > target {
			return
		}

		if sum == target {
			cp := make([]int, len(curr))
			copy(cp, curr)
			result = append(result, cp)
			return
		}

		dfs(i+1, sum, curr)

		curr = append(curr, candidates[i])
		dfs(i, sum+candidates[i], curr)
	}

	dfs(0, 0, []int{})
	return result
}

func combinationSumv12(candidates []int, target int) [][]int {
	result := make([][]int, 0)

	var dfs func(i, sum int, curr []int)
	dfs = func(i, sum int, curr []int) {
		if i >= len(candidates) || sum > target {
			return
		}

		if sum == target {
			cp := make([]int, len(curr))
			copy(cp, curr)
			result = append(result, cp)
			return
		}

		dfs(i+1, sum, curr)

		curr = append(curr, candidates[i])
		dfs(i, sum+candidates[i], curr)
	}
	dfs(0, 0, []int{})
	return result
}

func combinationSumv11(candidates []int, target int) [][]int {
	result := make([][]int, 0)

	var dfs func(i, sum int, curr []int)

	dfs = func(i, sum int, curr []int) {
		if i >= len(candidates) || sum > target {
			return
		}

		if sum == target {
			cp := make([]int, len(curr))
			copy(cp, curr)
			result = append(result, cp)
			return
		}

		curr = append(curr, candidates[i])
		dfs(i, sum+candidates[i], curr)

		curr = curr[:len(curr)-1]
		dfs(i+1, sum, curr)
	}
	dfs(0, 0, []int{})
	return result
}

func combinationSumv10(candidates []int, target int) [][]int {
	result := make([][]int, 0)

	var dfs func(i, sum int, curr []int)

	dfs = func(i, sum int, curr []int) {
		if i >= len(candidates) || sum > target {
			return
		}

		if sum == target {
			cp := make([]int, len(curr))
			copy(cp, curr)
			result = append(result, cp)
			return
		}

		curr = append(curr, candidates[i])
		dfs(i, sum+candidates[i], curr)

		curr = curr[:len(curr)-1]
		dfs(i+1, sum, curr)
	}
	dfs(0, 0, []int{})
	return result
}

func combinationSumv9(candidates []int, target int) [][]int {
	result := make([][]int, 0)

	var dfs func(i, sum int, curr []int)

	dfs = func(i, sum int, curr []int) {
		if i >= len(candidates) || sum > target {
			return
		}

		if sum == target {
			cp := make([]int, len(curr))
			copy(cp, curr)
			result = append(result, cp)
			return
		}

		curr = append(curr, candidates[i])
		dfs(i, sum+candidates[i], curr)

		curr = curr[:len(curr)-1]
		dfs(i+1, sum, curr)
	}

	dfs(0, 0, []int{})

	return result
}

func combinationSumv8(candidates []int, target int) [][]int {
	result := make([][]int, 0)

	var dfs func(i, sum int, curr []int)
	dfs = func(i, sum int, curr []int) {
		if i >= len(candidates) || sum > target {
			return
		}

		if sum == target {
			cp := make([]int, len(curr))
			copy(cp, curr)
			result = append(result, cp)
			return
		}

		curr = append(curr, candidates[i])
		dfs(i, sum+candidates[i], curr)

		curr = curr[:len(curr)-1]
		dfs(i+1, sum, curr)
	}
	dfs(0, 0, []int{})
	return result
}

func combinationSumv7(candidates []int, target int) [][]int {
	result := make([][]int, 0)

	var dfs func(i, sum int, curr []int)
	dfs = func(i, sum int, curr []int) {
		if i >= len(candidates) || sum > target {
			return
		}

		if sum == target {
			cp := make([]int, len(curr))
			copy(cp, curr)
			result = append(result, cp)
			return
		}

		curr = append(curr, candidates[i])
		dfs(i, sum+candidates[i], curr)

		curr = curr[:len(curr)-1]
		dfs(i+1, sum, curr)

	}

	dfs(0, 0, []int{})
	return result
}

func combinationSumv6(candidates []int, target int) [][]int {
	result := make([][]int, 0)

	var dfs func(i int, sum int, curr []int)
	dfs = func(i, sum int, curr []int) {
		if sum == target {
			cp := make([]int, len(curr))
			copy(cp, curr)
			result = append(result, cp)
			return
		}

		if i >= len(candidates) || sum > target {
			return
		}

		// picked itself
		curr = append(curr, candidates[i])
		dfs(i, sum+candidates[i], curr)

		// no picked
		curr = curr[:len(curr)-1]
		dfs(i+1, sum, curr)
	}

	dfs(0, 0, []int{})
	return result
}

func combinationSumv5(candidates []int, target int) [][]int {
	result := make([][]int, 0)

	var dfs func(i, sum int, curr []int)

	dfs = func(i, sum int, curr []int) {
		if sum == target {
			cp := make([]int, len(curr))
			copy(cp, curr)
			result = append(result, cp)
			return
		}

		if i >= len(candidates) || sum > target {
			return
		}

		curr = append(curr, candidates[i])
		dfs(i, sum+candidates[i], curr)

		curr = curr[:len(curr)-1]
		dfs(i+1, sum, curr)
	}

	dfs(0, 0, []int{})
	return result
}

func combinationSumv4(candidates []int, target int) [][]int {
	result := make([][]int, 0)
	var dfs func(i, sum int, curr []int)

	dfs = func(i, sum int, curr []int) {
		if sum == target {
			cp := make([]int, len(curr))
			copy(cp, curr)
			result = append(result, cp)
			return
		}

		if i >= len(candidates) || sum > target {
			return
		}

		curr = append(curr, candidates[i])
		dfs(i, sum+candidates[i], curr)

		curr = curr[:len(curr)-1]
		dfs(i+1, sum, curr)

	}

	dfs(0, 0, []int{})
	return result
}

func combinationSumv3(candidates []int, target int) [][]int {
	result := make([][]int, 0)

	var dfs func(i int, curr []int, sum int)
	dfs = func(i int, curr []int, sum int) {
		if sum == target {
			permutation := make([]int, len(curr))
			copy(permutation, curr)
			result = append(result, permutation)
			return
		}
		if i >= len(candidates) || sum > target {
			return
		}

		curr = append(curr, candidates[i])
		dfs(i, curr, sum+candidates[i])

		curr = curr[:len(curr)-1]
		dfs(i+1, curr, sum)
	}
	dfs(0, []int{}, 0)
	return result
}

func combinationSumv2(candidates []int, target int) [][]int {
	result := make([][]int, 0)

	var dfs func(i int, curr []int, sum int)
	dfs = func(i int, curr []int, sum int) {
		if sum == target {
			permutation := make([]int, len(curr))
			copy(permutation, curr)
			result = append(result, permutation)
			return
		}
		if i >= len(candidates) || sum > target {
			return
		}

		curr = append(curr, candidates[i])
		dfs(i, curr, sum+candidates[i])

		// if we picked i, we shouldn't put it into another branch
		curr = curr[:len(curr)-1]
		dfs(i+1, curr, sum)

	}

	dfs(0, []int{}, 0)
	return result
}

func combinationSumv1(candidates []int, target int) [][]int {
	result := make([][]int, 0)
	var dfs func(i int, currCandidates []int, sum int)
	dfs = func(i int, currCandidates []int, sum int) {

		if sum == target {
			permutation := make([]int, len(currCandidates))
			copy(permutation, currCandidates)
			result = append(result, permutation)
			return
		}

		if i >= len(candidates) || sum > target {
			return
		}

		// find the permutation including itself
		currCandidates = append(currCandidates, candidates[i])
		dfs(i, currCandidates, sum+candidates[i])

		// pop itself and try next elements
		currCandidates = currCandidates[:len(currCandidates)-1]
		dfs(i+1, currCandidates, sum)
	}

	dfs(0, []int{}, 0)
	return result
}
