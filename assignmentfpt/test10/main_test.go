package main

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	if !reflect.DeepEqual(Solution([]int{3, 2, 1, 1, 2, 3, 1}), 5) {
		t.Fatalf("FAILED TEST 1, Expected %d but result %d", 5, Solution([]int{3, 2, 1, 1, 2, 3, 1}))
	}
	if !reflect.DeepEqual(Solution([]int{4, 1, 4, 1}), 6) {
		t.Fatalf("FAILED TEST 2, Expected %d but result %d", 6, Solution([]int{4, 1, 4, 1}))
	}
	if !reflect.DeepEqual(Solution([]int{3, 3, 3}), 0) {
		t.Fatalf("FAILED TEST 3, Expected %d but result %d", 0, Solution([]int{3, 3, 3}))
	}
}
