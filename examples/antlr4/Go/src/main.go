// main.go
// bazel run //antlr4/Go/src:main
package main

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"

	"github.com/example/com/json/parser"
)

func demoLexer(input string) {
	// Setup the input
	is := antlr.NewInputStream(input)

	// Create the Lexer
	lexer := parser.NewJSONLexer(is)

	// Read all tokens
	for {
		t := lexer.NextToken()
		if t.GetTokenType() == antlr.TokenEOF {
			break
		}
		fmt.Printf("%s (%q)\n",
			lexer.SymbolicNames[t.GetTokenType()], t.GetText())
	}
}

type jsonListener struct {
	*parser.BaseJSONListener
}

func demoParser(input string) {
	// Setup the input
	is := antlr.NewInputStream(input)

	lexer := parser.NewJSONLexer(is)
	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewJSONParser(tokenStream)

	// Finally parse the expression (by walking the tree)
	var listener jsonListener
	antlr.ParseTreeWalkerDefault.Walk(&listener, p.Json())
}

func main() {
	input := `[1, {"a": [2, 3, 4, 4.53, {"b":  "c"}]}]`
	demoLexer(input)
	demoParser(input)
}
