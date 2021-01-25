package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

// Create slice. If line == "\n". Go to next slice id. Each slice = 1 passport. Regex for fields if valid or not.
// If only "cid" field missing = valid
// Required: byr,iyr,eyr,hgt,hcl,ecl,pid,cid
// if: byr,iyr,eyr,hgt,hcl,ecl,pid  = valid

var passportID int = 0

func main() {
	//passport := make([]string, 150)
	var passport [251]string
	file := readInputFile("./input")
	for _, line := range file {
		//fmt.Printf("len line: %d line contains: %s\n ", len(line), line)
		//	fmt.Println("len slice passport", len(passport))

		if len(line) <= 0 {
			passportID++
			//fmt.Println(line, passportID)
			//fmt.Println(len(line))
		} else if len(line) >= 1 {
			//passport = append(passport, line)
			passport[passportID] += line
		}
	}

	count := 0
	for _, val := range passport {
		//fmt.Printf("Index: %d w/ Value: %s\t\n", i, val)
		// if: byr,iyr,eyr,hgt,hcl,ecl,pid  = valid
		mtch1, err := regexp.MatchString("(byr)", val)
		mtch2, err := regexp.MatchString("(iyr)", val)
		mtch3, err := regexp.MatchString("(eyr)", val)
		mtch4, err := regexp.MatchString("(hgt)", val)
		mtch5, err := regexp.MatchString("(hcl)", val)
		mtch6, err := regexp.MatchString("(ecl)", val)
		mtch7, err := regexp.MatchString("(pid)", val)
		if err != nil {
			log.Fatal(err)
		}
		if mtch1 == true && mtch2 == true && mtch3 == true && mtch4 == true && mtch5 == true && mtch6 == true && mtch7 == true {
			fmt.Println(val)
			count++
		}
	}
	fmt.Println(count)
}
func readInputFile(filePath string) []string {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return strings.Split(strings.TrimSuffix(string(data), "\n"), "\n")
}
