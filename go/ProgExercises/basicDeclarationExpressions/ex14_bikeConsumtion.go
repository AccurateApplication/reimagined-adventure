package main

// Write a C program to calculate a bike’s average consumption from the given total distance (integer value) traveled (in km) and spent fuel
import (
	"fmt"
)

func main() {
	var distancE int = 350
	var literSpent int = 5
	avgConsumtion := distancE / literSpent
	fmt.Println(avgConsumtion)
}
