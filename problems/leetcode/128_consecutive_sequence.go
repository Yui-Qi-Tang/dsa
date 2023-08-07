package b75

import "fmt"

/*
128. Longest Consecutive Sequence
Medium

Given an unsorted array of integers nums, return the length of the longest consecutive elements sequence.

You must write an algorithm that runs in O(n) time.



Example 1:

Input: nums = [100,4,200,1,3,2]
Output: 4
Explanation: The longest consecutive elements sequence is [1, 2, 3, 4]. Therefore its length is 4.
Example 2:

Input: nums = [0,3,7,2,5,8,4,6,0,1]
Output: 9


Constraints:

0 <= nums.length <= 105
-109 <= nums[i] <= 109
*/

func longestConsecutivev25(nums []int) int {
	m := make(map[int]bool, len(nums))
	for _, num := range nums {
		m[num] = true
	}

	var longest int = 0
	for _, num := range nums {
		if !m[num-1] {
			length := 0
			for m[num+length] {
				length++
			}
			longest = max(longest, length)
		}
	}

	return longest
}

func longestConsecutivev24(nums []int) int {
	m := make(map[int]bool, len(nums))
	longest := 0

	for _, num := range nums {
		m[num] = true
	}

	for _, num := range nums {
		if !m[num-1] {
			length := 0
			for m[num+length] {
				length++
			}
			longest = max(longest, length)
		}
	}

	return longest
}

func longestConsecutivev23(nums []int) int {
	longest := 0
	m := make(map[int]bool, len(nums))
	for _, num := range nums {
		m[num] = true
	}

	for _, num := range nums {
		length := 0

		if !m[num-1] {
			for m[num+length] {
				length++
			}
		}

		longest = max(longest, length)
	}

	return longest
}

func longestConsecutivev22(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	result := 0
	m := make(map[int]bool)

	// need to save all of elements for computing
	for _, num := range nums {
		m[num] = true
	}

	for _, num := range nums {
		if !m[num-1] {
			longest := 0
			for m[num+longest] {
				longest++
			}
			m[num] = true
			result = max(result, longest)
		}
	}

	return result
}

func longestConsecutivev21(nums []int) int {
	m := make(map[int]bool, len(nums))

	for _, num := range nums {
		m[num] = true
	}

	var longest int = 0

	for _, num := range nums {
		if !m[num-1] {
			length := 0
			for m[num+length] {
				length++
			}

			if length > longest {
				longest = length
			}
		}
	}

	return longest
}

func longestConsecutivev20(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	var result int = 0
	m := make(map[int]bool)

	for _, num := range nums {
		m[num] = true
	}

	for _, num := range nums {
		if !m[num-1] {
			length := 0
			for m[num+length] {
				length++
			}
			result = max(result, length)
		}
	}

	return result
}

func longestConsecutivev19(nums []int) int {
	m := make(map[int]bool)

	for _, num := range nums {
		m[num] = true
	}

	var longest int = 1

	for _, num := range nums {
		if !m[num-1] {
			length := 0
			for m[num+length] {
				length++
			}
			longest = max(longest, length)
		}
	}

	return longest
}

func longestConsecutivev18(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return n
	}

	m := make(map[int]bool, n)
	for _, num := range nums {
		m[num] = true
	}

	var longest int = 1

	for _, num := range nums {
		if !m[num-1] {
			length := 0
			for m[num+length] {
				length++
			}
			longest = max(longest, length)
		}
	}

	return longest
}

func longestConsecutivev17(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return n
	}

	dp := make(map[int]bool)

	for _, num := range nums {
		dp[num] = true
	}

	var longest int = 1

	for _, num := range nums {
		if !dp[num-1] {
			length := 0

			for dp[num+length] {
				length++
			}

			longest = max(longest, length)
		}
	}

	return longest
}

func longestConsecutivev16(nums []int) int {
	dp := make(map[int]bool)

	for _, num := range nums {
		dp[num] = true
	}

	var longest int = 0

	for _, num := range nums {
		if !dp[num-1] {
			length := 0
			for dp[num+length] {
				length++
			}
			longest = max(longest, length)
		}
	}

	return longest

}

func longestConsecutivev15(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return n
	}

	dp := make(map[int]bool, n)
	for _, num := range nums {
		dp[num] = true
	}

	var longest int = 0

	for _, num := range nums {
		if !dp[num-1] {
			length := 0
			for dp[num+length] {
				length++
			}
			longest = max(length, longest)
		}

	}

	return longest

}

func longestConsecutivev14(nums []int) int {
	n := len(nums)
	if n <= 0 {
		return n
	}

	m := make(map[int]bool)

	for _, num := range nums {
		m[num] = true
	}

	var longest int = 0

	for _, num := range nums {
		if !m[num-1] {
			length := 0
			for m[num+length] {
				length++
			}
			longest = max(longest, length)
		}
	}

	return longest
}

func longestConsecutivev13(nums []int) int {
	m := make(map[int]bool, len(nums))

	for _, num := range nums {
		m[num] = true
	}

	var longest int = 0

	for _, num := range nums {
		if !m[num-1] {
			length := 0

			for m[num+length] {
				length++
			}

			longest = max(longest, length)
		}
	}

	return longest

}

func longestConsecutivev12(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	m := make(map[int]bool, n)

	for _, num := range nums {
		m[num] = true
	}

	longest := 0
	for _, num := range nums {
		if !m[num-1] {
			length := 0
			for m[num+length] {
				length++
			}

			longest = max(longest, length)
		}
	}

	return longest

}

func longestConsecutivev11(nums []int) int {
	m := make(map[int]bool, len(nums))

	for _, num := range nums {
		m[num] = true
	}

	longest := 0
	for _, num := range nums {
		if !m[num-1] {
			length := 0
			for m[num+length] {
				length++
			}
			longest = max(longest, length)
		}
	}

	return longest
}

func longestConsecutivev10(nums []int) int {
	m := make(map[int]bool, len(nums))

	// 1,2,100,2,3 -> 1,2,100,3 | 3,2,1,100 | 2,3,1,100
	for _, num := range nums {
		m[num] = true
	}

	longest := 0
	for _, num := range nums {
		if !m[num-1] {
			length := 0
			for m[num+length] {
				length++
			}
			longest = max(length, longest)
		}
	}

	return longest
}

func longestConsecutivev9(nums []int) int {
	m := make(map[int]bool, len(nums))

	for _, num := range nums {
		m[num] = true
	}

	longest := 0
	for _, num := range nums {
		if !m[num-1] {
			length := 0
			for m[num+length] {
				length++
			}

			longest = max(longest, length)
		}
	}

	return longest
}

func longestConsecutivev8(nums []int) int {
	m := make(map[int]bool, len(nums))

	for _, num := range nums {
		m[num] = true
	}

	longest := 0

	for _, num := range nums {
		if !m[num-1] {
			length := 0
			for m[num+length] {
				length++
			}
			longest = max(length, longest)
		}
	}

	return longest
}

func longestConsecutivev7(nums []int) int {
	set := make(map[int]bool, len(nums))

	for _, num := range nums {
		set[num] = true
	}

	longest := 0

	for _, num := range nums {
		if !set[num-1] {
			length := 0
			for set[num+length] {
				length++
			}

			longest = max(length, longest)
		}
	}

	return longest
}

func longestConsecutivev6(nums []int) int {

	m := make(map[int]bool, len(nums))
	for _, num := range nums {
		m[num] = true
	}

	longest := 0

	for _, num := range nums {
		if !m[num-1] {
			length := 0
			for m[num+length] {
				length++
			}
			longest = max(longest, length)
		}
	}

	return longest
}

func longestConsecutivev5(nums []int) int {
	set := make(map[int]bool, len(nums))

	for _, num := range nums {
		set[num] = true
	}

	var longest int = 0
	for _, num := range nums {
		if !set[num-1] {
			length := 0
			for set[num+length] {
				length++
			}
			longest = max(longest, length)
		}
	}

	return longest

}

func longestConsecutivev4(nums []int) int {
	m := make(map[int]bool)

	for _, num := range nums {
		m[num] = true
	}

	var longest int = 0

	for k := range m {
		if !m[k-1] {
			length := 0
			for {
				if m[k+length] {
					length++
				} else {
					break
				}
			}
			longest = max(longest, length)
		}
	}

	return longest
}

func longestConsecutivev3(nums []int) int {
	m := make(map[int]int) // value:cnt

	for _, n := range nums {
		m[n] = 1
	}

	longest := 0
	for _, num := range nums {
		if _, exist := m[num-1]; !exist {
			length := 0
			for {
				if _, ok := m[num+length]; ok {
					length++
				} else {
					break
				}
			}
			longest = max(longest, length)
		}
	}
	fmt.Println(m)
	return longest

}

func longestConsecutivev2(nums []int) int {
	m := make(map[int]bool)

	for _, n := range nums {
		m[n] = true
	}

	longest := 0
	for _, num := range nums {
		if _, exist := m[num-1]; !exist {
			length := 0
			for {
				if _, ok := m[num+length]; ok {
					length++
				} else {
					break
				}
			}
			longest = max(longest, length)
		}
	}

	return longest
}

func longestConsecutivev1(nums []int) int {
	m := make(map[int]bool) // value:cnt

	for _, n := range nums {
		m[n] = true
	}

	longest := 0
	for _, num := range nums {
		length := 1
		for {
			if _, ok := m[num-1]; ok {
				length++
				num--
			} else {
				break
			}
		}
		longest = max(longest, length)

	}

	/*
		m := make(map[int]int) // value:cnt

		result := 0

		for _, num := range nums {
			if _, exist := m[num]; !exist {
				m[num] = 1
			} else {
				//m[num] = result
				continue
			}

			if c, exist := m[num+1]; exist {
				m[num] += c
			}

			if c, exist := m[num-1]; exist {
				m[num] += c
			}

			if m[num] > result {
				result = m[num]
			}

			if _, exist := m[num+1]; exist {
				m[num+1] = result
			}

			if _, exist := m[num-1]; exist {
				m[num-1] = result
			}

		}
	*/

	return longest
}
