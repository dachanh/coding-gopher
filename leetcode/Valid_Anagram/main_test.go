package main

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	if !reflect.DeepEqual(isAnagram("anagram","nagaram"),true){
		t.Fatal("Error test 1")
	}
	if !reflect.DeepEqual(isAnagram("rat","car"),false){
		t.Fatal("Error test 1")
	}
	t.Fatal("PASSED")
}