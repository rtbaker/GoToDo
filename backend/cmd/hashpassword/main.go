package main

import (
	"fmt"
	"io"
	"os"

	"github.com/rtbaker/GoToDo/password"
)

// So tests can override
var out io.Writer = os.Stdout
var errOut io.Writer = os.Stderr

func main() {
	os.Exit(run())
}

func run() int {
	command := os.Args[0]
	argsWithoutProg := os.Args[1:]

	if len(argsWithoutProg) != 1 {
		fmt.Fprintf(errOut, "Usage: %s <password>", command)
		return 1
	}

	hash, err := password.HashPassword(argsWithoutProg[0])

	if err != nil {
		fmt.Fprintf(errOut, "Hashing error: %s\n", err)
		return 1
	}

	fmt.Fprintf(out, "%s\n", hash)

	return 0
}
