package engine

// FlagCondition represents ...
type FlagCondition struct {
	key   string
	value interface{}
	let   ConditionLet
}

// Check is ...
func (fc FlagCondition) Check(flag *Flag) bool {
	switch fc.let {
	case EQUAL:
		if flagValue, ok := flag.Get(fc.key); ok {
			return flagValue == fc.value
		}
		break
	case NEQUAL:
		if flagValue, ok := flag.Get(fc.key); ok {
			return flagValue != fc.value
		}
		break
	case EXIST:
		if _, ok := flag.Get(fc.key); ok {
			return true
		}
		break
	case NEXIST:
		if _, ok := flag.Get(fc.key); !ok {
			return true
		}
		break
	default:
		break
	}
	return false
}

// NewFlagCondition is ...
func NewFlagCondition(key string, value interface{}, let ConditionLet) *FlagCondition {
	return &FlagCondition{
		key:   key,
		value: value,
		let:   let,
	}
}

// Flag represents ...
type Flag struct {
	data map[string]interface{}
}

// Size is ...
func (flag *Flag) Size() int {
	return len(flag.data)
}

// Add is ...
func (flag *Flag) Add(name string) bool {
	if _, ok := flag.data[name]; !ok {
		flag.data[name] = nil
		return true
	}
	return false
}

// Delete is ...
func (flag *Flag) Delete(name string) bool {
	if _, ok := flag.data[name]; ok {
		delete(flag.data, name)
		return true
	}
	return false
}

// Set is ...
func (flag *Flag) Set(name string, value interface{}) bool {
	flag.data[name] = value
	return true
}

// Get is ...
func (flag *Flag) Get(name string) (interface{}, bool) {
	if value, ok := flag.data[name]; ok {
		return value, true
	}
	return nil, false
}

// CheckCondition is ...
func (flag *Flag) CheckCondition(cond interface{}) bool {
	return true
}

// NewFlag is ...
func NewFlag() *Flag {
	return &Flag{
		data: make(map[string]interface{}),
	}
}
