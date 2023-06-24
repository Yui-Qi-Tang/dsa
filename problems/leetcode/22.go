package b75

import "strings"

/*
22. Generate Parentheses
Given n pairs of parentheses,
write a function to generate all combinations of well-formed parentheses.



Example 1:

Input: n = 3
Output: ["((()))","(()())","(())()","()(())","()()()"]
Example 2:

Input: n = 1
Output: ["()"]
k

Constraints:
*/

func generateParenthesisv34(n int) []string {
	result, stack := []string{}, []string{}

	var backtrack func(open, close int)

	backtrack = func(open, close int) {
		if open == n && close == n {
			result = append(result, strings.Join(stack, ""))
		}

		if open < n {
			stack = append(stack, "(")
			backtrack(open+1, close)
			stack = stack[:len(stack)-1]
		}

		if close < open {
			stack = append(stack, ")")
			backtrack(open, close+1)
			stack = stack[:len(stack)-1]
		}
	}

	backtrack(0, 0)
	return result
}

func generateParenthesisv33(n int) []string {
	result, stack := []string{}, []string{}

	var backtrack func(open, close int)

	backtrack = func(open, close int) {
		if open == n && close == n {
			result = append(result, strings.Join(stack, ""))
		}

		if open < n {
			stack = append(stack, "(")
			backtrack(open+1, close)
			stack = stack[:len(stack)-1]
		}

		if close < open {
			stack = append(stack, ")")
			backtrack(open, close+1)
			stack = stack[:len(stack)-1]
		}
	}

	backtrack(0, 0)
	return result
}

func generateParenthesisv32(n int) []string {
	result, stack := []string{}, []string{}

	var backtrack func(open, close int)
	backtrack = func(open, close int) {
		if open == n && close == n {
			result = append(result, strings.Join(stack, ""))
		}

		if open < n {
			stack = append(stack, "(")
			backtrack(open+1, close)
			stack = stack[:len(stack)-1]
		}

		if close < open {
			stack = append(stack, ")")
			backtrack(open, close+1)
			stack = stack[:len(stack)-1]
		}
	}

	backtrack(0, 0)

	return result
}

func generateParenthesisv31(n int) []string {
	result, stack := make([]string, 0), make([]string, 0)

	var backtrack func(open, close int)

	backtrack = func(open, close int) {
		if open == n && close == n {
			result = append(result, strings.Join(stack, ""))
		}

		if open < n {
			stack = append(stack, "(")
			backtrack(open+1, close)
			stack = stack[:len(stack)-1]
		}

		if close < open {
			stack = append(stack, ")")
			backtrack(open, close+1)
			stack = stack[:len(stack)-1]
		}
	}

	backtrack(0, 0)

	return result
}

func generateParenthesisv30(n int) []string {
	result, stack := make([]string, 0), make([]string, 0)

	var backtrack func(open, close int)
	backtrack = func(open, close int) {
		if open == n && close == n {
			result = append(result, strings.Join(stack, ""))
		}

		if open < n {
			stack = append(stack, "(")
			backtrack(open+1, close)
			stack = stack[:len(stack)-1]
		}

		if close < open {
			stack = append(stack, ")")
			backtrack(open, close+1)
			stack = stack[:len(stack)-1]
		}
	}

	backtrack(0, 0)
	return result
}

func generateParenthesisv29(n int) []string {
	result, stack := make([]string, 0), make([]string, 0)

	var backtrack func(open, close int)

	backtrack = func(open, close int) {
		if open == n && close == n {
			result = append(result, strings.Join(stack, ""))
		}

		if open < n {
			stack = append(stack, "(")
			backtrack(open+1, close)
			stack = stack[:len(stack)-1]
		}

		if close < open {
			stack = append(stack, ")")
			backtrack(open, close+1)
			stack = stack[:len(stack)-1]
		}
	}

	backtrack(0, 0)

	return result
}

func generateParenthesisv28(n int) []string {
	result, stack := make([]string, 0), make([]string, 0)

	var backtrack func(open, close int)
	backtrack = func(open, close int) {
		if open == n && close == n {
			result = append(result, strings.Join(stack, ""))
		}

		if open < n {
			stack = append(stack, "(")
			backtrack(open+1, close)
			stack = stack[:len(stack)-1]
		}

		if close < open {
			stack = append(stack, ")")
			backtrack(open, close+1)
			stack = stack[:len(stack)-1]
		}
	}

	backtrack(0, 0)

	return result
}

func generateParenthesisv27(n int) []string {
	result, stack := make([]string, 0), make([]string, 0)

	var backtrack func(open, close int)

	backtrack = func(open, close int) {
		if open == n && close == n {
			result = append(result, strings.Join(stack, ""))
		}

		if open < n {
			stack = append(stack, "(")
			backtrack(open+1, close)
			stack = stack[:len(stack)-1]
		}

		if close < open {
			stack = append(stack, ")")
			backtrack(open, close+1)
			stack = stack[:len(stack)-1]
		}
	}

	backtrack(0, 0)

	return result
}

func generateParenthesisv26(n int) []string {
	result, stack := make([]string, 0), make([]string, 0)

	var backtrack func(open, close int)
	backtrack = func(open, close int) {
		if open == n && close == n {
			result = append(result, strings.Join(stack, ""))
		}

		if open < n {
			stack = append(stack, "(")
			backtrack(open+1, close)
			stack = stack[:len(stack)-1]
		}

		if close < open {
			stack = append(stack, ")")
			backtrack(open, close+1)
			stack = stack[:len(stack)-1]
		}
	}
	backtrack(0, 0)

	return result
}

func generateParenthesisv25(n int) []string {
	result, stack := make([]string, 0), make([]string, 0)

	var backtrack func(open, close int)

	backtrack = func(open, close int) {
		if open == n && close == n {
			result = append(result, strings.Join(stack, ""))
		}

		if open < n {
			stack = append(stack, "(")
			backtrack(open+1, close)
			stack = stack[:len(stack)-1]
		}

		if close < open {
			stack = append(stack, ")")
			backtrack(open, close+1)
			stack = stack[:len(stack)-1]
		}
	}
	backtrack(0, 0)
	return result
}

func generateParenthesisv24(n int) []string {
	result, stack := make([]string, 0), make([]string, 0)

	var backtrack func(open, close int)

	backtrack = func(open, close int) {
		if open == n && close == n {
			result = append(result, strings.Join(stack, ""))
		}

		if open < n {
			stack = append(stack, "(")
			backtrack(open+1, close)
			stack = stack[:len(stack)-1]
		}

		if close < open {
			stack = append(stack, ")")
			backtrack(open, close+1)
			stack = stack[:len(stack)-1]
		}
	}

	backtrack(0, 0)

	return result
}

func generateParenthesisv23(n int) []string {
	result, stack := make([]string, 0), make([]string, 0)

	var backtrack func(open, close int)

	backtrack = func(open, close int) {
		if open == n && close == n {
			result = append(result, strings.Join(stack, ""))
		}

		if open < n {
			stack = append(stack, "(")
			backtrack(open+1, close)
			stack = stack[:len(stack)-1]
		}

		if close < open {
			stack = append(stack, ")")
			backtrack(open, close+1)
			stack = stack[:len(stack)-1]
		}
	}

	backtrack(0, 0)

	return result
}

func generateParenthesisv22(n int) []string {
	result, stack := make([]string, 0), make([]string, 0)

	var backtrack func(open, close int)

	backtrack = func(open, close int) {
		if open == n && close == n {
			result = append(result, strings.Join(stack, ""))
		}

		if open < n {
			stack = append(stack, "(")
			backtrack(open+1, close)
			stack = stack[:len(stack)-1]
		}

		if close < open {
			stack = append(stack, ")")
			backtrack(open, close+1)
			stack = stack[:len(stack)-1]
		}
	}

	backtrack(0, 0)
	return result
}

func generateParenthesisv21(n int) []string {
	result, stack := make([]string, 0), make([]string, 0)

	var backtrack func(int, int)

	backtrack = func(open, close int) {
		if open == n && close == n {
			result = append(result, strings.Join(stack, ""))
		}

		if open < n {
			stack = append(stack, "(")
			backtrack(open+1, close)
			stack = stack[:len(stack)-1]
		}

		if close < open {
			stack = append(stack, ")")
			backtrack(open, close+1)
			stack = stack[:len(stack)-1]
		}
	}

	backtrack(0, 0)
	return result
}

func generateParenthesisv20(n int) []string {
	result, stack := make([]string, 0), make([]string, 0)

	var backtrack func(open, close int)
	backtrack = func(open, close int) {
		if open == n && close == n {
			result = append(result, strings.Join(stack, ""))
		}

		if open < n {
			stack = append(stack, "(")
			backtrack(open+1, close)
			stack = stack[:len(stack)-1]
		}

		if close < open {
			stack = append(stack, ")")
			backtrack(open, close+1)
			stack = stack[:len(stack)-1]
		}
	}
	backtrack(0, 0)

	return result
}

func generateParenthesisv19(n int) []string {
	result, stack := make([]string, 0), make([]string, 0)

	var backtrack func(open, close int)

	backtrack = func(open, close int) {
		if open == n && close == n {
			result = append(result, strings.Join(stack, ""))
		}

		if open < n {
			stack = append(stack, "(")
			backtrack(open+1, close)
			stack = stack[:len(stack)-1]
		}

		if close < open {
			stack = append(stack, ")")
			backtrack(open, close+1)
			stack = stack[:len(stack)-1]
		}
	}

	backtrack(0, 0)
	return result
}

func generateParenthesisv18(n int) []string {
	result, stack := make([]string, 0), make([]string, 0)

	var backtrack func(open, close int)

	backtrack = func(open, close int) {

		if open == n && close == n {
			result = append(result, strings.Join(stack, ""))
		}

		if open < n {
			stack = append(stack, "(")
			backtrack(open+1, close)
			stack = stack[:len(stack)-1]
		}

		if close < open {
			stack = append(stack, ")")
			backtrack(open, close+1)
			stack = stack[:len(stack)-1]
		}

	}

	backtrack(0, 0)

	return result
}

func generateParenthesisv17(n int) []string {
	result, stack := make([]string, 0), make([]string, 0)

	var backtrack func(open, close int)

	backtrack = func(open, close int) {
		if open == n && close == n {
			result = append(result, strings.Join(stack, ""))
		}

		if open < n {
			stack = append(stack, "(")
			backtrack(open+1, close)
			stack = stack[:len(stack)-1]
		}

		if close < open {
			stack = append(stack, ")")
			backtrack(open, close+1)
			stack = stack[:len(stack)-1]
		}
	}

	backtrack(0, 0)
	return result
}

func generateParenthesisv16(n int) []string {
	result, stack := make([]string, 0), make([]string, 0)

	var backtrack func(open, close int)
	backtrack = func(open, close int) {
		if open == n && close == n {
			result = append(result, strings.Join(stack, ""))
			return
		}

		if open < n {
			stack = append(stack, "(")
			backtrack(open+1, close)
			stack = stack[:len(stack)-1]
		}

		if close < open {
			stack = append(stack, ")")
			backtrack(open, close+1)
			stack = stack[:len(stack)-1]
		}
	}
	backtrack(0, 0)

	return result
}

func generateParenthesisv15(n int) []string {
	result, stack := make([]string, 0), make([]string, 0)

	var backtrack func(open, close int)
	backtrack = func(open, close int) {
		if open == n && close == n {
			result = append(result, strings.Join(stack, ""))
		}

		if open < n {
			stack = append(stack, "(")
			backtrack(open+1, close)
			stack = stack[:len(stack)-1]
		}
		if close < open {
			stack = append(stack, ")")
			backtrack(open, close+1)
			stack = stack[:len(stack)-1]
		}
	}
	backtrack(0, 0)
	return result
}

func generateParenthesisv14(n int) []string {
	result, stack := make([]string, 0), make([]string, 0)

	var backtrack func(open, close int)
	backtrack = func(open, close int) {
		if open == n && close == n {
			result = append(result, strings.Join(stack, ""))
		}

		if open < n {
			stack = append(stack, "(")
			backtrack(open+1, close)
			stack = stack[:len(stack)-1]
		}

		if close < open {
			stack = append(stack, ")")
			backtrack(open, close+1)
			stack = stack[:len(stack)-1]
		}
	}

	backtrack(0, 0)
	return result
}

func generateParenthesisv13(n int) []string {
	result, stack := make([]string, 0), make([]string, 0)

	var backtrack func(open, close int)
	backtrack = func(open, close int) {
		if open == n && close == n {
			result = append(result, strings.Join(stack, ""))
		}

		if open < n {
			stack = append(stack, "(")
			backtrack(open+1, close)
			stack = stack[:len(stack)-1]
		}

		if close < open {
			stack = append(stack, ")")
			backtrack(open, close+1)
			stack = stack[:len(stack)-1]
		}
	}
	backtrack(0, 0)
	return result
}

func generateParenthesisv12(n int) []string {
	result, stack := make([]string, 0), make([]string, 0)

	var backtrack func(int, int)
	backtrack = func(open, close int) {
		if open == n && close == n {
			result = append(result, strings.Join(stack, ""))
		}

		if open < n {
			stack = append(stack, "(")
			backtrack(open+1, close)
			stack = stack[:len(stack)-1]
		}

		if close < open {
			stack = append(stack, ")")
			backtrack(open, close+1)
			stack = stack[:len(stack)-1]
		}
	}

	backtrack(0, 0)

	return result
}

func generateParenthesisv11(n int) []string {
	result, stack := make([]string, 0), make([]string, 0)

	var backtrack func(open, close int)
	backtrack = func(open, close int) {
		if open == n && close == n {
			result = append(result, strings.Join(stack, ""))
		}

		// first add open: (
		if open < n {
			stack = append(stack, "(")
			backtrack(open+1, close)
			stack = stack[:len(stack)-1]
		}

		// second add close: )
		if close < open {
			stack = append(stack, ")")
			backtrack(open, close+1)
			stack = stack[:len(stack)-1]
		}
	}
	// The order in this backtracking function will be ({(/)}...
	// so, we can ingore the ")" starting case at the first input -> start with '){(|)}' case should not exist
	backtrack(0, 0)

	return result
}

func generateParenthesisv10(n int) []string {
	result := make([]string, 0)
	stack := make([]string, 0)

	var backtrack func(open, close int)
	backtrack = func(open, close int) {

		if open == n && close == n {
			result = append(result, strings.Join(stack, ""))
		}

		if open < n {
			stack = append(stack, "(")
			backtrack(open+1, close)
			stack = stack[:len(stack)-1]
		}

		if close < open {
			stack = append(stack, ")")
			backtrack(open, close+1)
			stack = stack[:len(stack)-1]
		}

	}

	backtrack(0, 0)
	return result
}

func generateParenthesisv9(n int) []string {
	result, stack := make([]string, 0), make([]string, 0)

	var backtrack func(open, close int)

	backtrack = func(open, close int) {
		if open == n && close == n {
			result = append(result, strings.Join(stack, ""))
		}

		if open < n {
			stack = append(stack, "(")
			backtrack(open+1, close)
			stack = stack[:len(stack)-1]
		}

		if close < open {
			stack = append(stack, ")")
			backtrack(open, close+1)
			stack = stack[:len(stack)-1]
		}
	}

	backtrack(0, 0)
	return result
}

func generateParenthesisv8(n int) []string {
	result, stack := make([]string, 0), make([]string, 0)

	var backtrack func(open, close int)

	backtrack = func(open, close int) {

		if open == n && close == n {
			result = append(result, strings.Join(stack, ""))
		}

		if open < n {
			stack = append(stack, "(")
			backtrack(open+1, close)
			stack = stack[:len(stack)-1]
		}

		if close < open {
			stack = append(stack, ")")
			backtrack(open, close+1)
			stack = stack[:len(stack)-1]
		}
	}
	backtrack(0, 0)
	return result
}

func generateParenthesisv7(n int) []string {
	result, stack := make([]string, 0), make([]string, 0)
	var backtrack func(open, close int)

	backtrack = func(open, close int) {
		if open == n && close == n {
			result = append(result, strings.Join(stack, ""))
		}

		if open < n {
			stack = append(stack, "(")
			backtrack(open+1, close)
			stack = stack[:len(stack)-1]
		}

		if close < open {
			stack = append(stack, ")")
			backtrack(open, close+1)
			stack = stack[:len(stack)-1]
		}
	}

	backtrack(0, 0)
	return result
}

func generateParenthesisv6(n int) []string {
	result := make([]string, 0)
	stack := make([]string, 0)

	var backtrack func(open, close int)

	backtrack = func(open, close int) {
		if open == n && close == n {
			result = append(result, strings.Join(stack, ""))
		}

		if open < n {
			stack = append(stack, "(")
			backtrack(open+1, close)
			stack = stack[:len(stack)-1]
		}

		if close < open {
			stack = append(stack, ")")
			backtrack(open, close+1)
			stack = stack[:len(stack)-1]
		}
	}

	backtrack(0, 0)
	return result
}

func generateParenthesisv5(n int) []string {
	result := make([]string, 0)

	stack := make([]string, 0)

	var backtrack func(open, close int)
	backtrack = func(open, close int) {

		if open == n && close == n {
			result = append(result, strings.Join(stack, ""))
		}

		if open < n {
			stack = append(stack, "(")
			backtrack(open+1, close)
			stack = stack[:len(stack)-1]
		}

		if close < open {
			stack = append(stack, ")")
			backtrack(open, close+1)
			stack = stack[:len(stack)-1]
		}

	}

	backtrack(0, 0)

	return result
}

func generateParenthesisv4(n int) []string {
	result := make([]string, 0)

	stack := make([]string, 0)
	var backtrack func(int, int)
	backtrack = func(open, close int) {
		if open == n && close == n {
			result = append(result, strings.Join(stack, ""))
		}

		if open < n {
			stack = append(stack, "(")
			backtrack(open+1, close)
			stack = stack[:len(stack)-1]
		}

		if close < open {
			stack = append(stack, ")")
			backtrack(open, close+1)
			stack = stack[:len(stack)-1]
		}
	}

	backtrack(0, 0)
	return result
}

func generateParenthesisv3(n int) []string {
	result := make([]string, 0)
	stack := make([]string, 0)

	var backtrack func(open, close int)
	backtrack = func(open, close int) {
		if open == n && close == n {
			result = append(result, strings.Join(stack, ""))
		}

		if open < n {
			stack = append(stack, "(")
			backtrack(open+1, close)
			stack = stack[:len(stack)-1]
		}

		if close < open {
			stack = append(stack, ")")
			backtrack(open, close+1)
			stack = stack[:len(stack)-1]
		}
	}

	backtrack(0, 0)
	return result
}

func generateParenthesisv2(n int) []string {
	result := make([]string, 0)
	stack := make([]string, 0)
	var backtrack func(open, close int)
	backtrack = func(open, close int) {
		if open == n && close == n {
			result = append(result, strings.Join(stack, ""))
		}

		if open < n {
			stack = append(stack, "(")
			backtrack(open+1, close)
			stack = stack[:len(stack)-1]
		}

		if close < open {
			stack = append(stack, ")")
			backtrack(open, close+1)
			stack = stack[:len(stack)-1]
		}
	}

	backtrack(0, 0)
	return result
}

func generateParenthesisv1(n int) []string {

	// ensure the ordering of the appending

	result, stack := make([]string, 0), make([]string, 0)
	var backtracing func(open, closed int)
	backtracing = func(open, closed int) {
		// base case
		if open == n && closed == n {
			result = append(result, strings.Join(stack, ""))
		}

		// add open first
		if open < n {
			stack = append(stack, "(")
			backtracing(open+1, closed)
			stack = stack[:len(stack)-1]
		}

		// add closed if number of closed is less than open
		if closed < open {
			stack = append(stack, ")")
			backtracing(open, closed+1)
			stack = stack[:len(stack)-1]
		}
	}

	backtracing(0, 0)
	return result
}
