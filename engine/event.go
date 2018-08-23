package engine

// ETime is ...
type ETime uint

// EventCallback represents ...
type EventCallback func(params ...interface{}) error

// IEvContent represents ...
type IEvContent interface {
}

// IEvent represents ...
type IEvent interface {
	GetName() string
	GetAtTime() ETime
	SetAtTime(ETime)
	SetCallback(EventCallback, ...interface{}) error
	GetContent() IEvContent
	SetContent(IEvContent)
	Exec() error
}

const (
	// ZeroTime is ...
	ZeroTime ETime = 0
)

// Event is ...
type Event struct {
	name     string
	atTime   ETime
	callback EventCallback
	cbParams []interface{}
	content  IEvContent
}

// GetName is ...
func (ev *Event) GetName() string {
	return ev.name
}

// GetAtTime is ...
func (ev *Event) GetAtTime() ETime {
	return ev.atTime
}

// SetAtTime is ...
func (ev *Event) SetAtTime(attime ETime) {
	ev.atTime = attime
}

// GetContent is ...
func (ev *Event) GetContent() IEvContent {
	return ev.content
}

// SetContent is ...
func (ev *Event) SetContent(content IEvContent) {
	ev.content = content
}

// String is ...
func (ev *Event) String() string {
	return ev.name
}

// SetCallback is ...
func (ev *Event) SetCallback(cb EventCallback, params ...interface{}) error {
	ev.callback = cb
	ev.cbParams = params
	return nil
}

// Exec is ...
func (ev *Event) Exec() error {
	if ev.callback != nil {
		return ev.callback(ev.cbParams...)
	}
	return nil
}

// NewEvent is ...
func NewEvent(name string, attime ETime) *Event {
	return &Event{
		name:   name,
		atTime: attime,
	}
}
