package novel

import (
	"fmt"
	"strings"

	"github.com/jrecuero/go-cli/parser"
	lexnovel "github.com/jrecuero/go-cli/parser/lex/novel"
	"github.com/jrecuero/go-cli/tools"
)

// Object is ...
type Object struct {
	Name string
}

// GoTo is ...
type GoTo struct {
	Name  string
	GType string
	Dest  *Room
}

// Room is ...
type Room struct {
	Name    string
	Objects []*Object
	Actors  []*Actor
	GoTos   []*GoTo
}

// Actor represents any actor in the app.
type Actor struct {
	Name     string
	Life     int
	Strength int
	IsPlayer bool
	IsNPC    bool
	IsEnemy  bool
}

// NewActor creates a new actor instance.
func NewActor(name string, life int, strength int) *Actor {
	return &Actor{
		Name:     name,
		Life:     life,
		Strength: strength,
	}
}

// ActionNames represents any action in the app.
type ActionNames struct {
	Origins []string
	Actions []string
	Targets []string
	Flags   []string
}

// AddOrigin adds a new origin
func (an *ActionNames) AddOrigin(in string) *ActionNames {
	an.Origins = append(an.Origins, in)
	return an
}

// AddAction adds a new action.
func (an *ActionNames) AddAction(in string) *ActionNames {
	an.Actions = append(an.Actions, in)
	return an
}

// AddTarget adds a new target.
func (an *ActionNames) AddTarget(in string) *ActionNames {
	an.Targets = append(an.Targets, in)
	return an
}

// AddFlags adds a new flag.
func (an *ActionNames) AddFlags(in string) *ActionNames {
	an.Flags = append(an.Flags, in)
	return an
}

// ActionCallback is ...
type ActionCallback func(origins []*Actor, target []*Actor, flags []string) error

// ActionSequence is ...
type ActionSequence struct {
	Origins []*Actor
	Actions []ActionCallback
	Targets []*Actor
	Flags   []string
}

// RunAction is ...
func (as *ActionSequence) RunAction() error {
	//tools.ToDisplay("RunAction: %#v\n", as)
	for _, action := range as.Actions {
		if err := action(as.Origins, as.Targets, as.Flags); err != nil {
			return err
		}
	}
	return nil
}

// CompileStatus represents the compile status.
type CompileStatus struct {
	actual int
}

// Next moves to the next status if condition is true.
func (cs *CompileStatus) Next(condition bool) *CompileStatus {
	if condition {
		cs.actual++
	}
	return cs
}

// Value returns compile status value.
func (cs *CompileStatus) Value() int {
	return cs.actual
}

// NewCompileStatus creates a new compile status instance.
func NewCompileStatus(start int) *CompileStatus {
	return &CompileStatus{start}
}

// ActionHit is ...
func ActionHit(novel *Novel) ActionCallback {
	return func(origins []*Actor, targets []*Actor, flags []string) error {
		//tools.ToDisplay("Hit: origins: %#v targets: %#v\n", origins, targets)
		if len(origins) != 1 {
			return tools.ERROR(nil, false, "Too many origins (%d) for action 'Hit'\n", len(origins))
		}
		for _, target := range targets {
			origin := origins[0]
			damage := origin.Strength
			target.Life -= damage
			tools.ToDisplay("%#v hits with %d damage to %#v: %d life points\n", origin.Name, origin.Strength, target.Name, target.Life)
		}
		return nil
	}
}

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

// Engine is ...
type Engine struct {
	Time   ETime
	Events []*Event
}

// Start is ...
func (eng *Engine) Start() error {
	return nil
}

// Stop is ...
func (eng *Engine) Stop() error {
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

// NewEngine is ...
func NewEngine() *Engine {
	return &Engine{}
}

// Novel represents the main object for the app.
type Novel struct {
	Actors      []*Actor
	ActionCalls map[string]ActionCallback
}

func (novel *Novel) nameToActor(names []string) []*Actor {
	var actors []*Actor
	for _, name := range names {
		if actor := novel.SearchByName(name); actor != nil {
			actors = append(actors, actor)
		}
	}
	return actors
}

func (novel *Novel) nameToActionCall(names []string) []ActionCallback {
	var actionCalls []ActionCallback
	for _, name := range names {
		actionCalls = append(actionCalls, novel.ActionCalls[name])
	}
	return actionCalls
}

// Parse parses the novel action.
func (novel *Novel) Parse(line string) *lexnovel.Syntax {
	parser := parser.NewParser(strings.NewReader(line), lexnovel.NewParser())
	result, _ := parser.Parse()
	return result.(*lexnovel.Syntax)
}

// Compile translate the parsing action to a struct.
func (novel *Novel) Compile(line string) *ActionNames {
	defer func() {
		if r := recover(); r != nil {
			tools.ToDisplay("Error: %#v\n", r)
		}
	}()
	const (
		parsingOrigin int = iota
		parsingAction
		parsingTarget
		parsingFlags
		parsingEnd
	)
	parsed := novel.Parse(line)
	//tools.ToDisplay("%#v\n", parsed)
	ae := &ActionNames{}
	bracketed := false
	status := NewCompileStatus(parsingOrigin)
	for i, token := range parsed.Tokens {
		str := parsed.Idents[i]
		switch token {
		case parser.IDENT:
			switch status.Value() {
			case parsingOrigin:
				ae.AddOrigin(str)
				status.Next(!bracketed)
				break
			case parsingAction:
				ae.AddAction(str)
				status.Next(!bracketed)
				break
			case parsingTarget:
				ae.AddTarget(str)
				status.Next(!bracketed)
			case parsingFlags:
				ae.AddFlags(str)
				status.Next(!bracketed)
			default:
				panic("bad action: unknown status")
				break
			}

		case lexnovel.OPENBRACKET:
			if bracketed {
				panic("bad action: too many brackets")
			}
			bracketed = true
			break
		case lexnovel.CLOSEBRACKET:
			if !bracketed {
				panic("bad action: no brackets")
			}
			bracketed = false
			status.Next(true)
			break
		default:
			panic("bad action: unknown token")
			break
		}
	}
	return ae
}

// RunAction is ...k
func (novel *Novel) RunAction(line string) error {
	//tools.ToDisplay("RunAction: %#v\n", line)
	if ae := novel.Compile(line); ae != nil {
		as := &ActionSequence{}
		as.Origins = novel.nameToActor(ae.Origins)
		as.Targets = novel.nameToActor(ae.Targets)
		as.Actions = novel.nameToActionCall(ae.Actions)
		as.Flags = as.Flags
		return as.RunAction()
	}
	return tools.ERROR(nil, false, "Invalid line: %#v\n", line)
}

// Execute is ...
func (novel *Novel) Execute(line string) error {
	if err := novel.Execute(line); err != nil {
		return err
	}
	novel.Update()
	return nil
}

// Update is ...
func (novel *Novel) Update() (bool, error) {
	var actors []*Actor
	isPlayerAlive := false
	for _, actor := range novel.Actors {
		if actor.Life > 0 {
			actors = append(actors, actor)
			if actor.IsPlayer {
				isPlayerAlive = true
			}
		}
	}
	novel.Actors = actors
	return isPlayerAlive, nil
}

// createActors is ...
func (novel *Novel) createActors() {
	novel.Actors = append(novel.Actors,
		&Actor{
			Name:     "Player",
			Life:     100,
			Strength: 10,
			IsPlayer: true,
		})
	for i := 0; i < 3; i++ {
		novel.Actors = append(novel.Actors, &Actor{
			Name:     fmt.Sprintf("Enemy-%d", i),
			Life:     15,
			Strength: 5,
			IsEnemy:  true,
		})
	}
}

// Run is ...
func (novel *Novel) Run() {
	novel.createActors()
}

// SearchByName is ...
func (novel *Novel) SearchByName(name string) *Actor {
	//tools.ToDisplay("actors: %#v\n", novel.Actors)
	for _, actor := range novel.Actors {
		//tools.ToDisplay("Search for %#v in %#v\n", name, actor.Name)
		if actor.Name == name {
			return actor
		}
	}
	return nil
}

// NewNovel is ...
func NewNovel() *Novel {
	novel := &Novel{
		ActionCalls: make(map[string]ActionCallback),
	}
	novel.ActionCalls["hit"] = ActionHit(novel)
	return novel
}

// TheNovel is ...
var TheNovel = NewNovel()
