package novel

// Engine is ...
type Engine struct {
	Time    ETime
	Events  []*Event
	Running bool
}

// Start is ...
func (eng *Engine) Start() error {
	eng.Running = true
	return nil
}

// Stop is ...
func (eng *Engine) Stop() error {
	eng.Running = true
	return nil
}

// AddEvent is ...
func (eng *Engine) AddEvent(ev *Event) error {
	eng.Events = append(eng.Events, ev)
	return nil
}

// AddEventFirst is ...
func (eng *Engine) AddEventFirst(ev *Event) error {
	ev.AtTime = ZeroTime
	eng.Events = append([]*Event{ev}, eng.Events...)
	return nil
}

// NextEvent is ...
func (eng *Engine) NextEvent() *Event {
	ev := eng.Events[0]
	evLen := len(eng.Events)
	eng.Events = eng.Events[1:evLen]
	return ev
}

// Next is ...
func (eng *Engine) Next() error {
	if ev := eng.NextEvent(); ev != nil {
		eng.Time = ev.AtTime
		return eng.ExecEvent(ev)
	}
	return nil
}

// ExecEvent is ...
func (eng *Engine) ExecEvent(ev *Event) error {
	return ev.Exec()
}

// Run is ...
func (eng *Engine) Run() error {
	if eng.Running {
		return eng.Next()
	}
	return nil
}

// NewEngine is ...
func NewEngine() *Engine {
	return &Engine{}
}
