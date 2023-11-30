package b75

import (
	"math"
)

/*
322. Coin Change
Medium
You are given an integer array coins representing coins of different denominations and an integer amount representing a total amount of money.

Return the fewest number of coins that you need to make up that amount. If that amount of money cannot be made up by any combination of the coins, return -1.

You may assume that you have an infinite number of each kind of coin.



Example 1:

Input: coins = [1,2,5], amount = 11
Output: 3
Explanation: 11 = 5 + 5 + 1
Example 2:

Input: coins = [2], amount = 3
Output: -1
Example 3:

Input: coins = [1], amount = 0
Output: 0


Constraints:

1 <= coins.length <= 12
1 <= coins[i] <= 231 - 1
0 <= amount <= 104
*/

func coinChangev32(coins []int, amount int) int {
	if len(coins) == 0 {
		return 0
	}
	dp := make([]int, amount+1)

	for i := 1; i <= amount; i++ {
		q := math.MaxInt
		for _, coin := range coins {
			if i-coin >= 0 && dp[i-coin] != -1 {
				q = min(q, 1+dp[i-coin])
			}
		}

		if q == math.MaxInt {
			q = -1
		}
		dp[i] = q
	}
	return dp[amount]
}

func coinChangev31(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	dp := make([]int, amount+1)

	for i := 1; i <= amount; i++ {
		q := math.MaxInt
		for _, coin := range coins {
			if i-coin >= 0 && dp[i-coin] != -1 {
				q = min(q, 1+dp[i-coin])
			}
		}

		if q == math.MaxInt {
			q = -1
		}
		dp[i] = q
	}

	return dp[amount]
}

func coinChangev30(coins []int, amount int) int {
	if amount < 1 {
		return 0
	}

	dp := make([]int, amount+1)

	for i := 1; i <= amount; i++ {
		q := math.MaxInt
		for _, coin := range coins {
			if i-coin >= 0 && dp[i-coin] != -1 {
				q = min(q, 1+dp[i-coin])
			}
		}

		if q == math.MaxInt {
			q = -1
		}

		dp[i] = q
	}

	return dp[amount]
}

func coinChangev29(coins []int, amount int) int {
	if amount < 1 {
		return 0
	}

	dp := make([]int, amount+1)

	for i := 1; i <= amount; i++ {
		q := math.MaxInt

		for _, coin := range coins {
			if i-coin >= 0 && dp[i-coin] != -1 {
				q = min(q, 1+dp[i-coin])
			}
		}

		if q == math.MaxInt {
			q = -1
		}

		dp[i] = q
	}

	return dp[amount]
}

func coinChangev28(coins []int, amount int) int {
	if amount < 1 {
		return 0
	}

	dp := make([]int, amount+1)

	for i := 1; i <= amount; i++ {
		q := math.MaxInt
		for _, coin := range coins {
			if i-coin >= 0 && dp[i-coin] != -1 {
				q = min(q, 1+dp[i-coin])
			}
		}

		if q == math.MaxInt {
			q = -1
		}

		dp[i] = q
	}

	return dp[amount]
}

func coinChangev27(coins []int, amount int) int {
	if amount < 1 {
		return 0
	}

	dp := make([]int, amount+1)

	for i := 1; i <= amount; i++ {
		q := math.MaxInt
		for _, coin := range coins {
			if i-coin >= 0 && dp[i-coin] != -1 {
				q = min(q, 1+dp[i-coin])
			}
		}

		if q == math.MaxInt {
			q = -1
		}
		dp[i] = q
	}

	return dp[amount]
}

func coinChangev26(coins []int, amount int) int {
	if amount < 1 {
		return 0
	}

	dp := make([]int, amount+1)

	for i := 1; i <= amount; i++ {
		q := math.MaxInt
		for _, coin := range coins {
			if i-coin >= 0 && dp[i-coin] != -1 {
				q = min(q, 1+dp[i-coin])
			}
		}

		if q == math.MaxInt {
			q = -1
		}
		dp[i] = q
	}

	return dp[amount]
}

func coinChangev25(coins []int, amount int) int {
	if amount < 1 {
		return 0
	}

	dp := make([]int, amount+1)

	for i := 1; i <= amount; i++ {
		q := math.MaxInt
		for _, coin := range coins {
			if i-coin >= 0 && dp[i-coin] != -1 {
				q = min(q, 1+dp[i-coin])
			}
		}

		if q == math.MaxInt {
			q = -1
		}
		dp[i] = q
	}

	return dp[amount]
}

func coinChangev24(coins []int, amount int) int {
	if amount < 1 {
		return 0
	}

	dp := make([]int, amount+1)

	for i := 1; i <= amount; i++ {
		q := math.MaxInt
		for _, coin := range coins {
			if i-coin >= 0 && dp[i-coin] != -1 {
				q = min(q, 1+dp[i-coin])
			}
		}

		if q == math.MaxInt {
			q = -1
		}
		dp[i] = q
	}

	return dp[amount]
}

func coinChangev23(coins []int, amount int) int {
	if amount < 1 {
		return 0
	}

	dp := make([]int, amount+1)

	for i := 1; i <= amount; i++ {
		q := math.MaxInt
		for _, coin := range coins {
			if i-coin >= 0 && dp[i-coin] != -1 {
				q = min(q, 1+dp[i-coin])
			}
		}

		if q == math.MaxInt {
			q = -1
		}

		dp[i] = q
	}

	return dp[amount]
}

func coinChangev22(coins []int, amount int) int {
	if amount < 1 {
		return 0
	}

	dp := make([]int, amount+1)

	for i := 1; i <= amount; i++ {
		q := math.MaxInt
		for _, coin := range coins {
			if i-coin >= 0 && dp[i-coin] != -1 {
				q = min(q, 1+dp[i-coin])
			}
		}

		if q == math.MaxInt {
			q = -1
		}
		dp[i] = q

	}

	return dp[amount]
}

func coinChangev21(coins []int, amount int) int {

	if amount < 1 {
		return 0
	}

	dp := make([]int, amount+1)

	for i := 1; i <= amount; i++ {
		q := math.MaxInt
		for _, coin := range coins {
			if i-coin >= 0 && dp[i-coin] != -1 {
				q = min(q, 1+dp[i-coin])
			}
		}

		if q == math.MaxInt {
			q = -1
		}

		dp[i] = q
	}

	return dp[amount]
}

func coinChangev20(coins []int, amount int) int {
	if amount < 1 {
		return 0
	}

	dp := make([]int, amount+1)

	for i := 1; i <= amount; i++ {
		q := math.MaxInt
		for _, coin := range coins {
			if i-coin >= 0 && dp[i-coin] != -1 {
				q = min(q, 1+dp[i-coin])
			}
		}

		if q == math.MaxInt {
			q = -1
		}
		dp[i] = q
	}

	return dp[amount]
}

func coinChangev19(coins []int, amount int) int {
	dp := make([]int, amount+1)

	for i := 1; i <= amount; i++ {
		q := math.MaxInt
		for _, coin := range coins {
			if i-coin >= 0 && dp[i-coin] != -1 {
				q = min(q, 1+dp[i-coin])
			}
		}

		if q == math.MaxInt {
			q = -1
		}
		dp[i] = q
	}

	return dp[amount]
}

func coinChangev18(coins []int, amount int) int {
	dp := make([]int, amount+1)

	for i := 1; i <= amount; i++ {
		q := math.MaxInt
		for _, coin := range coins {
			if i-coin >= 0 && dp[i-coin] != -1 {
				q = min(q, 1+dp[i-coin])
			}
		}

		if q == math.MaxInt {
			q = -1
		}

		dp[i] = q
	}

	return dp[amount]
}

func coinChangev17(coins []int, amount int) int {
	dp := make([]int, amount+1)

	for i := 1; i <= amount; i++ {
		q := math.MaxInt
		for _, coin := range coins {
			if i-coin >= 0 && dp[i-coin] != -1 {
				q = min(q, 1+dp[i-coin])
			}
		}

		if q == math.MaxInt {
			q = -1
		}

		dp[i] = q
	}

	return dp[amount]
}

func coinChangev16(coins []int, amount int) int {

	dp := make([]int, amount+1)

	for i := 1; i <= amount; i++ {
		q := math.MaxInt
		for _, coin := range coins {
			if i-coin >= 0 && dp[i-coin] != -1 {
				q = min(q, 1+dp[i-coin])
			}
		}

		if q == math.MaxInt {
			q = -1
		}

		dp[i] = q
	}

	return dp[amount]

}

func coinChangev15(coins []int, amount int) int {
	dp := make([]int, amount+1)

	for i := 1; i <= amount; i++ {
		q := math.MaxInt
		for _, coin := range coins {
			if i-coin >= 0 && dp[i-coin] != -1 {
				q = min(q, 1+dp[i-coin])
			}
		}

		if q == math.MaxInt {
			q = -1
		}

		dp[i] = q
	}

	return dp[amount]
}

func coinChangev14(coins []int, amount int) int {
	dp := make([]int, amount+1)

	for i := 1; i <= amount; i++ {
		q := math.MaxInt

		for _, coin := range coins {
			if i-coin >= 0 && dp[i-coin] != -1 {
				q = min(q, 1+dp[i-coin])
			}
		}

		if q == math.MaxInt {
			q = -1
		}

		dp[i] = q
	}

	return dp[amount]
}

func coinChangev13(coins []int, amount int) int {
	if coins == nil {
		return 0
	}

	dp := make([]int, amount+1)

	for i := 1; i <= amount; i++ {
		q := math.MaxInt
		for _, coin := range coins {
			if i-coin >= 0 && dp[i-coin] != -1 {
				q = min(q, 1+dp[i-coin])
			}
		}

		if q == math.MaxInt {
			q = -1
		}
		dp[i] = q

	}

	return dp[amount]
}

func coinChangev12(coins []int, amount int) int {
	dp := make([]int, amount+1)

	for i := 1; i <= amount; i++ {
		q := math.MaxInt
		for _, coin := range coins {
			if i-coin >= 0 && dp[i-coin] != -1 {
				q = min(q, 1+dp[i-coin])
			}
		}

		if q == math.MaxInt {
			q = -1
		}

		dp[i] = q
	}

	return dp[amount]
}

func coinChangev11(coins []int, amount int) int {

	dp := make([]int, amount+1)

	for i := 1; i <= amount; i++ {
		q := math.MaxInt

		for _, coin := range coins {
			if i-coin >= 0 && dp[i-coin] != -1 {
				q = min(q, 1+dp[i-coin])
			}
		}

		if q == math.MaxInt {
			q = -1
		}
		dp[i] = q

	}

	return dp[amount]

}

func cofinChangev10(coins []int, amount int) int {
	dp := make([]int, amount+1)

	for i := 1; i <= amount; i++ {

		q := math.MaxInt

		for _, coin := range coins {
			if i-coin >= 0 && dp[i-coin] != -1 {
				q = min(q, 1+dp[i-coin])
			}
		}

		if q == math.MaxInt {
			q = -1
		}

		dp[i] = q

	}

	return dp[amount]
}

func coinChangev9(coins []int, amount int) int {
	dp := make([]int, amount+1) // amount -> min. coin ways

	for i := 1; i <= amount; i++ {
		q := math.MaxInt
		for j := 0; j < len(coins); j++ {

			if i-coins[j] >= 0 && dp[i-coins[j]] != -1 {
				q = min(q, 1+dp[i-coins[j]])
			}
		}

		if q == math.MaxInt {
			q = -1
		}

		dp[i] = q
	}

	return dp[amount]
}

func coinChangev8(coins []int, amount int) int {
	dp := make([]int, amount+1)

	for i := 1; i <= amount; i++ {
		q := math.MaxInt
		for j := 0; j < len(coins); j++ {
			if i-coins[j] >= 0 && dp[i-coins[j]] != -1 {
				q = min(q, 1+dp[i-coins[j]])
			}
		}

		if q == math.MaxInt {
			q = -1
		}

		dp[i] = q
	}

	return dp[amount]
}

func coinChangev7(coins []int, amount int) int {
	dp := make([]int, amount+1)

	for i := 1; i <= amount; i++ {
		q := math.MaxInt

		for j := 0; j < len(coins); j++ {
			if i-coins[j] >= 0 && dp[i-coins[j]] != -1 {
				q = min(q, 1+dp[i-coins[j]])
			}
		}
		if q == math.MaxInt {
			q = -1
		}
		dp[i] = q
	}
	return dp[amount]
}

func coinChangev6(coins []int, amount int) int {
	m := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		q := math.MaxInt
		for _, coin := range coins {
			if i-coin >= 0 && m[i-coin] != -1 {
				q = min(q, 1+m[i-coin])
			}
		}
		if q == math.MaxInt {
			q = -1
		}
		m[i] = q
	}
	return m[amount]

}

func coinChangev5(coins []int, amount int) int {
	m := make([]int, amount+1)
	// m[0] = 0 -> coin = 0, it will return 0
	for i := 1; i <= amount; i++ {
		q := math.MaxInt
		for _, coin := range coins {
			if i-coin >= 0 && m[i-coin] != -1 {
				q = min(q, 1+m[i-coin])
			}
		}

		if q == math.MaxInt {
			q = -1
		}

		m[i] = q
	}
	return m[amount]
}

func coinChangev4(coins []int, amount int) int {

	m := make([]int, amount+1)

	for i := 1; i <= amount; i++ {
		q := math.MaxInt
		for _, coin := range coins {
			if i-coin >= 0 && m[i-coin] != -1 {
				// add 1 for i-coin
				q = min(q, 1+m[i-coin])
			}
		}

		if q == math.MaxInt {
			q = -1
		}

		m[i] = q
	}

	return m[amount]

}

func coinChangev3(coins []int, amount int) int {
	m := make([]int, amount+1)

	for i := 1; i <= amount; i++ {
		q := math.MaxInt
		for _, coin := range coins {
			// i-coin == 0 -> change complete; i-coin > 0 -> should change again
			if i-coin >= 0 && m[i-coin] != -1 {
				q = min(q, 1+m[i-coin])
			}
		}
		if q == math.MaxInt {
			q = -1
		}
		m[i] = q
	}
	return m[amount]
}

func CoinChange(coins []int, amount int) int {

	m := make([]int, amount+1)

	for i := 1; i <= amount; i++ {
		q := math.MaxInt

		for _, coin := range coins {
			if i-coin >= 0 && m[i-coin] != -1 {
				q = min(q, m[i-coin]+1)
			}
		}

		if q == math.MaxInt {
			m[i] = -1
		} else {
			m[i] = q
		}
	}

	return m[amount]

}

func coinChange(coins []int, amount int) int {

	m := make([]int, amount+1)

	for i := 1; i <= amount; i++ {
		q := math.MaxInt
		for _, coin := range coins {
			if i-coin >= 0 && m[i-coin] != -1 {
				q = min(q, 1+m[i-coin])
			}
		}
		if q != math.MaxInt {
			m[i] = q
		} else {
			m[i] = -1
		}
	}

	return m[amount]
}

func coinChangev2(coins []int, amount int) int {

	m := make([]int, amount+1)

	for i := 1; i <= amount; i++ {

		q := math.MaxInt
		for _, coin := range coins {
			if i-coin > -1 && m[i-coin] != -1 {
				q = min(q, 1+m[i-coin])
			}
		}

		if q == math.MaxInt {
			m[i] = -1
		} else {
			m[i] = q
		}

	}

	return m[amount]

}
