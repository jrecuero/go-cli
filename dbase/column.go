package dbase

// Column represents information required to identify any column in the table.
// Some of the information is related with the data and some other information
// is related with the column layout to be displayed if required.
type Column struct {
	Name  string
	Label string
	CType string
	Width int
	Align string
	Key   bool
}

// SetLayout assgins values to the column layout attributes.
func (col *Column) SetLayout(label interface{}, width int, align string) *Column {
	var _label string
	if label == nil {
		_label = col.Name
	} else {
		_label = label.(string)
	}
	col.Label = _label
	col.Width = width
	col.Align = align
	return col
}

// NewColumn creates a new Column instance. Only data related attributes are
// being initialized at this stage.
func NewColumn(name string, ctype string) *Column {
	return &Column{
		Name:  name,
		CType: ctype,
	}
}
