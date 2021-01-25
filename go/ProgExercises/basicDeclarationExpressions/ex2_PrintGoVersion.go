package main

// Write a C program to get the C version you are using. (go..)
import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	out, err := exec.Command("go", "version").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("go version is %s\n", string(out))
}
