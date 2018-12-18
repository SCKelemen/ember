package diag

// DiagnosticType is a reprsentation of the
// Diagnostic level. Info, Warning, or Error
type DiagnosticType int

const (
	// Error Diagnostic Level
	Error DiagnosticType = iota
	// Warning Diagnostic Level
	Warning
	// Informational Diagnostic Level
	Informational
)

// DiagnosticSource represents where the error was emitted
// Compiler, Linter, Static Code Analyser
type DiagnosticSource int

// DiagnosticPosition provides information on
// the source which triggered the diagnostic
type DiagnosticPosition struct { // eventually replace this with the standard Position class
	File   string
	Line   int
	Column int
	Length int
}

// Diagnostic is an informational object
// which provides feedback on errors which
// may have occured during the parsing,
// lexing, code generating, analysis, or
// other process during compilation
type Diagnostic struct {
	Type     DiagnosticType
	Position DiagnosticPosition
	Message  string
}
