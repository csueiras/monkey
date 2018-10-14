package main

import (
	"fmt"
	"github.com/csueiras/monkey/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! Type in your Monkey Language commands\n", user.Username)
	repl.Start(os.Stdin, os.Stdout)
}
