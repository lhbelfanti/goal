package main

import "fmt"

func main () {
	convert(binary, "James Moriarty")
}

func convert(to func(s string) string, s string) {
	fmt.Println(to(s))
}

func binary(s string) string {
	x := ""
	fmt.Println(s, "in binary is:")
	for _, v := range s {
		x += fmt.Sprintf("%b", v)
	}
	return x
}
