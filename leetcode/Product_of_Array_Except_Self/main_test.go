package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test(t *testing.T){
	if !reflect.DeepEqual(productExceptSelf([]int{1,2,3,4}),[]int{24,12,8,6}) {
		t.Fatal("FAILED TEST 1")
	}
	if !reflect.DeepEqual(productExceptSelf([]int{-1,1,0,-3,3}),[]int{0,0,9,0,0}) {
		t.Fatal("FAILED TEST 2")
	}
	if !reflect.DeepEqual(productExceptSelf([]int{0,1,0,-3,3}),[]int{0,0,0,0,0}) {
		t.Fatal("FAILED TEST 3")
	}
	fmt.Println("COMPLETED")
}
