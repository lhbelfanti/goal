package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	wg.Add(2)
	go sherlock()
	go watson()
	wg.Wait()

	fmt.Println("Exit")
}

func sherlock() {
	fmt.Println("I'm Sherlock Holmes")
	wg.Done()
}

func watson() {
	fmt.Println("I'm John Watson")
	wg.Done()
}
