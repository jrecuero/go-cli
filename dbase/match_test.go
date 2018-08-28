package dbase_test

import (
	"reflect"
	"testing"

	"github.com/jrecuero/go-cli/dbase"
)

// TestMatch_NewMatch is ...
func TestMatch_NewMatch(t *testing.T) {
	got := dbase.NewMatch("name", "EQUAL", "home", "AND")
	if got == nil {
		t.Errorf("NewMatch: return mismatch: <nil>")
	}
	exp := &dbase.Match{
		Key:     "name",
		Operand: "EQUAL",
		Value:   "home",
		Link:    "AND",
	}
	if !reflect.DeepEqual(exp, got) {
		t.Errorf("NewMatch output mismatch:exp:\t%#v\ngot:\t%#v\n", exp, got)
	}
}

// TestMatch_IsContinue is ...
func TestMatch_IsContinue(t *testing.T) {
	if match := dbase.NewMatch("name", "EQUAL", "home", "AND"); match != nil {
		if got := match.IsContinue(true); !got {
			t.Errorf("IsContinue: return mismatch: AND: exp: true got: %#v\n", got)
		}
		if got := match.IsContinue(false); got {
			t.Errorf("IsContinue: return mismatch: AND: exp: true got: %#v\n", got)
		}
	}
	if match := dbase.NewMatch("name", "EQUAL", "home", "OR"); match != nil {
		if got := match.IsContinue(true); !got {
			t.Errorf("IsContinue: return mismatch: OR: exp: true got: %#v\n", got)
		}
		if got := match.IsContinue(false); !got {
			t.Errorf("IsContinue: return mismatch: OR: exp: true got: %#v\n", got)
		}
	}
}

// TestMatch_IsMatch is ...
func TestMatch_IsMatch(t *testing.T) {
	if match := dbase.NewMatch("name", "EQUAL", "home", "AND"); match != nil {
		if got, conti := match.IsMatch("home"); !got || !conti {
			t.Errorf("IsMatch: return mismatch: exp true,true got: %#v,%#v\n", got, conti)
		}
		if got, conti := match.IsMatch("nothome"); got || conti {
			t.Errorf("IsMatch: return mismatch: exp false,false got: %#v,%#v\n", got, conti)
		}
	}
	if match := dbase.NewMatch("name", "EQUAL", "home", "OR"); match != nil {
		if got, conti := match.IsMatch("home"); !got || !conti {
			t.Errorf("IsMatch: return mismatch: exp true,true got: %#v,%#v\n", got, conti)
		}
		if got, conti := match.IsMatch("nothome"); got || !conti {
			t.Errorf("IsMatch: return mismatch: exp false,true got: %#v,%#v\n", got, conti)
		}
	}
	if match := dbase.NewMatch("name", "NOT EQUAL", "home", "AND"); match != nil {
		if got, conti := match.IsMatch("home"); got || conti {
			t.Errorf("IsMatch: return mismatch: exp false,false got: %#v,%#v\n", got, conti)
		}
		if got, conti := match.IsMatch("nothome"); !got || !conti {
			t.Errorf("IsMatch: return mismatch: exp true,true got: %#v,%#v\n", got, conti)
		}
	}
	if match := dbase.NewMatch("name", "NOT EQUAL", "home", "OR"); match != nil {
		if got, conti := match.IsMatch("home"); got || !conti {
			t.Errorf("IsMatch: return mismatch: exp false,true got: %#v,%#v\n", got, conti)
		}
		if got, conti := match.IsMatch("nothome"); !got || !conti {
			t.Errorf("IsMatch: return mismatch: exp true,true got: %#v,%#v\n", got, conti)
		}
	}
}
