package main

import (
	"fmt"
)

func main() {
	multipleof3And7(5)
	multipleof3And7(21)
	multipleof3And7(-21)
	multipleof3And7(-22)
}

// Write a C program to check if a given positive number is a multiple of 3 or a multiple of 7.

func multipleof3And7(n int) {
	if (n%3 == 0) && (n%7 == 0) {
		fmt.Println("oidjsaoidjsadjsa")
	} else {
		fmt.Println("False")
	}
}
