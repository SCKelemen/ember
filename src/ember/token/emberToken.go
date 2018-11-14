package token

import (
	"strconv"
)

type Token int

const (
	ILLEGAL Token = iota
	EOF
	COMMENT
	TRIVIA

	literal_beg
	IDENT  // main
	INT    // 12345
	FLOAT  // 123.45
	IMAG   // 123.45i
	CHAR   // 'a'
	STRING // "abc"
	literal_end

	arithmetic_op_beg
	ADD // +
	SUB // -
	MUL // *
	QUO // /
	REM // %
	EXP // **
	arithmetic_op_end

	arithmetic_asn_beg
	ADD_ASSIGN // +=
	SUB_ASSIGN // -=
	MUL_ASSIGN // *=
	QUO_ASSIGN // /=
	REM_ASSIGN // %=
	EXP_ASSIGN // **=
	arithmetic_asn_end

	bit_op_beg
	AND     // &
	OR      // |
	XOR     // ^
	SHL     // <<
	SHR     // >>
	AND_NOT // &^		perhaps this should be &!
	bit_op_end

	bit_asn_beg
	AND_ASSIGN     // &=
	OR_ASSIGN      // |=
	XOR_ASSIGN     // ^=
	SHL_ASSIGN     // <<=
	SHR_ASSIGN     // >>=
	AND_NOT_ASSIGN // &^= 		perhaps this should be &!=
	bit_asn_end

	logic_op_beg
	LAND // &&
	LOR  // ||
	INC  // ++
	DEC  // --
	logic_op_end

	equality_op_beg
	EQL // ==
	LT  // <
	GT  // >
	NOT // !

	NEQ // !=
	LTE // <=
	GTE // >=
	equality_op_end

	ops_beg
	ASSIGN   // =
	DEFINE   // :=
	ELLIPSIS // ...
	DBLDOT   // ..
	PIPE     // |>

	LPAREN // (
	RPAREN // )
	LBRACK // [
	RBRACK // ]
	LBRACE // {
	RBRACE // }

	DOT   // .
	COMMA // ,
	SEMI  // ;
	COLON // :
	ops_end

	keyword_beg
	// decl
	LET   // let
	CONST // const

	// conditionals
	IF   // if
	ELSE // else

	SWITCH  // switch
	CASE    // case
	DEFAULT // default

	// loops
	WHILE // while
	DO    // do
	FOR   // for
	BREAK // break
	RANGE // range

	// functions
	FUNC   // func
	RETURN // return
	YIELD  // yield

	// types
	TYPE      // type
	STRUCT    // struct
	CLASS     // class
	ACTOR     // actor
	INTERFACE // interface
	ABSTRACT  // abstract

	// access modifiers
	PUBLIC    // public
	PRIVATE   // private
	PROTECTED // protected
	DEFER     // defer
	keyword_end
)

var tokens = [...]string{

	ILLEGAL: "ILLEGAL",
	EOF:     "EOF",
	COMMENT: "COMMENT",
	TRIVIA:  "TRIVIA",

	IDENT:  "IDENT",
	INT:    "INT",
	FLOAT:  "FLOAT",
	IMAG:   "IMAG",
	CHAR:   "CHAR",
	STRING: "STRING",

	ADD: "+",
	SUB: "-",
	MUL: "*",
	QUO: "/",
	REM: "%",
	EXP: "**",

	ADD_ASSIGN: "+=",
	SUB_ASSIGN: "-=",
	MUL_ASSIGN: "*=",
	QUO_ASSIGN: "/=",
	REM_ASSIGN: "%=",
	EXP_ASSIGN: "**=",

	AND:     "&",
	OR:      "|",
	XOR:     "^",
	SHL:     "<<",
	SHR:     ">>",
	AND_NOT: "&^",

	AND_ASSIGN:     "&=",
	OR_ASSIGN:      "|=",
	XOR_ASSIGN:     "^=",
	SHL_ASSIGN:     "<<=",
	SHR_ASSIGN:     ">>=",
	AND_NOT_ASSIGN: "&^=",

	LAND: "&&",
	LOR:  "||",
	INC:  "++",
	DEC:  "--",

	EQL: "==",
	LT:  "<",
	GT:  ">",
	NOT: "!",

	NEQ: "!=",
	LTE: "<=",
	GTE: ">=",

	ASSIGN:   "=",
	DEFINE:   ":=",
	ELLIPSIS: "...",
	DBLDOT:   "..",
	PIPE:     "|>",

	LPAREN: "(",
	RPAREN: ")",
	LBRACK: "[",
	RBRACK: "]",
	LBRACE: "{",
	RBRACE: "}",

	DOT:   ".",
	COMMA: ",",
	SEMI:  ";",
	COLON: ":",

	LET:   "let",
	CONST: "const",

	IF:   "if",
	ELSE: "else",

	SWITCH:  "switch",
	CASE:    "case",
	DEFAULT: "default",

	WHILE: "while",
	DO:    "do",
	FOR:   "for",
	BREAK: "break",
	RANGE: "range",

	FUNC:   "func",
	RETURN: "return",
	YIELD:  "yield",

	TYPE:      "type",
	STRUCT:    "struct",
	CLASS:     "class",
	ACTOR:     "actor",
	INTERFACE: "interface",
	ABSTRACT:  "abstract",

	PUBLIC:    "public",
	PRIVATE:   "private",
	PROTECTED: "protected",
	DEFER:     "defer",
}

func (tok Token) String() string {
	s := ""
	if 0 <= tok && tok < Token(len(tokens)) {
		s = tokens[tok]
	}
	if s == "" {
		s = "token(" + strconv.Itoa(int(tok)) + ")"
	}
	return s
}

const (
	LowestPrec  = 0 // non ops
	UnaryPrec   = 7
	HighestPrec = 8
)

func (op Token) Precedence() int {
	switch op {
	case LOR:
		return 1
	case LAND:
		return 2
	case EQL, NEQ, LT, LTE, GT, GTE:
		return 3
	case ADD, SUB, OR, XOR:
		return 4
	case MUL, QUO, REM, SHL, SHR, AND, AND_NOT:
		return 5
	case EXP:
		return 6
	default:
		return LowestPrec
	}
}

var keywords map[string]Token

func init() {
	keywords = make(map[string]Token)
	for i := keyword_beg + 1; i < keyword_end; i++ {
		keywords[tokens[i]] = i
	}
}

func Lookup(ident string) Token {
	if tok, isKeyword := keywords[ident]; isKeyword {
		return tok
	}
	return IDENT
}

func (tok Token) IsLiteral() bool { return literal_beg < tok && tok < literal_end }

func (tok Token) IsOperator() bool { return arithmetic_op_beg < tok && tok < ops_end }

func (tok Token) IsKeyword() bool { return keyword_beg < tok && tok < keyword_end }
