package b75

/*
647. Palindromic Substrings
Given a string s, return the number of palindromic substrings in it.

A string is a palindrome when it reads the same backward as forward.

A substring is a contiguous sequence of characters within the string.



Example 1:

Input: s = "abc"
Output: 3
Explanation: Three palindromic strings: "a", "b", "c".
Example 2:

Input: s = "aaa"
Output: 6
Explanation: Six palindromic strings: "a", "a", "a", "aa", "aa", "aaa".


Constraints:

1 <= s.length <= 1000
s consists of lowercase English letters.
*/

func countSubstringsv31(s string) int {
	result := 0

	for i := range s {
		l, r := i, i
		for l >= 0 && r < len(s) && s[l] == s[r] {
			result++
			l--
			r++
		}

		l, r = i, i+1
		for l >= 0 && r < len(s) && s[l] == s[r] {
			result++
			l--
			r++
		}
	}

	return result
}

func countSubstringsv30(s string) int {
	result := 0

	for i := range s {
		l, r := i, i
		for l >= 0 && r < len(s) && s[l] == s[r] {
			result++
			l--
			r++
		}

		l, r = i, i+1
		for l >= 0 && r < len(s) && s[l] == s[r] {
			result++
			l--
			r++
		}
	}

	return result
}

func countSubstringsv29(s string) int {
	result := 0

	for i := range s {
		l, r := i, i
		for l >= 0 && r < len(s) && s[l] == s[r] {
			result++
			l--
			r++
		}

		l, r = i, i+1
		for l >= 0 && r < len(s) && s[l] == s[r] {
			result++
			l--
			r++
		}
	}

	return result
}

func countSubstringsv28(s string) int {
	result := 0

	for i := range s {
		l, r := i, i
		for l >= 0 && r < len(s) && s[l] == s[r] {
			result++
			l--
			r++
		}

		l, r = i, i+1
		for l >= 0 && r < len(s) && s[l] == s[r] {
			result++
			l--
			r++
		}

	}

	return result
}

func countSubstringsv27(s string) int {
	result := 0

	for i := range s {
		l, r := i, i
		for l >= 0 && r < len(s) && s[l] == s[r] {
			result++
			l--
			r++
		}

		l, r = i, i+1
		for l >= 0 && r < len(s) && s[l] == s[r] {
			result++
			l--
			r++
		}

	}

	return result
}

func countSubstringsv26(s string) int {
	result := 0
	for i := range s {
		l, r := i, i
		for l >= 0 && r < len(s) && s[l] == s[r] {
			result++
			l--
			r++
		}

		l, r = i, i+1
		for l >= 0 && r < len(s) && s[l] == s[r] {
			result++
			l--
			r++
		}
	}
	return result
}

func countSubstringsv25(s string) int {
	result := 0

	for i := range s {
		l, r := i, i
		for l >= 0 && r < len(s) && s[l] == s[r] {
			result++
			l--
			r++
		}

		l, r = i, i+1
		for l >= 0 && r < len(s) && s[l] == s[r] {
			result++
			l--
			r++
		}
	}

	return result
}

func countSubstringsv24(s string) int {
	result := 0
	for i := range s {
		l, r := i, i
		for l >= 0 && r < len(s) && s[l] == s[r] {
			result++
			l--
			r++
		}

		l, r = i, i+1
		for l >= 0 && r < len(s) && s[l] == s[r] {
			result++
			l--
			r++
		}
	}
	return result
}

func countSubstringsv23(s string) int {
	result := 0

	for i := range s {
		l, r := i, i
		for l >= 0 && r < len(s) && s[l] == s[r] {
			result++
			l--
			r++
		}

		l, r = i, i+1
		for l >= 0 && r < len(s) && s[l] == s[r] {
			result++
			l--
			r++
		}
	}

	return result
}

func countSubstringsv22(s string) int {
	result := 0

	for i := range s {
		l, r := i, i
		for l >= 0 && r < len(s) && s[l] == s[r] {
			l--
			r++
			result++
		}

		l, r = i, i+1
		for l >= 0 && r < len(s) && s[l] == s[r] {
			l--
			r++
			result++
		}

	}

	return result
}

func countSubstringsv21(s string) int {
	result := 0
	for i := range s {
		l, r := i, i
		for l >= 0 && r < len(s) && s[l] == s[r] {
			result++
			l--
			r++
		}

		l, r = i, i+1
		for l >= 0 && r < len(s) && s[l] == s[r] {
			result++
			l--
			r++
		}
	}
	return result
}

func countSubstringsv20(s string) int {
	result := 0
	for i := range s {
		l, r := i, i
		for l >= 0 && r < len(s) && s[l] == s[r] {
			result++
			l--
			r++
		}

		l, r = i, i+1
		for l >= 0 && r < len(s) && s[l] == s[r] {
			result++
			l--
			r++
		}
	}

	return result
}

func countSubstringsv19(s string) int {
	result := 0

	for i := range s {
		l, r := i, i
		for l >= 0 && r < len(s) && s[l] == s[r] {
			result++
			l--
			r++
		}

		l, r = i, i+1
		for l >= 0 && r < len(s) && s[l] == s[r] {
			result++
			l--
			r++
		}
	}

	return result
}

func countSubstringsv18(s string) int {
	result := 0
	for i := range s {
		l, r := i, i
		for l >= 0 && r < len(s) && s[l] == s[r] {
			result++
			l--
			r++
		}

		l, r = i, i+1
		for l >= 0 && r < len(s) && s[l] == s[r] {
			result++
			l--
			r++
		}
	}

	return result
}

func countSubstringsv17(s string) int {
	result := 0

	for i := range s {
		l, r := i, i
		for l >= 0 && r < len(s) && s[l] == s[r] {
			result++
			l--
			r++
		}

		l, r = i, i+1
		for l >= 0 && r < len(s) && s[l] == s[r] {
			result++
			l--
			r++
		}
	}

	return result
}

func countSubstringsv16(s string) int {
	result := 0

	for i := range s {
		l, r := i, i
		for l >= 0 && r < len(s) && s[l] == s[r] {
			result++
			l--
			r++
		}

		l, r = i, i+1
		for l >= 0 && r < len(s) && s[l] == s[r] {
			result++
			l--
			r++
		}
	}

	return result
}

func countSubstringsv15(s string) int {
	result := 0

	for i := range s {
		l, r := i, i
		for l >= 0 && r < len(s) && s[l] == s[r] {
			result++
			l--
			r++
		}

		l, r = i, i+1
		for l >= 0 && r < len(s) && s[l] == s[r] {
			result++
			l--
			r++
		}
	}

	return result
}

func countSubstringsv14(s string) int {
	result := 0

	for i := range s {
		l, r := i, i
		for l >= 0 && r < len(s) {
			if s[l] != s[r] {
				break
			}
			result++
			l--
			r++
		}

		l, r = i, i+1
		for l >= 0 && r < len(s) {
			if s[l] != s[r] {
				break
			}
			result++
			l--
			r++
		}
	}

	return result
}

func countSubstringsv13(s string) int {
	n := len(s)
	if n <= 1 {
		return n
	}
	result := 0

	for i := 0; i < n; i++ {
		l, r := i, i
		for l >= 0 && r < n && s[l] == s[r] {
			l--
			r++
			result++
		}
		l, r = i, i+1
		for l >= 0 && r < n && s[l] == s[r] {
			l--
			r++
			result++
		}

	}

	return result
}

func countSubstringsv12(s string) int {
	n := len(s)
	if n <= 1 {
		return n
	}

	result := 0

	for i := 0; i < n; i++ {
		l, r := i, i
		for l >= 0 && r < n && s[l] == s[r] {
			l--
			r++
			result++
		}
		l, r = i, i+1
		for l >= 0 && r < n && s[l] == s[r] {
			l--
			r++
			result++
		}

	}

	return result
}

func countSubstringsv11(s string) int {
	n := len(s)

	if n <= 1 {
		return n
	}

	result := 0

	for i := 0; i < n; i++ {
		l, r := i, i
		for l >= 0 && r < n && s[l] == s[r] {
			l--
			r++
			result++
		}
		l, r = i, i+1
		for l >= 0 && r < n && s[l] == s[r] {
			l--
			r++
			result++
		}

	}
	return result
}

func countSubstringsv10(s string) int {
	n := len(s)
	if n <= 1 {
		return n
	}

	var result int = 0

	for i := 0; i < n; i++ {
		l, r := i, i
		for l >= 0 && r < n && s[l] == s[r] {
			l--
			r++
			result++
		}
		l, r = i, i+1
		for l >= 0 && r < n && s[l] == s[r] {
			l--
			r++
			result++
		}

	}

	return result
}

func countSubstringsv9(s string) int {
	n := len(s)
	if n <= 1 {
		return n
	}

	var result int = 0

	for i := 0; i < n; i++ {
		l, r := i, i
		for l >= 0 && r < n && s[l] == s[r] {
			result++
			l--
			r++
		}
		l, r = i, i+1
		for l >= 0 && r < n && s[l] == s[r] {
			result++
			l--
			r++
		}

	}

	return result
}

func countSubstringsv8(s string) int {
	n := len(s)
	if n <= 1 {
		return n
	}

	var result int = 0
	for i := 0; i < n; i++ {
		l, r := i, i
		for l >= 0 && r < n && s[l] == s[r] {
			result++
			l--
			r++
		}

		l, r = i, i+1
		for l >= 0 && r < n && s[l] == s[r] {
			result++
			l--
			r++
		}
	}

	return result
}

func countSubstringsv7(s string) int {
	n := len(s)

	if n <= 1 {
		return n
	}

	var result int = 0
	for i := 0; i < n; i++ {
		l, r := i, i
		for l >= 0 && r < n && s[l] == s[r] {
			result++
			l--
			r++
		}
		l, r = i, i+1
		for l >= 0 && r < n && s[l] == s[r] {
			result++
			l--
			r++
		}

	}

	return result
}

func countSubstringsv6(s string) int {
	n := len(s)
	if n <= 1 {
		return n
	}

	var result int = 0

	for i := 0; i < n; i++ {
		l, r := i, i
		for l >= 0 && r < n && s[l] == s[r] {
			result++
			l--
			r++
		}
		l, r = i, i+1
		for l >= 0 && r < n && s[l] == s[r] {
			result++
			l--
			r++
		}

	}

	return result
}

func countSubstringsv5(s string) int {
	n := len(s)
	if n <= 1 {
		return n
	}

	var result int = 0

	for i := 0; i < n; i++ {
		l, r := i, i

		for l >= 0 && r < n && s[l] == s[r] {
			result++
			l--
			r++
		}

		l, r = i, i+1
		for l >= 0 && r < n && s[l] == s[r] {
			result++
			l--
			r++
		}
	}
	return result
}

func countSubstringsv4(s string) int {
	n := len(s)
	if n <= 1 {
		return n
	}

	var result int = 0

	for i := 0; i < n; i++ {
		l, r := i, i
		for l >= 0 && r < n && s[l] == s[r] {
			result++
			l--
			r++
		}

		l, r = i, i+1
		for l >= 0 && r < n && s[l] == s[r] {
			result++
			l--
			r++
		}
	}

	return result
}

func countSubstringsv3(s string) int {
	n := len(s)
	if n <= 1 {
		return n
	}

	var result int = 0

	for i := 0; i < n; i++ {
		// single charator
		l, r := i, i
		for l >= 0 && r < n && s[l] == s[r] {
			result++
			l--
			r++
		}

		// double charators
		l, r = i, i+1
		for l >= 0 && r < n && s[l] == s[r] {
			result++
			l--
			r++
		}

	}

	return result
}

func countSubstringsv2(s string) int {
	var result int = 0
	for i := range s {
		l, r := i, i

		for l >= 0 && r < len(s) && s[l] == s[r] {
			result++
			l--
			r++
		}
		l, r = i, i+1
		for l >= 0 && r < len(s) && s[l] == s[r] {
			result++
			l--
			r++
		}
	}

	return result
}

func countSubstringsv1(s string) int {

	var result int = 0

	for i := range s {
		l, r := i, i

		// one rune
		for l >= 0 && r < len(s) && s[l] == s[r] {
			result++
			l--
			r++
		}

		// two runes
		l, r = i, i+1
		for l >= 0 && r < len(s) && s[l] == s[r] {
			result++
			l--
			r++
		}
	}
	return result
}

func countSubstrings(s string) int {
	var result int = 0

	for i := range s {
		l, r := i, i

		// odd length
		for l >= 0 && r < len(s) && s[l] == s[r] {
			result++
			l--
			r++
		}

		// even length
		l, r = i, i+1
		for l >= 0 && r < len(s) && s[l] == s[r] {
			result++
			l--
			r++
		}

	}

	return result
}
