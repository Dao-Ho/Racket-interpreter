package lexer

import (
	"testing"
)

func TestNextToken_SingleSymbol(t *testing.T) {
	input := "+"

	lexer := NewLexer(input)

	token := lexer.NextToken()

	if token.Type != SYMBOL {
		t.Fatalf("token.Type wrong. expected=%q, got=%q", SYMBOL, token.Type)
	}

	if token.Literal != "+" {
		t.Fatalf("token.Literal wrong. expected=%q, got=%q", "+", token.Literal)
	}
}

func TestNextToken_Expression(t *testing.T) {
	input := "(+ 5 2)"

	lexer := NewLexer(input)

	token := lexer.NextToken()

	if token.Type != LEFT_PAREN {
		t.Fatalf("token.Type wrong. expected=%q, got=%q", LEFT_PAREN, token.Type)
	}

	if token.Literal != "(" {
		t.Fatalf("token.Literal wrong. expected=%q, got=%q", "(", token.Literal)
	}

	token = lexer.NextToken()

	if token.Type != SYMBOL {
		t.Fatalf("token.Type wrong. expected=%q, got=%q", SYMBOL, token.Type)
	}

	if token.Literal != "+" {
		t.Fatalf("token.Literal wrong. expected=%q, got=%q", "+", token.Literal)
	}

	token = lexer.NextToken()

	if token.Type != INTEGER {
		t.Fatalf("token.Type wrong. expected=%q, got=%q", INTEGER, token.Type)
	}

	if token.Literal != "5" {
		t.Fatalf("token.Literal wrong. expected=%q, got=%q", "5", token.Literal)
	}

	token = lexer.NextToken()

	if token.Type != INTEGER {
		t.Fatalf("token.Type wrong. expected=%q, got=%q", INTEGER, token.Type)
	}

	if token.Literal != "2" {
		t.Fatalf("token.Literal wrong. expected=%q, got=%q", "2", token.Literal)
	}

	token = lexer.NextToken()

	if token.Type != RIGHT_PAREN {
		t.Fatalf("token.Type wrong. expected=%q, got=%q", RIGHT_PAREN, token.Type)
	}

	if token.Literal != ")" {
		t.Fatalf("token.Literal wrong. expected=%q, got=%q", ")", token.Literal)
	}

	token = lexer.NextToken()
	if token.Type != EOF {
		t.Fatalf("token.Type wrong. expected=%q, got=%q", EOF, token.Type)
	}
}
