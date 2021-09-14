package main

import "fmt"

func maxProfit(prices []int) int {
	var res = 0
	for i := 0; i < len(prices)-1; i++ {
		for j := i + 1; j < len(prices); j++ {
			if prices[j]-prices[i] > res {
				res = prices[j] - prices[i]
			}
		}
	}
	return res
}

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}
