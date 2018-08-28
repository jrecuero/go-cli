package dbase

// FilterGetter represents any method that returns a value ot be used in a
// filter for the given key passed. The key is the value from the Match
// instance, and the return should feed Match.IsMatch call.
type FilterGetter func(key string) interface{}

// Filter represents a filter, which is composed by a list of possible matches.
type Filter struct {
	Matches []*Match
}

// Add is ...
func (f *Filter) Add(matches ...*Match) *Filter {
	for _, match := range matches {
		f.Matches = append(f.Matches, match)
	}
	return f
}

// IsMatch check if the filter matches for the given getter method. The getter
// will provide the value to feed for every March.IsMatch in the filter.
func (f *Filter) IsMatch(getter FilterGetter) bool {
	var gotMatch bool
	var conti bool
	var matchResult bool
	for _, match := range f.Matches {
		value := getter(match.Key)
		gotMatch, conti = match.IsMatch(value)
		if !gotMatch && !conti {
			return false
		}
		matchResult = matchResult || gotMatch
	}
	return matchResult
}

// NewFilter returns a new Filter instance.
func NewFilter(matches ...*Match) *Filter {
	f := &Filter{}
	return f.Add(matches...)
}
