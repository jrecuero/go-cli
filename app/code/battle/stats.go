package battle

// TStat represents ...
type TStat int

// IntToTStat is ...
func IntToTStat(val int) TStat {
	return TStat(val)
}

// InterfaceToTStat is ...
func InterfaceToTStat(val interface{}) TStat {
	return TStat(val.(int))
}

// Stats represents ...
type Stats struct {
	Str TStat
	Agi TStat
	Sta TStat
	Pre TStat
	Foc TStat
}

// NewStats is ...
func NewStats() *Stats {
	return &Stats{}
}

// UStatCb represents ...
type UStatCb func(TStat, IActor, ...interface{}) TStat

const (
	// StatStr is the Strength stat string representation.
	StatStr = "str"
	// StatAgi is the Agiity stat string representation.
	StatAgi = "agi"
	// StatSta is the Stamina stat string representation.
	StatSta = "sta"
	// StatPre is the Precision stat string representation.
	StatPre = "pre"
	// StatFoc is the Focus stat string representation.
	StatFoc = "foc"
)

// UStats represents ...
type UStats struct {
	UStr UStatCb
	UAgi UStatCb
	USta UStatCb
	UPre UStatCb
	UFoc UStatCb
}

// Set is ...
func (ustats *UStats) Set(name string, ustatCb UStatCb) *UStats {
	switch name {
	case StatStr:
		ustats.UStr = ustatCb
	case StatAgi:
		ustats.UAgi = ustatCb
	case StatSta:
		ustats.USta = ustatCb
	case StatPre:
		ustats.UPre = ustatCb
	case StatFoc:
		ustats.UFoc = ustatCb
	default:
		return nil
	}
	return ustats
}

// Update is ...
func (ustats *UStats) Update(name string, stat TStat, actor IActor, args ...interface{}) TStat {
	switch name {
	case StatStr:
		return ustats.UStr(stat, actor, args...)
	case StatAgi:
		return ustats.UAgi(stat, actor, args...)
	case StatSta:
		return ustats.USta(stat, actor, args...)
	case StatPre:
		return ustats.UPre(stat, actor, args...)
	case StatFoc:
		return ustats.UFoc(stat, actor, args...)
	default:
		panic("Unknown stat")
	}
	return 0
}

// plainStat is ...
func plainStat(stat TStat, actor IActor, args ...interface{}) TStat {
	return stat
}

// NewPlainUStats is ...
func NewPlainUStats() *UStats {
	return &UStats{
		UStr: plainStat,
		UAgi: plainStat,
		USta: plainStat,
		UPre: plainStat,
		UFoc: plainStat,
	}
}

// NewUStats is ...
func NewUStats(updates map[string]UStatCb) *UStats {
	ustats := NewPlainUStats()
	for k, ustatCb := range updates {
		ustats.Set(k, ustatCb)
	}
	return ustats
}
