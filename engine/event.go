package engine

// ETime is ...
type ETime uint

// EventCallback is ...
type EventCallback func(params ...interface{}) error

const (
	// ZeroTime is ...
	ZeroTime ETime = 0
)

// Event is ...
type Event struct {
	Name     string
	AtTime   ETime
	callback EventCallback
	params   []interface{}
}

// String is ...
func (ev *Event) String() string {
	return ev.Name
}

// SetCallback is ...
func (ev *Event) SetCallback(cb EventCallback, params ...interface{}) error {
	ev.callback = cb
	ev.params = params
	return nil
}

// Exec is ...
func (ev *Event) Exec() error {
	if ev.callback != nil {
		return ev.callback(ev.params...)
	}
	return nil
}

// NewEvent is ...
func NewEvent(name string, attime ETime) *Event {
	return &Event{
		Name:   name,
		AtTime: attime,
	}
}
