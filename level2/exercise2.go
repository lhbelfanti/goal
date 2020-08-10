package main

import "fmt"

func main () {
	n1 := 2
	n2 := 3

	var1 := (n1 == n2)
	var2 := (n1 <= n2)
	var3 := (n1 >= n2)
	var4 := (n1 != n2)
	var5 := (n1 > n2)
	var6 := (n1 < n2)

	fmt.Println(var1)
	fmt.Println(var2)
	fmt.Println(var3)
	fmt.Println(var4)
	fmt.Println(var5)
	fmt.Println(var6)
}
