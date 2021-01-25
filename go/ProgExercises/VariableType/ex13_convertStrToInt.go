package main

//  Write a C program to convert a string to an integer.
import (
	"fmt"
	"strconv"
)

func main() {
	var x string
	fmt.Println("enter \"int\" to convert to real int")
	fmt.Scanln(&x)
	fmt.Printf(" \n type before: %T \n", x)
	fmt.Println(x)
	strconv.Atoi(x)
	//strconv.ParseUint(x, 5, 123456789)
	fmt.Printf(" \n type after: %T \n", x)
	fmt.Println(x)

}
