package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		// just getting the working dir (e,g /home/something $ )
		dir, err := os.Getwd()
		if err != nil {
			fmt.Println("Error getting current directory:", err)
			dir = ""
		}
		fmt.Print(dir + " $ ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Some error occured :0 ", os.Stderr, err)
		} else {
			if err := executeCommand(input); err != nil {
				fmt.Println("Error:", err)
			}
		}

	}
}

func executeCommand(input string) error {

	input = strings.TrimSpace(input)

	args := strings.Split(input, " ")
	if len(args) == 0 || args[0] == "" {
		return errors.New("no command entered")
	}

	switch args[0] {
	case "cd":
		if len(args) < 2 {
			return errors.New("no path provided for cd")
		}
		return os.Chdir(args[1])
	case "whoami":
		currentuser, err := exec.Command("whoami").Output()
		if err != nil {
			return err
		}
		fmt.Println(strings.TrimSpace(string(currentuser)))
		return nil
	case "ls":
		files, err := os.ReadDir(".")
		if err != nil {
			return err
		}
		for _, file := range files {
			fmt.Println(file.Name())
		}
		return nil
	case "mkdir":
		if len(args) < 2 {
			return errors.New("no directory name provided for mkdir")
		}
		return os.Mkdir(args[1], 0755)
	case "rm":
		if len(args) < 2 {
			return errors.New("no file or directory name provided for rm")
		}
		return os.RemoveAll(args[1])
	case "exit":
		fmt.Println("Good bye :( ")
		os.Exit(0)
	}

	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd.Run()
}
