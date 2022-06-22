package main

import "fmt"

func Solution(N int, a []int, b []int) bool {
	graph := make([][]int, (N+1)*(N+1))
	var tempArray []int
	for i := 0; i <= N; i++ {
		tempArray = make([]int, N+1)
		for j := 0; j <= N; j++ {
			tempArray = append(tempArray, 0)
		}
		graph[i] = tempArray
	}

	for it := 0; it < len(b); it++ {
		graph[a[it]][b[it]] = 1
		graph[b[it]][a[it]] = 1
	}
	for i := 1; i <= N-1; i++ {
		if graph[i][i+1] == 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(Solution(4, []int{1, 2, 4, 4, 3}, []int{2, 3, 1, 3, 1}))
}
