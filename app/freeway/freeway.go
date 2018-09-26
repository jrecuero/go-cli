package freeway

import (
	"bytes"
	"fmt"
)

// Freeway represents ...
type Freeway struct {
	qsections []*QSection
}

// GetSection is ...
func (fway *Freeway) GetSection(isect int) ISection {
	return fway.qsections[isect].section
}

// AddSection is ...
func (fway *Freeway) AddSection(section ISection) *Freeway {
	fway.qsections = append(fway.qsections, NewQSection(section))
	return fway
}

// GetLen is ...
func (fway *Freeway) GetLen() int {
	return len(fway.qsections)
}

// LapLen is ...
func (fway *Freeway) LapLen() int {
	total := 0
	for i := 0; i < fway.GetLen(); i++ {
		total += fway.qsections[i].section.GetLen()
	}
	return total
}

// NextSectionIndex is ...
func (fway *Freeway) NextSectionIndex(isect int) (int, bool) {
	isect++
	if fway.GetLen() <= isect {
		return 0, true
	}
	return isect, false
}

// String is ...
func (fway *Freeway) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("len: %d\n", fway.GetLen()))
	for i, qs := range fway.qsections {
		buffer.WriteString(fmt.Sprintf("\t%d: %s\n", i, qs.GetSection()))
	}
	return buffer.String()
}

// NewFreeway is ...
func NewFreeway() *Freeway {
	return &Freeway{}
}
