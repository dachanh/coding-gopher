package main

import (
	"strings"
)

func isValid(s string) bool {
	if len(s)%2 != 0 || len(s) == 0{
		return false
	}
	liststr := strings.Split(s,"")
	var stack []int
	mapbracket := make(map[string]string,3)
	mapbracket["]"] =  	"["
	mapbracket["}"] =  	"{"
	mapbracket[")"] = 	"("
	for it := 0 ; it < len(liststr); it++{
		if val, ok := mapbracket[liststr[it]]; ok{
			if len(stack) == 0{
				return false
			}
			if val ==  liststr[stack[len(stack)-1]]{
				stack = stack[:len(stack)-1]
			}else {
				return  false
			}
		}else {
			stack = append(stack,it)
		}
	}
	if len(stack) != 0{
		return false
	}
	return true
}
