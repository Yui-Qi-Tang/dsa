package b75

/*
746. Min Cost Climbing Stairs

You are given an integer array cost where cost[i] is the cost of ith step on a staircase. Once you pay the cost, you can either climb one or two steps.

You can either start from the step with index 0, or the step with index 1.

Return the minimum cost to reach the top of the floor.



Example 1:

Input: cost = [10,15,20]
Output: 15
Explanation: You will start at index 1.
- Pay 15 and climb two steps to reach the top.
The total cost is 15.
Example 2:

Input: cost = [1,100,1,1,1,100,1,1,100,1]
Output: 6
Explanation: You will start at index 0.
- Pay 1 and climb two steps to reach index 2.
- Pay 1 and climb two steps to reach index 4.
- Pay 1 and climb two steps to reach index 6.
- Pay 1 and climb one step to reach index 7.
- Pay 1 and climb two steps to reach index 9.
- Pay 1 and climb one step to reach the top.
The total cost is 6.


Constraints:

2 <= cost.length <= 1000
0 <= cost[i] <= 999
*/

func minCostClimbingStairsv27(cost []int) int {
	a, b := cost[len(cost)-1], 0

	for i := len(cost) - 2; i >= 0; i-- {
		cost[i] += min(a, b)
		b = a
		a = cost[i]
	}

	return min(a, b)
}

func minCostClimbingStairsv26(cost []int) int {
	a, b := cost[len(cost)-1], 0

	for i := len(cost) - 2; i >= 0; i-- {
		cost[i] += min(a, b)
		b = a
		a = cost[i]
	}

	return min(a, b)
}

func minCostClimbingStairsv25(cost []int) int {
	a, b := cost[len(cost)-1], 0

	for i := len(cost) - 2; i >= 0; i-- {
		cost[i] += min(a, b)
		b = a
		a = cost[i]
	}

	return min(a, b)
}

func minCostClimbingStairsv24(cost []int) int {
	a, b := cost[len(cost)-1], 0

	for i := len(cost) - 2; i >= 0; i-- {
		cost[i] += min(a, b)
		b = a
		a = cost[i]
	}

	return min(a, b)
}

func minCostClimbingStairsv23(cost []int) int {
	a, b := cost[len(cost)-1], 0

	for i := len(cost) - 2; i >= 0; i-- {
		cost[i] += min(a, b)
		b = a
		a = cost[i]
	}

	return min(a, b)
}

func minCostClimbingStairsv22(cost []int) int {
	a, b := cost[len(cost)-1], 0

	for n := len(cost) - 2; n >= 0; n-- {
		cost[n] += min(a, b)
		b = a
		a = cost[n]
	}

	return min(a, b)
}

func minCostClimbingStairsv21(cost []int) int {
	a, b := cost[len(cost)-1], 0
	for n := len(cost) - 2; n >= 0; n-- {
		cost[n] += min(a, b)
		b = a
		a = cost[n]
	}
	return min(a, b)
}

func minCostClimbingStairsv20(cost []int) int {
	a, b := cost[len(cost)-1], 0
	for i := len(cost) - 2; i >= 0; i-- {
		cost[i] += min(a, b)
		b = a
		a = cost[i]
	}

	return min(a, b)
}

func minCostClimbingStairsv19(cost []int) int {
	a, b := cost[len(cost)-1], 0

	for i := len(cost) - 2; i >= 0; i-- {
		cost[i] += min(a, b)
		b = a
		a = cost[i]
	}

	return min(a, b)
}

func minCostClimbingStairsv18(cost []int) int {
	a, b := cost[len(cost)-1], 0

	for i := len(cost) - 2; i >= 0; i-- {
		cost[i] += min(a, b)
		b = a
		a = cost[i]
	}
	return min(a, b)
}

func minCostClimbingStairsv17(cost []int) int {
	a, b := cost[len(cost)-1], 0
	for i := len(cost) - 2; i >= 0; i-- {
		cost[i] += min(a, b)
		b = a
		a = cost[i]
	}
	return min(a, b)
}

func minCostClimbingStairsv16(cost []int) int {
	a, b := cost[len(cost)-1], 0

	for i := len(cost) - 2; i >= 0; i-- {
		cost[i] += min(a, b)
		b = a
		a = cost[i]
	}

	return min(a, b)
}

func minCostClimbingStairsv15(cost []int) int {
	n := len(cost)
	if n == 0 {
		return 0
	}

	a, b := cost[n-1], 0

	for i := n - 2; i >= 0; i-- {
		cost[i] += min(a, b)
		b = a
		a = cost[i]
	}

	return min(a, b)
}

func minCostClimbingStairsv14(cost []int) int {
	n := len(cost)
	if n == 0 {
		return 0
	}

	a, b := cost[n-1], 0

	for i := n - 2; i >= 0; i-- {
		cost[i] += min(a, b)
		b = a
		a = cost[i]
	}

	return min(a, b)
}

func minCostClimbingStairsv13(cost []int) int {
	if cost == nil {
		return 0
	}

	a, b := cost[len(cost)-1], 0

	for i := len(cost) - 2; i >= 0; i-- {
		cost[i] += min(a, b)
		b = a
		a = cost[i]
	}

	return min(a, b)
}

func minCostClimbingStairsv12(cost []int) int {
	if cost == nil {
		return 0
	}

	a, b := cost[len(cost)-1], 0

	for i := len(cost) - 2; i >= 0; i-- {
		cost[i] += min(a, b)
		b = a
		a = cost[i]
	}

	return min(a, b)
}

func minCostClimbingStairsv11(cost []int) int {
	if cost == nil {
		return 0
	}

	a, b := cost[len(cost)-1], 0

	for i := len(cost) - 2; i >= 0; i-- {
		cost[i] += min(a, b)
		b = a
		a = cost[i]
	}

	return min(a, b)
}

func minCostClimbingStairsv10(cost []int) int {
	n := len(cost)
	if n == 0 {
		return 0
	}

	if n == 1 {
		return cost[n-1]
	}

	a, b := cost[n-1], 0
	for i := n - 2; i >= 0; i-- {
		cost[i] += min(a, b)
		b = a
		a = cost[i]
	}

	return min(a, b)
}

func minCostClimbingStairsv9(cost []int) int {
	n := len(cost)
	if n == 0 {
		return 0
	}

	if n == 1 {
		return cost[0]
	}

	a, b := cost[n-1], 0

	for i := n - 2; i >= 0; i-- {
		cost[i] += min(a, b)
		b = a
		a = cost[i]
	}

	return min(a, b)
}

func minCostClimbingStairsv8(cost []int) int {
	if cost == nil {
		return 0
	}

	if len(cost) == 1 {
		return cost[0]
	}

	a, b := cost[len(cost)-1], 0

	for i := len(cost) - 2; i >= 0; i-- {
		cost[i] = min(cost[i]+a, cost[i]+b)
		b = a
		a = cost[i]

	}

	return min(a, b)
}

func minCostClimbingStairsv7(cost []int) int {

	if len(cost) == 0 {
		return 0
	}

	if len(cost) == 1 {
		return cost[0]
	}

	if len(cost) == 2 {
		return min(cost[0], cost[1])
	}

	cost = append(cost, 0)

	for i := len(cost) - 3; i >= 0; i-- {
		cost[i] += min(cost[i+1], cost[i+2])
	}

	return min(cost[0], cost[1])
}

func minCostClimbingStairsv6(cost []int) int {

	if cost == nil {
		return 0
	}

	if len(cost) == 1 {
		return cost[0]
	}

	if len(cost) == 2 {
		return min(cost[0], cost[1])
	}

	cost = append(cost, 0)

	for i := len(cost) - 3; i >= 0; i-- {
		cost[i] += min(cost[i+1], cost[i+2])
	}

	return min(cost[0], cost[1])
}

func minCostClimbingStairsv5(cost []int) int {
	n := len(cost)

	if n == 0 {
		return 0
	}

	if n == 1 {
		return cost[0]
	}

	cost = append(cost, 0)
	for i := 2; i <= n; i++ {
		cost[i] += min(cost[i-1], cost[i-2])
	}
	return cost[n]
}

func minCostClimbingStairsv4(cost []int) int {

	cost = append(cost, 0)

	for i := len(cost) - 3; i >= 0; i-- {
		cost[i] += min(cost[i+1], cost[i+2])
	}
	return min(cost[0], cost[1])
}

func minCostClimbingStairsv3(cost []int) int {
	cost = append(cost, 0)
	for i := 2; i < len(cost); i++ {
		cost[i] += min(cost[i-1], cost[i-2])
	}
	return cost[len(cost)-1]
}

// ref: youtube Neetcode
func MinCostClimbingStairs(cost []int) int {
	cost = append(cost, 0)
	for i := len(cost) - 3; i >= 0; i-- {
		cost[i] += min(cost[i+1], cost[i+2])
	}
	return min(cost[0], cost[1])
}

func MinCostClimbingStairsV2(cost []int) int {
	cost = append(cost, 0)

	// base case cost[n-1] = 0, cost[n-2] = m
	// so start n-3
	for i := len(cost) - 3; i >= 0; i-- {
		cost[i] += min(cost[i+1], cost[i+2])
	}

	return min(cost[0], cost[1])
}
