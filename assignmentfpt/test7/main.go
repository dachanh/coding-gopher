package main

import "fmt"

func Solution(N, K int) int {
	sum := (N - 1) * N / 2
	checkUsedGlass := make(map[int]bool, N+1)
	result := 0
	if sum == K {
		return N
	} else if N >= K {
		return 1
	} else if N < K && N == 1 {
		return -1
	} else {
		var fill int
		water := N
		maxContain := K
		for it := 1; it <= N; it++ {
			checkUsedGlass[it] = false
		}
		for maxContain > 0 {
			if maxContain > water {
				fill = water
			} else {
				fill = maxContain
			}
			for it := fill; it > 0; it-- {
				if !checkUsedGlass[it] {
					maxContain -= it
					checkUsedGlass[it] = true
					break
				}
			}
		}
	}
	for it := 1; it <= N; it++ {
		if checkUsedGlass[it] {
			result += 1
		}
	}
	return result
}

func main() {
	fmt.Println(Solution(5, 8))
}
