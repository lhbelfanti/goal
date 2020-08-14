package main

import (
	"fmt"
)

func main() {
	// Using buffered channel
	c := make(chan int, 1)
	c <- 42


	fmt.Println(<-c)
}
