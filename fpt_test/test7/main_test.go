package main

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	if !reflect.DeepEqual(Solution(5, 8), 2) {
		t.Fatalf("FAILED TEST 1, Expected %d but result %d", 2, Solution(5, 8))
	}
	if !reflect.DeepEqual(Solution(4, 10), 4) {
		t.Fatalf("FAILED TEST 2, Expected %d but result %d", 10, Solution(4, 10))
	}
	if !reflect.DeepEqual(Solution(1, 2), -1) {
		t.Fatalf("FAILED TEST 3, Expected %d but result %d", -1, Solution(1, 2))
	}
	if !reflect.DeepEqual(Solution(10, 5), 1) {
		t.Fatalf("FAILED TEST 4, Expected %d but result %d", 1, Solution(10, 5))
	}
}
