package dbase_test

import (
	"reflect"
	"testing"

	"github.com/jrecuero/go-cli/dbase"
)

// TestColumn_NewColumn is ...
func TestColumn_NewColumn(t *testing.T) {
	var tests = []struct {
		input *dbase.Column
		exp   *dbase.Column
	}{
		{
			input: &dbase.Column{
				Name:  "test",
				CType: "string",
			},
			exp: &dbase.Column{
				Name:  "test",
				CType: "string",
				Label: "",
				Align: "",
				Width: 0,
				Key:   false,
			},
		},
	}
	for i, tt := range tests {
		out := dbase.NewColumn(tt.input.Name, tt.input.CType)
		if !reflect.DeepEqual(tt.exp, out) {
			t.Errorf("%d. NewColumn output mismatch:\ninput:\t%#v\nexp:\t%#v\ngot:\t%#v\n", i, tt.input, tt.exp, out)
		}
	}
}

// TestColumn_SetLayout is ...
func TestColumn_SetLayout(t *testing.T) {
	var tests = []struct {
		input *dbase.Column
		exp   *dbase.Column
	}{
		{
			input: &dbase.Column{
				Name:  "test",
				CType: "string",
				Label: "TEST",
				Align: "RIGHT",
				Width: 10,
			},
			exp: &dbase.Column{
				Name:  "test",
				CType: "string",
				Label: "TEST",
				Align: "RIGHT",
				Width: 10,
				Key:   false,
			},
		},
	}
	for i, tt := range tests {
		out := dbase.NewColumn(tt.input.Name, tt.input.CType)
		out = out.SetLayout(tt.input.Label, tt.input.Width, tt.input.Align)
		if !reflect.DeepEqual(tt.exp, out) {
			t.Errorf("%d. SetLayout  output mismatch:\ninput:\t%#v\nexp:\t%#v\ngot:\t%#v\n", i, tt.input, tt.exp, out)
		}
	}
}
