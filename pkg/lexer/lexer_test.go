package lexer

import (
	"racket-interpreter/pkg/token"
	"testing"
)

func TestNextToken_SingleSymbol(t *testing.T) {
	input := "+"

	lexer := NewLexer(input)

	tok := lexer.NextToken()

	if tok.Type != token.PLUS {
		t.Fatalf("token.Type wrong. expected=%q, got=%q", token.PLUS, tok.Type)
	}

	if tok.Literal != "+" {
		t.Fatalf("token.Literal wrong. expected=%q, got=%q", "+", tok.Literal)
	}
}

func TestNextToken_Expression(t *testing.T) {
	input := "(+ 5 2)"

	lexer := NewLexer(input)

	tok := lexer.NextToken()

	if tok.Type != token.LEFT_PAREN {
		t.Fatalf("token.Type wrong. expected=%q, got=%q", token.LEFT_PAREN, tok.Type)
	}

	if tok.Literal != "(" {
		t.Fatalf("token.Literal wrong. expected=%q, got=%q", "(", tok.Literal)
	}

	tok = lexer.NextToken()

	if tok.Type != token.PLUS {
		t.Fatalf("token.Type wrong. expected=%q, got=%q", token.PLUS, tok.Type)
	}

	if tok.Literal != "+" {
		t.Fatalf("token.Literal wrong. expected=%q, got=%q", "+", tok.Literal)
	}

	tok = lexer.NextToken()

	if tok.Type != token.INT {
		t.Fatalf("token.Type wrong. expected=%q, got=%q", token.INT, tok.Type)
	}

	if tok.Literal != "5" {
		t.Fatalf("token.Literal wrong. expected=%q, got=%q", "5", tok.Literal)
	}

	tok = lexer.NextToken()

	if tok.Type != token.INT {
		t.Fatalf("token.Type wrong. expected=%q, got=%q", token.INT, tok.Type)
	}

	if tok.Literal != "2" {
		t.Fatalf("token.Literal wrong. expected=%q, got=%q", "2", tok.Literal)
	}

	tok = lexer.NextToken()

	if tok.Type != token.RIGHT_PAREN {
		t.Fatalf("token.Type wrong. expected=%q, got=%q", token.RIGHT_PAREN, tok.Type)
	}

	if tok.Literal != ")" {
		t.Fatalf("token.Literal wrong. expected=%q, got=%q", ")", tok.Literal)
	}

	tok = lexer.NextToken()
	if tok.Type != token.EOF {
		t.Fatalf("token.Type wrong. expected=%q, got=%q", token.EOF, tok.Type)
	}
	if tok.Literal != "" {
		t.Fatalf("token.Literal wrong. expected=%q, got=%q", "", tok.Literal)
	}
}

func testTokenIter(t *testing.T) {
	input := "+-)54"

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{{token.PLUS, "+"}, {token.MINUS, "-"}, {token.INT, "5"}, {token.INT, "4"}, {token.EOF, ""}}

	lexer := NewLexer(input)
	for i, test := range tests {
		token := lexer.NextToken()
		if token.Type != test.expectedType {
			t.Fatalf("tests[%d] - token.Type wrong. expected=%q, got=%q", i, test.expectedType, token.Type)
		}

		if token.Literal != test.expectedLiteral {
			t.Fatalf("tests[%d] - token.Literal wrong. expected=%q, got=%q", i, test.expectedLiteral, token.Literal)
		}
	}
}
