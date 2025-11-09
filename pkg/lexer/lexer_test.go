package lexer

import (
	"racket-interpreter/pkg/token"
	"testing"
)

func TestNextToken_SingleSymbol(t *testing.T) {
	input := "+"

	lexer := NewLexer(input)

	token := lexer.NextToken()

	if token.Type != token.PLUS {
		t.Fatalf("token.Type wrong. expected=%q, got=%q", token.PLUS, token.Type)
	}

	if token.Literal != "+" {
		t.Fatalf("token.Literal wrong. expected=%q, got=%q", "+", token.Literal)
	}
}

func TestNextToken_Expression(t *testing.T) {
	input := "(+ 5 2)"

	lexer := NewLexer(input)

	token := lexer.NextToken()

	if token.Type != token.LEFT_PAREN {
		t.Fatalf("token.Type wrong. expected=%q, got=%q", token.LEFT_PAREN, token.Type)
	}

	if token.Literal != "(" {
		t.Fatalf("token.Literal wrong. expected=%q, got=%q", "(", token.Literal)
	}

	token = lexer.NextToken()

	if token.Type != token.PLUS {
		t.Fatalf("token.Type wrong. expected=%q, got=%q", token.PLUS, token.Type)
	}

	if token.Literal != "+" {
		t.Fatalf("token.Literal wrong. expected=%q, got=%q", "+", token.Literal)
	}

	token = lexer.NextToken()

	if token.Type != token.INT {
		t.Fatalf("token.Type wrong. expected=%q, got=%q", token.INT, token.Type)
	}

	if token.Literal != "5" {
		t.Fatalf("token.Literal wrong. expected=%q, got=%q", "5", token.Literal)
	}

	token = lexer.NextToken()

	if token.Type != token.INT {
		t.Fatalf("token.Type wrong. expected=%q, got=%q", token.INT, token.Type)
	}

	if token.Literal != "2" {undefined: token.tokenType (but have TokenType) (compiler UndeclaredImportedName)undefined: token.tokenType (but have TokenType) (compiler UndeclaredImportedName)
		t.Fatalf("token.Literal wrong. expected=%q, got=%q", "2", token.Literal)
	}

	token = lexer.NextToken()

	if token.Type != token.RIGHT_PAREN {
		t.Fatalf("token.Type wrong. expected=%q, got=%q", token.RIGHT_PAREN, token.Type)
	}

	if token.Literal != ")" {
		t.Fatalf("token.Literal wrong. expected=%q, got=%q", ")", token.Literal)
	}

	token = lexer.NextToken()
	if token.Type != token.EOF {
		t.Fatalf("token.Type wrong. expected=%q, got=%q", token.EOF, token.Type)
	}
}

func testTokenIter(t *testing.T) {
	input := "+-)54"

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{{token.PLUS, "+"}, {token.MINUS, "-"}, {token.INT, "5"}, {token.INT, "4"}}
}
