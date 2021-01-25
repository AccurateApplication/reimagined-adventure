package main

import (
	"fmt"
)

func main() {
	// Variables:
	var num11 float64
	var num22 float64
	var operator string
	var result float64

	fmt.Println("Input num 1")
	fmt.Scanln(&num11)
	fmt.Println("Input num 2")
	fmt.Scanln(&num22)
	fmt.Println("Enter operand (+,- *)")
	fmt.Scanln(&operator)
	switch operator {
	case "*":
		result = num11 * num22
		break
	case "+":
		result = num11 + num22
		break
	case "-":
		result = num11 - num22
		break
	default:
		fmt.Println("Use - + * ")
	}

	fmt.Println("answer:\t:", result)

}

func takeinput() {
	fmt.Println("Takeinput func")
	var num1 float64
	fmt.Scanln(&num1)
	fmt.Print(num1)
	return
}
