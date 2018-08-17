package engine_test

import (
	"testing"

	"github.com/jrecuero/go-cli/engine"
)

// TestEvent_NewEvent is ...
func TestEvent_NewEvent(t *testing.T) {
	ev := engine.NewEvent("test-event", 1)
	if ev.Name != "test-event" {
		t.Errorf("NewEvent: Name mistmatch: exp: \"test-event\" got: %#v\n", ev.Name)
	}
	if ev.AtTime != 1 {
		t.Errorf("NewEvent: AtTime mistmatch: exp: 1 got: %d\n", ev.AtTime)
	}
}

// TestEvent_Callback is ...
func TestEvent_Callback(t *testing.T) {
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
	if err := ev.Exec(); err != nil {
		t.Errorf("Callback: Exec return mismatch: exp: nil got %#v\n", err)
	}
	if !called {
		t.Errorf("Callback: Exec called is false\n")
	}
}
