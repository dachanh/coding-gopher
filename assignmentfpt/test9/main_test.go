package main

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	if !reflect.DeepEqual(Solution("23"), 7) {
		t.Fatalf("FAILED TEST 1, Expected %d but result %d", 5, Solution("23"))
	}
	if !reflect.DeepEqual(Solution("0081"), 11) {
		t.Fatalf("FAILED TEST 2, Expected %d but result %d", 11, Solution("0081"))
	}
	if !reflect.DeepEqual(Solution("022"), 9) {
		t.Fatalf("FAILED TEST 3, Expected %d but result %d", 9, Solution("022"))
	}
}
