package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	file := readInputFile("./inputfile")
	fmt.Printf("type: %T\n", file)
	// Convert []string to int
	for _, line := range file {
		intLine, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		for _, line2 := range file {
			sum := 2020
			intLine2, err := strconv.Atoi(line2)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			if intLine+intLine2 == sum {
				fmt.Println(line, line2)
				answer := intLine * intLine2
				fmt.Printf("Answer: %d\n", answer)
			}
			for _, line3 := range file {
				sum := 2020
				intLine3, err := strconv.Atoi(line3)
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
				if intLine+intLine2+intLine3 == sum {
					fmt.Println(line, line2, line3)
					answer := intLine * intLine2 * intLine3
					fmt.Printf("Answer: %d\n", answer)
				}
			}
		}
	}
}

func readInputFile(filePath string) []string {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return strings.Split(strings.TrimSuffix(string(data), "\n"), "\n")
}
