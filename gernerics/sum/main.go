package main

import "fmt"

type Number interface {
	int | int64 | float64
}

func Sum[T Number](numbers []T) T {
	var total T
	for _, x := range numbers {
		total += x
	}
	return total
}

func main() {
	xs := []int{5, 1, 10}
	fmt.Println(Sum(xs))
}
