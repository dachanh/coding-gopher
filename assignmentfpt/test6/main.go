package main

import (
	"fmt"
	"sort"
	"strings"
)

func ADD(T string) string {
	return fmt.Sprintf("ADD %s", string(T[len(T)-1]))
}

func CheckSameString(S, T string) bool {
	Sslice := []string(strings.Split(S, ""))
	Tslice := []string(strings.Split(T, ""))
	sort.Strings(Sslice)
	sort.Strings(Tslice)
	newS := strings.Join(Sslice, "")
	newT := strings.Join(Tslice, "")
	return newS == newT
}

func Change(S, T string) string {
	countChangeLetter := 0
	var result string
	for i := 0; i < len(S); i++ {
		if S[i] != T[i] {
			countChangeLetter += 1
			result = fmt.Sprintf("CHANGE %s %s", string(S[i]), string(T[i]))
		}
	}
	if countChangeLetter > 1 {
		return "NOT_RESULT"
	}
	return result
}

func Solution(S string, T string) string {
	if len(T) == len(S) {
		if S == T {
			return "NOTHING"
		}
		if CheckSameString(S, T) {
			for i := 0; i < len(S); i++ {
				if S[i] != T[i] {
					return fmt.Sprintf("MOVE %s", string(S[i]))
				}
			}
		} else {
			result := Change(S, T)
			if result == "NOT_RESULT" {
				return "IMPOSSIBLE"
			}
			return result
		}
	} else if len(T)-len(S) == 1 {
		return ADD(T)
	}
	return "IMPOSSIBLE"
}
