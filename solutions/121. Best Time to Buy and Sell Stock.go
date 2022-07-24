package solutions

// Tags: Array, Dynamic Programming

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	min, max, best := prices[0], prices[0], 0

	for i := 1; i < len(prices)-1; i++ {
		// if local minimum
		if prices[i-1] >= prices[i] && prices[i] < prices[i+1] {
			if prices[i] < min {
				min = prices[i]
				max = prices[i]
			}

			// else if local maximum
		} else if prices[i-1] <= prices[i] && prices[i] > prices[i+1] {
			max = prices[i]
			b := max - min
			if b > best {
				best = b
			}
		}
	}
	if prices[len(prices)-1] > max {
		max = prices[len(prices)-1]
		if max-min > best {
			return max - min
		}
	}

	return best
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

//func maxProfitNotWorkingChunks(prices []int) int {
//
//	//minI, maxI := 0, 0
//	//dp := make([]int, 0, len(prices))
//	//
//	//for i := 1; i < len(prices); i++ {
//	//	// if рост
//	//	if prices[i-1] > prices[i] {
//	//		if prices[i] < min {
//	//			min = prices[i]
//	//			minI = i
//	//		}
//	//
//	//		// else if падение
//	//	} else if prices[i-1] < prices[i] {
//	//		if prices[i] > max {
//	//			max = prices[i]
//	//			maxI = i
//	//		}
//	//	}
//	//	if maxI > minI {
//	//		b := max - min
//	//		if b > best {
//	//			dp = append(dp, b)
//	//			best = b
//	//		}
//	//	}
//
//	/*		// if local minimum
//				if prices[i-1] > prices[i] && prices[i] < prices[i+1] {
//					min = prices[i]
//
//					// else if local maximum
//				} else if prices[i-1] < prices[i] && prices[i] > prices[i+1] {
//					max = prices[i]
//					b := max - min
//					if b > best {
//						best = b
//					}
//				}
//	}*/
//
//	/*
//		minI, maxI, best := 0, 1, 0
//
//		for i := 1; i < len(prices); i++ {
//			if minI < maxI && prices[i-1] < prices[minI] {
//				minI = i - 1
//			}
//			if prices[i] > prices[maxI] {
//				maxI = i
//			}
//			if prices[maxI]-prices[minI] > best {
//				best = prices[maxI] - prices[minI]
//			}
//		}
//	*/
//	//fmt.Println("<=", dp, ">=")
//}
//
//
//func maxProfitNOTWORK_2(prices []int) int {
//	//min, max, best := 1000000, -1, 0
//
//	/*	n := len(prices) * 2
//		for i := 0; i < n; i++ {
//			if maxI < minI {
//
//				maxI++
//			} else if maxI > minI {
//
//				minI++
//			}
//
//			// to do - don't forget >= and <=
//			if maxI < minI {
//				if prices[minI] < prices[minI-1] {
//					minI++
//				}
//			}
//			if prices[i] > max && maxI < minI {
//				max = prices[i]
//				maxI = i
//			}
//			b := max - min
//			if b > best {
//				best = b
//			}
//		}*/
//
//	//return best
//	return 0
//}
