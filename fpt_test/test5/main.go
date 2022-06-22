package main

import "fmt"

func Solution(A []int) int {
	currSum := 0
	result := 0
	var tempSum int
	for i := 0; i < len(A); i++ {
		tempSum = int((i + 2) * (i + 1) / 2)
		currSum += A[i]
		if currSum == tempSum {
			result += 1
		}

	}
	return result
}

func main() {
	fmt.Println(Solution([]int{1, 3, 4, 2, 5}))
}
