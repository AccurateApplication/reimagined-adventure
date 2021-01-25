package main

// . Write a C program to generate a random numbe

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Print(rand.Intn(100), ",")
	x := (rand.Intn(10))
	fmt.Println(rand.Float64())
	fmt.Println("enter random number see if correct")
	var y int
	fmt.Scanln(&y)
	if y == x {
		fmt.Printf("y (%d) is same as x (d%)", x, y)
	} else {
		fmt.Printf(" x %d is not as same  y %d ", x, y)
	}
}
