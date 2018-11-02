package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/sckelemen/ember/src/ember/lexer"
	"github.com/sckelemen/ember/src/ember/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	l := lexer.New("")
	l.NextToken()
	fmt.Printf("Hello %s! Welcome to Ember.\n", user.Username)
	fmt.Printf("Please type a command...\n")
	repl.Start(os.Stdin, os.Stdout)
}
