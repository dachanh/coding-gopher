package main

import "testing"

type question struct {
	one []int
}
type answer struct {
	one bool
}
type testcase struct {
	question
	answer
}

func Test(t *testing.T) {
	qs := []testcase{
		{
			question{[]int{1, 2, 3, 4, 1}},
			answer{true},
		},
		{
			question{[]int{1, 2, 3, 4}},
			answer{false},
		},
		{
			question{[]int{1, 2, 3, 4, 1, 2, 3}},
			answer{true},
		},
	}
	for i, q := range qs {
		if containsDuplicate(q.question.one) != q.answer.one {
			t.Fatal("wrong testcase ", i)
		}
	}
}
