package battle

import (
	"bytes"
	"fmt"
)

// TStat represents ...
type TStat int

const (
	// StatLix is the Life stat string representation.
	StatLix = "lix"
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

// IntToTStat is ...
func IntToTStat(val int) TStat {
	return TStat(val)
}

// InterfaceToTStat is ...
func InterfaceToTStat(val interface{}) TStat {
	return TStat(val.(int))
}

// StatPair represents ...
type StatPair struct {
	name string
	stat TStat
}

// Stats represents ...
type Stats struct {
	Lix TStat
	Str TStat
	Agi TStat
	Sta TStat
	Pre TStat
	Foc TStat
}

// Get is ...
func (stats *Stats) Get(name string) TStat {
	switch name {
	case StatLix:
		return stats.Lix
	case StatStr:
		return stats.Str
	case StatAgi:
		return stats.Agi
	case StatSta:
		return stats.Sta
	case StatPre:
		return stats.Pre
	case StatFoc:
		return stats.Foc
	default:
		panic(fmt.Sprintf("Unknown stat: %s", name))
	}
}

// GetString is ...
func (stats *Stats) GetString(name string) string {
	switch name {
	case StatLix:
		return "Life"
	case StatStr:
		return "Strength"
	case StatAgi:
		return "Agility"
	case StatSta:
		return "Stamina"
	case StatPre:
		return "Precision"
	case StatFoc:
		return "Focus"
	default:
		panic(fmt.Sprintf("Unknown stat: %s", name))
	}
}

// Set is ...
func (stats *Stats) Set(name string, stat TStat) *Stats {
	switch name {
	case StatLix:
		stats.Lix = stat
	case StatStr:
		stats.Str = stat
	case StatAgi:
		stats.Agi = stat
	case StatSta:
		stats.Sta = stat
	case StatPre:
		stats.Pre = stat
	case StatFoc:
		stats.Foc = stat
	default:
		panic(fmt.Sprintf("Unknown stat: %s", name))
	}
	return stats
}

// Sets is ...
func (stats *Stats) Sets(entries []*StatPair) *Stats {
	for _, statpair := range entries {
		stats.Set(statpair.name, statpair.stat)
	}
	return stats
}

// String is ...
func (stats *Stats) String() string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("-->Lix: %d\n", stats.Lix))
	buf.WriteString(fmt.Sprintf("-->Str: %d\n", stats.Str))
	buf.WriteString(fmt.Sprintf("-->Agi: %d\n", stats.Agi))
	buf.WriteString(fmt.Sprintf("-->Sta: %d\n", stats.Sta))
	buf.WriteString(fmt.Sprintf("-->Pre: %d\n", stats.Pre))
	buf.WriteString(fmt.Sprintf("-->Foc: %d\n", stats.Foc))
	return buf.String()
}

// NewStats is ...
func NewStats() *Stats {
	return &Stats{}
}

// UStatCb represents ...
type UStatCb func(TStat, IActor, ...interface{}) TStat

// UStats represents ...
type UStats struct {
	uLix UStatCb
	uStr UStatCb
	uAgi UStatCb
	uSta UStatCb
	uPre UStatCb
	uFoc UStatCb
}

// Get is ...
func (ustats *UStats) Get(name string) UStatCb {
	switch name {
	case StatLix:
		return ustats.uLix
	case StatStr:
		return ustats.uStr
	case StatAgi:
		return ustats.uAgi
	case StatSta:
		return ustats.uSta
	case StatPre:
		return ustats.uPre
	case StatFoc:
		return ustats.uFoc
	default:
		panic(fmt.Sprintf("Unknown stat: %s", name))
	}
}

// Set is ...
func (ustats *UStats) Set(name string, ustatCb UStatCb) *UStats {
	switch name {
	case StatLix:
		ustats.uLix = ustatCb
	case StatStr:
		ustats.uStr = ustatCb
	case StatAgi:
		ustats.uAgi = ustatCb
	case StatSta:
		ustats.uSta = ustatCb
	case StatPre:
		ustats.uPre = ustatCb
	case StatFoc:
		ustats.uFoc = ustatCb
	default:
		panic(fmt.Sprintf("Unknown stat: %s", name))
	}
	return ustats
}

// Call is ...
func (ustats *UStats) Call(name string, stat TStat, actor IActor, args ...interface{}) TStat {
	cb := ustats.Get(name)
	return cb(stat, actor, args...)
}

// plainStat is ...
func plainStat(stat TStat, actor IActor, args ...interface{}) TStat {
	return stat
}

// NewPlainUStats is ...
func NewPlainUStats() *UStats {
	return &UStats{
		uLix: plainStat,
		uStr: plainStat,
		uAgi: plainStat,
		uSta: plainStat,
		uPre: plainStat,
		uFoc: plainStat,
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
