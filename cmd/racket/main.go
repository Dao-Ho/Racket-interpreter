package main

import (
	"fmt"
	"os"
	"os/user"

	"racket-interpreter/pkg/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Hallo %s!\n", user.Name)
	fmt.Println("Provide some commands for the interpreter: ")
	repl.Start(os.Stdin, os.Stdout)
}
