// Package ast provides types and methods for creating and manipulating an abstract syntax
// tree representing a capuchin program.
package ast

import (
	"bytes"
	"capuchin/token"
)

// Node is a base interface for all AST elements.
type Node interface {
	TokenLiteral() string
	String() string
}

// Statement is the base interface for all AST nodes which represent statements.
type Statement interface {
	Node
	statementNode()
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

// String will return a string representation of the Program
func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
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
func (i *Identifier) String() string {
	return i.Value
}

// LetStatement represents a "let" token, storing its identifier and related expression.
type LetStatement struct {
	Token token.Token // The token.LET token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}
func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}

	out.WriteString(";")

	return out.String()
}

// ReturnStatement represents a "return" token.
type ReturnStatement struct {
	Token       token.Token // The 'return' token
	ReturnValue Expression  // The expression that will be returned
}

func (rs *ReturnStatement) statementNode() {}
func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}
func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")
	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteString(";")

	return out.String()
}

// ExpressionStatement represents an expression, such as "x + 5".
type ExpressionStatement struct {
	Token      token.Token // The first token of the expression
	Expression Expression
}

func (es *ExpressionStatement) statementNode() {}
func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}
