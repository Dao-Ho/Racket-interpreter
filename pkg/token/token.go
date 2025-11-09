package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT = "IDENT"
	INT   = "INT"

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	MULTIPLY = "*"
	DIVIDE   = "/"

	// Delimiters
	LEFT_PAREN  = "("
	RIGHT_PAREN = ")"

	// Keywords
	FUNCTION = "FUNCTION"
	DEFINE   = "DEFINE"
)
