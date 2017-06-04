package main

import (
	"capuchin/repl"
	"fmt"
	"os"
	"os/user"
)

func main() {

	user, err := user.Current()

	if err != nil {
		panic(err)
	}

	fmt.Printf("Welcome %s, this is the Capuchin REPL:\n", user.Username)
	repl.Start(os.Stdin, os.Stdout)
}
