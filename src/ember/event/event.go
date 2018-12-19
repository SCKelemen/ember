package event

type EventType int

const (
	INVALID EventType = iota
	HELLO
)

var eventStringMap = [...]string{
	"INVALID",
	"HELLO",
}

func (et EventType) String() string {
	if int(et) < len(eventStringMap) {
		return eventStringMap[et]
	}
	return eventStringMap[0]
}

type Event struct {
	Type   EventType
	Source interface{}
	Value  interface{}
}
