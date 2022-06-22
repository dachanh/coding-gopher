package main

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	if !reflect.DeepEqual(Solution([]int{3, 1, 2}, []int{2, 3, 1}), true) {
		t.Fatalf("FAILED TEST 1, Expected %v but result %v", true, Solution([]int{3, 1, 2}, []int{2, 3, 1}))
	}
	if !reflect.DeepEqual(Solution([]int{1, 2, 1}, []int{2, 3, 3}), false) {
		t.Fatalf("FAILED TEST 2, Expected %v but result %v", false, Solution([]int{1, 2, 1}, []int{2, 3, 3}))
	}
	if !reflect.DeepEqual(Solution([]int{1, 2, 3, 4}, []int{2, 1, 4, 4}), false) {
		t.Fatalf("FAILED TEST 3, Expected %v but result %v", false, Solution([]int{1, 2, 3, 4}, []int{2, 1, 4, 4}))
	}
	if !reflect.DeepEqual(Solution([]int{1, 2, 3, 4}, []int{2, 1, 4, 4}), false) {
		t.Fatalf("FAILED TEST 4, Expected %v but result %v", false, Solution([]int{1, 2, 3, 4}, []int{2, 1, 4, 4}))
	}
}
