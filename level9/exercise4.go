package main

import (
	"fmt"
	"sync"
)

func main() {
	var wgr sync.WaitGroup
	var mu sync.Mutex

	inc := 0

	times := 100
	wgr.Add(times)

	for i := 0; i < times; i++ {
		go func() {
			mu.Lock()
			inc++
			fmt.Println(inc)
			mu.Unlock()
			wgr.Done()
		}()
	}
	wgr.Wait()
	fmt.Println("Incrementer:", inc)

}
