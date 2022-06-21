package main

import "fmt"

func Solution(N int, a []int, b []int) bool {
	graph := make(map[int][]int)
	queue := make(chan int, (N+1)*(N+1))
	visit := make([]bool, N+1)
	result := true
	var nodeCurrent int
	for node := 1; node <= N; node++ {
		graph[node] = []int{}
		visit[node] = false
	}

	for it := 0; it < len(b); it++ {
		graph[a[it]] = append(graph[a[it]], b[it])
		graph[b[it]] = append(graph[b[it]], a[it])
	}
	queue <- 1
	visit[1] = true
	for len(queue) != 0 {
		nodeCurrent = <-queue
		for _, node := range graph[nodeCurrent] {
			if !visit[node] {
				visit[node] = true
			} else {
				queue <- node
			}
		}
	}
	for it := 1; it <= N; it++ {
		if !visit[it] {
			result = false
			return result
		}
	}
	return result
}

func main() {
	fmt.Println(Solution(4, []int{1, 2, 4, 4, 3}, []int{2, 3, 1, 3, 1}))
}
