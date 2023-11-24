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
