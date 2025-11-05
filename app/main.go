package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" and "os" imports in stage 1 (feel free to remove this!)
var _ = fmt.Fprint
var _ = os.Stdout

func main() {
	// TODO: Uncomment the code below to pass the first stage
	for {
		commands := []string{"echo", "exit", "type"}
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
		} else if args[0] == "type" {
			if len(args) < 2 {
				fmt.Println("Usage: type <command>")
			} else if slices.Contains(commands, args[1]) {
				fmt.Println(args[1] + " is a shell builtin")
			} else {
				fmt.Println(args[1] + ": not found")
			}
		} else {
			fmt.Println(args[0] + ": command not found")
		}
	}
}
