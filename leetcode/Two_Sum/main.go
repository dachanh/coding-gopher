package main

import "fmt"

func twoSum(nums []int, target int) []int {
	var res []int
	m := make(map[int]int, len(nums))

	for i, v := range nums {

		m[v] = i
	}
	for i, v := range nums {
		if nums[m[target-v]]+v == target && m[target-v] != i {
			res = []int{i, m[target-v]}
			break
		}
	}
	return res
}

func main() {
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
}
