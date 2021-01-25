package main

import (
	"fmt"
)

func main() {
	var employeeID string
	var hours float64
	var wage float64

	fmt.Println("Enter employee id:")
	fmt.Scanln(&employeeID)
	fmt.Println("enter working hrs")
	fmt.Scanln(&hours)
	fmt.Println("enter hour wage")
	fmt.Scanln(&wage)
	monthlyWage := wage * hours
	fmt.Printf("Employee ID: %d has a salary of %g/month.", employeeID, monthlyWage)

}
