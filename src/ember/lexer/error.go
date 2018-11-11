package lexer

import (
	"fmt"
)

type ScanError struct {
	error
	ErrorCode LexError
}

type LexError int

func (e *LexError) String() string {
	switch e {
	case Unknown:
		return "An unknown error has occured"
	default:
		return "An unknown error has occcured"
	}
}

const (
	Unknown LexError = iota
)

func (e *ScanError) String() string {
	return e.ErrorCode.String()
}

func (e *ScanError) Print() {
	fmt.Println(e)
}
