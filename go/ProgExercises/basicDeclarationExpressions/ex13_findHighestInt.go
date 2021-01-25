package main

// VWrite a C program that accepts three integers and find the maximum of three.

import (
	"fmt"
)

func main() {
	var a, b, c int
	fmt.Println("enter a  b  c")
	fmt.Scanln(&a, &b, &c)
	fmt.Printf("a: %d b: %d c: %d", a, b, c)
	if a >= b {
		fmt.Println("\n a is more than b", a)
	}

}
