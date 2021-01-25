// Each new term in the Fibonacci sequence is generated by adding the previous two terms. By starting with 1 and 2, the first 10 terms will be: 1, 2, 3, 5, 8... By considering the terms in the Fibonacci sequence whose values do not exceed four million, find the sum of the **even-valued** terms.
package main

import (
	"fmt"
)

var NumberSlice = []int{}

func main() {
	//var n int = 100
	t1 := 0
	t2 := 1
	nextT := 0
	for i := 1; i <= 3000; i++ {
		if i == 1 {
			continue
		}
		if i == 2 {
			continue
		}
		nextT = t1 + t2
		if nextT >= 4000000 {
			break
		}
		t1 = t2
		t2 = nextT
		if nextT%2 == 0 {
			NumberSlice = append(NumberSlice, nextT)
		}
		fmt.Printf("%d\n", nextT)

	}
	fmt.Printf("Final num: %d\n", nextT)
	sum := 0
	for _, val := range NumberSlice {
		sum += val
	}
	fmt.Printf("Sum of even numbers in slice: %d\n", sum)

	fmt.Printf("Sums of all numbers: %d\n", NumberSlice)

}
