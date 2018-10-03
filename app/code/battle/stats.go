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

// UStatCb represents ...
type UStatCb func(TStat) TStat

// UStats represents ...
type UStats struct {
	UStr UStatCb
	UAgi UStatCb
	USta UStatCb
	UPre UStatCb
	UFoc UStatCb
}

// NewStats is ...
func NewStats() *Stats {
	return &Stats{}
}

// plainStat is ...
func plainStat(stat TStat) TStat {
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
