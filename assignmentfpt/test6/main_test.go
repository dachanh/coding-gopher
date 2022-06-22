package main

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	if !reflect.DeepEqual(Solution("nice", "nicer"), "ADD r") {
		t.Fatalf("FAILED TEST 1, Expected %v but result %v", "ADD r", Solution("nice", "nicer"))
	}
	if !reflect.DeepEqual(Solution("test", "tent"), "CHANGE s n") {
		t.Fatalf("FAILED TEST 1, Expected %v but result %v", "CHANGE s n", Solution("test", "tent"))
	}
	if !reflect.DeepEqual(Solution("beans", "banes"), "MOVE e") {
		t.Fatalf("FAILED TEST 1, Expected %v but result %v", "MOVE e", Solution("beans", "banes"))
	}
	if !reflect.DeepEqual(Solution("o", "odd"), "IMPOSSIBLE") {
		t.Fatalf("FAILED TEST 1, Expected %v but result %v", "IMPOSSIBLE", Solution("o", "odd"))
	}
}
