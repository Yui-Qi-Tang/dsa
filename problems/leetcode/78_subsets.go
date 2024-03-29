package b75

/*
78. Subsets

Given an integer array nums of unique elements,
return all possible subsets (the power set).

The solution set must not contain duplicate subsets.
Return the solution in any order.

- Example 1:

Input: nums = [1,2,3]
Output: [[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]

- Example 2:

Input: nums = [0]
Output: [[],[0]]


- Constraints:

    1 <= nums.length <= 10
    -10 <= nums[i] <= 10
    All the numbers of nums are unique.
HINT: each element has 2 choices, 1 for picked up and 2 do nothing
*/

func subsetsv52(nums []int) [][]int {
	result := make([][]int, 0)
	var backtrace func(i int, curr []int)
	backtrace = func(i int, curr []int) {
		if i == len(nums) {
			cp := make([]int, len(curr))
			copy(cp, curr)
			result = append(result, cp)
			return
		}

		backtrace(i+1, curr)
		curr = append(curr, nums[i])
		backtrace(i+1, curr)

	}
	backtrace(0, []int{})
	return result
}

func subsetsv51(nums []int) [][]int {
	result := make([][]int, 0)
	var backtrace func(i int, curr []int)
	backtrace = func(i int, curr []int) {
		if i == len(nums) {
			cp := make([]int, len(curr))
			copy(cp, curr)
			result = append(result, cp)
			return
		}

		backtrace(i+1, curr)
		curr = append(curr, nums[i])
		backtrace(i+1, curr)
	}
	backtrace(0, []int{})
	return result
}

func subsetsv50(nums []int) [][]int {
	result := make([][]int, 0)
	var dfs func(i int, curr []int)
	dfs = func(i int, curr []int) {
		if i == len(nums) {
			cp := make([]int, len(curr))
			copy(cp, curr)
			result = append(result, cp)
			return
		}
		dfs(i+1, curr)
		curr = append(curr, nums[i])
		dfs(i+1, curr)
	}
	dfs(0, []int{})
	return result
}

func subsetsv49(nums []int) [][]int {
	result := make([][]int, 0)

	var dfs func(i int, curr []int)
	dfs = func(i int, curr []int) {
		if i == len(nums) {
			cp := make([]int, len(curr))
			copy(cp, curr)
			result = append(result, cp)
			return
		}

		dfs(i+1, curr)
		curr = append(curr, nums[i])
		dfs(i+1, curr)
	}
	dfs(0, []int{})
	return result
}

func subsetsv48(nums []int) [][]int {
	result := make([][]int, 0)

	var dfs func(i int, curr []int)
	dfs = func(i int, curr []int) {
		if i == len(nums) {
			cp := make([]int, len(curr))
			copy(cp, curr)
			result = append(result, cp)
			return
		}

		dfs(i+1, curr)
		curr = append(curr, nums[i])
		dfs(i+1, curr)
	}
	dfs(0, []int{})
	return result
}

func subsetsv47(nums []int) [][]int {
	result := make([][]int, 0)
	var dfs func(i int, curr []int)
	dfs = func(i int, curr []int) {
		if i == len(nums) {
			cp := make([]int, len(curr))
			copy(cp, curr)
			result = append(result, cp)
			return
		}

		dfs(i+1, curr)
		curr = append(curr, nums[i])
		dfs(i+1, curr)
	}
	dfs(0, []int{})
	return result
}

func subsetsv46(nums []int) [][]int {
	result := make([][]int, 0)

	var dfs func(i int, curr []int)
	dfs = func(i int, curr []int) {
		if i == len(nums) {
			cp := make([]int, len(curr))
			copy(cp, curr)
			result = append(result, cp)
			return
		}

		dfs(i+1, curr)
		curr = append(curr, nums[i])
		dfs(i+1, curr)
	}
	dfs(0, []int{})
	return result
}

func subsetsv45(nums []int) [][]int {
	result := make([][]int, 0)

	var dfs func(i int, curr []int)
	dfs = func(i int, curr []int) {
		if i == len(nums) {
			cp := make([]int, len(curr))
			copy(cp, curr)
			result = append(result, cp)
			return
		}

		dfs(i+1, curr)
		curr = append(curr, nums[i])
		dfs(i+1, curr)
	}
	dfs(0, []int{})
	return result
}

func subsetsv44(nums []int) [][]int {
	result := make([][]int, 0)

	var dfs func(i int, curr []int)
	dfs = func(i int, curr []int) {
		if i == len(nums) {
			cp := make([]int, len(curr))
			copy(cp, curr)
			result = append(result, cp)
			return
		}

		dfs(i+1, curr)
		curr = append(curr, nums[i])
		dfs(i+1, curr)
	}
	dfs(0, []int{})
	return result
}

func subsetsv43(nums []int) [][]int {
	result := make([][]int, 0)

	var dfs func(i int, curr []int)
	dfs = func(i int, curr []int) {
		if i == len(nums) {
			cp := make([]int, len(curr))
			copy(cp, curr)
			result = append(result, cp)
			return
		}

		dfs(i+1, curr)
		curr = append(curr, nums[i])
		dfs(i+1, curr)
	}
	dfs(0, []int{})
	return result
}

func subsetsv42(nums []int) [][]int {
	result := make([][]int, 0)
	var dfs func(i int, curr []int)
	dfs = func(i int, curr []int) {
		if i == len(nums) {
			cp := make([]int, len(curr))
			copy(cp, curr)
			result = append(result, cp)
			return
		}

		dfs(i+1, curr)
		curr = append(curr, nums[i])
		dfs(i+1, curr)
	}

	dfs(0, []int{})
	return result
}

func subsetsv41(nums []int) [][]int {
	result := make([][]int, 0)

	var dfs func(i int, curr []int)
	dfs = func(i int, curr []int) {
		if i >= len(nums) {
			cp := make([]int, len(curr))
			copy(cp, curr)
			result = append(result, cp)
			return
		}

		dfs(i+1, curr)
		curr = append(curr, nums[i])
		dfs(i+1, curr)
	}
	dfs(0, []int{})
	return result
}

func subsetsv40(nums []int) [][]int {
	result := make([][]int, 0)

	var dfs func(i int, curr []int)
	dfs = func(i int, curr []int) {
		if i >= len(nums) {
			cp := make([]int, len(curr))
			copy(cp, curr)
			result = append(result, cp)
			return
		}

		dfs(i+1, curr)
		curr = append(curr, nums[i])
		dfs(i+1, curr)
	}
	dfs(0, []int{})
	return result
}

func subsetsv39(nums []int) [][]int {
	result := make([][]int, 0)

	var dfs func(i int, curr []int)
	dfs = func(i int, curr []int) {
		if i >= len(nums) {
			cp := make([]int, len(curr))
			copy(cp, curr)
			result = append(result, cp)
			return
		}

		dfs(i+1, curr)
		curr = append(curr, nums[i])
		dfs(i+1, curr)
	}
	dfs(0, []int{})
	return result
}

func subsetsv38(nums []int) [][]int {
	result := make([][]int, 0)

	var dfs func(i int, curr []int)
	dfs = func(i int, curr []int) {
		if i == len(nums) {
			cp := make([]int, len(curr))
			copy(cp, curr)
			result = append(result, cp)
			return
		}

		dfs(i+1, curr)
		curr = append(curr, nums[i])
		dfs(i+1, curr)
	}
	dfs(0, []int{})
	return result
}

func subsetsv37(nums []int) [][]int {
	result := make([][]int, 0)
	var dfs func(i int, curr []int)
	dfs = func(i int, curr []int) {
		if i == len(nums) {
			cp := make([]int, len(curr))
			copy(cp, curr)
			result = append(result, cp)
			return
		}
		dfs(i+1, curr)
		curr = append(curr, nums[i])
		dfs(i+1, curr)
	}
	dfs(0, []int{})
	return result
}

func subsetsv36(nums []int) [][]int {
	result := make([][]int, 0)
	var dfs func(i int, curr []int)
	dfs = func(i int, curr []int) {
		if i >= len(nums) {
			cp := make([]int, len(curr))
			copy(cp, curr)
			result = append(result, cp)
			return
		}
		dfs(i+1, curr)
		curr = append(curr, nums[i])
		dfs(i+1, curr)
	}
	dfs(0, []int{})
	return result
}

func subsetsv35(nums []int) [][]int {
	result := make([][]int, 0)
	var dfs func(i int, curr []int)
	dfs = func(i int, curr []int) {
		if i >= len(nums) {
			cp := make([]int, len(curr))
			copy(cp, curr)
			result = append(result, cp)
			return
		}

		dfs(i+1, curr)
		curr = append(curr, nums[i])
		dfs(i+1, curr)
	}
	dfs(0, []int{})
	return result
}

func subsetsv34(nums []int) [][]int {
	result := make([][]int, 0)
	var dfs func(i int, curr []int)
	dfs = func(i int, curr []int) {
		if i >= len(nums) {
			cp := make([]int, len(curr))
			copy(cp, curr)
			result = append(result, cp)
			return
		}

		dfs(i+1, curr)
		curr = append(curr, nums[i])
		dfs(i+1, curr)
	}

	dfs(0, []int{})
	return result
}

func subsetsv33(nums []int) [][]int {
	result := make([][]int, 0)
	var dfs func(i int, curr []int)
	dfs = func(i int, curr []int) {
		if i >= len(nums) {
			cp := make([]int, len(curr))
			copy(cp, curr)
			result = append(result, cp)
			return
		}

		dfs(i+1, curr)
		curr = append(curr, nums[i])
		dfs(i+1, curr)
	}
	dfs(0, []int{})
	return result
}

func subsetsv32(nums []int) [][]int {
	result := make([][]int, 0)
	var dfs func(i int, curr []int)
	dfs = func(i int, curr []int) {
		if i >= len(nums) {
			cp := make([]int, len(curr))
			copy(cp, curr)
			result = append(result, cp)
			return
		}

		dfs(i+1, curr)
		curr = append(curr, nums[i])
		dfs(i+1, curr)
	}
	dfs(0, []int{})
	return result
}

func subsetsv31(nums []int) [][]int {
	result := make([][]int, 0)
	var dfs func(i int, curr []int)
	dfs = func(i int, curr []int) {
		if i >= len(nums) {
			cp := make([]int, len(curr))
			copy(cp, curr)
			result = append(result, cp)
			return
		}

		dfs(i+1, curr)
		curr = append(curr, nums[i])
		dfs(i+1, curr)
	}
	dfs(0, []int{})
	return result
}

func subsetsv30(nums []int) [][]int {
	result := make([][]int, 0)
	var dfs func(i int, curr []int)
	dfs = func(i int, curr []int) {
		if i >= len(nums) {
			cp := make([]int, len(curr))
			copy(cp, curr)
			result = append(result, cp)
			return
		}

		dfs(i+1, curr)
		curr = append(curr, nums[i])
		dfs(i+1, curr)
	}
	dfs(0, []int{})
	return result
}

func subsetsv29(nums []int) [][]int {
	result := make([][]int, 0)
	var dfs func(i int, curr []int)
	dfs = func(i int, curr []int) {
		if i >= len(nums) {
			cp := make([]int, len(curr))
			copy(cp, curr)
			result = append(result, cp)
			return
		}

		dfs(i+1, curr)
		curr = append(curr, nums[i])
		dfs(i+1, curr)
	}
	dfs(0, []int{})
	return result
}

func subsetsv28(nums []int) [][]int {
	result := make([][]int, 0)

	var dfs func(i int, curr []int)
	dfs = func(i int, curr []int) {
		if i >= len(nums) {
			cp := make([]int, len(curr))
			copy(cp, curr)
			result = append(result, cp)
			return
		}
		dfs(i+1, curr)
		curr = append(curr, nums[i])
		dfs(i+1, curr)
	}
	dfs(0, []int{})
	return result
}

func subsetsv27(nums []int) [][]int {
	result := make([][]int, 0)

	var dfs func(i int, curr []int)

	dfs = func(i int, curr []int) {
		if i >= len(nums) {
			cp := make([]int, len(curr))
			copy(cp, curr)
			result = append(result, cp)
			return
		}
		dfs(i+1, curr)
		curr = append(curr, nums[i])
		dfs(i+1, curr)
	}
	dfs(0, []int{})
	return result
}

func subsetsv26(nums []int) [][]int {
	result := make([][]int, 0)

	var dfs func(i int, curr []int)
	dfs = func(i int, curr []int) {
		if i >= len(nums) {
			cp := make([]int, len(curr))
			copy(cp, curr)
			result = append(result, cp)
			return
		}

		dfs(i+1, curr)
		curr = append(curr, nums[i])
		dfs(i+1, curr)
	}
	dfs(0, []int{})
	return result
}

func subsetsv25(nums []int) [][]int {
	result := make([][]int, 0)

	var dfs func(i int, curr []int)
	dfs = func(i int, curr []int) {
		if i >= len(nums) {
			cp := make([]int, len(curr))
			copy(cp, curr)
			result = append(result, cp)
			return
		}

		dfs(i+1, curr)
		curr = append(curr, nums[i])
		dfs(i+1, curr)
	}
	dfs(0, []int{})
	return result
}

func subsetsv24(nums []int) [][]int {
	result := make([][]int, 0)

	var dfs func(i int, curr []int)
	dfs = func(i int, curr []int) {
		if i >= len(nums) {
			cp := make([]int, len(curr))
			copy(cp, curr)
			result = append(result, curr)
			return
		}

		dfs(i+1, curr)
		curr = append(curr, nums[i])
		dfs(i+1, curr)
	}
	dfs(0, []int{})
	return result
}

func subsetsv23(nums []int) [][]int {
	result := make([][]int, 0)

	var dfs func(i int, curr []int)

	dfs = func(i int, curr []int) {
		if i >= len(nums) {
			cp := make([]int, len(curr))
			copy(cp, curr)
			result = append(result, cp)
			return
		}

		dfs(i+1, curr)
		curr = append(curr, nums[i])
		dfs(i+1, curr)
	}

	dfs(0, []int{})
	return result
}

func subsetsv22(nums []int) [][]int {
	result := make([][]int, 0)

	var dfs func(i int, curr []int)

	dfs = func(i int, curr []int) {
		if i >= len(nums) {
			cp := make([]int, len(curr))
			copy(cp, curr)
			result = append(result, cp)
			return
		}

		dfs(i+1, curr)
		curr = append(curr, nums[i])
		dfs(i+1, curr)
	}
	dfs(0, []int{})
	return result
}

func subsetsv21(nums []int) [][]int {
	result := make([][]int, 0)

	var dfs func(i int, curr []int)
	dfs = func(i int, curr []int) {
		if i >= len(nums) {
			cp := make([]int, len(curr))
			copy(cp, curr)
			result = append(result, cp)
			return
		}

		dfs(i+1, curr)
		curr = append(curr, nums[i])
		dfs(i+1, curr)
	}
	dfs(0, []int{})
	return result
}

func subsetsv20(nums []int) [][]int {
	result := make([][]int, 0)

	var dfs func(i int, curr []int)

	dfs = func(i int, curr []int) {
		if i >= len(nums) {
			cp := make([]int, len(curr))
			copy(cp, curr)
			result = append(result, cp)
			return
		}

		dfs(i+1, curr)
		curr = append(curr, nums[i])
		dfs(i+1, curr)
	}

	dfs(0, []int{})
	return result
}

func subsetsv19(nums []int) [][]int {
	result := make([][]int, 0)

	var dfs func(i int, curr []int)
	dfs = func(i int, curr []int) {
		if i >= len(nums) {
			cp := make([]int, len(curr))
			copy(cp, curr)
			result = append(result, cp)
			return
		}

		dfs(i+1, curr)
		curr = append(curr, nums[i])
		dfs(i+1, curr)
	}
	dfs(0, []int{})
	return result
}

func subsetsv18(nums []int) [][]int {
	result := make([][]int, 0)

	var dfs func(i int, curr []int)
	dfs = func(i int, curr []int) {
		if i >= len(nums) {
			cp := make([]int, len(curr))
			copy(cp, curr)
			result = append(result, cp)
			return
		}

		dfs(i+1, curr)
		curr = append(curr, nums[i])
		dfs(i+1, curr)
	}

	dfs(0, []int{})
	return result
}

func subsetsv17(nums []int) [][]int {
	result := make([][]int, 0)

	var dfs func(i int, curr []int)

	dfs = func(i int, curr []int) {
		if i >= len(nums) {
			cp := make([]int, len(curr))
			copy(cp, curr)
			result = append(result, cp)
			return
		}

		dfs(i+1, curr)

		curr = append(curr, nums[i])
		dfs(i+1, curr)
	}
	dfs(0, []int{})
	return result
}

func subsetsv16(nums []int) [][]int {
	result := make([][]int, 0)

	var dfs func(i int, curr []int)
	dfs = func(i int, curr []int) {
		if i >= len(nums) {
			cp := make([]int, len(curr))
			copy(cp, curr)
			result = append(result, cp)
			return
		}

		dfs(i+1, curr)

		curr = append(curr, nums[i])
		dfs(i+1, curr)
	}
	dfs(0, []int{})
	return result
}

func subsetsv15(nums []int) [][]int {
	result := make([][]int, 0)

	var dfs func(i int, curr []int)

	dfs = func(i int, curr []int) {
		if i >= len(nums) {
			cp := make([]int, len(curr))
			copy(cp, curr)
			result = append(result, cp)
			return
		}

		dfs(i+1, curr)

		curr = append(curr, nums[i])
		dfs(i+1, curr)
	}

	dfs(0, []int{})
	return result
}

func subsetsv14(nums []int) [][]int {
	result := make([][]int, 0)

	var dfs func(i int, curr []int)
	dfs = func(i int, curr []int) {
		if i >= len(nums) {
			cp := make([]int, len(curr))
			copy(cp, curr)
			result = append(result, cp)
			return
		}

		dfs(i+1, curr)

		curr = append(curr, nums[i])
		dfs(i+1, curr)
	}
	dfs(0, []int{})
	return result
}

func subsetsv13(nums []int) [][]int {
	result := make([][]int, 0)

	var dfs func(i int, curr []int)
	dfs = func(i int, curr []int) {
		if i >= len(nums) {
			cp := make([]int, len(curr))
			copy(cp, curr)
			result = append(result, cp)
			return
		}

		dfs(i+1, curr)

		curr = append(curr, nums[i])
		dfs(i+1, curr)
	}

	dfs(0, []int{})

	return result
}

func subsetsv12(nums []int) [][]int {
	result := make([][]int, 0)

	var dfs func(i int, curr []int)
	dfs = func(i int, curr []int) {
		if i >= len(nums) {
			cp := make([]int, len(curr))
			copy(cp, curr)
			result = append(result, cp)
			return
		}

		dfs(i+1, curr)

		curr = append(curr, nums[i])
		dfs(i+1, curr)
	}
	dfs(0, []int{})
	return result
}

func subsetsv11(nums []int) [][]int {
	result := make([][]int, 0)

	var dfs func(i int, curr []int)
	dfs = func(i int, curr []int) {
		if i >= len(nums) {
			cp := make([]int, len(curr))
			copy(cp, curr)
			result = append(result, curr)
			return
		}

		curr = append(curr, nums[i])
		dfs(i+1, curr)

		curr = curr[:len(curr)-1]
		dfs(i+1, curr)
	}
	dfs(0, []int{})
	return result
}

func subsetsv10(nums []int) [][]int {
	result := make([][]int, 0)

	var dfs func(i int, curr []int)
	dfs = func(i int, curr []int) {
		if i >= len(nums) {
			cp := make([]int, len(curr))
			copy(cp, curr)
			result = append(result, cp)
			return
		}

		curr = append(curr, nums[i])
		dfs(i+1, curr)

		curr = curr[:len(curr)-1]
		dfs(i+1, curr)
	}
	dfs(0, []int{})
	return result
}

func subsetsv9(nums []int) [][]int {
	result := make([][]int, 0)

	var dfs func(i int, curr []int)

	dfs = func(i int, curr []int) {
		if i >= len(nums) {
			cp := make([]int, len(curr))
			copy(cp, curr)
			result = append(result, cp)
			return
		}

		curr = append(curr, nums[i])
		dfs(i+1, curr)

		curr = curr[:len(curr)-1]
		dfs(i+1, curr)
	}

	dfs(0, []int{})
	return result
}

func subsetsv8(nums []int) [][]int {
	result := make([][]int, 0)

	var dfs func(i int, curr []int)

	dfs = func(i int, curr []int) {
		if i >= len(nums) {
			cp := make([]int, len(curr))
			copy(cp, curr)
			result = append(result, cp)
			return
		}

		curr = append(curr, nums[i])
		dfs(i+1, curr)

		curr = curr[:len(curr)-1]
		dfs(i+1, curr)
	}

	dfs(0, []int{})
	return result
}

func subsetsv7(nums []int) [][]int {
	result := make([][]int, 0)

	var dfs func(i int, curr []int)
	dfs = func(i int, curr []int) {
		if i >= len(nums) {
			cp := make([]int, len(curr))
			copy(cp, curr)
			result = append(result, cp)
			return
		}

		curr = append(curr, nums[i])
		dfs(i+1, curr)

		curr = curr[:len(curr)-1]
		dfs(i+1, curr)
	}

	dfs(0, []int{})
	return result
}

func subsetsv6(nums []int) [][]int {
	result := make([][]int, 0)

	var dfs func(i int, curr []int)

	dfs = func(i int, curr []int) {
		if i >= len(nums) {
			cp := make([]int, len(curr))
			copy(cp, curr)
			result = append(result, curr)
			return
		}

		curr = append(curr, nums[i])
		dfs(i+1, curr)

		curr = curr[:len(curr)-1]
		dfs(i+1, curr)
	}

	dfs(0, []int{})
	return result
}

func subsetsv5(nums []int) [][]int {
	result := make([][]int, 0)

	var dfs func(i int, curr []int)
	dfs = func(i int, curr []int) {
		if i >= len(nums) {
			cp := make([]int, len(curr))
			copy(cp, curr)
			result = append(result, cp)
			return
		}

		curr = append(curr, nums[i])
		dfs(i+1, curr)

		curr = curr[:len(curr)-1]
		dfs(i+1, curr)

	}
	dfs(0, []int{})
	return result
}

func subsetsv4(nums []int) [][]int {
	result := make([][]int, 0)

	var dfs func(i int, curr []int)
	dfs = func(i int, curr []int) {
		if i >= len(nums) {
			cp := make([]int, len(curr))
			copy(cp, curr)
			result = append(result, cp)
			return
		}

		curr = append(curr, nums[i])
		dfs(i+1, curr)

		curr = curr[:len(curr)-1]
		dfs(i+1, curr)
	}

	dfs(0, []int{})

	return result
}

func subsetsv3(nums []int) [][]int {
	result := make([][]int, 0)

	var dfs func(i int, curr []int)
	dfs = func(i int, curr []int) {
		if i >= len(nums) {
			elem := make([]int, len(curr))
			copy(elem, curr)
			result = append(result, elem)
			return
		}

		curr = append(curr, nums[i])
		dfs(i+1, curr)

		curr = curr[:len(curr)-1]
		dfs(i+1, curr)
	}

	dfs(0, []int{})
	return result
}

func subsetsv2(nums []int) [][]int {
	result := make([][]int, 0)

	var dfs func(i int, curr []int)
	dfs = func(i int, curr []int) {
		if i >= len(nums) {
			v := make([]int, len(curr))
			copy(v, curr)
			result = append(result, v)
			return
		}

		// picked it up
		curr = append(curr, nums[i])
		dfs(i+1, curr)

		// or do nothing
		curr = curr[:len(curr)-1]
		dfs(i+1, curr)

	}

	dfs(0, []int{})
	return result
}

func subsetsv1(nums []int) [][]int {
	result := make([][]int, 0)

	curr := make([]int, 0)
	var dfs func(i int)

	dfs = func(i int) {
		if i >= len(nums) {
			tmp := make([]int, len(curr))
			copy(tmp, curr)
			result = append(result, tmp)
			return
		}

		curr = append(curr, nums[i])
		dfs(i + 1)

		curr = curr[:len(curr)-1]
		dfs(i + 1)
	}

	dfs(0)
	return result
}
