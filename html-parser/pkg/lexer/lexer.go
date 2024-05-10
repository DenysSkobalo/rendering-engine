package lexer

import "rendering-engine/html-parser/pkg/token"

type Lexer struct {
	input        string
	position     int  // current position in input
	readPosition int  // current reading position
	CurrentCh    byte // current character
}

func NewLexer(input string) *Lexer {
	l := &Lexer{input: input}
	l.ReadChar()
	return l
}

// ReadChar reads the next character from the input data
func (l *Lexer) ReadChar() {
	if l.readPosition >= len(l.input) {
		l.CurrentCh = 0 // 0 is EOF
	} else {
		l.CurrentCh = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

// NextToken parses the next token and returns it
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.CurrentCh {
	case '<':
		tok = newToken(token.LEFT_ANGLE, l.CurrentCh)
	case '>':
		tok = newToken(token.RIGHT_ANGLE, l.CurrentCh)
	default:
		tok = newToken(token.ILLEGAL, l.CurrentCh)
	}

	l.ReadChar()
	return tok
}

// makes it easy to create a new token
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
