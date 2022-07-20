package solutions

func maxProfit(prices []int) int {
	i, j := 0, len(prices)
	if i == j {
		return 0
	}

	for i != j {
		begProf, begMin := profBeg(prices[i : i+2])
		endProf, endMax := profEnd(prices[j-2 : j])
	}

	return 0
}

func profBeg(prices []int) (int, int) {
	profit := prices[1] - prices[0]
	min := min(prices[0], prices[1])
	return profit, min
}

func profEnd(prices []int) (int, int) {
	profit := prices[1] - prices[0]
	max := max(prices[0], prices[1])
	return profit, max
}

//go:inline
func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

//go:inline
func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func maxProfitBruteforce(prices []int) int {
	max := 0
	for i := range prices {
		for j := i + 1; j < len(prices); j++ {
			if prices[j]-prices[i] > max {
				max = prices[j] - prices[i]
			}
		}
	}

	return max
}
