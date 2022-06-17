package main

func Solution(A []int) int {
	if len(A) == 0 {
		return -1
	}
	prefixSum := make([]int, len(A))
	prefixSum[0] = A[0]
	for it := 1; it < len(A); it++ {
		prefixSum[it] = prefixSum[it-1] + A[it]
	}
	result := 0
	for it := 0; it < len(A); it++ {
		if A[it] == 0 {
			result += prefixSum[len(A)-1] - prefixSum[it]
		}
		if result > 1000000000 {
			return -1
		}
	}
	return result
}
