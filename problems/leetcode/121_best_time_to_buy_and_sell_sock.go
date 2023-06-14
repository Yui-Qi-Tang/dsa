package b75

/*
You are given an array prices where prices[i] is the price of a given stock on the ith day.

You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.

Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.



Example 1:
    Input: prices = [7,1,5,3,6,4]
    Output: 5
    Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
    Note that buying on day 2 and selling on day 1 is not allowed because you must buy before you sell.

Example 2:
    Input: prices = [7,6,4,3,1]
    Output: 0
    Explanation: In this case, no transactions are done and the max profit = 0.

Constraints:

1 <= prices.length <= 105
0 <= prices[i] <= 104
*/

func maxProfitV15(prices []int) int {
	result := 0
	buy := prices[0]

	for i := 1; i < len(prices); i++ {
		if buy < prices[i] {
			result = max(result, prices[i]-buy)
		} else {
			buy = prices[i]
		}
	}

	return result
}

func maxProfitV14(prices []int) int {
	result := 0
	buy := prices[0]

	for i := 1; i < len(prices); i++ {
		if buy < prices[i] {
			result = max(result, prices[i]-buy)
		} else {
			buy = prices[i]
		}
	}

	return result
}

func maxProfitV13(prices []int) int {
	buy := prices[0] // choice the first price to buy
	result := 0      // result

	for i := 1; i < len(prices); i++ {
		if buy < prices[i] { // price[i] is greater than buy, so we can sell it
			sell := prices[i] - buy
			result = max(result, sell) // compare the max profit
		} else {
			buy = prices[i] // save the less price
		}
	}

	return result
}

func maxProfitV12(prices []int) int {
	p := prices[0]
	result := 0

	for i := 1; i < len(prices); i++ {
		if p < prices[i] {
			result = max(result, prices[i]-p)
		} else {
			p = prices[i]
		}
	}

	return result
}

func maxProfitV11(prices []int) int {
	n := len(prices)
	if n < 2 {
		return 0
	}

	l := 0
	result := 0
	for r := 1; r < len(prices); r++ {
		if prices[l] > prices[r] {
			l = r
		} else {
			result = max(result, prices[r]-prices[l])
		}
	}

	return result
}

func maxProfitV10(prices []int) int {
	n := len(prices)
	if n < 2 {
		return 0
	}

	l, r := 0, 1
	result := 0
	for r < n {
		if prices[l] > prices[r] {
			l = r
			r++
			continue
		} else {
			result = max(result, prices[r]-prices[l])
			r++
		}
	}

	return result
}

func maxProfitV9(prices []int) int {
	n := len(prices)
	if n < 2 {
		return 0
	}

	l, r := 0, 1

	result := 0
	for r < n {
		if prices[l] > prices[r] {
			l = r
			r++
		} else {
			result = max(result, prices[r]-prices[l])
			r++
		}

	}
	return result
}

func maxProfitV8(prices []int) int {
	n := len(prices)
	if n < 2 {
		return 0
	}

	l, r := 0, 1

	result := 0
	for r < n {
		if prices[l] > prices[r] {
			l = r
			r = l + 1
		} else {
			result = max(result, prices[r]-prices[l])
			r++
		}
	}

	return result
}

func maxProfitV7(prices []int) int {
	n := len(prices)
	if n < 2 {
		return 0
	}
	result := 0

	l, r := 0, 1
	for l < n && r < n {
		if prices[l] > prices[r] {
			l = r
			r = l + 1
		} else {
			result = max(result, prices[r]-prices[l])
			r++
		}

	}

	return result
}

func maxProfitV6(prices []int) int {
	n := len(prices)
	if n < 2 {
		return 0
	}

	l, r := 0, 1

	result := 0

	for l < n && r < n {

		if prices[l] > prices[r] {
			l++
			r = l + 1
		} else {
			result = max(result, prices[r]-prices[l])
			r++
		}

	}

	return result
}

func maxProfitV5(prices []int) int {
	n := len(prices)
	if n < 2 {
		return 0
	}

	maxProfit := 0
	l, r := 0, 1

	// buy small, sell big
	for l < n && r < n {

		if prices[l] > prices[r] {
			l = r
			r = l + 1
			continue
		}

		maxProfit = max(maxProfit, prices[r]-prices[l])
		r++
	}

	return maxProfit
}

func maxProfitV4(prices []int) int {
	n := len(prices)
	if n <= 1 {
		return 0
	}

	result := 0
	l := 0
	for r := 1; r < n; r++ {
		if prices[l] > prices[r] {
			l = r
			continue
		}
		result = max(result, prices[r]-prices[l])
	}

	return result
}

func maxProfitV3(prices []int) int {
	n := len(prices)
	if n <= 1 {
		return 0
	}

	maxp := 0

	l, r := 0, 1

	for l < n && r < n {
		if prices[l] < prices[r] {
			maxp = max(maxp, prices[r]-prices[l])
			r++
		} else {
			l++
			r = l + 1
		}

	}

	return maxp
}

func maxProfitV2(prices []int) int {
	n := len(prices)
	if n == 1 {
		// no transaction
		return 0
	}

	result := 0

	l, r := 0, 1

	for l < n && r < n {
		if prices[l] > prices[r] {
			l = r
			r = l + 1
		} else {
			result = max(result, prices[r]-prices[l])
			r++
		}
	}

	return result
}

func maxProfitV1(prices []int) int {

	if len(prices) <= 1 {
		return 0
	}

	bestBuy, maxProfit := prices[0], 0

	for i := 1; i < len(prices); i++ {
		if bestBuy >= prices[i] {
			bestBuy = prices[i]
		} else {
			if prices[i]-bestBuy > maxProfit {
				maxProfit = prices[i] - bestBuy
			}
		}
	}
	return maxProfit

}
