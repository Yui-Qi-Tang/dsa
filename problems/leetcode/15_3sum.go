package b75

import (
	"sort"
)

/*
15. 3Sum

Given an integer array nums, return all the triplets
[nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k,
and nums[i] + nums[j] + nums[k] == 0.

Notice that the solution set must not contain duplicate triplets.

Example 1:

Input: nums = [-1,0,1,2,-1,-4] => -4, -1, -1, 0, 1,2
Output: [[-1,-1,2],[-1,0,1]]
Explanation:

	nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0.
	nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0.
	nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0.

The distinct triplets are [-1,0,1] and [-1,-1,2].
Notice that the order of the output and the order of the triplets does not matter.

Example 2:

Input: nums = [0,1,1]
Output: []
Explanation: The only possible triplet does not sum up to 0.
Example 3:

Input: nums = [0,0,0]
Output: [[0,0,0]]
Explanation: The only possible triplet sums up to 0.

Constraints:

3 <= nums.length <= 3000
-105 <= nums[i] <= 105
*/

func ThreeSumv12(nums []int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0)
	n := len(nums)

	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		j := i + 1
		k := n - 1

		for j < k {
			if nums[i]+nums[j]+nums[k] == 0 {
				result = append(result, []int{nums[i], nums[j], nums[k]})

				k--
				for ; j < k && nums[k] == nums[k+1]; k++ {
				}
			} else if nums[i]+nums[j]+nums[k] > 0 {
				k--
			} else {
				j++
			}
		}
	}

	return result
}

func ThreeSumv11(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	result := make([][]int, 0)

	for i := 0; i < n-2; i++ {
		// because the nums are sorted, so we can skip the duplicated value
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		j := i + 1 // 2nd
		k := n - 1

		// nums[i] + nums[j...k] == 0
		for j < k {
			if nums[i]+nums[j]+nums[k] == 0 {
				result = append(result, []int{nums[i], nums[j], nums[k]})
				k--
				// reduce the duplicated from k
				for ; j < k && nums[k] == nums[k+1]; k-- {
				}
			} else if nums[i]+nums[k]+nums[k] > 0 {
				k--
			} else {
				j++
			}
		}
	}

	return result
}

func ThreeSumv10(nums []int) [][]int {

	result := make([][]int, 0)
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		j := i + 1
		k := len(nums) - 1

		for j < k {
			if nums[i]+nums[j]+nums[k] == 0 {
				result = append(result, []int{nums[i], nums[j], nums[k]})
				k--

				for ; j < k && nums[k] == nums[k+1]; k-- {
				}
			} else if nums[i]+nums[j]+nums[k] > 0 {
				k--
			} else {
				j++
			}
		}
	}

	return result
}

func ThreeSumv9(nums []int) [][]int {
	result := make([][]int, 0)
	for i := 0; i < len(nums)-2; i++ {

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		j := i + 1
		k := len(nums) - 1

		for j < k {
			if nums[i]+nums[j]+nums[k] == 0 {
				result = append(result, []int{nums[i], nums[j], nums[k]})
				k--

				for ; j < k && nums[k] == nums[k+1]; k-- {
					// remove duplicated value from k
				}
			} else if nums[i]+nums[j]+nums[k] > 0 {
				k--
			} else {
				j++
			}
		}

	}
	return result
}

func ThreeSumv8(nums []int) [][]int {
	result := make([][]int, 0)
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		j := i + 1
		k := len(nums) - 1

		for j < k {
			if nums[i]+nums[j]+nums[k] == 0 {
				result = append(result, []int{nums[i], nums[j], nums[k]})
				k--

				for ; j < k && nums[k] == nums[k+1]; k-- {
				}
			} else if nums[i]+nums[j]+nums[k] > 0 {
				k--
			} else {
				j++
			}
		}
	}
	return result
}

func ThreeSumv7(nums []int) [][]int {
	n := len(nums)
	result := make([][]int, 0)
	for i := 0; i < n-2; i++ {

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		j := i + 1
		k := n - 1

		for j < k {
			if nums[i]+nums[j]+nums[k] == 0 {
				result = append(result, []int{nums[i], nums[j], nums[k]})
				k--

				for ; j < k && nums[k] == nums[k+1]; k-- {
				}
			} else if nums[i]+nums[j]+nums[k] > 0 {
				k--
			} else {
				j++
			}
		}

	}
	return result
}

func ThreeSumv6(nums []int) [][]int {

	n := len(nums)

	result := make([][]int, 0)
	// -2: j,k
	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		j := i + 1
		k := n - 1

		for j < k {
			if nums[i]+nums[j]+nums[k] == 0 {
				result = append(result, []int{nums[i], nums[j], nums[k]})
				k--
				for ; j < k && nums[k] == nums[k-1]; k-- {
				}
			} else if nums[i]+nums[j]+nums[k] > 0 {
				// increase value
				j++
			} else {
				// reduce the value
				k--
			}

		}
	}
	return result
}

func ThreeSumv5(nums []int) [][]int {

	n := len(nums)

	result := make([][]int, 0)
	for i := 0; i < n; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			// skip duplicated value
			continue
		}

		j := i + 1
		k := n - 1

		for j < k {
			if nums[i]+nums[j]+nums[k] == 0 {
				result = append(result, []int{nums[i], nums[j], nums[k]})

				// skip duplicated value
				for ; j < k && nums[k] == nums[k-1]; k-- {

				}
			} else if nums[i]+nums[j]+nums[k] > 0 {
				// reduce the value
				k--
			} else {
				// increase the value
				j++
			}
		}
	}

	return result

}

func ThreeSumv4(nums []int) [][]int {

	// extend the 2 sum in sorted nums;
	// fixed the first i + sum(j, k)
	// one pointer points to j another one points k

	n := len(nums)
	// key: the nums should be sorted
	sort.Ints(nums)

	result := make([][]int, 0)
	for i := 0; i < n-2; i++ {
		// remove duplicated value from left
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		j := i + 1
		k := n - 1

		for j < k {
			if nums[i]+nums[j]+nums[k] == 0 {
				result = append(result, []int{nums[i], nums[j], nums[k]})
				k--

				// remove duplicated value from right until j+1
				for ; j < k && nums[k] == nums[k+1]; k-- {
				}
			} else if nums[i]+nums[j]+nums[k] > 0 {
				// reduce the sum
				k--
			} else {
				// increase the sum
				j++
			}
		}
	}

	return result
}

func ThreeSumv3(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)

	result := make([][]int, 0)
	for i := 0; i < n-2; i++ {
		// skip the duplicated value
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		j := i + 1
		k := n - 1

		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum == 0 {
				result = append(result, []int{nums[i], nums[j], nums[k]})
				k--

				// skip the duplicated value
				for ; j < k && nums[k] == nums[k+1]; k-- {
				}
			} else if sum > 0 {
				// reduce the value
				k--
			} else {
				// increase the value
				j++
			}
		}
	}
	return result
}

func ThreeSumv2(nums []int) [][]int {

	n := len(nums)
	sort.Ints(nums)

	result := make([][]int, 0)
	for i := 0; i < n-2; i++ {

		if i > 0 && nums[i] == nums[i-1] {
			// skip duplicated value from left
			continue
		}

		j := i + 1
		k := n - 1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum == 0 {
				result = append(result, []int{nums[i], nums[j], nums[k]})
				k--

				// skip duplicated value from right
				for ; j < k && nums[k] == nums[k+1]; k-- {
				}
			} else if sum > 0 {
				// reduce the sum value
				k--
			} else {
				// increase the sum value
				j++
			}
		}

	}

	return result

}

func ThreeSumv1(nums []int) [][]int {
	n := len(nums)
	if n < 3 {
		return nil
	}

	result := make([][]int, 0)
	keys := make(map[int]int)
	for i, num := range nums {
		keys[num] = i
	}

	processedIdx := make(map[int]bool)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			diff := 0 - (nums[i] + nums[j])
			if k, exist := keys[diff]; exist {
				if k != i && k != j {
					if processedIdx[i] && processedIdx[j] && processedIdx[k] {
						break
					}
					processedIdx[i] = true
					processedIdx[j] = true
					processedIdx[k] = true

					result = append(result, []int{nums[i], nums[j], nums[k]})
					break
				}
			}

		}

	}

	return result
}
