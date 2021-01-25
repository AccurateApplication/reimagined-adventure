package main

// f this
import (
	"fmt"
)

func main() {
	var money int
	fmt.Println("enter monies to convert to sedlar")
	fmt.Scanln(&money)
	var total int = money / 100
	fmt.Printf("antal 100 sedlar: %d\n", total)
	total = money / 50
	fmt.Printf("antal 50 sedlar: %d\n", total)

}
