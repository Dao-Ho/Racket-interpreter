package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL TokenType = "ILLEGAL"
	EOF     TokenType = "EOF"

	// Identifiers + literals
	IDENT  TokenType = "IDENT"
	INT    TokenType = "INT"
	STRING TokenType = "STRING"
	// Operators
	ASSIGN   TokenType = "="
	PLUS     TokenType = "+"
	MINUS    TokenType = "-"
	MULTIPLY TokenType = "*"
	DIVIDE   TokenType = "/"

	// Delimiters
	LEFT_PAREN  TokenType = "("
	RIGHT_PAREN TokenType = ")"

	// Keywords
	FUNCTION TokenType = "FUNCTION"
	DEFINE   TokenType = "DEFINE"
	LET      TokenType = "LET"
	COND     TokenType = "COND"
	IF       TokenType = "IF"
)

func NewToken(tokenType TokenType, ch byte) Token {
	return Token{Type: tokenType, Literal: string(ch)}
}
