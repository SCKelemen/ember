package parser

import (
	"github.com/sckelemen/ember/src/ember/ast"
	"github.com/sckelemen/ember/src/ember/lexer"
	"github.com/sckelemen/ember/src/ember/token"
)

type Parser struct {
	lexer *lexer.Lexer

	currentToken token.Token
	peekToken    token.Token
}

func New(lexer *lexer.Lexer) *Parser {
	p := &Parser{lexer: lexer}

	// load first two tokens
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.currentToken = p.peekToken
	p.peekToken = p.lexer.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	return nil
}

// page
/*

The Parser has three fields: l, curToken and peekToken. l is a pointer to an instance of the lexer, on which we repeatedly call NextToken() to get the next token in the input.  curToken and peekToken act exactly like the two “pointers” our lexer has: position and readPosition. But instead of pointing to a character in the input, they point to the current and the next token. Both are important: we need to look at the curToken, which is the current token under examination, to decide what to do next, and we also need  peekToken for this decision if curToken doesn’t give us enough information. Think of a single line only containing 5;. Then curToken is a token.INT and we need peekToken to decide whether we are at the end of the line or if we are at just the start of an arithmetic expression.
*/
