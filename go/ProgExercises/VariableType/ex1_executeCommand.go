package main

/*
1. Write a C program which will invoke the command processor to execute a command.
Expected Output :

Is command processor available?
Command processor available!
Executing command DIR
00c40280-5e27-11e6-bd4f-71e8825f8ea3
01691610-41e1-11e6-901d-35b72ececc72
...........
ff827330-443a-11e6-9820-23e2f60d924e
file.txt
logging_example.out
test.txt
Returned value is: 0.
*/

import (
	"fmt"
	"os/exec"
)

func main() {
	command := "pwd"
	cmd := exec.Command(command)
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Print(string(stdout))

}
