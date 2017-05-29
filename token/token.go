package token

// TokenType represents the particular type of source code token
type TokenType string

// Token represents a single token from the source code
type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers and literals
	IDENT = "IDENT" // add, foobar, x ,y...
	INT   = "INT"   // Integers 1,2,3,4...

	// Operators
	ASSIGN = "="
	PLUS   = "+"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

// keywords defines the language reserved keywords
var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

// lookupIdent takes the supplied string and check first if it is reserved
// keyword, if so it will return the corresponding token (stored in the
// keywords variable). If the supplied string is not a key word the IDENT token
// will be returned.
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
