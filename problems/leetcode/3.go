package b75

/*
3. Longest Substring Without Repeating Characters
Medium
Given a string s, find the length of
the "longest substring" without "repeating characters".



Example 1:

Input: s = "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.
Example 2:

Input: s = "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.
Example 3:

Input: s = "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.


Constraints:

0 <= s.length <= 5 * 104
s consists of English letters, digits, symbols and spaces.
*/

func lengthOfLongestSubstringv17(s string) int {
	dup := map[byte]bool{}
	result := 0
	l := 0
	for r := 0; r < len(s); r++ {
		for dup[s[r]] {
			delete(dup, s[l])
			l++
		}

		dup[s[r]] = true
		result = max(result, r-l+1)
	}

	return result
}

func lengthOfLongestSubstringv16(s string) int {
	result := 0

	dup := make(map[byte]bool)

	l := 0
	for r := range s {
		for dup[s[r]] {
			delete(dup, s[l])
			l++
		}

		dup[s[r]] = true
		result = max(result, r-l+1)
	}

	return result
}

func lengthOfLongestSubstringv15(s string) int {
	dup := map[byte]bool{}
	result := 0

	l := 0
	for r := range s {
		for dup[s[r]] {
			delete(dup, s[l])
			l++
		}

		dup[s[r]] = true
		result = max(result, r-l+1)

	}

	return result
}

func lengthOfLongestSubstringv14(s string) int {
	dup := make(map[byte]bool)

	result := 0

	l := 0
	for r := 0; r < len(s); r++ {
		for dup[s[r]] {
			delete(dup, s[l])
			l++
		}

		dup[s[r]] = true
		result = max(result, r-l+1)
	}

	return result
}

func lengthOfLongestSubstringv13(s string) int {
	result := 0

	duplicated := make(map[byte]bool)

	l := 0
	for r := range s {
		for duplicated[s[r]] {
			delete(duplicated, s[l])
			l++
		}

		duplicated[s[r]] = true
		result = max(result, r-l+1)
	}

	return result
}

func lengthOfLongestSubstringv12(s string) int {
	duplicated := map[byte]bool{}
	result := 0

	l := 0
	for r := 0; r < len(s); r++ {
		for duplicated[s[r]] {
			delete(duplicated, s[l])
			l++
		}

		duplicated[s[r]] = true
		result = max(result, r-l+1)
	}

	return result
}

func lengthOfLongestSubstringv11(s string) int {
	exists := make(map[byte]bool)
	result := 0

	l := 0
	for r := 0; r < len(s); r++ {
		for exists[s[r]] {
			delete(exists, s[l])
			l++
		}

		exists[s[r]] = true

		result = max(result, r-l+1)
	}

	return result
}

func lengthOfLongestSubstringv10(s string) int {
	n := len(s)
	if n < 2 {
		return n
	}

	duplicated := make(map[byte]bool)

	l := 0
	result := 0
	for r := 0; r < n; r++ {
		for duplicated[s[r]] {
			delete(duplicated, s[l])
			l++
		}

		duplicated[s[r]] = true
		result = max(result, r-l+1)
	}

	return result
}

func lengthOfLongestSubstringv9(s string) int {
	n := len(s)
	if n < 2 {
		return n
	}

	m := make(map[byte]bool)

	l, r := 0, 0

	result := 0
	for r < n {

		for m[s[r]] {
			delete(m, s[l])
			l++
		}

		m[s[r]] = true
		result = max(result, r-l+1)
		r++
	}

	return result
}

func lengthOfLongestSubstringv8(s string) int {
	n := len(s)
	if n <= 1 {
		return n
	}

	m := make(map[byte]bool)
	l := 0

	result := 0
	for r := 0; r < n; r++ {
		for m[s[r]] {
			delete(m, s[l])
			l++
		}

		m[s[r]] = true
		result = max(result, r-l+1)

	}

	return result
}

func lengthOfLongestSubstringv7(s string) int {
	n := len(s)
	if n < 2 {
		return n
	}

	exist := make(map[byte]bool)

	result := 0

	l, r := 0, 0

	for r < n {
		for exist[s[r]] {
			delete(exist, s[l])
			l++
		}

		exist[s[r]] = true

		result = max(result, r-l+1)
		r++
	}

	return result
}

func lengthOfLongestSubstringv6(s string) int {
	n := len(s)
	if n < 2 {
		return n
	}
	l, r, result := 0, 0, 0
	m := make(map[byte]bool)

	for r < n {

		for m[s[r]] {
			delete(m, s[l])
			l++
		}

		m[s[r]] = true
		result = max(result, r-l+1)
		r++
	}

	return result
}

func lengthOfLongestSubstringv5(s string) int {
	n := len(s)
	if n <= 1 {
		return n
	}

	m := make(map[byte]bool)

	result := 0
	l, r := 0, 0
	for r < n {

		for m[s[r]] {
			delete(m, s[l])
			l++
		}

		result = max(result, r-l+1)
		m[s[r]] = true
		r++

	}

	return result
}

func lengthOfLongestSubstringv4(s string) int {
	n := len(s)
	if n < 2 {
		return n
	}

	exist := make(map[byte]bool)
	result := 0
	l := 0

	for i := 0; i < n; i++ {
		for exist[s[i]] {
			delete(exist, s[l])
			l++
		}

		exist[s[i]] = true
		result = max(result, i-l+1)
	}

	return result
}

func lengthOfLongestSubstringv3(s string) int {
	n := len(s)
	if n < 2 {
		return n
	}
	l, r := 0, 0

	result := 0
	m := make(map[byte]bool)
	for l < n && r < n {

		for m[s[r]] {
			delete(m, s[l])
			l++
		}

		m[s[r]] = true
		result = max(result, r-l+1)
		r++
	}
	return result
}

func lengthOfLongestSubstringv2(s string) int {
	n := len(s)
	if n < 2 {
		return n
	}

	m := make(map[byte]bool)

	l := 0
	result := 0
	for r := 0; r < n; r++ {
		for m[s[r]] {
			delete(m, s[l])
			l++
		}

		m[s[r]] = true
		result = max(result, r-l+1)
	}

	return result
}

func lengthOfLongestSubstringv1(s string) int {

	n := len(s)
	if n < 2 {
		return n
	}

	charset := make(map[byte]bool)

	l := 0

	result := 0
	for r := 0; r < n; r++ {
		for charset[s[r]] {
			// move l and clean the duplicated value
			delete(charset, s[l])
			l++
		}

		charset[s[r]] = true
		result = max(result, r-l+1)
	}

	return result
}
