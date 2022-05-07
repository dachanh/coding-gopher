package main

import "math"

func maxSubArray(nums []int) int {
	allNegative := true

	for it := 0 ; it < len(nums);it++{
		if nums[it] > 0{
			allNegative = false
			break
		}
	}

	if allNegative{
		maxNegative := math.MinInt
		for it := 0 ; it < len(nums);it++{
			if maxNegative < nums[it]{
				maxNegative = nums[it]
			}
		}
		return maxNegative
	}

	maxSum := math.MinInt
	maxCurrent := 0
	for it := 0 ; it < len(nums);it++{
		maxCurrent = maxCurrent + nums[it]
		if maxCurrent < 0{
			maxCurrent = 0
		}
		if maxSum < maxCurrent{
			maxSum = maxCurrent
		}
	}
	return maxSum
}
