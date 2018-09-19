package monster

import (
	"math/rand"
	"time"

	"github.com/jrecuero/go-cli/tools"
)

// MTime represents ...
type MTime int

// MHandler represents ...
type MHandler struct {
	running   bool
	freeze    bool
	delay     int
	workerCbs []func(IActor)
	mtime     MTime
	actors    []IActor
	schedule  []IActor
	origin    int64
}

// tick is ...
func (mhdlr *MHandler) tick() {
	if mhdlr.mtime == 0 {
		mhdlr.origin = time.Now().UTC().UnixNano() / int64(time.Millisecond)
	}
	mhdlr.mtime++
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
func (mhdlr *MHandler) GetSchedule() []IActor {
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
func (mhdlr *MHandler) worker(actor IActor) {
	for mhdlr.running {
		time.Sleep(time.Duration(actor.GetSpeed()) * time.Millisecond)
		timeout := time.Now().UTC().UnixNano()/int64(time.Millisecond) - mhdlr.GetOrigin()
		tools.ToDisplay("[%#v] worker:  %#v\n", timeout, actor)
		mhdlr.schedule = append(mhdlr.schedule, actor)
		mhdlr.callWorkerCbs(actor)
	}
}

// Setup is ...
func (mhdlr *MHandler) Setup() {
}

// Start is ...
func (mhdlr *MHandler) Start() {
	mhdlr.running = true
	for _, actor := range mhdlr.actors {
		go mhdlr.worker(actor)
	}
	tools.ToDisplay("MHandler started ...\n")
	mhdlr.AddWorkerCb(func(actor IActor) {})
	for mhdlr.running {
		time.Sleep(time.Duration(mhdlr.delay) * time.Millisecond)
		mhdlr.tick()
	}
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

// NewMHandler is ...
func NewMHandler() *MHandler {
	rand.Seed(time.Now().UTC().UnixNano())
	return &MHandler{
		running: false,
		delay:   10,
	}
}
