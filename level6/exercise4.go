package main

import "fmt"

type person struct {
	first string
	last string
	age int
}

func (p person) speak() {
	fmt.Println("My name is", p.first, p.last, "and my age is:", p.age)
}

func main () {
	sh := person{
		first: "Sherlock",
		last: "Holmes",
		age: 2020 - 1854,
	}

	sh.speak()
}
