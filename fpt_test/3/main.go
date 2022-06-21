package main

import (
	"fmt"
	"math"
	"strconv"
)

func Solution(S string) int {
	nums := make([]int, 0)
	sum := 0
	for i := 0; i < len(S); i++ {
		intVar, err := strconv.Atoi(string(S[i]))
		if err == nil {
			nums = append(nums, intVar)
		}
	}
	for it := 0; it < len(nums); it++ {
		sum = sum + nums[it]
	}
	result := 0
	if math.Mod(float64(sum), 3.0) == 0 {
		result = 1
	}
	for it := 0; it < len(nums); it++ {
		tempSum := sum - nums[it]
		for j := 0; j < 10; j++ {
			if math.Mod(float64(tempSum+j), 3.0) == 0 && nums[it] != j {
				result += 1
			}
		}
	}
	return result
}

func main() {
	fmt.Println(Solution("022"))
}
