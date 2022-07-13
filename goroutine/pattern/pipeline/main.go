package main

import "fmt"

func main() {

	generator := func(done chan struct{}, num ...int) <-chan int {
		stream := make(chan int)
		go func() {
			defer close(stream)
			for _, i := range num {
				select {
				case <-done:
					return
				case stream <- i:
				}
			}
		}()
		return stream
	}
	multiply := func(done chan struct{}, inputstream <-chan int, v int) <-chan int {
		outputstream := make(chan int)
		go func() {
			defer close(outputstream)
			for i := range inputstream {
				select {
				case <-done:
					return
				case outputstream <- i * v:
				}
			}
		}()
		return outputstream
	}

	add := func(done chan struct{}, inputstream <-chan int, v int) <-chan int {
		outputstream := make(chan int)
		go func() {
			defer close(outputstream)
			for i := range inputstream {
				select {
				case <-done:
					return
				case outputstream <- i + v:
				}
			}
		}()
		return outputstream
	}
	done := make(chan struct{})
	for res := range add(done, multiply(done, add(done, generator(done, 1, 1, 2, 2, 4, 0, 2), 2), 4), 10) {
		fmt.Println(res)
	}
}
