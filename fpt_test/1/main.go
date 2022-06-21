package main

import (
	"math"
)

func Solution(A []int) int {
	dictNum := make(map[int]int)
	uniqueNum := make([]int, 0)
	sum := 0
	for _, element := range A {
		sum += element
		if _, ok := dictNum[element]; !ok {
			dictNum[element] = 1
			uniqueNum = append(uniqueNum, element)
		} else {
			dictNum[element] += 1
		}
	}
	if len(uniqueNum) == 1 {
		return 0
	}
	result := math.MaxInt
	for i := 0; i < len(uniqueNum); i++ {
		tempCost := 0
		for j := 0; j < len(uniqueNum); j++ {
			if uniqueNum[i] < uniqueNum[j] {
				tempCost = tempCost + dictNum[uniqueNum[j]]*(uniqueNum[j]-uniqueNum[i])
			} else {
				tempCost = tempCost + dictNum[uniqueNum[j]]*(uniqueNum[i]-uniqueNum[j])
			}
		}
		if result > tempCost {
			result = tempCost
		}
	}
	return result
}
