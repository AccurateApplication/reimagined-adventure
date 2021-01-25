package main

// Write a C program to sort the elements of an array
// sort slice instead
import (
	"fmt"
	"sort"
)

func main() {
	// slice
	scl2 := []int{55, 235, 25345, 765908, 1}
	sort.Ints(scl2)
	fmt.Println(scl2)
}
