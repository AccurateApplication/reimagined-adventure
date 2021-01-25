package main

// Write a C program to convert a given integer (in seconds) to hours, minutes and seconds

import (
	"fmt"
)

func main() {
	var secs int
	fmt.Println("input sec to convert")
	fmt.Scanln(&secs)
	hours := secs / 3600
	minutes := (secs - (3600 * hours)) / 60
	sec := (secs - (3600 * hours) - (minutes * 60))
	fmt.Printf("\nH:%d M:%d S:%d\n", hours, minutes, sec)

}
