package main

/* Write a C program to calculate the distance between the two points. Go to the editor
Test Data :
Input x1: 25
Input y1: 15
Input x2: 35
Input y2: 10
*/

import (
	"fmt"
	"math"
)

func main() {
	x1 := 25.0
	y1 := 15.0
	x2 := 35.0
	y2 := 10.0
	var gdistance float64 = ((x2 - x1) * (x2 - x1)) + ((y2 - y1) * (y2 - y1))
	fmt.Println(gdistance)
	var distance float64 = math.Sqrt(gdistance)
	fmt.Println(distance)
}
