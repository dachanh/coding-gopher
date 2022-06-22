package main

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	if !reflect.DeepEqual(Solution(4, []int{1, 2, 4, 4, 3}, []int{2, 3, 1, 3, 1}), true) {
		t.Fatalf("FAILED TEST 1, Expected %v but result %v", true, Solution(4, []int{1, 2, 4, 4, 3}, []int{2, 3, 1, 3, 1}))
	}
	if !reflect.DeepEqual(Solution(4, []int{1, 2, 1, 3}, []int{2, 4, 3, 4}), false) {
		t.Fatalf("FAILED TEST 2, Expected %v but result %v", false, Solution(4, []int{1, 2, 1, 3}, []int{2, 4, 3, 4}))
	}
	if !reflect.DeepEqual(Solution(6, []int{2, 4, 5, 3}, []int{3, 5, 6, 4}), false) {
		t.Fatalf("FAILED TEST 3, Expected %v but result %v", false, Solution(6, []int{2, 4, 5, 3}, []int{3, 5, 6, 4}))
	}
	if !reflect.DeepEqual(Solution(3, []int{1, 3}, []int{2, 2}), true) {
		t.Fatalf("FAILED TEST 4, Expected %v but result %v", true, Solution(3, []int{1, 3}, []int{2, 2}))
	}
}
