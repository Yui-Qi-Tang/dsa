package b75

/*
518. Coin Change II
You are given an integer array coins representing coins of different denominations and an integer amount representing a total amount of money.

Return the number of combinations that make up that amount. If that amount of money cannot be made up by any combination of the coins, return 0.

You may assume that you have an infinite number of each kind of coin.

The answer is guaranteed to fit into a signed 32-bit integer.



Example 1:

Input: amount = 5, coins = [1,2,5]
Output: 4
Explanation: there are four ways to make up the amount:
5=5
5=2+2+1
5=2+1+1+1
5=1+1+1+1+1
Example 2:

Input: amount = 3, coins = [2]
Output: 0
Explanation: the amount of 3 cannot be made up just with coins of 2.
Example 3:

Input: amount = 10, coins = [10]
Output: 1


Constraints:

1 <= coins.length <= 300
1 <= coins[i] <= 5000
All the values of coins are unique.
0 <= amount <= 5000
*/

func changev92(amount int, coins []int) int {
	dp := make(map[[2]int]int)
	var backtrace func(i, amt int) int
	backtrace = func(i, amt int) int {
		if i == len(coins) {
			if amt == amount {
				return 1
			}
			return 0
		}

		if amt > amount {
			return 0
		}

		if ans, exist := dp[[2]int{i, amt}]; exist {
			return ans
		}

		dp[[2]int{i, amt}] = backtrace(i, amt+coins[i]) + backtrace(i+1, amt)
		return dp[[2]int{i, amt}]
	}
	return backtrace(0, 0)
}

func changev91(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] += dp[i-coin]
		}
	}
	return dp[amount]
}

func changev90(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] += dp[i-coin]
		}
	}
	return dp[amount]
}

func changev89(amount int, coins []int) int {
	dp := make(map[[2]int]int)
	var backtrace func(i, amt int) int
	backtrace = func(i, amt int) int {
		if i == len(coins) {
			if amt == amount {
				return 1
			}
			return 0
		}

		if amt > amount {
			return 0
		}

		if ans, exist := dp[[2]int{i, amt}]; exist {
			return ans
		}

		dp[[2]int{i, amt}] = backtrace(i+1, amt) + backtrace(i, amt+coins[i])
		return dp[[2]int{i, amt}]
	}
	return backtrace(0, 0)
}

func changev88(amount int, coins []int) int {
	dp := make(map[[2]int]int)
	var backtrace func(i, amt int) int
	backtrace = func(i, amt int) int {
		if i == len(coins) {
			if amount == amt {
				return 1
			}
			return 0
		}

		if amt > amount {
			return 0
		}

		if ans, exist := dp[[2]int{i, amt}]; exist {
			return ans
		}

		dp[[2]int{i, amt}] = backtrace(i, amt+coins[i]) + backtrace(i+1, amt)
		return dp[[2]int{i, amt}]
	}
	return backtrace(0, 0)
}

func changev87(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] += dp[i-coin]
		}
	}

	return dp[amount]
}

func changev86(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1

	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] += dp[i-coin]
		}
	}
	return dp[amount]
}

func changev85(amount int, coins []int) int {
	dp := make(map[int]map[int]int)
	var backtrack func(i, amt int) int
	backtrack = func(i, amt int) int {
		if i == len(coins) {
			if amt == amount {
				return 1
			}
			return 0
		}

		if amt > amount {
			return 0
		}

		if ans, exist := dp[i][amt]; exist {
			return ans
		}

		dp[i] = make(map[int]int)
		dp[i][amt] = backtrack(i+1, amt) + backtrack(i, amt+coins[i])
		return dp[i][amt]
	}
	return backtrack(0, 0)
}

func changev84(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] += dp[i-coin]
		}
	}
	return dp[amount]
}

func changev83(amount int, coins []int) int {
	dp := make(map[int]map[int]int)
	var backtrace func(i, amt int) int
	backtrace = func(i, amt int) int {
		if i == len(coins) {
			if amt == amount {
				return 1
			}
			return 0
		}

		if amt > amount {
			return 0
		}

		if ans, exist := dp[i][amt]; exist {
			return ans
		}

		dp[i] = make(map[int]int)
		dp[i][amt] = backtrace(i, amt+coins[i]) + backtrace(i+1, amt)
		return dp[i][amt]
	}
	return backtrace(0, 0)
}

func changev82(amount int, coins []int) int {
	dp := make(map[int]map[int]int)
	var backtrace func(i, amt int) int
	backtrace = func(i, amt int) int {
		if i == len(coins) {
			if amt == amount {
				return 1
			}
			return 0
		}

		if amt > amount {
			return 0
		}

		dp[i] = make(map[int]int)
		dp[i][amt] = backtrace(i, amt+coins[i]) + backtrace(i+1, amt)
		return dp[i][amt]
	}
	return backtrace(0, 0)
}

func changev81(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] += dp[i-coin]
		}
	}

	return dp[amount]
}

func changev80(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] += dp[i-coin]
		}
	}
	return dp[amount]
}

func changev79(amount int, coins []int) int {
	dp := make(map[int]map[int]int)
	var backtrace func(i, amt int) int
	backtrace = func(i, amt int) int {
		if i == len(coins) {
			if amt == amount {
				return 1
			}
			return 0
		}

		if amt > amount {
			return 0
		}

		if ans, exist := dp[i][amt]; exist {
			return ans
		}

		dp[i] = make(map[int]int)
		dp[i][amt] = backtrace(i, amt+coins[i]) + backtrace(i+1, amt)
		return dp[i][amt]
	}
	return backtrace(0, 0)
}

func changev78(amount int, coins []int) int {
	dp := make(map[int]map[int]int)
	var backtrack func(i, amt int) int
	backtrack = func(i, amt int) int {
		if i == len(coins) {
			if amt == amount {
				return 1
			}
			return 0
		}

		if amt > amount {
			return 0
		}

		dp[i] = make(map[int]int)
		dp[i][amt] = backtrack(i+1, amt) + backtrack(i, amt+coins[i])
		return dp[i][amt]
	}
	return backtrack(0, 0)
}

func changev77(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1

	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] += dp[i-coin]
		}
	}

	return dp[amount]
}

func changev76(amount int, coins []int) int {
	dp := make(map[int]map[int]int)
	var backtrack func(i, amt int) int
	backtrack = func(i, amt int) int {
		if i == len(coins) {
			if amt == amount {
				return 1
			}
			return 0
		}

		if amt > amount {
			return 0
		}

		if ans, exist := dp[i][amt]; exist {
			return ans
		}

		dp[i] = make(map[int]int)
		dp[i][amt] = backtrack(i, amt+coins[i]) + backtrack(i+1, amt)
		return dp[i][amt]
	}
	return backtrack(0, 0)
}

func changev75(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] += dp[i-coin]
		}
	}
	return dp[amount]
}

func changev74(amount int, coins []int) int {
	dp := make(map[int]map[int]int)
	var backtrack func(i, amt int) int
	backtrack = func(i, amt int) int {
		if i == len(coins) {
			if amt == amount {
				return 1
			}
			return 0
		}

		if amt > amount {
			return 0
		}

		if ans, exist := dp[i][amt]; exist {
			return ans
		}

		dp[i] = make(map[int]int)
		dp[i][amt] = backtrack(i+1, amt) + backtrack(i, amt+coins[i])
		return dp[i][amt]
	}
	return backtrack(0, 0)
}

func changev73(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] += dp[i-coin]
		}
	}
	return dp[amount]
}

func changev72(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] += dp[i-coin]
		}
	}

	return dp[amount]
}

func changev71(amonut int, coins []int) int {
	dp := make(map[int]map[int]int)
	var backtrack func(i, amt int) int
	backtrack = func(i int, amt int) int {
		if i == len(coins) {
			if amonut == amt {
				return 1
			}
			return 0
		}

		if amt > amonut {
			return 0
		}

		if ans, exist := dp[i][amt]; exist {
			return ans
		}
		dp[i] = make(map[int]int)
		dp[i][amt] = backtrack(i, amt+coins[i]) + backtrack(i+1, amt)
		return dp[i][amt]
	}
	return backtrack(0, 0)
}

func changev70(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] += dp[i-coin]
		}
	}

	return dp[amount]
}

func changev69(amount int, coins []int) int {
	dp := make(map[int]map[int]int)
	var backtrack func(i, amt int) int
	backtrack = func(i, amt int) int {
		if i == len(coins) {
			if amt == amount {
				return 1
			}
			return 0
		}

		if amt > amount {
			return 0
		}

		if ans, exist := dp[i][amt]; exist {
			return ans
		}
		dp[i] = make(map[int]int)
		dp[i][amt] = backtrack(i+1, amt) + backtrack(i, amt+coins[i])
		return dp[i][amt]
	}
	return backtrack(0, 0)
}

func changev68(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] += dp[i-coin]
		}
	}
	return dp[amount]
}

func changev67(amount int, coins []int) int {
	dp := make(map[int]map[int]int)
	var backtrack func(i, amt int) int
	backtrack = func(i, amt int) int {
		if i == len(coins) {
			if amt == amount {
				return 1
			}
			return 0
		}

		if amt > amount {
			return 0
		}

		if ans, exist := dp[i][amt]; exist {
			return ans
		}

		dp[i] = make(map[int]int)
		dp[i][amt] = backtrack(i, amt+coins[i]) + backtrack(i+1, amt)
		return dp[i][amt]
	}
	return backtrack(0, 0)
}

func changev66(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] += dp[i-coin]
		}
	}

	return dp[amount]
}

func changev65(amount int, coins []int) int {
	dp := make(map[int]map[int]int)
	var backtrack func(i, amt int) int
	backtrack = func(i, amt int) int {
		if i == len(coins) {
			if amt == amount {
				return 1
			}
			return 0
		}

		if amt > amount {
			return 0
		}

		if ans, exist := dp[i][amt]; exist {
			return ans
		}

		dp[i] = make(map[int]int)
		dp[i][amt] = backtrack(i, amt+coins[i]) + backtrack(i+1, amt)
		return dp[i][amt]
	}

	return backtrack(0, 0)
}

func changev64(amount int, coins []int) int {
	dp := make(map[int]map[int]int)
	var backtrace func(i, amt int) int
	backtrace = func(i, amt int) int {
		if i == len(coins) {
			if amt == amount {
				return 1
			}
			return 0
		}

		if amt > amount {
			return 0
		}

		if ans, exist := dp[i][amt]; exist {
			return ans
		}

		dp[i] = make(map[int]int)
		dp[i][amt] = backtrace(i+1, amt) + backtrace(i, amt+coins[i])
		return dp[i][amt]
	}
	return backtrace(0, 0)
}

func changev63(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] += dp[i-coin]
		}
	}
	return dp[amount]
}

func changev62(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] += dp[i-coin]
		}
	}
	return dp[amount]
}

func changev61(amount int, coins []int) int {
	dp := make(map[int]map[int]int)
	var backtrace func(i, amt int) int

	backtrace = func(i, amt int) int {
		if i == len(coins) {
			if amt == amount {
				return 1
			}
			return 0
		}

		if amt > amount {
			return 0
		}

		if ans, exist := dp[i][amt]; exist {
			return ans
		}

		dp[i] = make(map[int]int)
		dp[i][amt] = backtrace(i, amt+coins[i]) + backtrace(i+1, amt)
		return dp[i][amt]
	}

	return backtrace(0, 0)
}

func changev60(amount int, coins []int) int {
	dp := make(map[int]map[int]int)
	var backtrace func(i, amt int) int

	backtrace = func(i, amt int) int {
		if i == len(coins) {
			if amt == amount {
				return 1
			}
			return 0
		}

		if amt > amount {
			return 0
		}

		if ans, exist := dp[i][amt]; exist {
			return ans
		}

		dp[i] = make(map[int]int)
		dp[i][amt] = backtrace(i, amt+coins[i]) + backtrace(i+1, amt)
		return dp[i][amt]
	}

	return backtrace(0, 0)
}

func changev59(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] += dp[i-coin]
		}
	}

	return dp[amount]
}

func changev58(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1

	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] += dp[i-coin]
		}
	}

	return dp[amount]
}

func changev57(amount int, coins []int) int {
	dp := make(map[int]map[int]int)
	var backtrace func(i, amt int) int
	backtrace = func(i, amt int) int {
		if i == len(coins) {
			if amt == amount {
				return 1
			}
			return 0
		}

		if amt > amount {
			return 0
		}

		if ans, exist := dp[i][amt]; exist {
			return ans
		}

		dp[i] = make(map[int]int)
		dp[i][amt] = backtrace(i, amt+coins[i]) + backtrace(i+1, amt)
		return dp[i][amt]
	}
	return backtrace(0, 0)
}

func changev56(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] += dp[i-coin]
		}
	}

	return dp[amount]
}

func changev55(amount int, coins []int) int {
	dp := make(map[int]map[int]int)
	var backtrace func(i, amt int) int
	backtrace = func(i, amt int) int {
		if i == len(coins) {
			if amt == amount {
				return 1
			}
			return 0
		}

		if amt > amount {
			return 0
		}

		if ans, exist := dp[i][amt]; exist {
			return ans
		}

		dp[i] = make(map[int]int)
		dp[i][amt] = backtrace(i, amt+coins[i]) + backtrace(i+1, amt)
		return dp[i][amt]
	}

	return backtrace(0, 0)
}

func changev54(amount int, coins []int) int {
	dp := make(map[int]map[int]int)
	var backtrace func(i, amt int) int
	backtrace = func(i, amt int) int {
		if i == len(coins) {
			if amt == amount {
				return 1
			}
			return 0
		}

		if amt > amount {
			return 0
		}

		if ans, exist := dp[i][amt]; exist {
			return ans
		}

		dp[i] = make(map[int]int)
		dp[i][amt] = backtrace(i, amt+coins[i]) + backtrace(i+1, amt)
		return dp[i][amt]
	}
	return backtrace(0, 0)
}

func changev53(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] += dp[i-coin]
		}
	}
	return dp[amount]
}

func changev52(amount int, coins []int) int {
	dp := make(map[int]map[int]int) // idx:amt -> ways
	var backtrace func(i, amt int) int
	backtrace = func(i, amt int) int {
		if i == len(coins) {
			if amt == amount {
				return 1
			}
			return 0
		}

		if amt > amount {
			return 0
		}

		if ans, exist := dp[i][amt]; exist {
			return ans
		}

		dp[i] = make(map[int]int)
		dp[i][amt] = backtrace(i, amt+coins[i]) + backtrace(i+1, amt)
		return dp[i][amt]
	}
	return backtrace(0, 0)
}

func changev51(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1 // amount = 0 -> only one way

	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] += dp[i-coin]
		}
	}

	return dp[amount]
}

func changev50(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] += dp[i-coin]
		}
	}
	return dp[amount]
}

func changev49(amount int, coins []int) int {
	dp := make(map[int]map[int]int)
	var backtrace func(i, amt int) int
	backtrace = func(i, amt int) int {
		if i == len(coins) {
			if amt == amount {
				return 1
			}
			return 0
		}

		if amt > amount {
			return 0
		}

		if ans, exist := dp[i][amt]; exist {
			return ans
		}

		dp[i] = make(map[int]int)
		dp[i][amt] = backtrace(i+1, amt) + backtrace(i, amt+coins[i])
		return dp[i][amt]
	}
	return backtrace(0, 0)
}

func changev48(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] += dp[i-coin]
		}
	}
	return dp[amount]
}

func changev47(amount int, coins []int) int {
	dp := make(map[int]map[int]int)
	var backtrace func(i, amt int) int
	backtrace = func(i, amt int) int {
		if i == len(coins) {
			if amt == amount {
				return 1
			}
			return 0
		}

		if amt > amount {
			return 0
		}

		if ans, exist := dp[i][amt]; exist {
			return ans
		}

		dp[i] = make(map[int]int)
		dp[i][amt] = backtrace(i+1, amt) + backtrace(i, amt+coins[i])
		return dp[i][amt]
	}
	return backtrace(0, 0)
}

func changev46(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] += dp[i-coin]
		}
	}

	return dp[amount]
}

func changev45(amount int, coins []int) int {
	dp := make(map[int]map[int]int)
	var backtrace func(i, amt int) int
	backtrace = func(i, amt int) int {
		if i == len(coins) {
			if amt == amount {
				return 1
			}
			return 0
		}

		if amt > amount {
			return 0
		}

		if ans, exist := dp[i][amt]; exist {
			return ans
		}

		dp[i] = make(map[int]int)
		dp[i][amt] = backtrace(i, amt+coins[i]) + backtrace(i+1, amt)
		return dp[i][amt]
	}
	return backtrace(0, 0)
}

func changev44(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] += dp[i-coin]
		}
	}

	return dp[amount]
}

func changev43(amount int, coins []int) int {
	dp := make(map[int]map[int]int)
	var backtrace func(i, amt int) int
	backtrace = func(i, amt int) int {
		if i == len(coins) {
			if amt == amount {
				return 1
			}
			return 0
		}

		if amt > amount {
			return 0
		}

		if ans, exist := dp[i][amt]; exist {
			return ans
		}

		dp[i] = make(map[int]int)
		dp[i][amt] = backtrace(i, amt+coins[i]) + backtrace(i+1, amt)
		return dp[i][amt]
	}

	return backtrace(0, 0)
}

func changev42(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1

	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] += dp[i-coin]
		}
	}
	return dp[amount]
}

func changev41(amount int, coins []int) int {
	dp := make(map[int]map[int]int) // idx: amt -> ways
	var backtrace func(i, amt int) int
	backtrace = func(i, amt int) int {
		if i == len(coins) {
			if amt == amount {
				return 1
			}
			return 0
		}

		if amt > amount {
			return 0
		}

		if v, exist := dp[i][amt]; exist {
			return v
		}

		dp[i] = make(map[int]int)
		dp[i][amt] = backtrace(i, amt+coins[i]) + backtrace(i+1, amt)
		return dp[i][amt]
	}
	return backtrace(0, 0)
}

func changev40(amount int, coins []int) int {
	var backtrack func(i, amt int) int
	dp := make(map[int]map[int]int) // idx: amt -> ways
	backtrack = func(i, amt int) int {
		if i == len(coins) {
			if amt == amount {
				return 1
			}
			return 0
		}

		if amt > amount {
			return 0
		}
		if v, exist := dp[i][amt]; exist {
			return v
		}

		dp[i] = make(map[int]int)
		dp[i][amt] = backtrack(i+1, amt) + backtrack(i, coins[i]+amt)
		return dp[i][amt]
	}
	return backtrack(0, 0)
}

func changev39(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] += dp[i-coin]
		}
	}
	return dp[amount]
}

func changev38(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] += dp[i-coin]
		}
	}
	return dp[amount]
}

// top-down
func changev37(amount int, coins []int) int {
	dp := make(map[int]map[int]int) // idx: amt -> ways
	var backtrack func(i, amt int) int
	backtrack = func(i, amt int) int {
		if i == len(coins) {
			if amt == amount {
				return 1
			}
			return 0
		}

		if amt > amount {
			return 0
		}

		if v, exist := dp[i][amt]; exist {
			return v
		}

		dp[i] = make(map[int]int)
		dp[i][amt] = backtrack(i+1, amt) + backtrack(i, amt+coins[i])
		return dp[i][amt]
	}
	return backtrack(0, 0)
}

// bottom up
func changev36(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] += dp[i-coin]
		}
	}

	return dp[amount]
}

// top-down
func changev35(amount int, coins []int) int {
	var dfs func(i, amt int) int
	dp := make(map[int]map[int]int) // idx: amt -> ways
	dfs = func(i, amt int) int {
		if i == len(coins) {
			if amt == amount {
				return 1
			}
			return 0
		}

		if amt > amount {
			return 0
		}

		if v, exist := dp[i][amt]; exist {
			return v
		}
		dp[i] = make(map[int]int)
		dp[i][amt] = dfs(i+1, amt) + dfs(i, amt+coins[i])
		return dp[i][amt]
	}
	return dfs(0, 0)
}

func changev34(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1

	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] += dp[i-coin]
		}
	}
	return dp[amount]
}

// dp: top-down
func changev33(amount int, coins []int) int {
	var dfs func(i, amt int) int
	dp := make(map[int]map[int]int) // idx: amt -> ways

	dfs = func(i, amt int) int {
		// first, check the table
		if ans, exist := dp[i][amt]; exist {
			return ans
		}
		// second, compute the answer
		if i == len(coins) { // picking the coins is done
			if amt == amount {
				return 1
			}
			return 0
		}

		if amt > amount {
			return 0
		}
		dp[i] = make(map[int]int)
		// reuse the same coin + pick next coin
		dp[i][amt] = dfs(i, amt+coins[i]) + dfs(i+1, amt)
		return dp[i][amt]
	}
	return dfs(0, 0)
}

// bottom up
func changev32(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1 // base case: if amount == 0 -> exchange done -> return 1

	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] += dp[i-coin]
		}
	}

	return dp[amount]
}

// top-down
func changev31(amount int, coins []int) int {
	var dfs func(i, amt int) int
	dp := make(map[int]map[int]int)
	dfs = func(i, amt int) int {
		if i == len(coins) {
			if amt == amount {
				return 1
			}
			return 0
		}

		if amt > amount {
			return 0
		}

		if v, exist := dp[i][amt]; exist {
			return v
		}
		dp[i] = make(map[int]int)
		dp[i][amt] = dfs(i, amt+coins[i]) + dfs(i+1, amt)
		return dp[i][amt]
	}
	return dfs(0, 0)
}

// bottom-up
func changev30(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1

	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] += dp[i-coin]
		}
	}
	return dp[amount]
}

// bottom-up
func changev29(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] += dp[i-coin]
		}
	}
	return dp[amount]
}

// top-down
func changev28(amount int, coins []int) int {
	var dfs func(i, amt int) int
	dp := make(map[int]map[int]int)
	dfs = func(i, amt int) int {
		if i == len(coins) {
			if amt == amount {
				return 1
			}
			return 0
		}

		if amt > amount {
			return 0
		}

		if v, exist := dp[i][amt]; exist {
			return v
		}
		dp[i] = make(map[int]int)
		dp[i][amt] = dfs(i, amt+coins[i]) + dfs(i+1, amt)
		return dp[i][amt]
	}

	return dfs(0, 0)
}

// bottom-up
func changev27(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] += dp[i-coin]
		}
	}

	return dp[amount]
}

// top-down
func changev26(amount int, coins []int) int {
	var dfs func(i, amt int) int
	dp := make(map[int]map[int]int)
	dfs = func(i, amt int) int {
		if i == len(coins) {
			if amt == amount {
				return 1
			}
			return 0
		}

		if amt > amount {
			return 0
		}

		if v, exist := dp[i][amt]; exist {
			return v
		}
		dp[i] = make(map[int]int)
		dp[i][amt] = dfs(i+1, amt) + dfs(i, amt+coins[i])
		return dp[i][amt]
	}

	return dfs(0, 0)
}

// top-down
func changev25(amount int, coins []int) int {

	dp := make(map[int]map[int]int) // idx: amt -> ways
	var dfs func(i, amt int) int
	dfs = func(i, amt int) int {
		if i == len(coins) {
			if amt == amount {
				return 1
			} else {
				return 0
			}
		}

		if amt > amount {
			return 0
		}

		if v, ok := dp[i][amt]; ok {
			return v
		}

		dp[i] = make(map[int]int)
		dp[i][amt] = dfs(i, amt+coins[i]) + dfs(i+1, amt)
		return dp[i][amt]
	}

	return dfs(0, 0)
}

// bottom up
func changev24(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1

	for _, coin := range coins {
		for j := coin; j <= amount; j++ {
			dp[j] += dp[j-coin]
		}
	}

	return dp[amount]
}

// bottom-up way
func changev23(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for _, coin := range coins {
		for j := coin; j <= amount; j++ {
			dp[j] += dp[j-coin]
		}
	}
	return dp[amount]
}

// top-down way
func changev22(amount int, coins []int) int {
	dp := make(map[int]map[int]int) // idx: amt -> ways

	var dfs func(i, amt int) int
	dfs = func(i, amt int) int {
		if i == len(coins) {
			if amt == amount {
				return 1
			} else {
				return 0
			}
		}

		if amt > amount {
			return 0
		}

		if v, ok := dp[i][amt]; ok {
			return v
		}

		dp[i] = make(map[int]int)
		dp[i][amt] = dfs(i, amt+coins[i]) + dfs(i+1, amt)
		return dp[i][amt]
	}
	return dfs(0, 0)
}

func changev21(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1

	for _, coin := range coins {
		for j := coin; j <= amount; j++ {
			dp[j] += dp[j-coin]
		}
	}

	return dp[amount]
}

func changev20(amount int, coins []int) int {
	dp := make(map[int]map[int]int) // coin idxe: amt -> ways

	var dfs func(i, amt int) int
	dfs = func(i, amt int) int {
		if i == len(coins) {
			if amt == amount {
				return 1
			} else {
				return 0
			}
		}

		if amt > amount {
			return 0
		}

		if v, exist := dp[i][amt]; exist {
			return v
		}

		dp[i] = make(map[int]int)
		dp[i][amt] = dfs(i, amt+coins[i]) + dfs(i+1, amt)
		return dp[i][amt]
	}

	return dfs(0, 0)
}

func changev19(amount int, coins []int) int {
	dp := make(map[int]map[int]int) // coin idx: amt -> ways

	var dfs func(i, amt int) int

	dfs = func(i, amt int) int {

		if i >= len(coins) {
			if amt == amount {
				return 1
			}
			return 0
		}

		if amt > amount {
			return 0
		}

		if v, exist := dp[i][amt]; exist {
			return v
		}

		dp[i] = make(map[int]int)
		dp[i][amt] = dfs(i, amt+coins[i]) + dfs(i+1, amt)
		return dp[i][amt]

	}

	return dfs(0, 0)
}

func changev18(amount int, coins []int) int {
	dp := make(map[int]map[int]int) // coin idx: amt -> ways

	var dfs func(i, amt int) int
	dfs = func(i, amt int) int {
		if amt == amount {
			return 1
		}
		if amt > amount {
			return 0
		}
		if i >= len(coins) {
			return 0
		}
		if v, exist := dp[i][amt]; exist {
			return v
		}
		dp[i] = make(map[int]int)
		dp[i][amt] = dfs(i, amt+coins[i]) + dfs(i+1, amt)
		return dp[i][amt]
	}

	return dfs(0, 0)
}

func changev17(amount int, coins []int) int {
	dp := make(map[int]map[int]int) // coin index: amt -> ways

	var dfs func(i, amt int) int

	dfs = func(i, amt int) int {

		if amt == amount {
			return 1
		}

		if i >= len(coins) || amt > amount {
			return 0
		}

		if v, exist := dp[i][amt]; exist {
			return v
		}

		dp[i] = map[int]int{} //make(map[int]int)
		dp[i][amt] = dfs(i, amt+coins[i]) + dfs(i+1, amt)
		return dp[i][amt]

	}

	return dfs(0, 0)
}

func changev16(amount int, coins []int) int {
	dp := make(map[int]map[int]int) // coin index: amt -> ways

	var dfs func(i, amt int) int

	dfs = func(i, amt int) int {

		if amt == amount {
			return 1
		}

		if amt > amount {
			return 0
		}

		if i >= len(coins) {
			return 0
		}

		if v, exist := dp[i][amt]; exist {
			return v
		}

		dp[i] = make(map[int]int)
		dp[i][amt] = dfs(i, amt+coins[i]) + dfs(i+1, amt)
		return dp[i][amt]
	}

	return dfs(0, 0)
}

func changev15(amount int, coins []int) int {
	dp := make(map[int]map[int]int) // coin index: amt -> ways

	var dfs func(i, amt int) int

	dfs = func(i, amt int) int {
		if amt == amount {
			return 1
		}

		if amt > amount || i >= len(coins) {
			return 0
		}

		if v, exist := dp[i][amt]; exist {
			return v
		}

		dp[i] = make(map[int]int)
		dp[i][amt] = dfs(i, amt+coins[i]) + dfs(i+1, amt)
		return dp[i][amt]
	}

	return dfs(0, 0)
}

func changev14(amount int, coins []int) int {
	dp := make(map[int]map[int]int) // coin_index: amt -> ways

	var dfs func(i, amt int) int

	dfs = func(i, amt int) int {

		if amt == amount {
			return 1
		}

		if amt > amount {
			return 0
		}

		if i >= len(coins) {
			return 0
		}

		if v, exist := dp[i][amt]; exist {
			return v
		}

		dp[i] = make(map[int]int)
		dp[i][amt] = dfs(i, amt+coins[i]) + dfs(i+1, amt)
		return dp[i][amt]
	}

	return dfs(0, 0)
}

// dfs way
func changev13(amount int, coins []int) int {
	dp := make(map[int]map[int]int) // coin_index: amt -> ways

	var dfs func(i, amt int) int
	dfs = func(i, amt int) int {
		if amt == amount {
			return 1
		}

		if amt > amount || i >= len(coins) {
			return 0
		}

		if v, exist := dp[i][amt]; exist {
			return v
		}

		dp[i] = make(map[int]int)
		dp[i][amt] = dfs(i, amt+coins[i]) + dfs(i+1, amt)
		return dp[i][amt]
	}

	return dfs(0, 0)
}

func changev12(amount int, coins []int) int {

	//fmt.Println("=====>")
	dp := make([]int, amount+1)
	dp[0] = 1
	for _, c := range coins {
		// coin - amount -> ways
		for j := c; j <= amount; j++ {
			// fmt.Println(j, j-c)
			dp[j] += dp[j-c]
		}
	}
	//fmt.Println(dp)
	//fmt.Println("<=====")
	return dp[amount]
}

// dfs way
func changev11(amont int, coins []int) int {
	dp := make(map[int]map[int]int) // coin_index: amount -> ways

	var dfs func(i, amt int) int

	dfs = func(i, amt int) int {

		if amt == amont {
			return 1
		}

		if amt > amont {
			return 0
		}

		if i >= len(coins) {
			return 0
		}

		if v, exist := dp[i][amt]; exist {
			return v
		}

		dp[i] = make(map[int]int)
		dp[i][amt] = dfs(i, amt+coins[i]) + dfs(i+1, amt)
		return dp[i][amt]
	}

	return dfs(0, 0)
}

// dfs way
func changev10(amount int, coins []int) int {

	dp := make(map[int]map[int]int) // coin index: amt -> ways

	var dfs func(i, amt int) int

	dfs = func(i, amt int) int {

		if amt == amount {
			return 1
		}

		if amt > amount {
			return 0
		}

		if i >= len(coins) {
			return 0
		}

		if v, exist := dp[i][amt]; exist {
			return v
		}

		dp[i] = make(map[int]int)
		dp[i][amt] = dfs(i, amt+coins[i]) + dfs(i+1, amt)
		return dp[i][amt]

	}

	return dfs(0, 0)
}

func changev9(amount int, coins []int) int {

	if coins == nil {
		return 0
	}

	// m[coin_index][amt] -> ways
	m := make(map[int]map[int]int)
	var dfs func(i, amt int) int

	dfs = func(i, amt int) int {

		if amt == amount {
			return 1
		}

		if amt > amount {
			return 0
		}

		if i >= len(coins) {
			return 0
		}

		if v, exist := m[i][amt]; exist {
			return v
		}

		m[i] = make(map[int]int)
		m[i][amt] = dfs(i, amt+coins[i]) + dfs(i+1, amt)

		return m[i][amt]

	}

	return dfs(0, 0)

}

func changev8(amount int, coins []int) int {
	dp := make(map[int]map[int]int) // coin_index: amount -> way

	var dfs func(i, amt int) int

	dfs = func(i, amt int) int {

		if amt == amount {
			return 1
		}

		if amt > amount {
			return 0
		}

		if i >= len(coins) {
			return 0
		}

		if v, exist := dp[i][amt]; exist {
			return v
		}

		dp[i] = make(map[int]int)
		dp[i][amt] = dfs(i, amt+coins[i]) + dfs(i+1, amt)
		return dp[i][amt]

	}

	return dfs(0, 0)

}

func changev7(amount int, coins []int) int {
	dp := make(map[int]map[int]int) // coin_index, amt -> ways

	var dfs func(i, amt int) int

	dfs = func(i, amt int) int {
		if amt == amount {
			return 1
		}
		if amt > amount {
			return 0
		}
		if i >= len(coins) {
			return 0
		}

		if v, exist := dp[i][amt]; exist {
			return v
		}

		dp[i] = make(map[int]int)
		dp[i][amt] = dfs(i, amt+coins[i]) + dfs(i+1, amt)
		return dp[i][amt]
	}

	return dfs(0, 0)
}

func changev6(amount int, coins []int) int {
	dp := make(map[int]map[int]int) // coin_index:amount -> ways

	var dfs func(i, amt int) int
	dfs = func(i, amt int) int {
		if amt == amount {
			return 1
		}
		if amt > amount {
			return 0
		}
		if i >= len(coins) {
			return 0
		}

		if v, exsit := dp[i][amt]; exsit {
			return v
		}

		dp[i] = make(map[int]int)
		dp[i][amt] = dfs(i, amt+coins[i]) + dfs(i+1, amt)
		return dp[i][amt]

	}

	return dfs(0, 0)

}

func changev5(amount int, coins []int) int {
	dp := make(map[int]map[int]int) // coin_index:amt -> ways

	var dfs func(i, amt int) int

	dfs = func(i, amt int) int {
		if amt == amount {
			return 1
		}

		if amt > amount {
			return 0
		}

		if i >= len(coins) {
			return 0
		}

		if v, exist := dp[i][amt]; exist {
			return v
		}

		dp[i] = make(map[int]int)
		dp[i][amt] = dfs(i, amt+coins[i]) + dfs(i+1, amt)
		return dp[i][amt]
	}

	return dfs(0, 0)
}

func changev4(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for _, coin := range coins {
		for j := coin; j <= amount; j++ {
			dp[j] += dp[j-coin]
		}
	}

	return dp[amount]
}

func changev3(amount int, coins []int) int {
	dp := make(map[int]map[int]int) // (coin_index:amt) -> way

	var dfs func(i, amt int) int
	dfs = func(i, amt int) int {
		if amt == amount {
			return 1
		}

		if amt > amount {
			return 0
		}
		if i == len(coins) {
			return 0
		}

		if v, exists := dp[i][amt]; exists {
			return v
		}

		dp[i] = make(map[int]int)
		dp[i][amt] = dfs(i, amt+coins[i]) + dfs(i+1, amt)
		return dp[i][amt]

	}

	return dfs(0, 0)
}

func changev2(amount int, coins []int) int {

	dp := make([]int, amount+1)
	dp[0] = 1
	for _, c := range coins {
		for j := c; j <= amount; j++ {
			dp[j] += dp[j-c]
		}
	}
	return dp[amount]
}

func changev1(amount int, coins []int) int {

	// index, amt, counter
	dp := make(map[int]map[int]int)
	var dfs func(i, amt int) int
	dfs = func(i, amt int) int {
		if amt == amount {
			return 1
		}
		if amt > amount {
			return 0
		}
		if i == len(coins) {
			return 0
		}
		if cnt, exist := dp[i][amt]; exist {
			return cnt
		}
		dp[i] = make(map[int]int)
		dp[i][amt] = dfs(i, amt+coins[i]) + dfs(i+1, amt)
		return dp[i][amt]
	}
	return dfs(0, 0)
}
