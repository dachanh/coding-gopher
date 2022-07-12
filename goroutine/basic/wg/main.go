package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for _, element := range []string{"test1", "test2", "test3"} {
		wg.Add(1)
		go func(value string) {
			defer wg.Done()
			fmt.Println(value)
		}(element)
		wg.Wait()
	}

}
