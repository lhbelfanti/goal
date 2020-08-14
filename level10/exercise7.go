package main

import (
	"fmt"
	"sync"
)

func main() {

	var s sync.WaitGroup

	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			s.Add(1)
			go func(i int) {
				for j := 0; j < 10; j++ {
					c <- j*10 + i
				}
				s.Done()
			}(i)
		}
		s.Wait()
		close(c)
	}()


	for v := range c {
		fmt.Println(v)
	}
}

