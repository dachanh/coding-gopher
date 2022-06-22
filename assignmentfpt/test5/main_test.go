package main

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	if !reflect.DeepEqual(Solution([]int{2, 1, 3, 5, 4}), 3) {
		t.Fatalf("FAILED TEST 1, Expected %v but result %v", 3, Solution([]int{2, 1, 3, 5, 4}))
	}
	if !reflect.DeepEqual(Solution([]int{2, 3, 4, 1, 5}), 2) {
		t.Fatalf("FAILED TEST 1, Expected %v but result %v", 2, Solution([]int{2, 3, 4, 1, 5}))
	}
	if !reflect.DeepEqual(Solution([]int{1, 3, 4, 2, 5}), 3) {
		t.Fatalf("FAILED TEST 1, Expected %v but result %v", 3, Solution([]int{1, 3, 4, 2, 5}))
	}
}
