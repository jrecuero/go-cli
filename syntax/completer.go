package syntax

const _joint = ""
const _rooty = "ROOT"
const _sink = "SINK"
const _start = "START"
const _end = "END"
const _loop = "LOOP"

// Completer represents any generic completer.
type Completer struct {
	label   string
	Content IContent
}

// NewCompleter returns a new Completer instance.
func NewCompleter(content IContent) *Completer {
	c := &Completer{
		label:   content.GetLabel(),
		Content: content,
	}
	return c
}

// GetContent returns the content for any completer.
func (c *Completer) GetContent() IContent {
	return c.Content
}

// GetLabel returns the label for any completer.
func (c *Completer) GetLabel() string {
	return c.label
}

// Validate checks if the content is value for the given line.
func (c *Completer) Validate(ctx *Context, line interface{}) bool {
	return true
}

// Help returns the help for any node completer.
func (c *Completer) Help(ctx *Context, line interface{}) (interface{}, bool) {
	return nil, false
}

// Query returns the query for any node completer.
func (c *Completer) Query(ctx *Context, line interface{}) (interface{}, bool) {
	return nil, false
}

// Complete returns the complete match for any node completer.
func (c *Completer) Complete(ctx *Context, line interface{}) (interface{}, bool) {
	return nil, false
}

// CCommand represents the completer for a command node.
type CCommand struct {
	*Completer
}

// Match returns the match for a command node completer.
func (cc *CCommand) Match(ctx *Context, line interface{}, index int) (int, bool) {
	tokens := line.([]string)
	if tokens[0] == cc.GetLabel() {
		return index + 1, true
	}
	return index, false
}

// CIdent represents the completer for any CIdent node.
type CIdent struct {
	*Completer
}

// NewCIdent returns a new CIdent instance.
func NewCIdent(content IContent) *CIdent {
	i := &CIdent{
		&Completer{
			label:   content.GetLabel(),
			Content: content,
		},
	}
	return i
}

// Match returns the match for any ident node completer.
func (i *CIdent) Match(ctx *Context, line interface{}, index int) (int, bool) {
	tokens := line.([]string)
	if tokens[0] == i.GetLabel() {
		return index + 1, true
	}
	return index, false
}

//// Help returns the help for any ident node completer.
//func (i *CIdent) Help(ctx *Context, line interface{}) (interface{}, bool) {
//    return nil, false
//}

//// Query returns the query for any ident node completer.
//func (i *CIdent) Query(ctx *Context, line interface{}) (interface{}, bool) {
//    return nil, false
//}

//// Complete returns the complete match for any ident node completer.
//func (i *CIdent) Complete(ctx *Context, line interface{}) (interface{}, bool) {
//    return nil, false
//}

// CAny represents the completer for CAny character sequence node.
type CAny struct {
	*Completer
}

// NewCAny returns a new CAny instance.
func NewCAny(content IContent) *CAny {
	a := &CAny{
		&Completer{
			label:   content.GetLabel(),
			Content: content,
		},
	}
	return a
}

// Match returns the match for CAny node completer.
func (a *CAny) Match(ctx *Context, line interface{}, index int) (int, bool) {
	tokens := line.([]string)
	if tokens[0] == CR.GetLabel() {
		return index, false
	}
	return index + 1, true
}

//// Query returns the query for CAny node completer.
//func (a *CAny) Query(ctx *Context, line interface{}) (interface{}, bool) {
//    return nil, false
//}

//// Help returns the help for CAny node completer.
//func (a *CAny) Help(ctx *Context, line interface{}) (interface{}, bool) {
//    return nil, false
//}

//// Complete returns the complete match for CAny node completer.
//func (a *CAny) Complete(ctx *Context, line interface{}) (interface{}, bool) {
//    return nil, false
//}

// CCustom represents the completer for CCustom node.
type CCustom struct {
	*Completer
}

// NewCCustom returns a new CCustom instance.
func NewCCustom(content IContent) *CCustom {
	c := &CCustom{
		&Completer{
			label:   content.GetLabel(),
			Content: content,
		},
	}
	return c
}

// Match returns the match for CCustom node completer.
func (c *CCustom) Match(ctx *Context, line interface{}, index int) (int, bool) {
	tokens := line.([]string)
	if tokens[0] == c.GetContent().GetLabel() {
		return index + 1, true
	}
	return index, false
}

//// Query returns the query for CCustom node completer.
//func (c *CCustom) Query(ctx *Context, line interface{}) (interface{}, bool) {
//    return nil, false
//}

//// Help returns the help for CCustom node completer.
//func (c *CCustom) Help(ctx *Context, line interface{}) (interface{}, bool) {
//    return nil, false
//}

//// Complete returns the complete match for CCustom node completer.
//func (c *CCustom) Complete(ctx *Context, line interface{}) (interface{}, bool) {
//    return nil, false
//}

// CJoint represents the completer for any joint node.
type CJoint struct {
	*Completer
}

// NewCJoint returns a new CJoint instance.
func NewCJoint(label string) *CJoint {
	if label == "" {
		label = "JOINT"
	}
	j := &CJoint{
		&Completer{
			label: label,
		},
	}
	return j
}

// Match returns the match for any joint node completer.
func (j *CJoint) Match(ctx *Context, line interface{}, index int) (int, bool) {
	if j.GetContent() == nil {
		return index, true
	}
	tokens := line.([]string)
	if tokens[0] == j.GetContent().GetLabel() {
		return index + 1, true
	}
	return index, false
}

//// Query returns the query for any joint node completer.
//func (j *CJoint) Query(ctx *Context, line interface{}) (interface{}, bool) {
//    return nil, false
//}

//// Help returns the help for any joint node completer.
//func (j *CJoint) Help(ctx *Context, line interface{}) (interface{}, bool) {
//    return nil, false
//}

//// Complete returns the complete match for any joint node completer.
//func (j *CJoint) Complete(ctx *Context, line interface{}) (interface{}, bool) {
//    return nil, false
//}

// NewCStart returns a new Start instance.
func NewCStart() *CJoint {
	return NewCJoint(_start)
}

// NewCEnd returns a new End instance.
func NewCEnd() *CJoint {
	return NewCJoint(_end)
}

// NewCLoop returns a new Loop instance.
func NewCLoop() *CJoint {
	return NewCJoint(_loop)
}

// NewCSink returns a new sinkn completer instance.
func NewCSink() *CJoint {
	j := &CJoint{
		&Completer{
			label:   _sink,
			Content: CR,
		},
	}
	return j
}
