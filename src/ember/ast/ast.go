package ast

import (
	"bytes"
	"strings"

	"github.com/sckelemen/ember/src/ember/token"
)

type Node interface {
	TokenLiteral() string
	String() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}

	return ""
}

func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

type LetStatement struct {
	Token token.Token // token.LET
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.String() }
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

type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.String() }
func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")

	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteString(";")

	return out.String()
}

type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

func (es *ExpressionStatement) statementNode()       {}
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.String() }
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

type BlockStatement struct {
	Token      token.Token // the { token
	Statements []Statement
}

func (bs *BlockStatement) statementNode()       {}
func (bs *BlockStatement) TokenLiteral() string { return bs.Token.String() }
func (bs *BlockStatement) String() string {
	var out bytes.Buffer

	for _, s := range bs.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

type Identifier struct {
	Token token.Token // token.IDENT
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.String() }
func (i *Identifier) String() string {
	return i.Value
}

// literals

type Boolean struct {
	Token token.Token
	Value bool
}

func (b *Boolean) expressionNode()      {}
func (b *Boolean) TokenLiteral() string { return b.Token.String() }
func (b *Boolean) String() string       { return b.Token.String() }

type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (il *IntegerLiteral) expressionNode()      {}
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.String() }
func (il *IntegerLiteral) String() string       { return il.Token.String() }

type PrefixExpression struct {
	Token    token.Token
	Operator string
	Right    Expression
}

func (pe *PrefixExpression) expressionNode()      {}
func (pe *PrefixExpression) TokenLiteral() string { return pe.Token.String() }
func (pe *PrefixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")

	return out.String()
}

type InfixExpression struct {
	Token    token.Token
	Left     Expression
	Operator string
	Right    Expression
}

func (e *InfixExpression) expressionNode()      {}
func (e *InfixExpression) TokenLiteral() string { return e.Token.String() }
func (e *InfixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(e.Left.String())
	out.WriteString(" " + e.Operator + " ")
	out.WriteString(e.Right.String())
	out.WriteString(")")

	return out.String()
}

type IfExpression struct {
	Token       token.Token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (e *IfExpression) expressionNode()      {}
func (e *IfExpression) TokenLiteral() string { return e.Token.String() }
func (e *IfExpression) String() string {
	var out bytes.Buffer

	out.WriteString("if")
	out.WriteString(e.Condition.String())
	out.WriteString(" ")
	out.WriteString(e.Consequence.String())

	if e.Alternative != nil {
		out.WriteString("else ")
		out.WriteString(e.Alternative.String())
	}

	return out.String()
}

type FunctionLiteral struct {
	Token      token.Token
	Parameters []*Identifier
	Body       *BlockStatement
}

func (l *FunctionLiteral) expressionNode()      {}
func (l *FunctionLiteral) TokenLiteral() string { return l.Token.String() }
func (l *FunctionLiteral) String() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range l.Parameters {
		params = append(params, p.String())
	}

	out.WriteString(l.TokenLiteral())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") ")
	out.WriteString(l.Body.String())

	return out.String()
}

type InvocationExpression struct {
	Token     token.Token
	Function  Expression
	Arguments []Expression
}

func (e *InvocationExpression) expressionNode()      {}
func (e *InvocationExpression) TokenLiteral() string { return e.Token.String() }
func (e *InvocationExpression) String() string {
	var out bytes.Buffer

	args := []string{}
	for _, a := range e.Arguments {
		args = append(args, a.String())
	}

	out.WriteString(e.Function.String())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")

	return out.String()
}

type StringLiteral struct {
	Token token.Token
	Value string
}

func (l *StringLiteral) expressionNode()      {}
func (l *StringLiteral) TokenLiteral() string { return l.Token.String() }
func (l *StringLiteral) String() string       { return l.Token.String() }

type ArrayLiteral struct {
	Token    token.Token
	Elements []Expression
}

func (l *ArrayLiteral) expressionNode()      {}
func (l *ArrayLiteral) TokenLiteral() string { return l.Token.String() }
func (l *ArrayLiteral) String() string {
	var out bytes.Buffer

	elements := []string{}
	for _, el := range l.Elements {
		elements = append(elements, el.String())
	}

	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")

	return out.String()
}

type IndexExpression struct {
	Token token.Token
	Left  Expression
	Index Expression
}

func (e *IndexExpression) expressionNode()      {}
func (e *IndexExpression) TokenLiteral() string { return e.Token.String() }
func (e *IndexExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(e.Left.String())
	out.WriteString("[")
	out.WriteString(e.Index.String())
	out.WriteString("])")

	return out.String()
}

type HashLiteral struct {
	Token token.Token
	Pairs map[Expression]Expression
}

func (l *HashLiteral) expressionNode()      {}
func (l *HashLiteral) TokenLiteral() string { return l.Token.String() }
func (l *HashLiteral) String() string {
	var out bytes.Buffer

	pairs := []string{}
	for key, value := range l.Pairs {
		pairs = append(pairs, key.String()+":"+value.String())
	}

	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")

	return out.String()
}
