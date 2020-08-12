package main

import "fmt"

func main () {
	defer foo2()
	bar2()
}

func foo2() {
	fmt.Println("Foo")
}

func bar2() {
	fmt.Println("Bar")
}
