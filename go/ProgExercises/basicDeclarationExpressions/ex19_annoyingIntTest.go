package main

import "fmt"

func main() {
	var p, q, r, s int
	fmt.Println("enter p")
	fmt.Scanln(&p)
	fmt.Println("enter q")
	fmt.Scanln(&q)
	fmt.Println("enter r")
	fmt.Scanln(&r)
	fmt.Println("enter s")
	fmt.Scanln(&s)
	if (q >= r) && (s >= p) && (r+s >= p+q) {
		fmt.Println("Correct values")
	} else {
		fmt.Println("Wrong values")
	}
}
