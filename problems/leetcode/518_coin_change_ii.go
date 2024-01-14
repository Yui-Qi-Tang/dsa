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