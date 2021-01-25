package main

import (
	"fmt"
	"os"
)

func main() {

	dirs := os.Args[1:]
	if len(dirs) <= 1 {
		for _, dir := range dirs {
			os.Mkdir(dir, 0744)
			fmt.Println("created directory: ", dir)
		}
	} else {
		for _, dir := range dirs {
			fmt.Println("created dir: ", dir)
			os.Mkdir(dir, 0744)
		}
	}
	//fmt.Println(len(dirs))
}
