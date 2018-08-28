package dbase_test

import (
	"testing"

	"github.com/jrecuero/go-cli/dbase"
)

// TestRow_NewRow is ...
func TestRow_NewRow(t *testing.T) {
	var row *dbase.Row
	row = dbase.NewRow("a")
	if len(row.Data) != 1 {
		t.Errorf("NewRow: length mismatch: got: %d exp: 1\n", len(row.Data))
	}
	if row.Data[0] != "a" {
		t.Errorf("NewRow: data mismatch: got: %#v exp: \"a\"\n", row.Data[0])
	}

	row = dbase.NewRow(1, 2, 3, 4, 5)
	if len(row.Data) != 5 {
		t.Errorf("NewRow: length mismatch: got: %d exp: 5\n", len(row.Data))
	}
	for i := 0; i < 5; i++ {
		if row.Data[i] != i+1 {
			t.Errorf("NewRow: data mismatch: got: %#v exp: %d\n", row.Data[i], i+1)
		}
	}
}
