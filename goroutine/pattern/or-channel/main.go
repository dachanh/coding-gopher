package main

import (
	"fmt"
	"time"
)

type checkResult struct {
	Test string
}

func main() {
	var or func(channels ...<-chan checkResult) <-chan checkResult
	or = func(channels ...<-chan checkResult) <-chan checkResult {

		switch len(channels) {
		case 0:
			return nil
		case 1:
			return channels[0]
		}

		orDone := make(chan checkResult)
		go func() {
			defer close(orDone)

			switch len(channels) {
			case 2:

				select {
				case msg := <-channels[0]:
					orDone <- msg
				case msg := <-channels[1]:
					orDone <- msg
				}
			default:
				select {
				case msg := <-channels[0]:
					orDone <- msg
				case msg := <-channels[1]:
					orDone <- msg
				case msg := <-channels[2]:
					orDone <- msg
				case msg := <-or(append(channels[3:], orDone)...):
					orDone <- msg
				}
			}
		}()
		return orDone
	}
	sig := func(test string, after time.Duration) <-chan checkResult {
		c := make(chan checkResult)
		go func() {
			defer close(c)
			time.Sleep(after)
			c <- checkResult{Test: test}
		}()
		return c
	}

	res := <-or(
		sig("test 3", time.Minute*1),
		sig("test 4", time.Microsecond*20),
		sig("test 1", time.Microsecond*50),
		sig("test 2", time.Hour*2),
	)
	fmt.Println(res.Test)
}
