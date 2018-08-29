package dbase

var tranID int

// Transaction represents ...
type Transaction struct {
	ID     int
	Shadow *TableMap
	active bool
}

// Start is ...
func (tran *Transaction) Start() error {
	tran.active = true
	return nil
}

// close is ...
func (tran *Transaction) close() {
	tran.active = false
	tran.Shadow = nil
}

// Commit is ...
func (tran *Transaction) Commit() error {
	tran.close()
	return nil
}

// Abort is ...
func (tran *Transaction) Abort() error {
	tran.close()
	return nil
}

// NewTransaction is ...
func NewTransaction() *Transaction {
	tranID++
	return &Transaction{
		ID:     tranID,
		Shadow: NewTableMap(),
	}
}
