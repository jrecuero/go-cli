package engine

// Engine is ...
type Engine struct {
	Time    ETime
	Events  []IEvent
	Running bool
	pipe    chan bool
	waiting bool
}

// Start is ...
func (eng *Engine) Start() error {
	eng.Running = true
	return nil
}

// Stop is ...
func (eng *Engine) Stop() error {
	eng.Running = false
	return nil
}

// AddEventAtTime is ...
func (eng *Engine) AddEventAtTime(ev IEvent) error {
	//tools.ToDisplay("Add event: %s\n", ev)
	eng.Events = append(eng.Events, ev)
	if eng.waiting {
		eng.pipe <- true
	}
	return nil
}

// AddEvent is ...
func (eng *Engine) AddEvent(ev IEvent) error {
	ev.SetAtTime(ev.GetAtTime() + eng.Time)
	return eng.AddEventAtTime(ev)
}

// AddEventFirst is ...
func (eng *Engine) AddEventFirst(ev IEvent) error {
	ev.SetAtTime(eng.Time)
	eng.Events = append([]IEvent{ev}, eng.Events...)
	if eng.waiting {
		eng.pipe <- true
	}
	return nil
}

// PeekNext is ...
func (eng *Engine) PeekNext() IEvent {
	return eng.Events[0]
}

// NextEvent is ...
func (eng *Engine) NextEvent() IEvent {
	ev := eng.Events[0]
	evLen := len(eng.Events)
	eng.Events = eng.Events[1:evLen]
	return ev
}

// Next is ...
func (eng *Engine) Next() error {
	if ev := eng.NextEvent(); ev != nil {
		eng.Time = ev.GetAtTime()
		return eng.ExecEvent(ev)
	}
	return nil
}

// ExecEvent is ...
func (eng *Engine) ExecEvent(ev IEvent) error {
	return ev.Exec()
}

// Run is ...
func (eng *Engine) Run() error {
	if eng.Running {
		return eng.Next()
	}
	return nil
}

// tick is ...
func (eng *Engine) tick() {
}

// loop is ...
func (eng *Engine) loop(done chan bool) {
	eng.Start()
	for eng.Running {
		//tools.ToDisplay("looping here... %#v\n", len(eng.Events))
		if len(eng.Events) == 0 {
			//tools.ToDisplay("waiting here...\n")
			eng.waiting = true
			<-eng.pipe
			eng.waiting = false
			//tools.ToDisplay("wake up here...\n")
		} else {
			next := eng.PeekNext()
			if next.GetAtTime() == eng.Time {
				//tools.ToDisplay("running here... %d\n", len(eng.Events))
				if err := eng.Run(); err != nil {
					break
				}
			} else {
				eng.tick()
			}
		}

	}
	//tools.ToDisplay("exit here...\n")
	eng.endloop()
	done <- true
}

// endloop is ...
func (eng *Engine) endloop() {
	eng.waiting = false
	eng.Stop()
}

// Loop is ...
func (eng *Engine) Loop() {
	done := make(chan bool, 1)
	go eng.loop(done)
	<-done
}

// EndLoop is ...
func (eng *Engine) EndLoop() {
	eng.pipe <- true
	eng.Stop()
}

// NewEngine is ...
func NewEngine() *Engine {
	return &Engine{
		pipe: make(chan bool),
	}
}
