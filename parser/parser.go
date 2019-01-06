package parser

import (
	"capuchin/ast"
	"capuchin/lexer"
	"capuchin/token"
	"fmt"
)

// Parser reads tokens from the supplied lexer into an abstract syntax tree.
type Parser struct {
	lex       *lexer.Lexer
	errors    []string
	curToken  token.Token
	peekToken token.Token
}

// New creates a new Parser which reads tokens from the supplied lexer.
func New(lex *lexer.Lexer) *Parser {
	p := &Parser{
		lex:    lex,
		errors: []string{},
	}

	//Read two tokens, so curToken and peekToken are both set
	p.nextToken()
	p.nextToken()

	return p
}

// Errors returns an array of error messages for errors encountered by the parser.
func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.lex.NextToken()
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	default:
		return nil
	}
}

// ParseProgram loops through the tokens provided by the lexer this parsers is attached to
// and creates an abstract syntax tree of the program the tokens describe.
func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for !p.curTokenIs(token.EOF) {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}

	return program
}

// parseLetStatement handles the creation of LetStatement nodes in the abstract syntax
// tree.
func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.curToken}

	// If the next token is not an identifier (eg the x in "let x = 6;") then this is
	// not a valid let statement.
	if !p.expectPeek(token.IDENT) {
		return nil
	}

	// If the next token was an identifier then create an Identifier node for the AST.
	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	// If the next statement is not an assignment token (eg "=") then this is not a
	// valid let statement.
	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	// TODO: for now we just skip to the end.
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt

}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}

func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, but got %s instead.",
		t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}
