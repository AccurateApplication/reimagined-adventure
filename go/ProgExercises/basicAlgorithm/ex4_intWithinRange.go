package main

import (
	"fmt"
)

func main() {
	within100or200(55)
	within100or200(150)
}

// Write a C program to check a given integer and return true if it is within 10 of 100 or 200.

func within100or200(a int) {
	if (a >= 100) && (a <= 200) {
		fmt.Print("True")
	} else {
		fmt.Println("false")
	}

}
