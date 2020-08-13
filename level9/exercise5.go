package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var wgr sync.WaitGroup

	var inc int64 = 0

	times := 100
	wgr.Add(times)

	for i := 0; i < times; i++ {
		go func() {
			atomic.AddInt64(&inc, 1)
			fmt.Println(atomic.LoadInt64(&inc))
			wgr.Done()
		}()
	}
	wgr.Wait()
	fmt.Println("Incrementer:", inc)
}
