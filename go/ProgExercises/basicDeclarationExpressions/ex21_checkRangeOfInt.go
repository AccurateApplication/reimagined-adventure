package main

import (
	"fmt"
)

func main() {
	var x int
	fmt.Println("enter x to test")
	fmt.Scanln(&x)
	if x <= 20 && x >= 0 {
		fmt.Println("range 0-19")
	} else if x <= 40 && x >= 20 {
		fmt.Println("range 20-40")
	} else if x <= 0 {
		fmt.Println("less than 0")
	} else if x >= 80 {
		fmt.Println("more than 80")
	} else {
		fmt.Println("not specfiied range")
	}
}
