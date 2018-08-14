package novel

// ETime is ...
type ETime uint

const (
	// ZeroTime is ...
	ZeroTime ETime = 0
)

// Event is ...
type Event struct {
	Name   string
	AtTime ETime
}

// Exec is ...
func (ev *Event) Exec() error {
	return nil
}

// NewEvent is ...
func NewEvent(name string, attime ETime) *Event {
	return &Event{
		Name:   name,
		AtTime: attime,
	}
}
