package main

import (
	"fmt"
)

// Write a C program to compute the sum of the two given integer values. If the two values are the same, then return triple their sum.
func main() {
	twoints(10, 11)

}

func twoints(a int, b int) {
	var result int
	if a == b {
		result = (a + b) * 3
	} else {
		result = a + b
	}
	fmt.Println(result)

}
