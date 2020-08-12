package main

import "fmt"

type person struct {
	first string
	last string
}

func main() {
	p1 := person{
		first: "Sherlock",
		last: "Holmes",
	}

	fmt.Println(p1)
	changeMe(&p1)
	fmt.Println(p1)
}

func changeMe(p *person) {
	p.first = "John"
	p.last = "Watson"
}
