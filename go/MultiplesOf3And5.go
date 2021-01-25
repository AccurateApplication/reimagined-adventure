// If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23. Find the sum of all the multiples of 3 or 5 below 1000.
package main

import "fmt"

var NumberList = []int{}

func main() {
	for i := 1; i < 1000; i++ {
		if i%3 == 0 {
			NumberList = append(NumberList, i)
		} else if i%5 == 0 {
			NumberList = append(NumberList, i)
		}
	}
	fmt.Println(NumberList)
	sum := 0
	for _, val := range NumberList {
		sum += val
	}
	fmt.Printf("Sum of multiples 3 5 is %d\n", sum)

}
