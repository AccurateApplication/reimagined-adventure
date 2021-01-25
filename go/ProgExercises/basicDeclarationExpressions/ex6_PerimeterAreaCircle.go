package main

// Write a C program to compute the perimeter and area of a circle with a radius of 6 inches

import (
	"fmt"
)

func main() {
	radius := 6.0
	pi := 3.14
	area := pi * radius * radius
	fmt.Print(area)
	perimeter := 2 * pi * radius
	fmt.Printf(" \nthis is perimeter: #this didnt work!!! %d \n \n  ", perimeter)
	fmt.Println(perimeter)
}
