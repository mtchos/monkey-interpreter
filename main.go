package main

import (
	"fmt"
	"monkey-interpreter/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		fmt.Printf("error getting current user: %s\n", err)
	}
	fmt.Printf("hello %s! this is the Monkey programming language!\n", user.Username)
	fmt.Printf("feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
