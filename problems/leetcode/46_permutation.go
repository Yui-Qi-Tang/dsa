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

func permutev29(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{{nums[0]}}
	}

	n := len(nums)
	result := make([][]int, 0)
	for n > 0 {
		n--
		v := nums[0]
		nums = nums[1:]
		perms := permutev29(nums)
		for i := range perms {
			perms[i] = append(perms[i], v)
		}
		nums = append(nums, v)
		result = append(result, perms...)
	}
	return result
}

func permutev28(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{{nums[0]}}
	}

	result := make([][]int, 0)
	n := len(nums)
	for n > 0 {
		n--
		v := nums[0]
		nums = nums[1:]
		perms := permutev28(nums)
		for i := range perms {
			perms[i] = append(perms[i], v)
		}
		result = append(result, perms...)
		nums = append(nums, v)
	}
	return result
}

func permutev27(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{{nums[0]}}
	}

	result := make([][]int, 0)
	n := len(nums)
	for n > 0 {
		n--
		v := nums[0]
		nums = nums[1:]
		perms := permutev27(nums)
		for i := range perms {
			perms[i] = append(perms[i], v)
		}
		nums = append(nums, v)
		result = append(result, perms...)
	}

	return result
}

func permutev26(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{{nums[0]}}
	}

	result := make([][]int, 0)
	n := len(nums)

	for n > 0 {
		n--
		v := nums[0]
		nums = nums[1:]
		perms := permutev26(nums)
		for i := range perms {
			perms[i] = append(perms[i], v)
		}
		result = append(result, perms...)
		nums = append(nums, v)
	}

	return result
}

func permutev25(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{{nums[0]}}
	}

	n := len(nums)
	result := make([][]int, 0)
	for n > 0 {
		n--
		v := nums[0]
		nums = nums[1:]
		perms := permutev25(nums)
		for i := range perms {
			perms[i] = append(perms[i], v)
		}
		result = append(result, perms...)
		nums = append(nums, v)
	}
	return result
}

func permutev24(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{{nums[0]}}
	}
	result := make([][]int, 0)
	n := len(nums)
	for n > 0 {
		n--
		v := nums[0]
		nums = nums[1:]
		perms := permutev24(nums)
		for i := range perms {
			perms[i] = append(perms[i], v)
		}
		result = append(result, perms...)
		nums = append(nums, v)
	}
	return result
}

func permutev23(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{{nums[0]}}
	}

	result := make([][]int, 0)
	n := len(nums)
	for n > 0 {
		n--
		v := nums[0]
		nums = nums[1:]
		perms := permutev23(nums)
		for i := range perms {
			perms[i] = append(perms[i], v)
		}
		result = append(result, perms...)
		nums = append(nums, v)
	}

	return result
}

func permutev22(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{{nums[0]}}
	}

	result := make([][]int, 0)
	n := len(nums)

	for n > 0 {
		n--
		v := nums[0]
		nums = nums[1:]
		perms := permutev22(nums)
		for i := range perms {
			perms[i] = append(perms[i], v)
		}
		result = append(result, perms...)
		nums = append(nums, v)
	}

	return result
}

func permutev21(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{{nums[0]}}
	}

	n := len(nums)
	result := make([][]int, 0)
	for n > 0 {
		n--
		v := nums[0]
		nums = nums[1:]
		perms := permutev21(nums)
		for i := range perms {
			perms[i] = append(perms[i], v)
		}
		result = append(result, perms...)
		nums = append(nums, v)
	}

	return result
}

func permutev20(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{{nums[0]}}
	}

	result := make([][]int, 0)
	n := len(nums)

	for n > 0 {
		n--
		num := nums[0]
		nums = nums[1:]
		perms := permutev20(nums)
		for i := range perms {
			perms[i] = append(perms[i], num)
		}
		result = append(result, perms...)
		nums = append(nums, num)
	}
	return result
}

func permutev19(nums []int) [][]int {

	if len(nums) == 1 {
		return [][]int{{nums[0]}}
	}

	result := make([][]int, 0)
	n := len(nums)

	for n > 0 {
		n--
		v := nums[0]
		nums = nums[1:]
		perms := permutev19(nums)
		for i := range perms {
			perms[i] = append(perms[i], v)
		}
		result = append(result, perms...)
		nums = append(nums, v)
	}

	return result
}

func permutev18(nums []int) [][]int {

	if len(nums) == 1 {
		return [][]int{{nums[0]}}
	}

	result := make([][]int, 0)
	n := len(nums)

	for n > 0 {
		v := nums[0]
		nums = nums[1:]
		perms := permutev18(nums)
		for i := range perms {
			perms[i] = append(perms[i], v)
		}
		result = append(result, perms...)
		nums = append(nums, v)
		n--
	}

	return result
}

func permutev17(nums []int) [][]int {

	if len(nums) == 1 {
		return [][]int{{nums[0]}}
	}

	result := make([][]int, 0)
	n := len(nums)

	for n > 0 {
		v := nums[0]
		nums = nums[1:]
		perms := permutev17(nums)
		for i := range perms {
			perms[i] = append(perms[i], v)
		}
		result = append(result, perms...)
		nums = append(nums, v)
		n--
	}

	return result
}

func permutev16(nums []int) [][]int {

	if len(nums) == 1 {
		return [][]int{{nums[0]}}
	}

	result := make([][]int, 0)
	n := len(nums)

	for n > 0 {
		v := nums[0]
		nums = nums[1:]
		perms := permutev16(nums)
		for i := range perms {
			perms[i] = append(perms[i], v)
		}
		result = append(result, perms...)
		nums = append(nums, v)
		n--
	}

	return result
}

func permutev15(nums []int) [][]int {

	if len(nums) == 1 {
		return [][]int{{nums[0]}}
	}

	result := make([][]int, 0)

	n := len(nums)

	for n > 0 {
		v := nums[0]
		nums = nums[1:]
		perms := permutev15(nums)
		for i := range perms {
			perms[i] = append(perms[i], v)
		}
		result = append(result, perms...)
		nums = append(nums, v)
		n--
	}

	return result
}

func permutev14(nums []int) [][]int {

	if len(nums) == 1 {
		return [][]int{{nums[0]}}
	}

	result := make([][]int, 0)
	n := len(nums)
	for n > 0 {
		v := nums[0]
		nums = nums[1:]
		perms := permutev14(nums)
		for i := range perms {
			perms[i] = append(perms[i], v)
		}
		result = append(result, perms...)
		nums = append(nums, v)
		n--
	}

	return result
}

func permutev13(nums []int) [][]int {

	if len(nums) == 1 {
		return [][]int{{nums[0]}}
	}

	result := make([][]int, 0)
	n := len(nums)

	for n > 0 {
		v := nums[0]
		nums = nums[1:]
		perms := permutev13(nums)
		for i := range perms {
			perms[i] = append(perms[i], v)
		}
		result = append(result, perms...)
		nums = append(nums, v)
		n--
	}

	return result
}

func permutev12(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{{nums[0]}}
	}

	result := make([][]int, 0)

	n := len(nums)
	for n > 0 {

		v := nums[0]
		nums = nums[1:]
		perms := permutev12(nums)
		for i := range perms {
			perms[i] = append(perms[i], v)
		}
		result = append(result, perms...)
		nums = append(nums, v)
		n--
	}

	return result
}

func permutev11(nums []int) [][]int {

	if len(nums) == 1 {
		v := nums[0]
		return [][]int{{v}}
	}

	result := make([][]int, 0)
	n := len(nums)
	for n > 0 {

		v := nums[0]
		nums = nums[1:]
		perms := permutev11(nums)
		for i := range perms {
			perms[i] = append(perms[i], v)
		}
		result = append(result, perms...)
		nums = append(nums, v)

		n--
	}

	return result
}

func permutev10(nums []int) [][]int {
	if len(nums) == 1 {
		v := nums[0]
		return [][]int{{v}}
	}

	n := len(nums)
	result := make([][]int, 0)

	for n > 0 {
		v := nums[0]
		nums = nums[1:]
		perms := permutev10(nums)
		for i := range perms {
			perms[i] = append(perms[i], v)
		}
		result = append(result, perms...)
		nums = append(nums, v)
		n--
	}

	return result
}

func permutev9(nums []int) [][]int {

	n := len(nums)
	if n == 1 {
		v := nums[0]
		return [][]int{{v}}
	}

	result := make([][]int, 0)

	for n > 0 {

		v := nums[0]
		nums = nums[1:]
		perms := permutev9(nums)
		for i := range perms {
			perms[i] = append(perms[i], v)
		}
		result = append(result, perms...)
		nums = append(nums, v)
		n--
	}

	return result
}

func permutev8(nums []int) [][]int {

	if len(nums) == 1 {
		cp := make([]int, 1)
		copy(cp, nums)
		return [][]int{cp}
	}

	result := make([][]int, 0)
	n := len(nums)
	for n > 0 {
		num := nums[0]
		nums = nums[1:]
		permu := permutev8(nums)
		for i := range permu {
			permu[i] = append(permu[i], num)
		}
		result = append(result, permu...)
		nums = append(nums, num)
		n--
	}
	return result
}

func permutev7(nums []int) [][]int {
	n := len(nums)
	if n == 1 {
		cp := make([]int, n)
		copy(cp, nums)
		return [][]int{cp}
	}

	result := make([][]int, 0)

	for n > 0 {
		v := nums[0]
		nums = nums[1:]
		perm := permutev7(nums)
		for i := range perm {
			perm[i] = append(perm[i], v)
		}
		nums = append(nums, v)
		result = append(result, perm...)
		n--
	}

	return result
}

func permutev6(nums []int) [][]int {

	if len(nums) == 1 {
		cp := make([]int, len(nums))
		copy(cp, nums)
		return [][]int{cp}
	}

	result := make([][]int, 0)

	n := len(nums)
	for n > 0 {
		v := nums[0]
		nums = nums[1:]
		perms := permutev6(nums)

		for i := range perms {
			perms[i] = append(perms[i], v)
		}
		result = append(result, perms...)
		nums = append(nums, v)
		n--
	}

	return result
}

func permutev5(nums []int) [][]int {

	if len(nums) == 1 {
		cp := make([]int, len(nums))
		copy(cp, nums)
		return [][]int{cp}
	}

	result := make([][]int, 0)
	n := len(nums)
	for n > 0 {
		v := nums[0]
		nums = nums[1:]
		perms := permutev5(nums)
		for i := range perms {
			perms[i] = append(perms[i], v)
		}
		result = append(result, perms...)
		nums = append(nums, v)
		n--
	}

	return result
}

func permutev4(nums []int) [][]int {
	if len(nums) == 1 {
		cp := make([]int, len(nums))
		copy(cp, nums)
		return [][]int{cp}
	}

	result := make([][]int, 0)
	n := len(nums)

	for n > 0 {
		v := nums[0]
		nums = nums[1:]
		perm := permutev4(nums)
		for i := range perm {
			perm[i] = append(perm[i], v)
		}
		result = append(result, perm...)
		nums = append(nums, v)
		n--
	}

	return result
}

func permutev3(nums []int) [][]int {
	n := len(nums)
	cp := make([]int, n)
	copy(cp, nums)
	if n == 1 {
		return [][]int{cp}
	}

	result := make([][]int, 0)
	for n > 0 {
		v := nums[0]
		nums = nums[1:]
		perm := permutev3(nums)
		for i := range perm {
			perm[i] = append(perm[i], v)
		}
		result = append(result, perm...)
		nums = append(nums, v)
		n--
	}

	return result
}

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
