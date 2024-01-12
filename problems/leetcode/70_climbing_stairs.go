package b75

import (
	"math"
)

/*
Example 1:

Input: n = 2
Output: 2
Explanation: There are two ways to climb to the top.
1. 1 step + 1 step
2. 2 steps
Example 2:

Input: n = 3
Output: 3
Explanation: There are three ways to climb to the top.
1. 1 step + 1 step + 1 step
2. 1 step + 2 steps
3. 2 steps + 1 step


Constraints:

1 <= n <= 45
*/

func ClimbingStairsv42(n int) int {
	if n < 2 {
		return n
	}

	n1, n2 := 1, 2
	for i := 3; i <= n; i++ {
		tmp := n1 + n2
		n1 = n2
		n2 = tmp
	}

	return n2
}

func ClimbingStairsv41(n int) int {
	if n < 2 {
		return n
	}

	n1, n2 := 1, 2
	for i := 3; i <= n; i++ {
		tmp := n1 + n2
		n1 = n2
		n2 = tmp
	}

	return n2
}

func ClimbingStairsv40(n int) int {
	if n < 2 {
		return n
	}

	n1, n2 := 1, 2
	for i := 3; i <= n; i++ {
		tmp := n1 + n2
		n1 = n2
		n2 = tmp
	}
	return n2
}

func ClimbingStairsv39(n int) int {
	if n < 2 {
		return n
	}

	n1, n2 := 1, 2
	for i := 3; i <= n; i++ {
		tmp := n1 + n2
		n1 = n2
		n2 = tmp
	}

	return n2
}

func ClimbingStairsv38(n int) int {
	if n < 2 {
		return n
	}

	n1, n2 := 1, 2
	for i := 3; i <= n; i++ {
		tmp := n1 + n2
		n1 = n2
		n2 = tmp
	}
	return n2
}

func ClimbingStairsv37(n int) int {
	if n < 2 {
		return n
	}

	n1, n2 := 1, 2

	for i := 3; i <= n; i++ {
		tmp := n1 + n2
		n1 = n2
		n2 = tmp
	}

	return n2
}

func ClimbingStairsv36(n int) int {
	if n < 2 {
		return n
	}

	n1, n2 := 1, 2
	for i := 3; i <= n; i++ {
		tmp := n1 + n2
		n1 = n2
		n2 = tmp
	}
	return n2
}

func ClimbingStairsv35(n int) int {
	if n < 2 {
		return n
	}

	n1, n2 := 1, 2

	for i := 3; i <= n; i++ {
		tmp := n1 + n2
		n1 = n2
		n2 = tmp
	}

	return n2
}

func ClimbingStairsv34(n int) int {
	if n < 2 {
		return n
	}

	n1, n2 := 1, 2
	for i := 3; i <= n; i++ {
		tmp := n1 + n2
		n1 = n2
		n2 = tmp
	}

	return n2
}

func ClimbingStairsv33(n int) int {
	if n < 2 {
		return n
	}

	n1, n2 := 1, 2

	for i := 3; i <= n; i++ {
		tmp := n1 + n2
		n1 = n2
		n2 = tmp
	}

	return n2
}

func ClimbingStairsv32(n int) int {
	if n < 2 {
		return n
	}

	n1, n2 := 1, 2
	for i := 3; i <= n; i++ {
		tmp := n1 + n2
		n1 = n2
		n2 = tmp
	}

	return n2
}

func ClimbingStairsv31(n int) int {
	if n < 2 {
		return n
	}

	n1, n2 := 1, 2

	for i := 3; i <= n; i++ {
		tmp := n1 + n2
		n1 = n2
		n2 = tmp
	}

	return n2
}

func ClimbingStairsv30(n int) int {
	if n < 2 {
		return n
	}

	n1, n2 := 1, 2

	for i := 3; i <= n; i++ {
		tmp := n1 + n2
		n1 = n2
		n2 = tmp
	}

	return n2
}

func ClimbingStairsv29(n int) int {
	if n < 2 {
		return n
	}

	n1, n2 := 1, 2

	for i := 3; i <= n; i++ {
		tmp := n1 + n2
		n1 = n2
		n2 = tmp
	}
	return n2
}

func ClimbingStairsv28(n int) int {
	if n < 3 {
		return n
	}

	n1, n2 := 1, 2
	for i := 3; i <= n; i++ {
		tmp := n1 + n2
		n1 = n2
		n2 = tmp
	}

	return n2
}

func ClimbingStairsv27(n int) int {
	if n <= 2 {
		return n
	}

	n1, n2 := 1, 2
	for i := 3; i <= n; i++ {
		tmp := n1 + n2
		n1 = n2
		n2 = tmp
	}
	return n2
}

func ClimbingStairsv26(n int) int {
	if n <= 2 {
		return n
	}

	a, b := 1, 2
	for i := 3; i <= n; i++ {
		tmp := a + b
		a = b
		b = tmp
	}

	return b
}

func ClimbingStairsv25(n int) int {
	if n <= 2 {
		return n
	}

	a, b := 1, 2
	for i := 3; i <= n; i++ {
		tmp := a + b
		a = b
		b = tmp
	}

	return b
}

func ClimbingStairsv24(n int) int {
	if n < 3 {
		return n
	}

	a, b := 1, 2
	for i := 3; i <= n; i++ {
		tmp := a + b
		a = b
		b = tmp
	}
	return b
}

func ClimbingStairsv23(n int) int {
	if n <= 2 {
		return n
	}

	a, b := 1, 2

	for i := 3; i <= n; i++ {
		tmp := a + b
		a = b
		b = tmp
	}

	return b
}

func ClimbingStairsv22(n int) int {
	if n <= 2 {
		return n
	}

	a, b := 1, 2
	for i := 3; i <= n; i++ {
		tmp := a + b
		a = b
		b = tmp
	}

	return b
}

func ClimbingStairsv21(n int) int {
	if n <= 2 {
		return n
	}

	a, b := 1, 2

	for i := 3; i <= n; i++ {
		tmp := a + b
		a = b
		b = tmp
	}

	return b
}

func ClimbingStairsv20(n int) int {
	if n <= 2 {
		return n
	}

	a, b := 1, 2

	for i := 3; i <= n; i++ {
		tmp := a + b
		a = b
		b = tmp
	}

	return b
}

func ClimbingStairsv19(n int) int {
	if n <= 2 {
		return n
	}

	a, b := 1, 2

	for i := 3; i <= n; i++ {
		tmp := a + b
		a = b
		b = tmp
	}

	return b
}

func ClimbingStairsv18(n int) int {
	if n <= 2 {
		return n
	}

	a, b := 1, 2

	for i := 3; i <= n; i++ {
		tmp := a + b
		a = b
		b = tmp
	}

	return b
}

func ClimbingStairsv17(n int) int {
	if n <= 2 {
		return n
	}

	f1, f2 := 1, 2

	for i := 3; i <= n; i++ {
		tmp := f1 + f2
		f1 = f2
		f2 = tmp
	}

	return f2
}

func ClimbingStairsv16(n int) int {

	if n <= 2 {
		return n
	}

	a, b := 1, 2

	for i := 3; i <= n; i++ {
		tmp := a + b
		a = b
		b = tmp
	}

	return b
}

func ClimbingStairsv15(n int) int {
	if n < 3 {
		return n
	}

	a, b := 1, 2

	for i := 3; i <= n; i++ {
		tmp := a + b
		a = b
		b = tmp
	}

	return b
}

func ClimbingStairsv14(n int) int {
	if n <= 2 {
		return n
	}

	a, b := 1, 2

	for i := 3; i <= n; i++ {
		tmp := a + b
		a = b
		b = tmp
	}

	return b
}

func ClimbingStairsv13(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	a, b := 1, 2

	for i := 3; i <= n; i++ {
		tmp := a + b
		a = b
		b = tmp
	}

	return b
}

func ClimbingStairsv12(n int) int {
	if n == 0 {
		return 1
	}

	a, b := 0, 1

	for i := 0; i < n; i++ {
		tmp := a + b
		a = b
		b = tmp
	}

	return b
}

func ClimbingStairsv11(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	a, b := 0, 1

	for i := 0; i < n; i++ {
		tmp := a + b
		a = b
		b = tmp
	}

	return b
}

func ClimbingStairsv10(n int) int {

	if n == 0 {
		return 0
	}

	a, b := 0, 1

	for i := 0; i < n; i++ {
		tmp := a + b
		a = b
		b = tmp
	}

	return b
}

func ClimbingStairsv9(n int) int {
	if n == 0 {
		return 0
	}

	a, b := 0, 1

	for i := 0; i < n; i++ {
		tmp := a + b
		a = b
		b = tmp
	}

	return b
}

func ClimbingStairsv8(n int) int {
	a, b := 0, 1

	for i := 0; i < n; i++ {
		tmp := a + b
		a = b
		b = tmp
	}

	return b
}

func ClimbingStairsv7(n int) int {
	a, b := 0, 1

	for i := 0; i < n; i++ {
		tmp := a + b
		a = b
		b = tmp
	}

	return b

}

func ClimbingStairsv6(n int) int {
	a, b := 0, 1

	for i := 0; i < n; i++ {
		tmp := a + b
		a = b
		b = tmp
	}

	return b
}

func ClimbingStairsv5(n int) int {

	if n == 1 || n == 2 {
		return n
	}

	r1, r2 := 1, 2

	for i := 3; i <= n; i++ {
		tmp := r1 + r2
		r1 = r2
		r2 = tmp
	}

	return r2
}

func ClimbingStairsv4(n int) int {

	q := make([][]int, 2)
	for i := 0; i < 2; i++ {
		q[i] = make([]int, 2)
	}

	q[0][0] = 1 // fn
	q[0][1] = 1 // fn-1
	q[1][0] = 1 // fn-1
	q[1][1] = 0 // fn-2

	for i := 0; i < n-1; i++ {
		a, b, c, d := q[0][0], q[0][1], q[1][0], q[1][1]
		q[0][0] = a*1 + b*1
		q[0][1] = a*1 + b*0
		q[1][0] = c*1 + d*1
		q[1][1] = c*1 + d*0
	}

	return q[0][0]
}

func ClimbingStairsv3(n int) int {
	if n <= 2 {
		return n
	}

	phi := (1 + math.Sqrt(5)) / 2
	return int(math.Round((math.Pow(phi, float64(n+1)) / math.Sqrt(5))))
}

func ClimbingStairsv2(n int) int {
	if n <= 2 {
		return n
	}

	a, b := 1, 2
	for i := 3; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

func ClimbingStairsv1(n int) int {

	if n <= 0 {
		return 0
	}

	if n == 1 || n == 2 {
		return n
	}

	m := make([]int, n+1)
	// base cases
	// 0: no stairs => 0 step
	// becasue we can move 1 or 2 steps so
	// 1: 1 stair => 1 step
	// 2: 2 stairs => move by 1 step (1 step + 1 step) + move by 2 steps (2 steps)
	m[0], m[1], m[2] = 0, 1, 2

	for i := 3; i <= n; i++ {
		m[i] = m[i-1] + m[i-2]
	}

	return m[n]

}

func ClimbingStairs(n int) int {

	if n <= 0 {
		return 0
	}

	if n == 1 || n == 2 {
		return n
	}

	m := make([]int, n+1)
	m[0] = 0
	m[1] = 1
	m[2] = 2
	for i := 3; i <= n; i++ {
		m[i] = m[i-1] + m[i-2]
	}

	return m[n]
}
