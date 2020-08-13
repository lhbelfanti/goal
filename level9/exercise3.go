package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wgr sync.WaitGroup

	inc := 0

	times := 100
	wgr.Add(times)

	for i := 0; i < times; i++ {
		go func() {
			v := inc
			runtime.Gosched()
			v++
			inc = v
			fmt.Println(inc)
			wgr.Done()
		}()
	}
	wgr.Wait()
	fmt.Println("Incrementer:", inc)

	// There are data race conditions
}
