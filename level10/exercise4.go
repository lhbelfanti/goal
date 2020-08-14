package main

import (
	"fmt"
)

func main() {
	q := make(chan int)
	c := generate(q)

	receive(c, q)

	fmt.Println("about to exit")
}

func generate(q chan<- int) <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		q <- 0
		close(c)
	}()


	return c
}

func receive(c, q <-chan int) {
	for {
		select {
		case v := <-c:
			fmt.Println("c channel:", v)
		case <-q:
			fmt.Println("q channel")
			return
		}
	}
}
