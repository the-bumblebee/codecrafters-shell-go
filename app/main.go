package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" and "os" imports in stage 1 (feel free to remove this!)
var _ = fmt.Fprint
var _ = os.Stdout

func main() {
	// TODO: Uncomment the code below to pass the first stage
	for {
		fmt.Fprint(os.Stdout, "$ ")
		command, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			return
		}
		command = strings.TrimSpace(command)
		args := strings.Split(command, " ")
		if args[0] == "exit" {
			os.Exit(0)
		} else if args[0] == "echo" {
			fmt.Println(strings.Join(args[1:], " "))
		} else {
			fmt.Println(args[0] + ": command not found")
		}
	}
}
