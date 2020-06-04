package runner

import (
	"fmt"
	"io/ioutil"

	"github.com/uesteibar/lainoa/pkg/evaluator"
	"github.com/uesteibar/lainoa/pkg/lexer"
	"github.com/uesteibar/lainoa/pkg/object"
	"github.com/uesteibar/lainoa/pkg/parser"
)

func Start(filepath string) {
	data, err := ioutil.ReadFile(filepath)

	if err != nil {
		fmt.Println("Error reading file", err)
	}

	l := lexer.New(string(data), filepath)
	p := parser.New(l)
	program := p.ParseProgram()

	if len(p.Errors()) > 0 {
		fmt.Println("Oops! Something is wrong here:")
		fmt.Println("  parser errors:")
		for _, err := range p.Errors() {
			fmt.Println(fmt.Sprintf("- %s\n", err.String()))
		}
	} else {
		env := object.NewEnvironment()
		evaluated := evaluator.Eval(program, env)
		if object.IsError(evaluated) {
			fmt.Println(evaluated.Inspect())
		}
	}
}
