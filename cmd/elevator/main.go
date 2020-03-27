package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("elevator: need at least command to execute")
	}
	i := 1
	if strings.Compare(os.Args[i], "-p") == 0 {
		i++
		if len(os.Args) < 3 {
			fmt.Println("elevator: need at least command to execute")
		}
	}
	cmdName := os.Args[i]
	fmt.Printf("elevator: cmdName: %s\n", cmdName)
	var cmdArgs []string = nil
	if len(os.Args) > (i + 1) {
		cmdArgs = os.Args[i+1:]
		fmt.Printf("elevator: cmdArgs: %v\n", cmdArgs)
	}
	var cmd *exec.Cmd = nil
	if cmdArgs != nil {
		cmd = exec.Command(cmdName, cmdArgs...)
	} else {
		cmd = exec.Command(cmdName)
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Printf("elevator: %s\n", err)
	}
	if i > 1 {
		fmt.Print("elevator: press 'Enter' to continue...")
		bufio.NewReader(os.Stdin).ReadBytes('\n')
	}
}
