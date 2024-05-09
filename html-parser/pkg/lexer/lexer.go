package lexer

import "rendering-engine/html-parser/pkg/token"

type Lexer struct {
	input        string
	position     int  // current position in input
	readPosition int  // current reading position
	ch           byte // current character
}

func NewLexer(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// reads the next character from the input data
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 // 0 is EOF
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

// NextToken parses the next token and returns it
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.ch {
	case '<':
		tok = newToken(token.LEFT_ANGLE, l.ch)
	case '>':
		tok = newToken(token.RIGHT_ANGLE, l.ch)
	default:
		tok = newToken(token.ILLEGAL, l.ch)
	}

	l.readChar()
	return tok
}

// makes it easy to create a new token
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
