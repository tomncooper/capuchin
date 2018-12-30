// Package ast provides types and methods for creating and manipulating an abstract syntax
// tree representing a capuchin program.
package ast

import "capuchin/token"

// Node is a base interface for all AST elements.
type Node interface {
	TokenLiteral() string
}

// Statement is the base interface for all AST nodes which represent statements.
type Statement interface {
	Node
	statmentNode()
}

// Expression is the base interface for all AST nodes which represent expressions.
type Expression interface {
	Node
	expressionNode()
}

// Program is the root node for a capuchin program AST.
type Program struct {
	Statements []Statement
}

// TokenLiteral will return the Token Literal of the first statement in the program which
// should be the programs name.
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}

	return ""
}

// Identifier represents an IDENT token and its value.
type Identifier struct {
	Token token.Token // The token.IDENT token
	Value string
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

// LetStatement represents a "let" token, storing its identifier and related expression.
type LetStatement struct {
	Token token.Token // The token.LET token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statmentNode() {}
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}
