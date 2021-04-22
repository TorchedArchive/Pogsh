package main

import (
	"fmt"
	"bufio"
	"os"
	"os/exec"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSuffix(input, "\n")
		cmdArgs := strings.Fields(input)

		cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)
		if err := cmd.Run() ; err != nil {
			fmt.Println("Not Pog")
			continue
		}
		fmt.Println("Pog")
	}
}
