package engine

// ListaCondition represents ...
type ListaCondition struct {
	key   string
	value interface{}
	let   ConditionLet
}

// Check is ...
func (cc ListaCondition) Check(lista *Lista) bool {
	switch cc.let {
	case EQUAL:
		if listaValue, ok := lista.Get(cc.key); ok {
			return listaValue == cc.value
		}
		break
	case NEQUAL:
		if listaValue, ok := lista.Get(cc.key); ok {
			return listaValue != cc.value
		}
		break
	case EXIST:
		if _, ok := lista.Get(cc.key); ok {
			return true
		}
		break
	case NEXIST:
		if _, ok := lista.Get(cc.key); !ok {
			return true
		}
		break
	default:
		break
	}
	return false
}

// NewListaCondition is ...
func NewListaCondition(key string, value interface{}, let ConditionLet) *ListaCondition {
	return &ListaCondition{
		key:   key,
		value: value,
		let:   let,
	}
}

// Lista represents ...
type Lista struct {
	data map[string]interface{}
}

// Size is ...
func (lista *Lista) Size() int {
	return len(lista.data)
}

// Add is ...
func (lista *Lista) Add(name string) bool {
	if _, ok := lista.data[name]; !ok {
		lista.data[name] = nil
		return true
	}
	return false
}

// Delete is ...
func (lista *Lista) Delete(name string) bool {
	if _, ok := lista.data[name]; ok {
		delete(lista.data, name)
		return true
	}
	return false
}

// Set is ...
func (lista *Lista) Set(name string, value interface{}) bool {
	lista.data[name] = value
	return true
}

// Get is ...
func (lista *Lista) Get(name string) (interface{}, bool) {
	if value, ok := lista.data[name]; ok {
		return value, true
	}
	return nil, false
}

// CheckCondition is ...
func (lista *Lista) CheckCondition(cond interface{}) bool {
	condition := cond.(*ListaCondition)
	return condition.Check(lista)
}

// NewLista is ...
func NewLista() *Lista {
	return &Lista{
		data: make(map[string]interface{}),
	}
}
