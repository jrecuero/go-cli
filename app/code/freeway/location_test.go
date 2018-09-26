package freeway_test

import (
	"testing"

	"github.com/jrecuero/go-cli/app/code/freeway"
)

func TestLocation_Location(t *testing.T) {
	loc := freeway.NewLocation(nil)
	if loc == nil {
		t.Errorf("NewLocation: return code: nil\n")
	}
}
