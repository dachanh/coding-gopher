package main

import "fmt"

func maxProfit(prices []int) int {
	var minVar = int(1e4)
	var res = 0
	for i := 0; i < len(prices); i++ {
		if prices[i] <= minVar {
			minVar = prices[i]
		} else if prices[i]-minVar > res {
			res = prices[i] - minVar
		}
	}
	return res
}

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}
