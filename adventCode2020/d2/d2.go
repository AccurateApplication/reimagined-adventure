package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file := readInputFile("./input")
	fmt.Printf("type: %T\n", file)
	// Convert []string to int
	NumberOfMatches := 0
	for _, line := range file {
		fmt.Println(line)
		//test := strings.Split(line, ":")
		//myRegex, err = regexp.Compile("([0-9][0-9]|[0-9]|$[0-9][0-9]|$[0-9])")
		myRegex, err := regexp.Compile("^([0-9]+)-([0-9]+) ([a-z]): (.*)$")
		if err != nil {
			fmt.Println("err:", err)
		}
		match := myRegex.FindStringSubmatch(line)
		fmt.Printf("found=%s\n", match[0])
		fmt.Printf("found=%s\n", match[3])
		fr := match[1]
		to := match[2]
		frInt, err := strconv.Atoi(fr)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		toInt, err := strconv.Atoi(to)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(toInt, frInt)
		char := match[3]
		password := match[4]
		fmt.Printf("fr=%s\n", fr)
		fmt.Printf("to=%s\n", to)
		fmt.Printf("char=%s\n", char)
		fmt.Printf("pw=%s\n", password)
		numMatch := (strings.Count(password, char))
		if numMatch >= frInt && numMatch <= toInt {
			fmt.Println("correct")
			NumberOfMatches++

		}

	}
	fmt.Println(NumberOfMatches)
}

func readInputFile(filePath string) []string {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return strings.Split(strings.TrimSuffix(string(data), "\n"), "\n")
}
