package lexer

import "racket-interpreter/pkg/token"

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

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
}

func (l *Lexer) NextToken() token.Token {
	var token token.Token
	switch l.ch {
	case '(':
		tok = NewToken(token.LEFT_PAREN, l.ch)
	}
}
