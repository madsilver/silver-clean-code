package entity

import (
	"math"
	"time"
)

type Transaction struct {
	TransactionID   uint64
	AccountID       uint64
	Amount          float64
	OperationTypeID int
	EventDate       string
}

func NewTransactionEntity() *Transaction {
	return &Transaction{}
}

func (te *Transaction) NormalizeAmount() *Transaction {
	te.Amount = math.Abs(te.Amount)
	if te.OperationTypeID != OperationType.Payment {
		te.Amount = te.Amount * -1
	}
	return te
}

func (te *Transaction) TimeNow() *Transaction {
	t := time.Now()
	te.EventDate = t.Format(time.RFC3339Nano)
	return te
}
