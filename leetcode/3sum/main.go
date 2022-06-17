package main

import "fmt"

type flag struct {
	x, y int
}

func threeSum(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	dict := make(map[int]int, len(nums))
	flagDict := make(map[flag]bool)
	for it := 0; it < len(nums); it++ {
		dict[-nums[it]] = it
	}
	var res [][]int
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if val, ok := dict[nums[i]+nums[j]]; ok && val != i && val != j {

				if _, checkExist := flagDict[flag{x: nums[i], y: nums[j]}]; checkExist {
					continue
				}

				res = append(res, []int{nums[i], nums[j], -(nums[i] + nums[j])})
				flagDict[flag{x: nums[i], y: -(nums[i] + nums[j])}] = true
				flagDict[flag{x: -(nums[i] + nums[j]), y: nums[i]}] = true

				flagDict[flag{x: nums[i], y: nums[j]}] = true
				flagDict[flag{x: nums[j], y: nums[i]}] = true

				flagDict[flag{x: nums[j], y: -(nums[i] + nums[j])}] = true
				flagDict[flag{x: -(nums[i] + nums[j]), y: nums[j]}] = true
			}
		}
	}
	return res
}

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
}
