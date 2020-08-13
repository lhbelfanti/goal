package main

import "fmt"

type person struct {
	first string
	last string
	age int
}

type human interface {
	speak()
}

func (p *person) speak() {
	fmt.Printf("Name: %s %s - Age: %d\n", p.first, p.last, p.age)
}

func main() {
	p1 := person{
		first: "Sherlock",
		last: "Holmes",
		age: 33,
	}

	saySomething(&p1)
	//saySomething(p1) // It doesn't work because `func (p *person) speak() {...} ` only accepts a pointer as parameter
	p1.speak()
}

func saySomething(h human) {
	h.speak()
}
