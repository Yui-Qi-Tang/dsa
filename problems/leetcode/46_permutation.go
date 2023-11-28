package b75

/*
46. Permutations
Given an array nums of distinct integers,
return all the possible permutations.
You can return the answer in any order.



Example 1:

Input: nums = [1,2,3]
Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
Example 2:

Input: nums = [0,1]
Output: [[0,1],[1,0]]
Example 3:

Input: nums = [1]
Output: [[1]]


Constraints:

1 <= nums.length <= 6
-10 <= nums[i] <= 10
All the integers of nums are unique.

HINT: pop itself and make permutation on another elements
*/

// permutev2 is another way to recursive the subproblems to solve the problem
// this is more efficient than v1
func permutev2(nums []int) [][]int {
	n := len(nums)
	ans := make([][]int, 0)
	curr := make([]int, 0, n)
	vis := make(map[int]int)
	var backtrack func(idx int)
	backtrack = func(idx int) {
		if len(curr) == n {
			ans = append(ans, append([]int{}, curr...))
		}
		for i := 0; i < n; i++ {
			if vis[i] == 0 {
				vis[i]++
				curr = append(curr, nums[i])
				backtrack(i + 1)
				curr = curr[:len(curr)-1]
				vis[i]--
			}
		}
	}
	backtrack(0)
	return ans
}

// permutev1 is the recursive way to solve the problems
func permutev1(nums []int) [][]int {
	cp := make([]int, len(nums))
	copy(cp, nums)

	if len(nums) == 1 {
		return [][]int{cp}
	}

	result := make([][]int, 0)

	n := len(cp)
	for i := 0; i < n; i++ {
		v := cp[0]
		cp = cp[1:]
		restPermute := permutev1(cp)
		for j := range restPermute {
			restPermute[j] = append(restPermute[j], v)
		}
		result = append(result, restPermute...)
		cp = append(cp, v)
	}

	return result
}
