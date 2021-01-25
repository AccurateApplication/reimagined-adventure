package main

import (
	"fmt"
)

func main() {
	sumArray(5, 5)
}

func sumArray(ar []int) {
	for i := 0; i < len(ar); i++ {
		fmt.Println(ar)
		fmt.Printf("Hej %d and", i)
	}
	fmt.Println(ar)
}
