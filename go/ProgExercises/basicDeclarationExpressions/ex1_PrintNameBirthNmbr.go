package main

// Write a C program to print your name, date of birth. and mobile number.

import "fmt"

func main() {
	printl()
}

func printl() {
	var name string
	var month string
	var year string

	fmt.Println("enter name")
	fmt.Scanln(&name, " enter name")
	fmt.Println("enter month born")
	fmt.Scanln(&month)
	fmt.Println("enter year born")
	fmt.Scanln(&year)
	fmt.Printf("year: %s \nname: %s \nmonth: %s", year, name, month)
}
