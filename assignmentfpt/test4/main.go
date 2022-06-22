package main

import (
	"sort"
	"strings"
)

func Solution(A, B []string, P string) string {
	indexIng := make(map[string]int)
	for it := 0; it < len(A); it++ {
		indexIng[A[it]] = it
	}
	sort.Strings(A)
	for it := 0; it < len(A); it++ {
		if strings.Contains(B[indexIng[A[it]]], P) {
			return A[it]
		}
	}
	return "NO CONTACT"
}
