package main

import (
	"math"
)

func maxProduct(nums []int) int {
	maxSub := float64(nums[0])
	minSub :=float64(nums[0])
	maxRes := float64(nums[0])
	for it := 1; it < len(nums);it++{
		temp := math.Max(float64(nums[it]),math.Max(maxSub*float64(nums[it]),minSub*float64(nums[it])))
		minSub = math.Min(float64(nums[it]),math.Min(maxSub*float64(nums[it]),minSub*float64(nums[it])))
		maxSub = temp
		maxRes = math.Max(maxRes,maxSub)
	}
	return int(maxRes)
}
