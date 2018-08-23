package engine_test

import (
	"testing"

	"github.com/jrecuero/go-cli/engine"
)

// TestEngine_NewEngine is ...
func TestEngine_NewEngine(t *testing.T) {
	eng := engine.NewEngine()
	if eng.Time != 0 {
		t.Errorf("NewEngine: Time mistmatch: exp: 0 got: %d\n", eng.Time)
	}
	if len(eng.Events) != 0 {
		t.Errorf("NewEngine: Events length mismatch: exp: 0 got: %d\n", len(eng.Events))
	}
	if eng.Running {
		t.Errorf("NewEngine: Running flag mistmatch: exp: false got: %v\n", eng.Running)
	}
}

// TestEngine_StartStop is ...
func TestEngine_StartStop(t *testing.T) {
	eng := engine.NewEngine()
	eng.Start()
	if !eng.Running {
		t.Errorf("Start: Running flag mistmatch: exp: true got: %v\n", eng.Running)
	}
	eng.Stop()
	if eng.Running {
		t.Errorf("Stop: Running flag mistmatch: exp: true got: %v\n", eng.Running)
	}
}

// TestEngine_Events is ...
func TestEngine_Events(t *testing.T) {
	eng := engine.NewEngine()
	if err := eng.AddEvent(engine.NewEvent("1", 1)); err != nil {
		t.Errorf("AddEvent: return code error: %#v\n", err)
	}
	if err := eng.AddEvent(engine.NewEvent("2", 2)); err != nil {
		t.Errorf("AddEvent: return code error: %#v\n", err)
	}
	if len(eng.Events) != 2 {
		t.Errorf("AddEvent: Events length mismatch: exp: %d got: %d\n", 2, len(eng.Events))
	}
	if ev := eng.NextEvent(); ev == nil {
		t.Errorf("NextEvent: event is <nil>\n")
	} else {
		if len(eng.Events) != 1 {
			t.Errorf("NextEvent: Events length mismatch: exp: %d got: %d\n", 1, len(eng.Events))
		}
		if ev.GetName() != "1" {
			t.Errorf("NextEvent: Events name mismatch: exp: %#v got: %#v\n", "1", ev.GetName())
		}
	}
	eng.AddEventFirst(engine.NewEvent("first", 0))
	if len(eng.Events) != 2 {
		t.Errorf("AddEventFirst: Events length mismatch: exp: 2 got: %d\n", len(eng.Events))
	}
	if eng.Events[0].GetName() != "first" {
		t.Errorf("AddEventFirst: First event name mismatch: exp: %#v got: %#v\n", "first", eng.Events[0].GetName())
	}
}

// TestEngine_Run is ...
func TestEngine_Run(t *testing.T) {
	called := false
	name := "test callback"
	ev := engine.NewEvent("test-event", 1)
	ev.SetCallback(func(params ...interface{}) error {
		if len(params) != 1 {
			t.Errorf("Callback: length params mismatch: exp: 1 got %#v\n", len(params))
		}
		if params[0] != "test callback" {
			t.Errorf("Callback: params mismatch: exp: \"test callback\" got %#v\n", params[0])
		}
		called = true
		return nil
	}, name)
	eng := engine.NewEngine()
	eng.AddEvent(ev)
	eng.AddEvent(engine.NewEvent("second", 1))
	if err := eng.Run(); err != nil {
		t.Errorf("Run: return code mistmatch: %#v\n", err)
	} else {
		if called {
			t.Errorf("Run:Stopped Engine:: Exec called is not false\n")
		}
	}
	eng.Start()
	if err := eng.Run(); err != nil {
		t.Errorf("Run: return code mistmatch: %#v\n", err)
	} else {
		if !called {
			t.Errorf("Run:Started Engine:: Exec called is false\n")
		}
	}
}
