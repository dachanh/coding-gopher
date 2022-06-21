package main

import (
	"fmt"
	"math"
	"strings"
)

func Solution(S string) int {
	var tempString string
	var count, minlen int
	minlen = math.MaxInt
	uniqueString := []string{}
	for i := 0; i < len(S)-1; i++ {
		for j := i + 1; j < len(S); j++ {
			tempString = S[i:j]
			count = strings.Count(S, tempString)
			if count == 1 {
				uniqueString = append(uniqueString, tempString)
			}
		}
	}

	if len(uniqueString) == 0 {
		return len(S)
	}

	for it := 0; it < len(uniqueString); it++ {
		if len(uniqueString[it]) < minlen {
			minlen = len(uniqueString[it])
		}
	}
	return minlen
}

func main() {
	fmt.Println(Solution("xxxx"))
}
