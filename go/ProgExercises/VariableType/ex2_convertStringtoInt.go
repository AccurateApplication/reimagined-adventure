package main

import (
	"fmt"
	"strconv"
)

func main() {
	stringz := "hahaiduhsaiudhasuidhuiashd"
	var hejhej int = 3849057394857
	fmt.Printf("stringz type:  %T \n", stringz)
	fmt.Printf("hejhej type:  %T \n", hejhej)
	i, err := strconv.Atoi("-34444")
	if err != nil {
		fmt.Print(err)
	}
	s := strconv.Itoa(234)
	fmt.Print("S: ", s)
	fmt.Print(" \n i: ", i)
	fmt.Println("\n\n\n ")
	fmt.Printf("i type:  %T \n", i)
	fmt.Printf("S type:  %T \n", s)
}
