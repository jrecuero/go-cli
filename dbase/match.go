package dbase

import "fmt"

// Match represents any match entry in a filter. Any Match is composed for a
// kye, which will be usd to identify the data to search, an operand, a value
// to match and a link with further Matches instances.
type Match struct {
	Key     string
	Operand string
	Value   interface{}
	Link    string
}

// IsContinue returns if the match should allows to continue matching further
// Match instances.
func (m *Match) IsContinue(match bool) bool {
	var conti bool
	if m.Link == "AND" && match {
		conti = true
	} else if m.Link == "AND" && !match {
		conti = false
	} else if m.Link == "OR" {
		conti = true
	} else {
		panic(fmt.Sprintf("unknown link %#v", m.Link))
	}
	return conti
}

// IsMatch returns if there is a match and if the Match allows to continue
// matching further.
func (m *Match) IsMatch(value interface{}) (bool, bool) {
	var match bool
	switch m.Operand {
	case "EQUAL":
		match = value == m.Value
		break
	case "NOT EQUAL":
		match = value != m.Value
		break
	default:
		panic(fmt.Sprintf("unknown operand %#v", m.Operand))
	}
	return match, m.IsContinue(match)
}

// NewMatch returns a new Match instance.
func NewMatch(key string, operand string, value interface{}, link string) *Match {
	return &Match{
		Key:     key,
		Operand: operand,
		Value:   value,
		Link:    link,
	}
}
