package lexer

import "github.com/sckelemen/ember/src/ember/token"

type DLexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

/* special chars

+		ADD
+=		ADD_ASSIGN
++		INC

-		SUB
-=		SUB_ASSIGN
--		DEC

*		MUL
*=		MUL_ASSIGN
**		EXP
**=		EXP_ASSIGN
*\/		MLCOMMENT_END

/		QUO
/=		QUO_ASSIGN
//		LNCOMMENT
/* 		MLCOMMENT_BEG

%		REM
%=		REM_ASSIGN

&		AND
&^		AND_NOT
&=		AND_ASSIGN
&^= 	AND_NOT_ASSIGN
&&		LAND

|		OR
|=		OR_ASSIGN
||		LOR
|>		PIPE

^		XOR
^=		XOR_ASSIGN

<		LT
<=		LTE
<<		SHL
<<=		SHL_ASSIGN

>		GT
>=		GTE
>>		SHR
>>=		SHR_ASSIGN

=		ASSIGN
==		EQL

!		NOT
!=		NEQL

. 		DOT
.. 		DBLDOT
... 	ELIPSIS

,		COMMA
;		SEMI

:		COLON
:=		DEFINE

(		LP
)		RP
[		LB
]		RB
{		LBR
}		RBR

*/

func (l *Lexer) ScanNext() token.Token {
	var tok token.Token

	switch l.ch {
	case '+':
		l.ScanPlus()
		break
	case '-':
		l.ScanMinus()
		break
	case '*':
		l.ScanStar()
		break
	case '/':
		l.ScanSlash()
		break
	case '%':
		l.ScanPercent()
		break
	case '&':
		l.ScanAmpersand()
		break
	case '|':
		l.ScanPipeBar()
		break
	case '^': 
		l.ScanCarat()
		break
	case '<':
		l.ScanLeftChevron()
		break
	case '>':
		l.ScanRightChevron()
		break
	case '=':
		l.ScanEqual()
		break
	case '!':
		l.ScanBang()
		break
	case '.':
		l.ScanDot()
		break
	case ':': 
		l.ScanColon() 
		break
	case ';':
		CreateLexeme(token.SEMI, l.ch)
		break
	case ',':
		CreateLexeme(token.COMMA, l.ch)
		break
	case '(':
		CreateLexeme(token.LPAREN, l.ch)
		break
	case ')':
		CreateLexeme(token.RPAREN, l.ch)
		break
	case '[':
		CreateLexeme(token.LBRACK, l.ch)
		break
	case ']':
		CreateLexeme(token.RBRACK, l.ch)
		break
	case '{':
		CreateLexeme(token.LBRACE, l.ch)
		break
	case '}':
		CreateLexeme(token.RBRACE, l.ch)
		break
	}
}

func (l *Lexer) ScanPlus() token.Token {
	/*
		+		ADD
		++		INC
		+=		ADD_ASSIGN
	*/
	switch l.peekChar() {
	case '+':
		l.readChar()
		return token.INC
	case '=':
		l.readChar()
		return token.ADD_ASSIGN
	default:
		return token.ADD
	}
}

func (l *Lexer) ScanMinus() token.Token {
	/*
		-		SUB
		--		DEC
		-=		SUB_ASSIGN
	*/
	switch l.peekChar() {
	case '-':
		l.readChar()
		return token.DEC
	case '=':
		l.readChar()
		return token.SUB_ASSIGN
	default:
		return token.SUB
	}
}

func (l *Lexer) ScanStar() token.Token {
	/*
		*		MUL
		**		EXP
		**= 	EXP_ASSIGN
		*=		MUL_ASSIGN
	*/
	switch l.peekChar() {
	case '*': // **
		l.readChar()
		nchar := l.peekChar()
		if nchar == '=' // is it **= ?
			l.readChar()
			return token.EXP_ASSIGN 
		return token.EXP // nope, it's **
	case '=': // *=
		l.readChar()
		return token.MUL_ASSIGN
		break
	//case '/':
	//	break
	default:
		return token.MUL
	}
}

func (l *Lexer) ScanSlash() token.Token {
	/*
		/		QUO
		/=		QUO_ASSIGN
		//		LNCOMMENT
		\/* 	MLCOMMENT_BEG
	*/
	switch l.peekChar() {
	case '=':
		l.readChar()
		return token.QUO_ASSIGN
	case '/', '*':
		// consume this char from ScanComment
		return l.ScanComment()
	default:
		return token.QUO
	}

}

func (l *Lexer) ScanPercent() token.Token {
	/*
		%		REM
		%=		REM_ASSIGN
	*/
	if l.peekChar() == '=' {
		l.readChar()
		return token.REM_ASSIGN
	}
	return token.REM
}

func (l *Lexer) ScanAmpersand() token.Token {
	/*
		&		AND
		&^		AND_NOT
		&=		AND_ASSIGN
		&^= 	AND_NOT_ASSIGN
		&&		LAND 
	*/
	switch l.peekChar() {
	case '^':
		l.readChar() // &^
		if l.peekChar() == '=' {
			l.readChar() // &^=
			return token.AND_NOT_ASSIGN
		}
		return token.AND_NOT
	case '=':
		l.readChar()
		return token.AND_ASSIGN
	case '&':
		l.readChar()
		return token.LAND
	default:
		return token.AND
	}
}

func (l *Lexer) ScanPipeBar() token.Token {
	/*
		|		OR
		|=		OR_ASSIGN
		||		LOR
		|>		PIPE
	*/
	switch l.peekChar() {
	case '=':
		l.readChar()
		return token.OR_ASSIGN
	case '|':
		l.readChar()
		return token.LOR
	case '>':
		l.readChar()
		return l.ScanPipeline()
	default:
		return token.OR
	}
}

func (l *Lexer) ScanCarat() token.Token {
	/*
		^		XOR
		^=		XOR_ASSIGN 
	*/
	if l.peekChar() == '=' {
		l.readChar()
		return token.XOR_ASSIGN
	}
	return token.XOR
}

func (l *Lexer) ScanLeftChevron() token.Token {
	/*
		<		LT
		<=		LTE
		<<		SHL
		<<=		SHL_ASSIGN
	*/
	switch l.peakChar() {
	case '<':
		l.readChar()
		if l.peakChar() == '=' {
			l.readChar()
			return token.SHL_ASSIGN
		}
		return token.SHL
	case '=':
		l.readChar()
		return token.SHL
	default:
		return token.LT
	}
}

func (l *Lexer) ScanRightChevron() token.Token {
	/*
		>		GT
		>=		GTE
		>>		SHR
		>>=		SHR_ASSIGN
	*/
	switch l.peakChar() {
	case '>':
		l.readChar()
		if l.peakChar() == '=' {
			l.readChar()
			return token.SHR_ASSIGN
		}
		return token.SHL
	case '=':
		l.readChar()
		return token.SHR
	default:
		return token.GT
	}
}

func (l *Lexer) ScanEqual() token.Token {
	/*
		=		ASSIGN
		==		EQL
	*/
	if l.peekChar() == "=" {
		l.readChar()
		return token.EQL
	}
	return token.ASSIGN
}

func (l *Lexer) ScanBang() token.Token {
	/*
		!		NOT
		!=		NEQL
	*/
	if l.peekChar() == '=' {
		l.readChar() 
		return token.NEQL
	}
	return token.NOT
}

func (l *Lexer) ScanDot() token.Token {
	/*
		.		DOT
		..		DBLDOT
		...		ELIPSIS
	*/
	if l.peekChar() == '.' {
		l.readChar()
		if l.peekChar() == '.'{
			l.readChar()
			return token.ELIPSIS
		}
		return token.DBLDOT
	}
	return token.DOT
}

func (l *Lexer) ScanColon()	token.Token {
	/*
		:		COLON
		:=		DEFINE
	*/
	if l.peekChar() == '=' {
		l.readChar() 
		return token.DEFINE
	}
	return token.COLON
}

func (l *Lexer) ScanComment() token.Token {

}

func (l *Lexer) ScanPipeline() token.Token {

}











func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.EQ, Literal: literal}
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.NOT_EQ, Literal: literal}
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) skipWhitespace() token.Token {
	var trivia []byte
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		append(trvia, l.ch)
		l.readChar()
	}
	return token.Trivia
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}

	return l.input[l.readPosition]
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
