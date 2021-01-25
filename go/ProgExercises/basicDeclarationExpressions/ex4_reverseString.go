package main

import (
	"fmt"
)

func main() {
	reverseItBaby()
}

func reverseItBaby() {
	var stringen string
	fmt.Println("inpuit string")
	fmt.Scanln(&stringen)
	//fmt.Println(stringen)
	/* doesnt work:
	Runes := []rune(stringen)
	for i, j := 0, len(stringen)-1; i < j; i, j = i+1, j-1 {
		Runes[i], Runes[j] = Runes[j], Runes[i]
		fmt.Println(stringen)
	}
	//return string(stringen)
	fmt.Println(stringen)
	*/
	var result string
	for _, v := range stringen {
		result = string(v) + result
	}
	fmt.Println(result)

}
