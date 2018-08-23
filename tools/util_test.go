package tools_test

import (
	"testing"

	"github.com/jrecuero/go-cli/tools"
)

// TestUtil_Sorting is ...
func TestUtil_Sorting(t *testing.T) {
	compa := func(a interface{}, b interface{}) int {
		_a := a.(int)
		_b := b.(int)
		if _a > _b {
			return 1
		} else if _a < _b {
			return -1
		}
		return 0
	}
	var tests = []struct {
		input  []interface{}
		elem   interface{}
		result int
	}{
		{
			input:  []interface{}{1, 2, 3, 4, 5},
			elem:   4,
			result: 4,
		},
		{
			input:  []interface{}{1, 2, 3, 4, 5},
			elem:   0,
			result: 0,
		},
		{
			input:  []interface{}{1, 2, 3, 4, 5},
			elem:   6,
			result: 5,
		},
		{
			input:  []interface{}{1, 1, 1, 1, 1},
			elem:   1,
			result: 5,
		},
		{
			input:  []interface{}{1, 4, 4, 4, 5},
			elem:   4,
			result: 4,
		},
		{
			input:  []interface{}{},
			elem:   4,
			result: 0,
		},
		{
			input:  []interface{}{1},
			elem:   4,
			result: 1,
		},
		{
			input:  []interface{}{4},
			elem:   4,
			result: 1,
		},
		{
			input:  []interface{}{5},
			elem:   4,
			result: 0,
		},
	}
	for i, test := range tests {
		if got := tools.Sorting(test.input, test.elem, compa); got != test.result {
			t.Errorf("[%d] Sorting: index mismatch: %#v %#v exp: %d got: %d\n", i, test.input, test.elem, test.result, got)
		}
	}
}
