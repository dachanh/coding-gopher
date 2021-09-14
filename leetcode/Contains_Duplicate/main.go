package main

import "fmt"

func containsDuplicate(nums []int) bool {
	var m = make(map[int]bool, len(nums))
	for _, n := range nums {
		if _, ok := m[n]; ok {
			return true
		}
		m[n] = true
	}
	return false
}

func main() {
	fmt.Println(containsDuplicate([]int{1, 2, 3}))
}
