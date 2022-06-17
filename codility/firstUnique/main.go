package main

import "fmt"

func Solution(A []int) int {
	// write your code in Go 1.4
	var result []int
	dict := make(map[int]bool)
	firstNum := -1
	for it := 0; it < len(A); it++ {
		if _, ok := dict[A[it]]; !ok {
			dict[A[it]] = true
			result = append(result, A[it])

		} else {
			dict[A[it]] = false
		}
	}
	for it := 0; it < len(result); it++ {
		if dict[result[it]] {
			firstNum = result[it]
			break
		}
	}
	return firstNum
}

func main() {
	fmt.Println(Solution([]int{4, 10, 5, 4, 2, 10}))
}
