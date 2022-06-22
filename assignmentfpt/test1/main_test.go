package main

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	if !reflect.DeepEqual(Solution(14), 19) {
		t.Fatalf("FAILED TEST 1, Expected %d but result %d", 19, Solution(14))
	}
	if !reflect.DeepEqual(Solution(10), 11) {
		t.Fatalf("FAILED TEST 2, Expected %d but result %d", 11, Solution(10))
	}
	if !reflect.DeepEqual(Solution(99), 9999) {
		t.Fatalf("FAILED TEST 3, Expected %d but result %d", 9999, Solution(99))
	}
}
