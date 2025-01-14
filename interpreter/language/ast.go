package language

import (
	"bytes"
	"strings"
)

// DynamoExpression the root node of the AST
type DynamoExpression struct {
	Statement Statement
}

func (de *DynamoExpression) String() string {
	var out bytes.Buffer

	out.WriteString(de.Statement.String())

	return out.String()
}

// TokenLiteral returns the literal token of the node
func (de *DynamoExpression) TokenLiteral() string {
	return de.Statement.TokenLiteral()
}

// Node the AST node type
type Node interface {
	TokenLiteral() string
	String() string
}

// Statement represents the node type statement
type Statement interface {
	Node
	statementNode()
}

// Expression represents the node type expression
type Expression interface {
	Node
	expressionNode()
}

// Identifier identifier expression node
type Identifier struct {
	Token Token // the token.IDENT token
	Value string
}

func (i *Identifier) expressionNode() {
	_ = 1 // HACK for passing coverage
}

// TokenLiteral returns the literal token of the node
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

func (i *Identifier) String() string { return i.Value }

// ExpressionStatement is the expression node
type ExpressionStatement struct {
	Token      Token // the return token
	Expression Expression
}

func (es *ExpressionStatement) statementNode() {
	_ = 1 // HACK for passing coverage
}

// TokenLiteral returns the literal token of the node
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }

func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}

	return ""
}

// PrefixExpression prefix operator expression
type PrefixExpression struct {
	Token    Token // The prefix token, e.g. NOT
	Operator string
	Right    Expression
}

func (pe *PrefixExpression) expressionNode() {
	_ = 1 // HACK for passing coverage
}

// TokenLiteral returns the literal token of the node
func (pe *PrefixExpression) TokenLiteral() string {
	return pe.Token.Literal
}

func (pe *PrefixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")

	return out.String()
}

// InfixExpression infix operator expression
type InfixExpression struct {
	Token    Token // The operator token, e.g. =
	Left     Expression
	Operator string
	Right    Expression
}

func (oe *InfixExpression) expressionNode() {
	_ = 1 // HACK for passing coverage
}

// TokenLiteral returns the literal token of the node
func (oe *InfixExpression) TokenLiteral() string {
	return oe.Token.Literal
}

func (oe *InfixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(oe.Left.String())
	out.WriteString(" " + oe.Operator + " ")
	out.WriteString(oe.Right.String())
	out.WriteString(")")

	return out.String()
}

// CallExpression function call expression
type CallExpression struct {
	Token    Token // The '(' token
	Function Expression
	// Identifier or FunctionLiteral
	Arguments []Expression
}

func (ce *CallExpression) expressionNode() {
	_ = 1 // HACK for passing coverage
}

// TokenLiteral returns the literal token of the node
func (ce *CallExpression) TokenLiteral() string {
	return ce.Token.Literal
}

func (ce *CallExpression) String() string {
	var out bytes.Buffer

	args := []string{}
	for _, a := range ce.Arguments {
		args = append(args, a.String())
	}

	out.WriteString(ce.Function.String())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")

	return out.String()
}

// BetweenExpression function between expression
type BetweenExpression struct {
	Token Token // The 'BETWEEN' token
	Left  Expression
	// Identifiers
	Range [2]Expression
}

func (ce *BetweenExpression) expressionNode() {
	_ = 1 // HACK for passing coverage
}

// TokenLiteral returns the literal token of the node
func (ce *BetweenExpression) TokenLiteral() string {
	return ce.Token.Literal
}

func (ce *BetweenExpression) String() string {
	var out bytes.Buffer

	out.WriteString(ce.Left.String())
	out.WriteString(" BETWEEN ")
	out.WriteString(ce.Range[0].String())
	out.WriteString(" AND ")
	out.WriteString(ce.Range[1].String())

	return out.String()
}
