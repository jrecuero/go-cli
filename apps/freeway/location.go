package freeway

import "fmt"

// Location represents ...
type Location struct {
	freeway *Freeway
	isect   int
	pos     int
	laps    int
}

// GetFreeway is ...
func (loc *Location) GetFreeway() *Freeway {
	return loc.freeway
}

// GetSection is ...
func (loc *Location) GetSection() ISection {
	return loc.freeway.GetSection(loc.isect)
}

// GetPos is ...
func (loc *Location) GetPos() int {
	return loc.pos
}

// SetPos is ...
func (loc *Location) SetPos(pos int) {
	loc.pos = pos
}

// GetLocationIndex is ...
func (loc *Location) GetLocationIndex() (int, int) {
	return loc.isect, loc.GetPos()
}

// GetLocation is ...
func (loc *Location) GetLocation() (ISection, int) {
	return loc.GetSection(), loc.GetPos()
}

// GetLaps is ...
func (loc *Location) GetLaps() int {
	return loc.laps
}

// NextSection is ...
func (loc *Location) NextSection() (ISection, int) {
	lap := false
	if loc.isect, lap = loc.freeway.NextSectionIndex(loc.isect); lap {
		loc.laps++
	}
	loc.SetPos(0)
	return loc.GetLocation()
}

// String is ...
func (loc *Location) String() string {
	return fmt.Sprintf("sec: %d pos: %d\n", loc.isect, loc.pos)
}

// NewLocation is ...
func NewLocation(freeway *Freeway) *Location {
	return &Location{
		freeway: freeway,
	}
}
