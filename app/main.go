package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
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
		pathDirs := strings.Split(os.Getenv("PATH"), ":")
		fmt.Fprint(os.Stdout, "$ ")
		command, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			os.Exit(1)
		}
		command = strings.TrimSpace(command)
		args := strings.Split(command, " ")
		switch args[0] {
		case "exit":
			os.Exit(0)
		case "type":
			if len(args) < 2 {
				fmt.Println("Usage: type <command>")
			} else if slices.Contains(commands, args[1]) {
				fmt.Println(args[1] + " is a shell builtin")
			} else {
				flag := false
				for _, dir := range pathDirs {
					file := filepath.Join(dir, args[1])
					info, err := os.Stat(file)
					if err == nil && info.Mode()&0111 != 0 {
						fmt.Println(args[1] + " is " + file)
						flag = true
						break
					}
				}
				if !flag {
					fmt.Println(args[1] + ": not found")
				}
			}
		case "echo":
			fmt.Println(strings.Join(args[1:], " "))
		default:
			flag := false
			for _, dir := range pathDirs {
				file := filepath.Join(dir, args[0])
				info, err := os.Stat(file)
				if err == nil && info.Mode()&0111 != 0 {
					flag = true
					cmd := exec.Command(args[0], args[1:]...)
					var out bytes.Buffer
					cmd.Stdout = &out
					err := cmd.Run()
					if err != nil {
						fmt.Fprintln(os.Stderr, "Command execution failed: ", err)
						break
					}
					fmt.Print(out.String())
					break
				}
			}
			if !flag {
				fmt.Println(args[0] + ": command not found")
			}
		}
	}
}
