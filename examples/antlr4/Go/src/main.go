// main.go
// bazel run //antlr4/Go/src:main
package main

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"

	"github.com/example/com/parser/json"
)

func demoLexer(input string) {
	// Setup the input
	is := antlr.NewInputStream(input)

	// Create the Lexer
	lexer := json.NewJSONLexer(is)

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
	*json.BaseJSONListener
}

func (l *jsonListener) EnterObj(ctx *json.ObjContext) {
	fmt.Printf("Object: %s\n", ctx.GetText())
}

func (l *jsonListener) EnterPair(ctx *json.PairContext) {
	fmt.Printf("Pair: %s\n", ctx.GetText())
}

func (l *jsonListener) EnterArray(ctx *json.ArrayContext) {
	fmt.Printf("Array: %s\n", ctx.GetText())
}

func demoParser(input string) {
	// Setup the input
	is := antlr.NewInputStream(input)

	lexer := json.NewJSONLexer(is)
	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := json.NewJSONParser(tokenStream)

	// Finally parse the expression (by walking the tree)
	var listener jsonListener
	tree := p.Json()
	antlr.ParseTreeWalkerDefault.Walk(&listener, tree)
}

func main() {
	input := `[1, {"a": [2, 3, 4, 4.53, {"b":  "c"}]}]`
	demoLexer(input)
	demoParser(input)
}
