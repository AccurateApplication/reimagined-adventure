package main

import (
	"fmt"
)

func main() {
	checkingInts(32, 2)
}

//Write a C program to check two given integers, and return true if one of them is 30 or if their sum is 30.

func checkingInts(a int, b int) {
	var result int
	result = a + b
	if result == 30 {
		fmt.Print("True")
	} else if a == 30 {
		fmt.Print("True")
	} else if b == 30 {
		fmt.Print("True")
	} else {
		fmt.Print(result)
	}
}
