package b75

/*
5. Longest Palindromic Substring
Given a string s, return the longest palindromic substring in s.

A string is called a palindrome string if the reverse of that string is the same as the original string.

Example 1:

Input: s = "babad"
Output: "bab"
Explanation: "aba" is also a valid answer.
Example 2:

Input: s = "cbbd"
Output: "bb"

Constraints:

1 <= s.length <= 1000
s consist of only digits and English letters.
*/

func longestPalindromev40(s string) string {
	ll, lr := 0, 1

	for i := range s {
		l, r := i, i
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if r-l+1 > lr-ll {
				ll, lr = l, r+1
			}
			l--
			r++
		}

		l, r = i, i+1
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if r-l+1 > lr-ll {
				ll, lr = l, r+1
			}
			l--
			r++
		}

	}

	return s[ll:lr]
}

func longestPalindromev39(s string) string {
	ll, lr := 0, 1 // length

	for i := range s {
		l, r := i, i
		for l >= 0 && r < len(s) && s[l] == s[r] {
			// r-l+1: length
			if r-l+1 > lr-ll {
				ll, lr = l, r+1
			}
			l--
			r++
		}

		l, r = i, i+1
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if r-l+1 > lr-ll {
				ll, lr = l, r+1
			}
			l--
			r++
		}

	}

	return s[ll:lr]
}

func longestPalindromev38(s string) string {
	ll, lr := 0, 0

	for i := range s {
		l, r := i, i
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if r-l+1 > lr-ll+1 {
				ll, lr = l, r
			}

			l--
			r++
		}

		l, r = i, i+1
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if r-l+1 > lr-ll+1 {
				ll, lr = l, r
			}

			l--
			r++
		}

	}

	return s[ll : lr+1]
}

func longestPalindromev37(s string) string {
	ll, lr := 0, 0

	for i := range s {
		l, r := i, i
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if r-l+1 > lr-ll+1 {
				ll, lr = l, r
			}
			l--
			r++
		}

		l, r = i, i+1
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if r-l+1 > lr-ll+1 {
				ll, lr = l, r
			}
			l--
			r++
		}

	}

	return s[ll : lr+1]
}

func longestPalindromev36(s string) string {
	ll, lr := 0, 0

	for i := range s {
		l, r := i, i
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if r-l+1 > lr-ll+1 {
				ll, lr = l, r
			}
			l--
			r++
		}

		l, r = i, i+1
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if r-l+1 > lr-ll+1 {
				ll, lr = l, r
			}
			l--
			r++
		}

	}

	return s[ll : lr+1]
}

func longestPalindromev35(s string) string {
	ll, lr := 0, 0

	for i := range s {
		l, r := i, i
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if r-l+1 > lr-ll+1 {
				ll, lr = l, r
			}

			l--
			r++
		}

		l, r = i, i+1
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if r-l+1 > lr-ll+1 {
				ll, lr = l, r
			}

			l--
			r++
		}

	}

	return s[ll : lr+1]
}

func longestPalindromev34(s string) string {
	ll, lr := 0, 0

	for i := range s {
		l, r := i, i
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if r-l+1 > lr-ll+1 {
				ll, lr = l, r
			}

			l--
			r++
		}

		l, r = i, i+1
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if r-l+1 > lr-ll+1 {
				ll, lr = l, r
			}

			l--
			r++
		}

	}

	return s[ll : lr+1]
}

func longestPalindromev33(s string) string {
	ll, lr := 0, 0

	for i := range s {
		l, r := i, i
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if r-l+1 > lr-ll+1 {
				ll, lr = l, r
			}
			l--
			r++
		}

		l, r = i, i+1
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if r-l+1 > lr-ll+1 {
				ll, lr = l, r
			}
			l--
			r++
		}

	}

	return s[ll : lr+1]
}

func longestPalindromev32(s string) string {
	ll, lr := 0, 0

	for i := range s {
		l, r := i, i
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if r-l+1 > lr-ll+1 {
				ll, lr = l, r
			}
			l--
			r++
		}

		l, r = i, i+1
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if r-l+1 > lr-ll+1 {
				ll, lr = l, r
			}
			l--
			r++
		}

	}

	return s[ll : lr+1]
}

func longestPalindromev31(s string) string {
	ll, lr := 0, 0

	for i := range s {
		l, r := i, i
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if r-l+1 > lr-ll+1 {
				ll, lr = l, r
			}
			l--
			r++
		}

		l, r = i, i+1
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if r-l+1 > lr-ll+1 {
				ll, lr = l, r
			}
			l--
			r++
		}
	}

	return s[ll : lr+1]
}

func longestPalindromev30(s string) string {
	ll, lr := 0, 0

	for i := range s {
		l, r := i, i
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if r-l+1 > lr-ll+1 {
				ll, lr = l, r
			}
			l--
			r++
		}

		l, r = i, i+1
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if r-l+1 > lr-ll+1 {
				ll, lr = l, r
			}
			l--
			r++
		}

	}

	return s[ll : lr+1]
}

func longestPalindromev29(s string) string {
	ll, lr := 0, 0

	for i := range s {
		l, r := i, i
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if r-l+1 > lr-ll+1 {
				ll, lr = l, r
			}
			l--
			r++
		}
		l, r = i, i+1
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if r-l+1 > lr-ll+1 {
				ll, lr = l, r
			}
			l--
			r++
		}

	}

	return s[ll : lr+1]
}

func longestPalindromev28(s string) string {
	ll, lr := 0, 0

	for i := range s {
		l, r := i, i
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if r-l+1 > lr-ll+1 {
				ll, lr = l, r
			}
			l--
			r++
		}
		l, r = i, i+1
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if r-l+1 > lr-ll+1 {
				ll, lr = l, r
			}
			l--
			r++
		}

	}

	return s[ll : lr+1]
}

func longestPalindromev27(s string) string {
	ll, lr := 0, 0

	for i := range s {
		l, r := i, i
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if r-l+1 > lr-ll+1 {
				ll, lr = l, r
			}

			l--
			r++
		}

		l, r = i, i+1
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if r-l+1 > lr-ll+1 {
				ll, lr = l, r
			}

			l--
			r++
		}

	}

	return s[ll : lr+1]
}

func longestPalindromev26(s string) string {
	ll, lr := 0, 0

	for i := range s {
		l, r := i, i
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if r-l+1 > lr-ll+1 {
				ll, lr = l, r
			}
			l--
			r++
		}
		l, r = i, i+1
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if r-l+1 > lr-ll+1 {
				ll, lr = l, r
			}
			l--
			r++
		}

	}

	return s[ll : lr+1]
}

func longestPalindromev25(s string) string {
	ll, lr := 0, 0

	for i := range s {
		l, r := i, i
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if r-l+1 > lr-ll+1 {
				ll, lr = l, r
			}

			l--
			r++
		}

		l, r = i, i+1
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if r-l+1 > lr-ll+1 {
				ll, lr = l, r
			}

			l--
			r++
		}

	}

	return s[ll : lr+1]
}

func longestPalindromev24(s string) string {
	ll, lr := 0, 0

	for i := range s {
		l, r := i, i
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if r-l+1 > lr-ll+1 {
				ll, lr = l, r
			}
			l--
			r++
		}
		l, r = i, i+1
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if r-l+1 > lr-ll+1 {
				ll, lr = l, r
			}
			l--
			r++
		}

	}

	return s[ll : lr+1]
}

func longestPalindromev23(s string) string {
	ll, lr := 0, 0

	for i := range s {
		l, r := i, i
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if r-l+1 > lr-ll+1 {
				ll, lr = l, r
			}
			l--
			r++
		}
		l, r = i, i+1
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if r-l+1 > lr-ll+1 {
				ll, lr = l, r
			}
			l--
			r++
		}

	}

	return s[ll : lr+1]
}

func longestPalindromev22(s string) string {
	ll, lr := 0, 0

	for i := range s {
		l, r := i, i
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if r-l+1 > lr-ll+1 {
				ll, lr = l, r
			}
			l--
			r++
		}

		l, r = i, i+1
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if r-l+1 > lr-ll+1 {
				ll, lr = l, r
			}
			l--
			r++
		}

	}

	return s[ll : lr+1]
}

func longestPalindromev21(s string) string {
	ll, lr := 0, 0

	for i := range s {
		l, r := i, i
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if r-l+1 > lr-ll+1 {
				ll, lr = l, r
			}
			l--
			r++
		}

		l, r = i, i+1
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if r-l+1 > lr-ll+1 {
				ll, lr = l, r
			}
			l--
			r++
		}

	}

	return s[ll : lr+1]
}

func longestPalindromev20(s string) string {
	ll, lr := 0, 0

	for i := range s {
		l, r := i, i
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if r-l+1 > lr-ll+1 {
				ll, lr = l, r
			}
			l--
			r++
		}

		l, r = i, i+1
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if r-l+1 > lr-ll+1 {
				ll, lr = l, r
			}
			l--
			r++
		}

	}

	return s[ll : lr+1]
}

func longestPalindromev19(s string) string {
	ll, lr := 0, 0

	for i := range s {
		l, r := i, i
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if r-l+1 > lr-ll+1 {
				ll, lr = l, r
			}
			l--
			r++
		}
		l, r = i, i+1
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if r-l+1 > lr-ll+1 {
				ll, lr = l, r
			}
			l--
			r++
		}

	}

	return s[ll : lr+1]
}

func longestPalindromev18(s string) string {

	ll, lr := 0, 0

	for i := range s {
		l, r := i, i
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if r-l+1 > lr-ll+1 {
				ll, lr = l, r
			}

			l--
			r++
		}
		l, r = i, i+1
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if r-l+1 > lr-ll+1 {
				ll, lr = l, r
			}

			l--
			r++
		}

	}

	return s[ll : lr+1]
}

func longestPalindromev17(s string) string {
	ll, lr := 0, 0

	for i := range s {
		l, r := i, i
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if r-l+1 > lr-ll+1 {
				ll, lr = l, r
			}

			l--
			r++
		}

		l, r = i, i+1
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if r-l+1 > lr-ll+1 {
				ll, lr = l, r
			}

			l--
			r++
		}
	}

	return s[ll : lr+1]
}

func longestPalindromev16(s string) string {
	ll, lr := 0, 0

	for i := range s {
		l, r := i, i
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if r-l+1 > lr-ll+1 {
				ll, lr = l, r
			}

			l--
			r++
		}

		l, r = i, i+1
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if r-l+1 > lr-ll+1 {
				ll, lr = l, r
			}

			l--
			r++
		}
	}

	return s[ll : lr+1]
}

func longestPalindromev15(s string) string {
	ll, lr := 0, 0

	for i := range s {
		l, r := i, i
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if r-l+1 > lr-ll+1 {
				ll, lr = l, r
			}

			l--
			r++
		}
		l, r = i, i+1
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if r-l+1 > lr-ll+1 {
				ll, lr = l, r
			}

			l--
			r++
		}

	}

	return s[ll : lr+1]
}

func longestPalindromev14(s string) string {
	ll, lr := 0, 0

	for i := range s {
		l, r := i, i
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if r-l > lr-ll {
				ll, lr = l, r
			}
			l--
			r++
		}

		l, r = i, i+1
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if r-l > lr-ll {
				ll, lr = l, r
			}
			l--
			r++
		}
	}

	return s[ll : lr+1]
}

func longestPalindromev13(s string) string {

	ll, lr := 0, 0

	for i := range s {
		l, r := i, i
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if r-l+1 > lr-ll+1 {
				ll, lr = l, r
			}
			l--
			r++
		}
		l, r = i, i+1
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if r-l+1 > lr-ll+1 {
				ll, lr = l, r
			}

			l--
			r++
		}

	}

	return s[ll : lr+1]
}

func longestPalindromev12(s string) string {
	n := len(s)
	if n <= 1 {
		return s
	}

	ll, lr := 0, 0

	for i := range s {
		l, r := i, i
		for l >= 0 && r < n && s[l] == s[r] {
			if r-l+1 > lr-ll+1 {
				ll, lr = l, r
			}
			l--
			r++
		}

		l, r = i, i+1
		for l >= 0 && r < n && s[l] == s[r] {
			if r-l+1 > lr-ll+1 {
				ll, lr = l, r
			}
			l--
			r++
		}
	}

	return s[ll : lr+1]
}

func longestPalindromev11(s string) string {
	n := len(s)
	if n <= 1 {
		return s
	}

	ll, lr := 0, 0

	for i := range s {
		l, r := i, i
		for l >= 0 && r < n && s[l] == s[r] {
			if r-l+1 > lr-ll+1 {
				ll, lr = l, r
			}
			l--
			r++
		}

		l, r = i, i+1
		for l >= 0 && r < n && s[l] == s[r] {
			if r-l+1 > lr-ll+1 {
				ll, lr = l, r
			}
			l--
			r++
		}
	}

	return s[ll : lr+1]
}

func longestPalindromev10(s string) string {
	n := len(s)
	if n <= 1 {
		return s
	}

	// at least one repeat char in the string s
	ll, lr := 0, 0

	for i := range s {
		l, r := i, i
		for l >= 0 && r < n && s[l] == s[r] {
			if r-l+1 > lr-ll+1 {
				ll = l
				lr = r
			}
			l--
			r++
		}

		l, r = i, i+1
		for l >= 0 && r < n && s[l] == s[r] {
			if r-l+1 > lr-ll+1 {
				ll = l
				lr = r
			}
			l--
			r++
		}
	}

	return s[ll : lr+1]
}

func longestPalindromev9(s string) string {
	n := len(s)
	if n <= 1 {
		return s
	}

	ll, lr := 0, 0 // ll:longest left index, lr: longest right index

	for i := range s {
		l, r := i, i
		for l >= 0 && r < n && (s[l] == s[r]) {
			if r-l+1 > lr-ll {
				ll = l
				lr = r + 1
			}
			l--
			r++
		}

		l, r = i, i+1
		for l >= 0 && r < n && (s[l] == s[r]) {
			if r-l+1 > lr-ll {
				ll = l
				lr = r + 1
			}
			l--
			r++
		}

	}

	return s[ll:lr]
}

func longestPalindromev8(s string) string {
	n := len(s)
	if n <= 1 {
		return s
	}

	result := ""

	for i := range s {
		l, r := i, i
		for l >= 0 && r < n && s[l] == s[r] {
			substring := s[l : r+1]
			if len(substring) > len(result) {
				result = substring
			}
			l--
			r++
		}
		l, r = i, i+1
		for l >= 0 && r < n && s[l] == s[r] {
			substring := s[l : r+1]
			if len(substring) > len(result) {
				result = substring
			}
			l--
			r++
		}
	}

	return result
}

func longestPalindromev7(s string) string {
	n := len(s)
	if n < 2 {
		return s
	}

	result := ""
	for i := range s {
		l, r := i, i
		for l >= 0 && r < n && s[r] == s[l] {
			substring := s[l : r+1]
			if len(substring) > len(result) {
				result = substring
			}
			l--
			r++
		}

		l, r = i, i+1
		for l >= 0 && r < n && s[r] == s[l] {
			substring := s[l : r+1]
			if len(substring) > len(result) {
				result = substring
			}
			l--
			r++
		}

	}

	return result
}

func longestPalindromev6(s string) string {
	n := len(s)
	if n <= 1 {
		return s
	}

	result := ""

	for i := range s {
		l, r := i, i
		for l >= 0 && r < n && s[l] == s[r] {
			substring := s[l : r+1]
			if len(substring) > len(result) {
				result = substring
			}

			l--
			r++
		}

		l, r = i, i+1
		for l >= 0 && r < n && s[l] == s[r] {
			substring := s[l : r+1]
			if len(substring) > len(result) {
				result = substring
			}

			l--
			r++
		}

	}

	return result
}

func longestPalindromev5(s string) string {
	n := len(s)
	if n <= 1 {
		return s
	}

	result := ""

	for i := range s {
		l, r := i, i
		for l >= 0 && r < n && s[l] == s[r] {
			substring := s[l : r+1]
			if len(substring) > len(result) {
				result = substring
			}
			l--
			r++
		}
		l, r = i, i+1
		for l >= 0 && r < n && s[l] == s[r] {
			substring := s[l : r+1]
			if len(substring) > len(result) {
				result = substring
			}
			l--
			r++
		}

	}

	return result
}

func longestPalindromev4(s string) string {
	n := len(s)
	if n <= 1 {
		return s
	}

	result := ""

	for i := range s {
		l, r := i, i
		for l >= 0 && r < n && s[l] == s[r] {
			substring := s[l : r+1]
			if len(substring) > len(result) {
				result = substring
			}
			r++
			l--
		}

		l, r = i, i+1
		for l >= 0 && r < n && s[l] == s[r] {
			substring := s[l : r+1]
			if len(substring) > len(result) {
				result = substring
			}
			r++
			l--
		}
	}

	return result
}

func longestPalindromev3(s string) string {
	n := len(s)

	result := ""

	for i := range s {
		l, r := i, i
		for l >= 0 && r < n && s[l] == s[r] {
			substring := s[l : r+1]
			if len(substring) > len(result) {
				result = substring
			}
			r++
			l--
		}

		l, r = i, i+1
		for l >= 0 && r < n && s[l] == s[r] {
			substring := s[l : r+1]
			if len(substring) > len(result) {
				result = substring
			}
			r++
			l--
		}
	}

	return result
}

func longestPalindromev2(s string) string {
	n := len(s)
	if n <= 1 {
		return s
	}

	var result string

	for i := range s {
		l, r := i, i
		for l >= 0 && r < n && s[l] == s[r] {
			substring := s[l : r+1]
			if len(substring) > len(result) {
				result = substring
			}
			r++
			l--
		}

		l, r = i, i+1
		for l >= 0 && r < n && s[l] == s[r] {
			substring := s[l : r+1]
			if len(substring) > len(result) {
				result = substring
			}
			r++
			l--
		}

	}

	return result
}

func longestPalindromev1(s string) string {
	var res string = ""

	for i := range s {
		l, r := i, i

		for l >= 0 && r < len(s) && s[l] == s[r] {
			v := s[l : r+1] // [index:number of elements]
			if len(v) > len(res) {
				res = v
			}
			l--
			r++
		}
		l, r = i, i+1
		for l >= 0 && r < len(s) && s[l] == s[r] {
			v := s[l : r+1]
			if len(v) > len(res) {
				res = v
			}
			l--
			r++
		}

	}

	return res
}

func longestPalindrome(s string) string {
	var result string = ""
	for i := range s {
		// odd length
		l, r := i, i
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if r-l+1 > len(result) {
				result = s[l : r+1]
			}
			l--
			r++
		}
		// even length
		l, r = i, i+1
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if r-l+1 > len(result) {
				result = s[l : r+1]
			}
			l--
			r++
		}
	}

	return result
}
