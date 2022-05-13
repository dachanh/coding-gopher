package main

import (
	"reflect"
	"testing"
)

func Test(t *testing.T){
	if !reflect.DeepEqual(maxProduct([]int{2,3,-2,4}),6){
		t.Fatal("Failed Test 1")
	}
	if !reflect.DeepEqual(maxProduct([]int{2,0,1}),2){
		t.Fatal("Failed Test 2")
	}
	if !reflect.DeepEqual(maxProduct([]int{-2,0,-1}),0){
		t.Fatal("Failed Test 3")
	}
	if !reflect.DeepEqual(maxProduct([]int{2,-3,-2}),12){
		t.Fatal("Failed Test 4")
	}
	if !reflect.DeepEqual(maxProduct([]int{1,2,4,0,2,-3,-2}),12){
		t.Fatal("Failed Test 5")
	}
}