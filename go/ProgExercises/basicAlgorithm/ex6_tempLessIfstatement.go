package main

import (
	"fmt"
)

func main() {
	lessthan0(5)
}

func lessthan0(n int) {
	if n < 0 {
		fmt.Println("less than 0?")
	} else if n > 100 {
		fmt.Println("more than 100")
	} else {
		fmt.Println("between 0 and 100")
	}
}
