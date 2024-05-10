package cmd

import (
	"fmt"
	"rendering-engine/html-parser/pkg/lexer"
)

func Init() {
	fmt.Println("HTML Parser")

	input := "<!DOCTYPE html>"
	lex := lexer.NewLexer(input)

	for lex.CurrentCh != 0 {
		fmt.Printf("Current character: %c\n", lex.CurrentCh)
		lex.ReadChar()
	}
}
