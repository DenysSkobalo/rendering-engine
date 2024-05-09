package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL TokenType = "ILLEGAL"
	EOF     TokenType = "EOF"

	// Single characters
	LEFT_ANGLE  TokenType = "<"
	RIGHT_ANGLE TokenType = ">"

	// Datas
	TEXT TokenType = "TEXT"

	// Tags
	START_TAG        TokenType = "START_TAG"
	END_TAG          TokenType = "END_TAG"
	SELF_CLOSING_TAG TokenType = "SELF_CLOSING_TAG"

	// Attributes
	ATTRIBUTE_NAME  TokenType = "ATTRIBUTE_NAME"
	ATTRIBUTE_VALUE TokenType = "ATTRIBUTE_VALUE"
)
