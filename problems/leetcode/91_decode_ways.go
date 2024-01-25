package b75

/*
91. Decode Ways
Medium

A message containing letters from A-Z can be encoded into numbers using the following mapping:

'A' -> "1"
'B' -> "2"
...
'Z' -> "26"
To decode an encoded message, all the digits must be grouped then
mapped back into letters using the reverse of the mapping above (there may be multiple ways).



For example, "11106" can be mapped into:

"AAJF" with the grouping (1 1 10 6)
"KJF" with the grouping (11 10 6)
Note that the grouping (1 11 06) is invalid because "06" cannot be mapped into 'F' since "6" is different from "06".

Given a string s containing only digits, return the number of ways to decode it.

The test cases are generated so that the answer fits in a 32-bit integer.



Example 1:

Input: s = "12"
Output: 2
Explanation: "12" could be decoded as "AB" (1 2) or "L" (12).
Example 2:

Input: s = "226"
Output: 3
Explanation: "226" could be decoded as "BZ" (2 26), "VF" (22 6), or "BBF" (2 2 6).
Example 3:

Input: s = "06"
Output: 0
Explanation: "06" cannot be mapped to "F" because of the leading zero ("6" is different from "06").


Constraints:

1 <= s.length <= 100
s contains only digits and may contain leading zero(s).
*/

func numDecodingsv47(s string) int {
	if len(s) == 0 {
		return 0
	}

	if s[0] == '0' {
		return 0
	}

	dp := make([]int, len(s)+1)
	dp[0] = 1
	for i := 1; i <= len(s); i++ {
		if s[i-1] != '0' {
			dp[i] = dp[i-1]
		}

		if i > 1 && ((s[i-2] == '1') || (s[i-2] == '2' && s[i-1] <= '6')) {
			dp[i] += dp[i-2]
		}
	}

	return dp[len(s)]
}

func numDecodingsv46(s string) int {
	if len(s) == 0 {
		return 0
	}

	if s[0] == '0' {
		return 0
	}

	dp := make([]int, len(s)+1)
	dp[0] = 1
	for i := 1; i <= len(s); i++ {
		if s[i-1] != '0' {
			dp[i] = dp[i-1]
		}

		if i > 1 && ((s[i-2] == '1') || (s[i-2] == '2' && s[i-1] <= '6')) {
			dp[i] += dp[i-2]
		}
	}
	return dp[len(s)]
}

func numDecodingsv45(s string) int {
	if len(s) == 0 {
		return 0
	}

	if s[0] == '0' {
		return 0
	}

	dp := make([]int, len(s)+1)
	dp[0] = 1
	for i := 1; i <= len(s); i++ {
		if s[i-1] != '0' {
			dp[i] = dp[i-1]
		}

		if i > 1 && ((s[i-2] == '1') || (s[i-2] == '2' && s[i-1] <= '6')) {
			dp[i] += dp[i-2]
		}
	}

	return dp[len(s)]
}

func numDecodingsv44(s string) int {
	if len(s) == 0 {
		return 0
	}

	if s[0] == '0' {
		return 0
	}

	dp := make([]int, len(s)+1)
	dp[0] = 1
	for i := 1; i <= len(s); i++ {
		if s[i-1] != '0' {
			dp[i] = dp[i-1]
		}

		if i > 1 && ((s[i-2] == '1') || (s[i-2] == '2' && s[i-1] <= '6')) {
			dp[i] += dp[i-2]
		}
	}

	return dp[len(s)]
}

func numDecodingsv43(s string) int {
	if len(s) == 0 {
		return 0
	}

	if s[0] == '0' {
		return 0
	}

	dp := make([]int, len(s)+1)
	dp[0] = 1
	for i := 1; i <= len(s); i++ {
		if s[i-1] != '0' {
			dp[i] = dp[i-1]
		}

		if i > 1 && ((s[i-2] == '1') || (s[i-2] == '2' && s[i-1] <= '6')) {
			dp[i] += dp[i-2]
		}
	}
	return dp[len(s)]
}

func numDecodingsv42(s string) int {
	if len(s) == 0 {
		return 0
	}

	if s[0] == '0' {
		return 0
	}

	dp := make([]int, len(s)+1)
	dp[0] = 1
	for i := 1; i <= len(s); i++ {
		if s[i-1] != '0' {
			dp[i] = dp[i-1]
		}

		if i > 1 && ((s[i-2] == '1') || (s[i-2] == '2' && s[i-1] <= '6')) {
			dp[i] += dp[i-2]
		}
	}

	return dp[len(s)]
}

func numDecodingsv41(s string) int {
	if len(s) == 0 {
		return 0
	}

	if s[0] == '0' {
		return 0
	}

	dp := make([]int, len(s)+1)
	dp[0] = 1
	for i := 1; i <= len(s); i++ {
		if s[i-1] != '0' {
			dp[i] = dp[i-1]
		}

		if i > 1 && ((s[i-2] == '1') || (s[i-2] == '2' && s[i-1] <= '6')) {
			dp[i] += dp[i-2]
		}
	}
	return dp[len(s)]
}

func numDecodingsv40(s string) int {
	if len(s) == 0 {
		return 0
	}

	if s[0] == '0' {
		return 0
	}

	dp := make([]int, len(s)+1)
	dp[0] = 1
	for i := 1; i <= len(s); i++ {
		if s[i-1] != '0' {
			dp[i] = dp[i-1]
		}

		if i > 1 && ((s[i-2] == '1') || (s[i-2] == '2' && s[i-1] <= '6')) {
			dp[i] += dp[i-2]
		}
	}

	return dp[len(s)]
}

func numDecodingsv39(s string) int {
	if len(s) == 0 {
		return 0
	}

	if s[0] == '0' {
		return 0
	}

	dp := make([]int, len(s)+1)
	dp[0] = 1
	for i := 1; i <= len(s); i++ {
		if s[i-1] != '0' {
			dp[i] = dp[i-1]
		}

		if i > 1 && ((s[i-2] == '1') || (s[i-2] == '2' && s[i-1] <= '6')) {
			dp[i] += dp[i-2]
		}
	}

	return dp[len(s)]
}

func numDecodingsv38(s string) int {
	if len(s) == 0 {
		return 0
	}

	if s[0] == '0' {
		return 0
	}

	dp := make([]int, len(s)+1)
	dp[0] = 1
	for i := 1; i <= len(s); i++ {
		if s[i-1] != '0' {
			dp[i] = dp[i-1]
		}

		if i > 1 && ((s[i-2] == '1') || (s[i-2] == '2' && s[i-1] <= '6')) {
			dp[i] += dp[i-2]
		}
	}
	return dp[len(s)]
}

func numDecodingsv37(s string) int {
	if len(s) == 0 {
		return 0
	}

	if s[0] == '0' {
		return 0
	}

	dp := make([]int, len(s)+1)
	dp[0] = 1

	for i := 1; i <= len(s); i++ {
		if s[i-1] != '0' {
			dp[i] = dp[i-1]
		}

		if i > 1 && ((s[i-2] == '1') || (s[i-2] == '2' && s[i-1] <= '6')) {
			dp[i] += dp[i-2]
		}
	}
	return dp[len(s)]
}

func numDecodingsv36(s string) int {
	if len(s) == 0 {
		return 0
	}

	if s[0] == '0' {
		return 0
	}

	dp := make([]int, len(s)+1)
	dp[0] = 1
	for i := 1; i <= len(s); i++ {
		if s[i-1] != '0' {
			dp[i] = dp[i-1]
		}

		if i > 1 && (s[i-2] == '1' || (s[i-2] == '2' && s[i-1] <= '6')) {
			dp[i] += dp[i-2]
		}
	}

	return dp[len(s)]
}

func numDecodingsv35(s string) int {

	if len(s) == 0 {
		return 0
	}

	if s[0] == '0' {
		return 0
	}

	dp := make([]int, len(s)+1)
	dp[0] = 1

	for i := 1; i <= len(s); i++ {
		if s[i-1] != '0' {
			dp[i] = dp[i-1]
		}

		if i > 1 && (s[i-2] == '1' || s[i-2] == '2' && s[i-1] <= '6') {
			dp[i] += dp[i-2]
		}
	}

	return dp[len(s)]
}

func numDecodingsv34(s string) int {
	if len(s) == 0 {
		return 0
	}

	if s[0] == '0' {
		return 0
	}

	dp := make([]int, len(s)+1)
	dp[0] = 1
	for i := 1; i <= len(s); i++ {
		if s[i-1] != '0' {
			dp[i] = dp[i-1]
		}
		if i > 1 && ((s[i-2] == '1') || (s[i-2] == '2' && s[i-1] <= '6')) {
			dp[i] += dp[i-2]
		}
	}

	return dp[len(s)]
}

func numDecodingsv33(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}

	if s[0] == '0' {
		return 0
	}

	dp := make([]int, n+1)
	dp[0] = 1

	for i := 1; i <= n; i++ {
		if s[i-1] != '0' {
			dp[i] = dp[i-1]
		}

		if i > 1 && (s[i-2] == '1' || (s[i-2] == '2' && s[i-1] <= '6')) {
			dp[i] += dp[i-2]
		}
	}

	return dp[n]
}

func numDecodingsv32(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}

	if s[0] == '0' {
		return 0
	}

	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		if s[i-1] != '0' {
			dp[i] = dp[i-1]
		}

		if i > 1 && (s[i-2] == '1' || (s[i-2] == '2' && s[i-1] <= '6')) {
			dp[i] += dp[i-2]
		}
	}

	return dp[n]
}

func numDecodingsv31(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}

	if s[0] == '0' {
		return 0
	}

	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		if s[i-1] != '0' {
			dp[i] = dp[i-1]
		}

		if i > 1 && (s[i-2] == '1' || (s[i-2] == '2' && s[i-1] <= '6')) {
			dp[i] += dp[i-2]
		}
	}
	return dp[n]
}

func numDecodingsv30(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}

	if s[0] == '0' {
		return 0
	}

	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		if s[i-1] != '0' {
			dp[i] = dp[i-1]
		}

		if i > 1 && (s[i-2] == '1' || (s[i-2] == '2' && s[i-1] <= '6')) {
			dp[i] += dp[i-2]
		}
	}

	return dp[n]
}

func numDecodingsv29(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}
	if s[0] == '0' {
		return 0
	}

	dp := make([]int, n+1)
	dp[0] = 1

	for i := 1; i <= n; i++ {
		if s[i-1] != '0' {
			dp[i] = dp[i-1]
		}

		if i > 1 && (s[i-2] == '1' || (s[i-2] == '2' && s[i-1] <= '6')) {
			dp[i] += dp[i-2]
		}
	}

	return dp[n]
}

func numDecodingsv28(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}

	if s[0] == '0' {
		return 0
	}

	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		if s[i-1] != '0' {
			dp[i] = dp[i-1]
		}

		if i > 1 && (s[i-2] == '1' || (s[i-2] == '2' && s[i-1] <= '6')) {
			dp[i] += dp[i-2]
		}
	}

	return dp[n]
}

func numDecodingsv27(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}

	if s[0] == '0' {
		return 0
	}

	dp := make([]int, n+1)
	dp[0] = 1

	for i := 1; i <= n; i++ {
		if s[i-1] != '0' {
			dp[i] = dp[i-1]
		}

		if i > 1 && (s[i-2] == '1' || (s[i-2] == '2' && s[i-1] <= '6')) {
			dp[i] += dp[i-2]
		}
	}

	return dp[n]
}

func numDecodingsv26(s string) int {
	n := len(s)

	if n == 0 {
		return 0
	}

	if s[0] == '0' {
		return 0
	}

	dp := make([]int, n+1)
	dp[0] = 1

	for i := 1; i <= n; i++ {
		if s[i-1] != '0' {
			dp[i] = dp[i-1]
		}

		if i > 1 && (s[i-2] == '1' || (s[i-2] == '2' && s[i-1] <= '6')) {
			dp[i] += dp[i-2]
		}
	}

	return dp[n]
}

func numDecodingsv25(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}

	if s[0] == '0' {
		return 0
	}

	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i <= n; i++ {

		if s[i-1] != '0' {
			dp[i] = dp[i-1]
		}

		if i > 1 && (s[i-2] == '1' || (s[i-2] == '2' && s[i-1] <= '6')) {
			dp[i] += dp[i-2]
		}
	}

	return dp[n]
}

func numDecodingsv24(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}

	if s[0] == '0' {
		return 0
	}

	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		if s[i-1] != '0' {
			dp[i] = dp[i-1]
		}

		if i > 1 && (s[i-2] == '1' || s[i-2] == '2' && s[i-1] <= '6') {
			dp[i] += dp[i-2]
		}
	}

	return dp[n]
}

func numDecodingsv23(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}

	if s[0] == '0' {
		return 0
	}

	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i <= n; i++ {

		if s[i-1] != '0' {
			dp[i] = dp[i-1]
		}

		if i > 1 && (s[i-2] == '1' || (s[i-2] == '2' && s[i-1] <= '6')) {
			dp[i] += dp[i-2]
		}

	}

	return dp[n]
}

func numDecodingsv22(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}

	if s[0] == '0' {
		return 0
	}

	dp := make([]int, n+1)
	dp[0] = 1

	for i := 1; i <= n; i++ {
		if s[i-1] != '0' {
			dp[i] = dp[i-1]
		}

		if i > 1 && (s[i-2] == '1' || (s[i-2] == '2' && s[i-1] <= '6')) {
			dp[i] += dp[i-2]
		}
	}

	return dp[n]
}

func numDecodingsv21(s string) int {
	if s[0] == '0' {
		return 0
	}

	dp := make([]int, len(s)+1)
	dp[0] = 1

	for i := 1; i <= len(s); i++ {
		if s[i-1] != '0' {
			dp[i] = dp[i-1]
		}

		if i > 1 && (s[i-2] == '1' || (s[i-2] == '2' && s[i-1] <= '6')) {
			dp[i] += dp[i-2]
		}
	}

	return dp[len(s)]
}

func numDecodingsv20(s string) int {
	n := len(s)

	if n == 0 {
		return 0
	}

	if s[0] == '0' {
		return 0
	}

	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		if s[i-1] != '0' {
			dp[i] = dp[i-1]
		}

		if i > 1 && (s[i-2] == '1' || (s[i-2] == '2' && s[i-1] <= '6')) {
			dp[i] += dp[i-2]
		}
	}

	return dp[n]
}

func numDecodingsv19(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}

	if s[0] == '0' {
		return 0
	}

	dp := make([]int, n+1)
	dp[0] = 1

	for i := 1; i <= n; i++ {
		if s[i-1] != '0' {
			dp[i] = dp[i-1]
		}

		if i > 1 && (s[i-2] == '1' || (s[i-2] == '2' && s[i-1] <= '6')) {
			dp[i] += dp[i-2]
		}
	}

	return dp[n]
}

func numDecodingsv18(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}

	if s[0] == '0' {
		return 0
	}

	dp := make([]int, n+1)
	dp[0] = 1

	for i := 1; i <= n; i++ {
		if s[i-1] != '0' {
			dp[i] = dp[i-1]
		}

		if i > 1 && (s[i-2] == '1' || (s[i-2] == '2' && s[i-1] <= '6')) {
			dp[i] += dp[i-2]
		}
	}

	return dp[n]
}

func numDecodingsv17(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}

	if s[0] == '0' {
		return 0
	}

	dp := make([]int, n+1)
	dp[0] = 1

	for i := 1; i <= n; i++ {
		if s[i-1] != '0' {
			dp[i] = dp[i-1]
		}

		if i > 1 && (s[i-2] == '1' || (s[i-2] == '2' && s[i-1] <= '6')) {
			dp[i] += dp[i-2]
		}
	}

	return dp[n]
}

func numDecodingsv16(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}

	if s[0] == '0' {
		return 0
	}

	dp := make([]int, n+1)
	dp[0] = 1

	for i := 1; i <= n; i++ {
		if s[i-1] != '0' {
			dp[i] = dp[i-1]
		}

		if i > 1 && (s[i-2] == '1' || (s[i-2] == '2' && s[i-1] <= '6')) {
			dp[i] += dp[i-2]
		}

	}

	return dp[n]
}

func numDecodingsv15(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}

	if n == 1 && s[0] == '0' {
		return 0
	}

	dp := make([]int, n+1)
	dp[0] = 1

	for i := 1; i <= n; i++ {
		if s[i-1] != '0' {
			dp[i] = dp[i-1]
		}

		if i > 1 && (s[i-2] == '1' || (s[i-2] == '2' && s[i-1] <= '6')) {
			dp[i] += dp[i-2]
		}
	}

	return dp[n]
}

func numDecodingsv14(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}
	if n == 1 && s[0] == '0' {
		return 0
	}

	dp := make([]int, n+1)
	dp[0] = 1

	for i := 1; i <= n; i++ {
		if s[i-1] != '0' {
			dp[i] = dp[i-1]
		}

		if i > 1 && (s[i-2] == '1' || (s[i-2] == '2' && s[i-1] <= '6')) {
			dp[i] += dp[i-2]
		}
	}

	return dp[n]
}

func numDecodingsv13(s string) int {
	n := len(s)

	if n == 0 {
		return 0
	}

	if s[0] == '0' {
		return 0
	}

	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		if s[i-1] != '0' {
			dp[i] = dp[i-1]
		}

		if i > 1 && (s[i-2] == '1' || (s[i-2] == '2' && s[i-1] <= '6')) {
			dp[i] += dp[i-2]
		}
	}

	return dp[n]
}

func numDecodingsv12(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}

	if s[0] == '0' {
		return 0
	}

	dp := make([]int, n+1)
	dp[0] = 1

	for i := 1; i <= n; i++ {
		if s[i-1] != '0' {
			dp[i] = dp[i-1]
		}

		if i > 1 && (s[i-2] == '1' || s[i-2] == '2' && s[i-1] <= '6') {
			dp[i] += dp[i-2]
		}
	}

	return dp[n]
}

func numDecodingsv11(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}

	if s[0] == '0' {
		return 0
	}

	dp := make([]int, n+1)
	dp[0] = 1

	for i := 1; i <= n; i++ {
		if s[i-1] != '0' {
			dp[i] = dp[i-1]
		}

		if i > 1 && (s[i-2] == '1' || s[i-2] == '2' && s[i-1] <= '6') {
			dp[i] += dp[i-2]
		}
	}

	return dp[n]
}

func numDecodingsv10(s string) int {
	n := len(s)

	if n == 0 {
		return 0
	}

	if s[0] == '0' {
		return 0
	}

	dp := make([]int, n+1)
	dp[0] = 1

	for i := 1; i <= n; i++ {
		if s[i-1] != '0' {
			dp[i] = dp[i-1] // base case: if s[i-1] != '0' -> got 1 way for decoding
		}

		if i > 1 && (s[i-2] == '1' || s[i-2] == '2' && s[i-1] <= '6') {
			dp[i] += dp[i-2]
		}
	}

	return dp[n]
}

func numDecodingsv9(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}

	if s[0] == '0' {
		return 0
	}

	dp := make([]int, n+1)
	dp[0] = 1

	for i := 1; i <= n; i++ {
		if s[i-1] != '0' {
			dp[i] = dp[i-1]
		}

		if i > 1 && (s[i-2] == '1' || s[i-2] == '2' && s[i-1] <= '6') {
			dp[i] += dp[i-2]
		}
	}

	return dp[n]
}

func numDecodingsv8(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}

	if s[0] == '0' {
		return 0
	}

	dp := make([]int, n+1)
	dp[0] = 1 // first valid
	for i := 1; i <= n; i++ {
		if s[i-1] != '0' {
			dp[i] = dp[i-1]
		}

		if i > 1 && (s[i-2] == '1' || s[i-2] == '2' && s[i-1] <= '6') {
			dp[i] += dp[i-2]
		}
	}
	return dp[n]
}

func numDecodingsv7(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}

	if s[0] == '0' {
		return 0
	}

	dp := make([]int, n+1)
	dp[0] = 1

	for i := 1; i <= n; i++ {
		if s[i-1] != '0' {
			dp[i] = dp[i-1]
		}

		if i > 1 && (s[i-2] == '1' || s[i-2] == '2' && s[i-1] <= '6') {
			dp[i] += dp[i-2]
		}
	}

	return dp[n]
}

func numDecodingsv6(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}

	if s[0] == '0' {
		return 0
	}

	dp := make([]int, n+1)
	dp[0] = 1

	for i := 1; i <= n; i++ {

		if s[i-1] != '0' {
			dp[i] = dp[i-1]
		}

		if i > 1 && (s[i-2] == '1' || s[i-2] == '2' && s[i-1] <= '6') {
			dp[i] += dp[i-2]
		}

	}

	return dp[n]
}

func numDecodingsv5(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}

	if s[0] == '0' {
		return 0
	}

	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		if s[i-1] != '0' {
			dp[i] = dp[i-1]
		}

		if i > 1 && (s[i-2] == '1' || s[i-2] == '2' && s[i-1] <= '6') {
			dp[i] += dp[i-2]
		}
	}
	return dp[n]
}

func numDecodingsv4(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}

	if s[0] == '0' {
		return 0
	}

	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		if s[i-1] != '0' {
			dp[i] = dp[i-1]
		}
		if i > 1 && (s[i-2] == '1' || s[i-2] == '2' && s[i-1] <= '6') {
			dp[i] += dp[i-2]
		}
	}

	return dp[n]
}

func numDecodingsv3(s string) int {

	if len(s) == 0 {
		return 0
	}

	if s[0] == '0' {
		return 0
	}

	m := make([]int, len(s)+1)
	// len(s) == 1 the m[0], m[1] are valid so set them to 1
	m[0], m[1] = 1, 1
	for i := 2; i <= len(s); i++ {
		if s[i-1] != '0' {
			m[i] = m[i-1]
		}

		if s[i-2] == '1' || (s[i-2] == '2' && s[i-1] <= '6') {
			m[i] += m[i-2]
		}
	}

	return m[len(s)]
}

func numDecodingsv2(s string) int {

	if len(s) == 0 {
		return 0
	}

	if len(s) > 0 && s[0] == '0' {
		return 0
	}

	m := make([]int, len(s)+1)
	m[0], m[1] = 1, 1 // it denotes the first two rune is valid

	for i := 2; i <= len(s); i++ {
		if s[i-1] != '0' {
			m[i] = m[i-1]
		}

		if s[i-2] == '1' || (s[i-2] == '2' && s[i-1] >= '0' && s[i-1] <= '6') {
			m[i] += m[i-2]
		}
	}

	return m[len(s)]
}

func numDecodingsv1(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}

	if s[0] == '0' {
		return 0
	}

	m := make([]int, n+1)
	// base case is valid
	m[0], m[1] = 1, 1

	for i := 2; i <= n; i++ {
		if s[i-1] != '0' {
			m[i] = m[i-1]
		}

		if s[i-2] == '1' || (s[i-2] == '2' && s[i-1] >= '0' && s[i-1] <= '6') {
			m[i] += m[i-2]
		}
	}

	return m[n]
}

func NumDecodings(s string) int {

	n := len(s)
	if n == 0 || s[0] == '0' {
		return 0
	}

	m := make([]int, n+1)
	m[0] = 1

	for i := 1; i <= n; i++ {
		if s[i-1] != '0' {
			m[i] = m[i-1]
		}
		if i >= 2 && (s[i-2] == '1' || (s[i-2] == '2' && s[i-1] <= '6')) {
			m[i] += m[i-2]
		}
	}
	return m[n]

}

func numDecodings(s string) int {

	n := len(s)
	if n == 0 {
		return 0
	}

	// first rune shouldn't be '0'
	if s[0] == '0' {
		return 0
	}

	m := make([]int, n+1)
	m[0] = 1 // because we checked the first rune is not '0', so we say the s[0] is valid
	for i := 1; i <= n; i++ {
		if s[i-1] != '0' {
			// check the single rune
			// start from position 1 and s[i-1] != '0' which implies s[1] is valid(0-9)
			// so, just clone the record from m[i-1] to m[i]
			// consider 12, we have (1,2)
			m[i] = m[i-1]
		}

		// consider 2 runes
		// if s[i-2] == 1, the s[i-1] can be 0-9
		// if s[i-2] == 2, the s[i-1] can be 0-6
		// count m[i-2], because it is denoted the record of m[i-2] && m[i-1]
		if i >= 2 && (s[i-2] == '1' || (s[i-2] == '2' && s[i-1] >= '0' && s[i-1] <= '6')) {
			m[i] += m[i-2]
		}
	}

	return m[n]
}
