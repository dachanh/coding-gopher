package main

import "testing"

func Test(t *testing.T) {
	if maxProfit([]int{7, 1, 5, 3, 6, 4}) != 5 {
		t.Fatal("Wrong test 1")
	}
	if maxProfit([]int{7, 6, 4, 3, 1}) != 0 {
		t.Fatal("Wrong test 2")
	}
	if maxProfit([]int{7, 6, 4, 3, 1, 15}) != 11 {
		t.Fatal("Wrong test 3")
	}
}
