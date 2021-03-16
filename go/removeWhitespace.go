// Replaces all empty spaces and replaces with underscore ("_")
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	files, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}
	//regex := regexp.MustCompile(" ")
	//underScore := ("_")
	for _, fil := range files {
		found, err := regexp.MatchString(" ", fil.Name())
		if err != nil {
			log.Fatal(err)
		}
		if found {
			fmt.Printf("found whitespace, file: %s . will rename \n", fil.Name())
			newFileName := strings.Replace(fil.Name(), " ", "_", -1) // -1 = replace all occurence
			os.Rename(fil.Name(), newFileName)

		}

		//regex.ReplaceAllString(fil.Name(), underScore)
	}

}
