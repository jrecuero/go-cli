package battle

// IListable represents ...
type IListable interface {
	GetName() string
}

// IListHandler represents ..
type IListHandler interface {
	GetAll() []IListable
	Add(...IListable) bool
	Remove(...IListable) bool
	RemoveByName(string) bool
	GetByName(string) IListable
}

// ListHandler represents ...
type ListHandler struct {
	lista []IListable
}

// GetAll is ...
func (lh *ListHandler) GetAll() []IListable {
	return lh.lista
}

// Add is ...
func (lh *ListHandler) Add(entries ...IListable) bool {
	lh.lista = append(lh.lista, entries...)
	return true
}

// Remove is ...
func (lh *ListHandler) Remove(entries ...IListable) bool {
	for _, entry := range entries {
		if !lh.RemoveByName(entry.GetName()) {
			return false
		}
	}
	return true
}

// RemoveByName is ...
func (lh *ListHandler) RemoveByName(name string) bool {
	for i, entry := range lh.lista {
		if entry.GetName() == name {
			lh.lista = append(lh.lista[:i], lh.lista[i+1:]...)
			return true
		}
	}
	return false
}

// GetByName is ...
func (lh *ListHandler) GetByName(name string) IListable {
	for _, entry := range lh.lista {
		if entry.GetName() == name {
			return entry
		}
	}
	return nil
}
