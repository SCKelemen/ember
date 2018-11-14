package kw

type Keyword int

const (
	INVALID Keyword = iota
	ACTOR
	CLASS
	INTERFACE
	ABSTRACT
)

func (kw Keyword) String() string {
	switch kw {
	case ACTOR:
		return "actor"
	case CLASS:
		return "class"
	case INTERFACE:
		return "interface"
	case ABSTRACT:
		return "abstract"
	default:
		return "INVALID"
	}
}
