package engine

// Engine is ...
type Engine struct {
	Time         ETime
	Events       []IEvent
	cachedEvents map[string][]IEvent
	flagedEvents map[string][]IEvent
	listedEvents map[string][]IEvent
	Running      bool
	pipe         chan bool
	waiting      bool
	caches       map[string]*Cache
	flags        map[string]*Flag
	listas       map[string]*Lista
}

// GetCacheForApp is ...
func (eng *Engine) GetCacheForApp(app string) *Cache {
	return eng.caches[app]
}

// GetFlagForApp is ...
func (eng *Engine) GetFlagForApp(app string) *Flag {
	return eng.flags[app]
}

// GetListaForApp is ...
func (eng *Engine) GetListaForApp(app string) *Lista {
	return eng.listas[app]
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
				eng.Time = next.GetAtTime()
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

// UpdateCachedEvents is ...
func (eng *Engine) UpdateCachedEvents() bool {
	return false
}

// UpdateFlagedEvents is ...
func (eng *Engine) UpdateFlagedEvents() bool {
	return false
}

// UpdateListedEvents is ...
func (eng *Engine) UpdateListedEvents() bool {
	return false
}

// NewEngine is ...
func NewEngine() *Engine {
	return &Engine{
		pipe:         make(chan bool),
		caches:       make(map[string]*Cache),
		flags:        make(map[string]*Flag),
		listas:       make(map[string]*Lista),
		cachedEvents: make(map[string][]IEvent),
		flagedEvents: make(map[string][]IEvent),
		listedEvents: make(map[string][]IEvent),
	}
}
