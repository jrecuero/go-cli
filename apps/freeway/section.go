package freeway

import "fmt"

// Spec represents ...
type Spec string

// Traverse represents ...
type Traverse func(Spec) int

// Approach represents ...
type Approach func(int) int

// ISection represents ...
type ISection interface {
	GetLen() int
	GetWidth() int
	GetSpec() Spec
	Traversing() int
	Entering(int) int
	Exiting(int) int
}

// Section represents ...
type Section struct {
	length     int
	width      int
	spec       Spec
	traversing Traverse
	entering   Approach
	exiting    Approach
}

// GetLen is ...
func (section *Section) GetLen() int {
	return section.length
}

// GetWidth is ..,
func (section *Section) GetWidth() int {
	return section.width
}

// GetSpec is ...
func (section *Section) GetSpec() Spec {
	return section.spec
}

// Traversing is ...
func (section *Section) Traversing() int {
	if section.traversing != nil {
		return section.traversing(section.spec)
	}
	return 1
}

// Entering is ...
func (section *Section) Entering(speed int) int {
	if section.entering != nil {
		return section.entering(speed)
	}
	return 1
}

// Exiting is ...
func (section *Section) Exiting(speed int) int {
	if section.exiting != nil {
		return section.exiting(speed)
	}
	return 1
}

// String is ...
func (section *Section) String() string {
	return fmt.Sprintf("length/width: %d/%d spec: %s\n", section.GetLen(), section.GetWidth(), section.GetSpec())
}

// NewSection is ...
func NewSection(length int, width int, spec Spec, traversing Traverse, entering Approach, exiting Approach) *Section {
	return &Section{
		length:     length,
		width:      width,
		spec:       spec,
		traversing: traversing,
		entering:   entering,
		exiting:    exiting,
	}
}

// QSection represents ...
type QSection struct {
	section ISection
	queue   []IDevice
}

// GetSection is ...
func (qs *QSection) GetSection() ISection {
	return qs.section
}

// NewQSection is ...
func NewQSection(section ISection) *QSection {
	return &QSection{
		section: section,
	}
}
