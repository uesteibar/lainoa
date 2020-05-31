package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/uesteibar/lainoa/pkg/repl"
	"github.com/uesteibar/lainoa/pkg/runner"
)

func printHelp() {
	fmt.Println(`
The following commands are available:

	run		run a file
	repl	start the lainoa REPL (interactive console)
	help	print this nice little help`)
}

func run() {
	if len(os.Args) < 3 {
		fmt.Println("You need to tell me what file to run:")
		fmt.Println("\tlainoa run path/to/file.ln")
		return
	}

	filepath := os.Args[2]
	runner.Start(filepath)
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
	case "run":
		run()
	case "repl":
		startRepl()
	case "help":
		printHelp()
	default:
		fmt.Printf("Command %s not supported\n", action)
		printHelp()
	}
}
