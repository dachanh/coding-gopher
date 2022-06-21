package main

import "fmt"

func Solution(A, B []int) bool {
	N := len(A)
	graph := make(map[int][]int)
	queue := make(chan int, (N+1)*(N+1))
	visit := make([]bool, N+1)
	var nodeCurrent int
	result := false
	for node := 1; node <= N; node++ {
		graph[node] = []int{}
		visit[node] = false

	}
	for it := 0; it < len(B); it++ {
		graph[A[it]] = append(graph[A[it]], B[it])
	}

	queue <- 1
	visit[1] = true
	for len(queue) != 0 {
		nodeCurrent = <-queue
		for _, node := range graph[nodeCurrent] {
			if !visit[node] {
				visit[node] = true
				queue <- node
			} else {
				if node == 1 {
					result = true
					close(queue)
					break
				}
			}
		}
	}

	return result
}

func main() {
	fmt.Println(Solution([]int{1, 2, 3, 4}, []int{2, 3, 4, 2}))
}
