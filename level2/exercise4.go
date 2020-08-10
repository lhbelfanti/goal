package main

import "fmt"

func main () {
	n := 123
	fmt.Printf("Decimal: %v | Binary: %b | Hex: %x\n", n, n, n)
	n = n << 1
	fmt.Printf("Decimal: %v | Binary: %b | Hex: %x", n, n, n)
}
