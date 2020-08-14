package main

import (
	"fmt"
)

func main() {
	// Using func literal
	c := make(chan int)

	go func(){
		c <- 42
	}()

	fmt.Println(<-c)
}
