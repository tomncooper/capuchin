package lexer

import (
	"capuchin/token"
)

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

	// Advance the lexer past any whitespace characters
	l.skipWhitespace()

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
	default:
		// If non of the switch statements are triggered see if this is
		// a letter.
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else {
			// If it is not a switch or letter then it is not valid
			tok = newToken(token.ILLEGAL, l.ch)
		}
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

// readIdentifier will iterate trough the characters in the lexer's input
// string until it the isLetter function returns false. It will then return the
// characters from the initial position in the input string up to the index
// where isLetter returned false.
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// isLetter will test if the supplied character byte is a letter (upper or
// lower case as well as '_').
func isLetter(ch byte) bool {

	letter := false

	if 'a' <= ch && ch <= 'z' {
		letter = true
	} else if 'A' <= ch && ch <= 'Z' {
		letter = true
	} else if ch == '_' {
		letter = true
	}

	return letter
}

// isWhitespace checks weather the supplied character bytes correspond to a
// whitespace character.
func isWhitespace(ch byte) bool {

	whitespace := false

	if ch == ' ' {
		whitespace = true
	} else if ch == '\t' {
		whitespace = true
	} else if ch == '\n' {
		whitespace = true
	} else if ch == '\r' {
		whitespace = true
	}

	return whitespace
}

// skipWhitespace will advance the lexer along its input string until a non
// whitespace character is found.
func (l *Lexer) skipWhitespace() {

	for isWhitespace(l.ch) {
		l.readChar()
	}
}
