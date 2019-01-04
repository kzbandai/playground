package token

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// identifier, literals
	IDENTIFIER = "IDENTIFIER" // add, foobar, x, y, ...
	INT        = "INT"

	// operators
	ASSIGN   = "="
	ASTERISK = "*"
	BANG     = "!"
	MINUS    = "-"
	PLUS     = "+"
	SLASH    = "/"

	LT = "<"
	GT = ">"

	EQUAL     = "=="
	NOT_EQUAL = "!="

	// delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// keywords
	ELSE     = "ELSE"
	FALSE    = "FALSE"
	FUNCTION = "FUNCTION"
	IF       = "IF"
	LET      = "LET"
	RETURN   = "RETURN"
	TRUE     = "TRUE"
)

type TokenType string

type Token struct {
	Literal string
	Type    TokenType
}

var keywords = map[string]TokenType{
	"else":   ELSE,
	"false":  FALSE,
	"fn":     FUNCTION,
	"if":     IF,
	"let":    LET,
	"return": RETURN,
	"true":   TRUE,
}

func LookupIdentifier(identifier string) TokenType {
	if tok, ok := keywords[identifier]; ok {
		return tok
	}
	return IDENTIFIER
}
