package lexer

import (
	"github.com/kzbandai/playground/go/src/interpreter/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `let five = 5;
let ten = 10;

let add = fn(x, y) {
	x + y;
};

let result = add(five, ten);
!-/*5;
5 < 10 > 5;

if (5 < 10) {
	return true;
} else {
	return false;
}

10 == 10;
10 != 9;
`

	tests := []struct {
		expectedLiteral string
		expectedType    token.TokenType
	}{
		{"let", token.LET},
		{"five", token.IDENTIFIER},
		{"=", token.ASSIGN},
		{"5", token.INT},
		{";", token.SEMICOLON},

		{"let", token.LET},
		{"ten", token.IDENTIFIER},
		{"=", token.ASSIGN},
		{"10", token.INT},
		{";", token.SEMICOLON},

		{"let", token.LET},
		{"add", token.IDENTIFIER},
		{"=", token.ASSIGN},
		{"fn", token.FUNCTION},
		{"(", token.LPAREN},
		{"x", token.IDENTIFIER},
		{",", token.COMMA},
		{"y", token.IDENTIFIER},
		{")", token.RPAREN},
		{"{", token.LBRACE},
		{"x", token.IDENTIFIER},
		{"+", token.PLUS},
		{"y", token.IDENTIFIER},
		{";", token.SEMICOLON},
		{"}", token.RBRACE},
		{";", token.SEMICOLON},

		{"let", token.LET},
		{"result", token.IDENTIFIER},
		{"=", token.ASSIGN},
		{"add", token.IDENTIFIER},
		{"(", token.LPAREN},
		{"five", token.IDENTIFIER},
		{",", token.COMMA},
		{"ten", token.IDENTIFIER},
		{")", token.RPAREN},
		{";", token.SEMICOLON},

		{"!", token.BANG},
		{"-", token.MINUS},
		{"/", token.SLASH},
		{"*", token.ASTERISK},
		{"5", token.INT},
		{";", token.SEMICOLON},

		{"5", token.INT},
		{"<", token.LT},
		{"10", token.INT},
		{">", token.GT},
		{"5", token.INT},
		{";", token.SEMICOLON},

		{"if", token.IF},
		{"(", token.LPAREN},
		{"5", token.INT},
		{"<", token.LT},
		{"10", token.INT},
		{")", token.RPAREN},
		{"{", token.LBRACE},

		{"return", token.RETURN},
		{"true", token.TRUE},
		{";", token.SEMICOLON},

		{"}", token.RBRACE},
		{"else", token.ELSE},
		{"{", token.LBRACE},

		{"return", token.RETURN},
		{"false", token.FALSE},
		{";", token.SEMICOLON},

		{"}", token.RBRACE},

		{"10", token.INT},
		{"==", token.EQUAL},
		{"10", token.INT},
		{";", token.SEMICOLON},

		{"10", token.INT},
		{"!=", token.NOT_EQUAL},
		{"9", token.INT},
		{";", token.SEMICOLON},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - token type is wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal is wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}
