package main

import (
	"fmt"
)

func main() {
	var days int
	fmt.Println("Enter days to convert")
	fmt.Scanln(&days)
	var years int = (days / 365)
	var weeks int = (days % 365) / 7
	var dayz int = days - ((years * 365) + (weeks * 7))
	fmt.Printf("%d days is\n Years: %d Weeks: %d Days: %d \n \n\n ", days, years, weeks, dayz)

}
