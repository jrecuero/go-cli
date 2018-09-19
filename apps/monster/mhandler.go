package monster

import (
	"math/rand"
	"sync"
	"time"

	"github.com/jrecuero/go-cli/tools"
)

// MTime represents ...
type MTime int

// Token represents ...
type Token struct {
	Actor IActor
	Timed int
}

// MHandler represents ...
type MHandler struct {
	running   bool
	freeze    bool
	delay     int
	workerCbs []func(IActor)
	mtime     MTime
	timed     int
	actors    []IActor
	schedule  []*Token
	origin    int64
}

// tick is ...
func (mhdlr *MHandler) tick() {
	if mhdlr.mtime == 0 {
		mhdlr.origin = time.Now().UTC().UnixNano() / int64(time.Millisecond)
	}
	mhdlr.mtime++
}

// resetTick is ...
func (mhdlr *MHandler) resetTick() {
	mhdlr.mtime = 0
	mhdlr.timed = 0
}

// GetOrigin is ...
func (mhdlr *MHandler) GetOrigin() int64 {
	return mhdlr.origin
}

// SetDelay is ...
func (mhdlr *MHandler) SetDelay(delay int) *MHandler {
	mhdlr.delay = delay
	return mhdlr
}

// AddActor is ...
func (mhdlr *MHandler) AddActor(actor IActor) *MHandler {
	mhdlr.actors = append(mhdlr.actors, actor)
	return mhdlr
}

// GetActors is ...
func (mhdlr *MHandler) GetActors() []IActor {
	return mhdlr.actors
}

// GetSchedule is ...
func (mhdlr *MHandler) GetSchedule() []*Token {
	return mhdlr.schedule
}

// AddWorkerCb is ...
func (mhdlr *MHandler) AddWorkerCb(cb func(IActor)) *MHandler {
	mhdlr.workerCbs = append(mhdlr.workerCbs, cb)
	return mhdlr
}

// callWorkerCbs is ...
func (mhdlr *MHandler) callWorkerCbs(actor IActor) {
	for _, cb := range mhdlr.workerCbs {
		if cb != nil {
			cb(actor)
		}
	}
}

// worker is ...
func (mhdlr *MHandler) worker(wg *sync.WaitGroup, actor IActor) {
	defer wg.Done()
	tools.ToDisplay("worker %#v starts\n", actor)
	for mhdlr.running {
		time.Sleep(time.Duration(actor.GetNext()) * time.Millisecond)
		timeout := time.Now().UTC().UnixNano()/int64(time.Millisecond) - mhdlr.GetOrigin()
		tools.ToDisplay("[%#v] worker:  %#v\n", timeout, actor)
		mhdlr.schedule = append(mhdlr.schedule, &Token{actor, actor.GetNext()})
		actor.SetNext(actor.GetSpeed())
		mhdlr.callWorkerCbs(actor)
	}
	tools.ToDisplay("worker %#v ends\n", actor.GetName())
}

// ticker is ...
func (mhdlr *MHandler) ticker(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Millisecond)
	for mhdlr.running {
		mhdlr.tick()
		time.Sleep(time.Duration(mhdlr.delay) * time.Millisecond)
	}
	//tools.ToDisplay("ticker ends at %d\n", time.Now().UTC().UnixNano()/int64(time.Millisecond)-mhdlr.GetOrigin())
}

// Setup is ...
func (mhdlr *MHandler) Setup() {
}

// Start is ...
func (mhdlr *MHandler) Start() {
	var wg sync.WaitGroup
	mhdlr.schedule = nil
	mhdlr.running = true
	tools.ToDisplay("MHandler started ...\n")
	mhdlr.resetTick()
	wg.Add(1)
	go mhdlr.ticker(&wg)
	mhdlr.AddWorkerCb(func(actor IActor) {})
	for _, actor := range mhdlr.actors {
		wg.Add(1)
		go mhdlr.worker(&wg, actor)
	}
	wg.Wait()
	tools.ToDisplay("MHandler stopped\n")
}

// Stop is ...
func (mhdlr *MHandler) Stop() {
	mhdlr.running = false
}

// SetFreeze is ...
func (mhdlr *MHandler) SetFreeze(freeze bool) {
	mhdlr.freeze = freeze
}

// Next is ...
func (mhdlr *MHandler) Next() IActor {
	token := mhdlr.schedule[0]
	mhdlr.schedule = mhdlr.schedule[1:]
	timed := token.Timed
	for _, actor := range mhdlr.actors {
		//tools.ToDisplay("next:  %#v %d %d\n", actor, timed, mhdlr.timed)
		newNext := actor.GetProcessing() - (timed - mhdlr.timed)
		actor.SetProcessing(newNext)
	}
	mhdlr.timed = timed
	token.Actor.SetProcessing(token.Actor.GetSpeed())
	return token.Actor
}

// NewMHandler is ...
func NewMHandler() *MHandler {
	rand.Seed(time.Now().UTC().UnixNano())
	return &MHandler{
		running: false,
		delay:   10,
	}
}
