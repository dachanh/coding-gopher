package main

import "fmt"

func testSliceChannel(channels ...<-chan interface{}) []<-chan interface{} {
	return channels
}

func main() {

	a := make(chan interface{})
	b := make(chan interface{}, 9)
	channels := testSliceChannel(b)
	fmt.Println(channels)
	fmt.Println(testSliceChannel(append(channels[0:], a)...))
}
