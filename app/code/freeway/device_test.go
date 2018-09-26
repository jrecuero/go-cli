package freeway_test

import (
	"testing"

	"github.com/jrecuero/go-cli/app/code/freeway"
)

func TestDevice_Device(t *testing.T) {
	device := freeway.NewFullDevice("dev", "dev-class", "dev-sub", 100)
	if device == nil {
		t.Errorf("NewDevice: null pointer\n")
	}
	if device.GetName() != "dev" {
		t.Errorf("NewDevice: name mismatch: exp: %#v got: %#v\n", "dev", device.GetName())
	}
	if device.GetClass() != "dev-class" {
		t.Errorf("NewDevice: class mismatch: exp: %#v got: %#v\n", "dev-class", device.GetClass())
	}
	if device.GetSubClass() != "dev-sub" {
		t.Errorf("NewDevice: subclass mismatch: exp: %#v got: %#v\n", "dev-sub", device.GetSubClass())
	}
	if device.GetPower() != 100 {
		t.Errorf("NewDevice: power mismatch: exp %d got : %d\n", 100, device.GetPower())
	}
}
