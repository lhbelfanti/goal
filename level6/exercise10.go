package main

import "fmt"

func main () {
	ia := "Irene Adler"
	m := discover(ia)
	fmt.Println(m())
	fmt.Println(m())
	fmt.Println(m())
	fmt.Println(m())
	fmt.Println(m())
	fmt.Println(m())
	fmt.Println(m())
	fmt.Println(m())
	fmt.Println(m())
	fmt.Println(m())
	fmt.Println(m())
}

func discover(s string) func() string{
	i := 1
	return func() string {
		str := ""
		for idx, r := range s {
			if i == idx {
				i++
				break
			}
			str += fmt.Sprintf("%c", r)
		}

		return str
	}
}
