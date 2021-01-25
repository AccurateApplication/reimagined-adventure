package main

import (
	"fmt"
)

func main() {
	fmt.Println(absolute(-5435))
}

func absolute(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}
