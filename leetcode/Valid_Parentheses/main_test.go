package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test(t *testing.T){
	if !reflect.DeepEqual(isValid("()"),true){
		t.Fatal("Failed test 1")
	}
	if !reflect.DeepEqual(isValid("()[]{}"),true){
		t.Fatal("Failed test 2")
	}
	if !reflect.DeepEqual(isValid("(]"),false){
		t.Fatal("Failed test 3")
	}
	if !reflect.DeepEqual(isValid("()["),false){
		t.Fatal("Failed test 4")
	}
	if !reflect.DeepEqual(isValid(""),false){
		t.Fatal("Failed test 5")
	}
	if !reflect.DeepEqual(isValid("(((("),false){
		t.Fatal("Failed test 6")
	}
	if !reflect.DeepEqual(isValid("}{}{}"),false){
		t.Fatal("Failed test 7")
	}
	if !reflect.DeepEqual(isValid("){"),false){
		t.Fatal("Failed test 8")
	}
	if !reflect.DeepEqual(isValid("([}}])"),false){
		t.Fatal("Failed test 9")
	}
	fmt.Println("PASSED")
}
