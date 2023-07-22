package entity

type Transaction struct {
	TransactionID   uint64
	AccountID       uint64
	Amount          float64
	OperationTypeID int
	EventDate       string
}
