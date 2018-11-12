package lexer

import (
	"unicode"
	"unicode/utf8"

	"github.com/sckelemen/ember/src/ember/token"
)

type Scanner struct {
	File      string
	Directory string
	Errors    []ScanError
	Lexemes   []Lexeme
	src []byte

	ch       rune
	offset   int
	rdOffset int
	lnOffset int
}

func (s *Scanner) Scan() {
	while s.rdOffset < len(s.src) {
		
		
	}
}


func (s *Scanner) next() {

}

// base types

func (s *Scanner) scanRune() string {

}

func (s *Scanner) scanString() string {

}

func (s *Scanner) scanNumber() string {

}

func (s *Scanner) scanComment() string {
	// scan comments

	// scan for `//` comments

}

func (s *Scanner) scanIdentifier() string {
	// identifiers start with
	// _
	// A-Z
	// a-z
	// unicode
	// and continues to the next whitespace or op

	offs := s.offset                      // mark the start range of the ident string
	for isLetter(s.ch) || isDigit(s.ch) { // scan letters or digits
		s.next() // advance through letters or digits
	}
	return string(s.src[offs:s.offset]) // return the range when its not a digit or letter
}

func (s *Scanner) consumeWhiteSpace() {
	for s.ch == ' ' || s.ch == '\t' || s.ch == '\n' && !s.insertSemi || s.ch == '\r' {
		s.next()
	}
}

// helpers

func isLetter(ch rune) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_' || ch >= utf8.RuneSelf && unicode.IsLetter(ch)
}

func isDigit(ch rune) bool {
	return '0' <= ch && ch <= '9' || ch >= utf8.RuneSelf && unicode.IsDigit(ch)
}

func digitVal(ch rune) int {
	switch {
	case '0' <= ch && ch <= '9':
		return int(ch - '0')
	case 'a' <= ch && ch <= 'f':
		return int(ch - 'a' + 10)
	case 'A' <= ch && ch <= 'F':
		return int(ch - 'A' + 10)
	}
	return 16 // larger than any legal digit val
}

// built in types
// null
// false
// true
// numbers

/*
cases:
 current	next	 	state
 =			space		assign
 =			=			eql
 =			>			lambda
 !			ident		negate
 !			= 			neq
 /			space		quo
 /			ident		quo
 /			literal		quo
 /			/			lncomment
 /			*			mlcomment



*/

type Lexeme struct {
	Token token.Token
	Value string
	Source   []byte
}

func CreateLexeme(token: token.Token, value: string) Lexeme {
	return CreateLexemeWithSource(token: token, value: value, source: value)
}

func CreateLexemeWithSource(token: token.Token, value: string, source: []byte) Lexeme {
	return Lexeme{Token: token, Value: value, Source: source}
}