package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/uesteibar/lainoa/pkg/repl"
)

func printHelp() {
	fmt.Println(`
The following commands are available:

	repl	start the lainoa REPL (interactive console)
	help	print this nice little help`)
}

func startRepl() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Lainoa programming language.\n",
		user.Username)
	fmt.Println("\nGo ahead and enter some code!")

	repl.Start()
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("You need to tell me what to do!")
		printHelp()
		return
	}
	action := os.Args[1]

	switch action {
	case "repl":
		startRepl()
	case "help":
		printHelp()
	default:
		fmt.Printf("Command %s not supported\n", action)
		printHelp()
	}
}
