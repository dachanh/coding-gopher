package main

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	if !reflect.DeepEqual(Solution("abaaba"), 2) {
		t.Fatalf("FAILED TEST 1, Expected %d but result %d", 2, Solution("abaaba"))
	}
	if !reflect.DeepEqual(Solution("zyzyzyz"), 5) {
		t.Fatalf("FAILED TEST 2, Expected %d but result %d", 5, Solution("zyzyzyz"))
	}
	if !reflect.DeepEqual(Solution("aabbbabaaa"), 3) {
		t.Fatalf("FAILED TEST 3, Expected %d but result %d", 3, Solution("aabbbabaaa"))
	}

}
