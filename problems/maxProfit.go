package problems

func MaxProfit(prices []int) int {

	var profit int = 0
	var minPrice int = prices[0]
	var n int = len(prices)

	for i := 1; i < n; i++ {
		minPrice = min(minPrice, prices[i])
		profit = max(profit, prices[i]-minPrice)
	}

	return profit

}
