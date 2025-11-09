package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL TokenType = "ILLEGAL"
	EOF     TokenType = "EOF"

	IDENT TokenType = "IDENT"
	INT   TokenType = "INT"
	// Operators
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

// mapping given word to tokenType
var keywords = map[string]TokenType{
	"FUNCTION": FUNCTION,
	"DEFINE":   DEFINE,
	"LET":      LET,
	"COND":     COND,
	"IF":       IF,
}

func LookupWord(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
