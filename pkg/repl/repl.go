package repl

import (
	"fmt"

	"github.com/chzyer/readline"
	"github.com/uesteibar/lainoa/pkg/evaluator"
	"github.com/uesteibar/lainoa/pkg/lexer"
	"github.com/uesteibar/lainoa/pkg/object"
	"github.com/uesteibar/lainoa/pkg/parser"
)

const PROMPT = "⛅️ >> "

func Start() {
	l, err := readline.NewEx(&readline.Config{
		Prompt:          PROMPT,
		HistoryFile:     "/tmp/lainoa_repl_history.tmp",
		InterruptPrompt: "^C",
		EOFPrompt:       "exit",

		HistorySearchFold: true,
	})

	if err != nil {
		panic(err)
	}
	defer l.Close()

	env := object.NewEnvironment()

	for {
		line, err := l.Readline()
		if err != nil {
			return
		}

		l := lexer.New(line, "repl")
		p := parser.New(l)
		program := p.ParseProgram()

		if len(p.Errors()) > 0 {
			fmt.Println("Oops! Something is wrong here:")
			fmt.Println("  parser errors:")
			for _, err := range p.Errors() {
				fmt.Println(fmt.Sprintf("- %s\n", err.String()))
			}
		} else {
			evaluated := evaluator.Eval(program, env)

			if evaluated != nil {
				fmt.Println(evaluated.Inspect())
			}
		}
	}
}
