package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/uesteibar/lainoa/pkg/lexer"
	"github.com/uesteibar/lainoa/pkg/parser"
)

const PROMPT = "⛅️ >> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprint(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)
		program := p.ParseProgram()

		if len(p.Errors()) == 0 {
			fmt.Fprintf(out, "%s\n", program)
		} else {
			fmt.Fprintln(out, "Oops! Something is wrong here:")
			fmt.Fprintln(out, "  parser errors:")
			for _, err := range p.Errors() {
				fmt.Fprintf(out, "- %s\n", err)
			}
		}
	}
}
