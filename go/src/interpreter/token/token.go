package token

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// identifier, literals
	IDENTIFIER = "IDENTIFIER" // add, foobar, x, y, ...
	INT        = "INT"

	// operators
	ASSIGN = "="
	PLUS   = "+"

	// delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

type TokenType string

type Token struct {
	Literal string
	Type    TokenType
}
