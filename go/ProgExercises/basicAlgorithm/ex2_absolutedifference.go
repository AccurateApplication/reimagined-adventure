package main

import (
	"fmt"
)

func main() {
	absolutedifference(50)
}

//Write a C program to get the absolute difference between n and 51. If n is greater than 51 return triple the absolute difference.
func absolutedifference(a int) {
	var value int
	var totales int
	n := 51
	value = a - n
	if value > 51 {
		totales = value * 3
		fmt.Println(totales, "if statement true")
	} else {
		fmt.Println(value)
	}

}
