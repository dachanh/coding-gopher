package main

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
