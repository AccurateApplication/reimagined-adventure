package main

import (
	"fmt"
)

func main() {
	checkInt(5, 15, 60)
}

// Write a C program to check whether three given integer values are in the range 20..50 inclusive. Return true if 1 or more of them are in the said range otherwise return false

func checkInt(a, b, c int) string {
	x := "hej"
	if a >= 20 && a <= 50 {
		fmt.Println("first int true")
	} else if b >= 20 && b <= 50 {
		fmt.Println("second int true")
	} else if c >= 20 && c <= 50 {
		fmt.Println("third int true")
	} else {
		x := "hejhej"
		fmt.Print(x)
	}
	return x
}
