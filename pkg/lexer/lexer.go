package lexer

import (
	"racket-interpreter/pkg/token"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func NewLexer(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// process the token and advance the read cursor
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

// Process the token and advance to the next
func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	// skip all whitespace ahead
	l.skipWhitespace()
	switch l.ch {
	case '(':
		tok = token.NewToken(token.LEFT_PAREN, l.ch)
	case ')':
		tok = token.NewToken(token.RIGHT_PAREN, l.ch)
	case '+':
		tok = token.NewToken(token.PLUS, l.ch)
	case '-':
		tok = token.NewToken(token.MINUS, l.ch)
	case '*':
		tok = token.NewToken(token.MULTIPLY, l.ch)
	case '/':
		tok = token.NewToken(token.DIVIDE, l.ch)
	case 0:
		tok.Type = token.EOF
		tok.Literal = ""
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupWord(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Literal = l.readNumber()
			tok.Type = token.INT
			return tok
		} else {
			tok = token.NewToken(token.ILLEGAL, l.ch)
		}
	}
	l.readChar()
	return tok
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	// continue to read char while it's a valid char to process as one
	// word
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// continue reading number until end of number found
func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// check if a token is a legal letter
func isLetter(char byte) bool {
	return 'a' <= char && char <= 'z' || 'A' <= char && char <= 'Z' || char == '_'
}

// check if valid digit
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

// skip whitespace
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}
