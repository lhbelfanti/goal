package main

import (
	"fmt"
)

type customErr struct {
	name string
}

func (ce customErr) Error() string {
	return fmt.Sprintf("There was an error with the custom error: %s", ce.name)
}

func main() {
	c1 := customErr{"Super Error"}
	foo(c1)
}

func foo(e error) {
	fmt.Println(e)
}
