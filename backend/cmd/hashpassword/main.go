package main

import (
	"fmt"
	"os"

	"github.com/rtbaker/GoToDo/password"
)

func main() {
	command := os.Args[0]
	argsWithoutProg := os.Args[1:]

	if len(argsWithoutProg) != 1 {
		fmt.Printf("Usage: %s <password>", command)
		os.Exit(1)
	}

	hash, err := password.HashPassword(argsWithoutProg[0])

	if err != nil {
		fmt.Printf("Hashing error: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("%s\n", hash)

	os.Exit(0)
}
