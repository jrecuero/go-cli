package battle

import (
	"bytes"
	"fmt"
	"strings"
)

// toShortName converts a string to a valid short name.
func toShortName(name string) string {
	if len(name) > 6 {
		return strings.ToUpper(name)[0:6]
	}
	return strings.ToUpper(name)
}

// IBase represents all basic and common interface methods to be used for
// techniques, styles and stances.
type IBase interface {
	GetName() string
	GetShortName() string
	GetDN() string
	GetDescription() string
	Enabled() bool
	Learned() bool
	Active() bool
	GetUpdateStats() *UStats
	IsDefault() bool
	SetShortName(string)
	SetDescription(string)
	SetEnabled(bool)
	SetLearned(bool)
	SetActive(bool)
	SetUpdateStats(*UStats)
	SetAsDefault(bool)
}

// Base represents all common arguments to be used in other structures like
// techniques, styles and stances.
type Base struct {
	name        string
	shortn      string
	desc        string
	enabled     bool
	learned     bool
	active      bool
	updatestats *UStats
	isdefault   bool
}

// GetName is ...
func (base *Base) GetName() string {
	return base.name
}

// GetShortName is ...
func (base *Base) GetShortName() string {
	return base.shortn
}

// GetDN is ...
func (base *Base) GetDN() string {
	return base.GetShortName()
}

// SetShortName is ...
func (base *Base) SetShortName(shortname string) {
	base.shortn = toShortName(shortname)
}

// GetDescription is ...
func (base *Base) GetDescription() string {
	return base.desc
}

// SetDescription is ...
func (base *Base) SetDescription(desc string) {
	base.desc = desc
}

// Enabled is ...
func (base *Base) Enabled() bool {
	return base.enabled
}

// SetEnabled is ...
func (base *Base) SetEnabled(enabled bool) {
	base.enabled = enabled
}

// Learned is ...
func (base *Base) Learned() bool {
	return base.learned
}

// SetLearned is ...
func (base *Base) SetLearned(learned bool) {
	base.learned = learned
}

// Active is ...
func (base *Base) Active() bool {
	return base.active
}

// SetActive is ...
func (base *Base) SetActive(active bool) {
	base.active = active
}

// GetUpdateStats is ...
func (base *Base) GetUpdateStats() *UStats {
	return base.updatestats
}

// SetUpdateStats is ...
func (base *Base) SetUpdateStats(ustats *UStats) {
	base.updatestats = ustats
}

// IsDefault is ...
func (base *Base) IsDefault() bool {
	return base.isdefault
}

// SetAsDefault is ...
func (base *Base) SetAsDefault(isdefault bool) {
	base.isdefault = isdefault
}

// String is ...
func (base *Base) String() string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%s\t[%t | %t | %t]",
		base.GetName(), base.Enabled(), base.Learned(), base.Active()))
	return buf.String()
}

// NewBase is ...
func NewBase(name string) *Base {
	return &Base{
		name:        name,
		shortn:      toShortName(name),
		updatestats: NewPlainUStats(),
	}
}

// NewFullBase is ...
func NewFullBase(name string, desc string, ustats *UStats) *Base {
	return &Base{
		name:        name,
		shortn:      toShortName(name),
		desc:        desc,
		updatestats: ustats,
	}
}
