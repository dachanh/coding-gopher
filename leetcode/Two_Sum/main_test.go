package main

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	if !reflect.DeepEqual(twoSum([]int{7, 2, 3, 1}, 9), []int{0, 1}) {
		t.Fatal("Error test 1")
	}
	if !reflect.DeepEqual(twoSum([]int{3, 3}, 6), []int{0, 1}) {
		t.Fatal("Error test 2", twoSum([]int{3, 3}, 6))

	}
	if !reflect.DeepEqual(twoSum([]int{3, 2, 4}, 6), []int{1, 2}) {
		t.Fatal("Error test 3", twoSum([]int{3, 2, 4}, 6))
	}
}
