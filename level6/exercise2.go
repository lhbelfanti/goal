package main

import "fmt"

func main () {
	x := []int {1, 3, 5, 7}
	a := foo1(x...)
	b := bar1(x)

	fmt.Println(a)
	fmt.Println(b)
}

func foo1(nums ...int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}

	return sum
}

func bar1(nums []int) int{
	sum := 0
	for _, v := range nums {
		sum += v
	}

	return sum
}
