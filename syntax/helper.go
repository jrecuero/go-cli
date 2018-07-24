package syntax

// CompleteHelp represents complete and help together.
type CompleteHelp struct {
	Complete interface{} // token complete value.
	Help     interface{} // token help value.
}

// NewCompleteHelp returns a new CompleteHelp instance.
func NewCompleteHelp(complete interface{}, help interface{}) *CompleteHelp {
	return &CompleteHelp{
		Complete: complete,
		Help:     help,
	}
}

// GetCompletes returns all complete values for an array of CompleteHelp
// isntances.
func GetCompletes(in []*CompleteHelp) []interface{} {
	var completes []interface{}
	for _, entry := range in {
		completes = append(completes, entry.Complete)
	}
	return completes
}

// GetHelps returns all help values for an array of CompleteHelp instances.
func GetHelps(in []*CompleteHelp) []interface{} {
	var helps []interface{}
	for _, entry := range in {
		helps = append(helps, entry.Help)
	}
	return helps
}
