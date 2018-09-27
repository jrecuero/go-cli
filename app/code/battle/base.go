package battle

// IBase represents all basic and common interface methods to be used for
// techniques, styles and stances.
type IBase interface {
	GetName() string
	GetDescription() string
	Enabled() bool
	SetEnabled(bool)
	Learned() bool
	SetLearned(bool)
	Active() bool
	SetActive(bool)
	GetUpdateStats() *UStats
	SetUpdateStats(*UStats)
}

// Base represents all common arguments to be used in other structures like
// techniques, styles and stances.
type Base struct {
	name        string
	desc        string
	enabled     bool
	learned     bool
	active      bool
	updatestats *UStats
}

// GetName is ...
func (base *Base) GetName() string {
	return base.name
}

// GetDescription is ...
func (base *Base) GetDescription() string {
	return base.desc
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

// NewBase is ...
func NewBase(name string) *Base {
	return &Base{
		name:        name,
		updatestats: NewPlainUStats(),
	}
}
