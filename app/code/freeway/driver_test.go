package freeway_test

import (
	"testing"

	"github.com/jrecuero/go-cli/app/code/freeway"
)

func TestDriver_Driver(t *testing.T) {
	driver := freeway.NewDriver("me")
	if driver == nil {
		t.Errorf("NewDriver: return code: nil\n")
	}
	if driver.GetName() != "me" {
		t.Errorf("NewDriver: name mismatch: exp: %#v got: %#v\n", "me", driver.GetName())
	}
}
