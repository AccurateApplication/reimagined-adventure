package main

import (
	"fmt"
	"sort"
)

func main() {
	// array
	var ad [5]int
	ad[1] = 5
	fmt.Println(ad, "\tarray")
	// slice
	numbers := []int{
		2,
		555,
		67,
		320,
		2,
		2,
		3,
		4,
		4,
		44,
	}
	fmt.Println(numbers)
	// sort slice of ints
	sort.Sort(sort.IntSlice(numbers))
	fmt.Println(numbers)
	// get sum of slice
	sum := 0
	for i := range numbers {
		sum += numbers[i]
	}
	fmt.Println(sum, "<- SUM ")
	// see matching int in SUM
	/* didnt work
	var matching int = 0
	for i := 0; i < len(numbers)-1; i++ {
		if numbers[i] == numbers[+i] {
			matching++
			fmt.Println(numbers[i])
			fmt.Println(numbers[i++])
		}
	}
	fmt.Println("\n matching int:", matching)
	*/
	//
	/*
			only prints out 1 2 3 4.....
		//	var matching int = 0
			for x, y := range numbers {
				y++
				fmt.Println(x)
				if y == x {
					matching++
					fmt.Println("x = y   ", x, y)
				}
			}
			fmt.Println("match:", matching)
	*/
	//for _, itemCopy := range numbers {
	//	itemCopy = itemCopy + 0
	//}

	fmt.Println(numbers, " before loop")
	for index, itemCopy := range numbers {
		fmt.Printf("%v: %v\n", index, itemCopy)
	}
}
