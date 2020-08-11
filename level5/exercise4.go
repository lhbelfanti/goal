package main

import "fmt"

func main () {

	x := struct {
		name string
		age int
		friends []string
	}{
		name: "John",
		age: 48,
		friends: []string {"Peter", "Mary"},
	}

	fmt.Println(x.name)
	fmt.Println(x.age)
	for _, v := range x.friends {
		fmt.Println(v)
	}
}
