package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	file := readInputFile("./input")
	//fmt.Println(file)
	//stringPos := 0
	//numMatches var int = 0
	var karta [1024]string

	mapHeight := 0
	mapWidth := 31
	count := 0
	//count1 := 0
	//count2 := 0
	//count3 := 0

	for i, line := range file {
		karta[i] = line
		mapHeight++
		//fmt.Println(line)
		//fmt.Println(subLine)
		//fmt.Printf("match: %b fr: %d to: %d\n\n", numMatch, fr, to)
	}
	//fmt.Println(karta)
	fmt.Println(string(karta[0][0]) == ".")
	//numMatch := (strings.Count("#", subLine))

	var posX int = 0
	var posY int = 0
	for posY < mapHeight {
		fakeX := posX % mapWidth
		isTree := (string(karta[posY][fakeX]) == "#")
		posY += 1 //down
		posX += 3 // right
		//fmt.Println(isTree, fakeX, posY)
		if isTree == true {
			count++
		}
	}
	fmt.Println("Answer:", count)

}

func readInputFile(filePath string) []string {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return strings.Split(strings.TrimSuffix(string(data), "\n"), "\n")
}
