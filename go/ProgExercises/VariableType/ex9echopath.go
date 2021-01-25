package main

/*
1. Write a C program which will invoke the command processor to execute a command.
// did a copy of ex 1 to display $PATH variagble
*/
import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	command := "echo"
	lss := "$PATH"
	cmd := exec.Command(command, lss)
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Print(string(stdout))
	fmt.Println(os.Getenv("SHELL"))

}
