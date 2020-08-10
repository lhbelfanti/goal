package main

import "fmt"

func main () {
	const (
		_ = iota
		a = 2020 + iota
		b = 2020 + iota
		c = 2020 + iota
		d = 2020 + iota
	)

	fmt.Printf("%v - %v - %v - %v", a, b, c, d)
}
