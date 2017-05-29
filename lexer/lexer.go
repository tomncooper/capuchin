package lexer

import "capuchin/token"

// Lexer is the lexical analyser for the Capuchin programming language
type Lexer struct {

	// input is the string being analysed by the lexer.
	input string

	// position is the current location in the input string and points to the
	// char currently being processed.
	position int

	// readPosition points to the next char to be read by the lexer.
	readPosition int

	// ch holds the current char being processed
	ch byte
}

// New creates an Lexer instance using the supplied input string.
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar() // Set the initial values within the Lexer
	return l
}

// readChar reads the next character in the lexer's input string and advances
// the position and readPosition variables. If the end of the input string is
// reached then the ch variable will be set to 0 (zero).
func (l *Lexer) readChar() {

	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++

}

// NextToken reads the next character in the Lexer's input strings and returns
// a token.Token instance for the associated token.
func (l *Lexer) NextToken() token.Token {

	var tok token.Token

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	}

	// Advance the lexer to the next character of the input string
	l.readChar()
	return tok
}

// newToken is a helper function which initialises the token.Token instances
// using the supplied Type and character bytes.
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
