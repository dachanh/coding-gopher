package main

import (
	"fmt"
	"sync"
)

func main() {
	var count int
	var lock sync.Mutex
	var wg sync.WaitGroup
	descrease := func() {
		lock.Lock()
		defer lock.Unlock()
		count--
		fmt.Println("- ", count)
	}

	increase := func() {
		lock.Lock()
		defer lock.Unlock()
		count++
		fmt.Println("+ ", count)
	}

	for it := 0; it <= 5; it++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			increase()
		}()
	}

	for it := 0; it <= 10; it++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			descrease()
		}()
	}
	wg.Wait()
	fmt.Println("The Result :", count)

}
