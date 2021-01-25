package main

import (
	"fmt"
)

func main() {
	height := 7
	width := 5
	area := width * height
	perimeter := (width * 2) + (height * 2)
	fmt.Println(area, perimeter)
}
