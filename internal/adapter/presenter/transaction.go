package presenter

import "silver-clean-code/internal/entity"

type Transaction struct {
	TransactionID   uint64  `json:"transaction_id"`
	AccountID       uint64  `json:"account_id"`
	Amount          float64 `json:"amount"`
	OperationTypeID int     `json:"operation_type_id"`
	EventDate       string  `json:"eventDate"`
}

func NewTransactionPresenter(ent *entity.Transaction) *Transaction {
	transaction := &Transaction{}
	if ent != nil {
		return transaction.Parse(ent)
	}
	return transaction
}

func (tp *Transaction) Parse(ent *entity.Transaction) *Transaction {
	tp.TransactionID = ent.TransactionID
	tp.AccountID = ent.AccountID
	tp.Amount = ent.Amount
	tp.OperationTypeID = ent.OperationTypeID
	tp.EventDate = ent.EventDate
	return tp
}

func (tp *Transaction) ToEntity() *entity.Transaction {
	return &entity.Transaction{
		TransactionID:   tp.TransactionID,
		AccountID:       tp.AccountID,
		Amount:          tp.Amount,
		OperationTypeID: tp.OperationTypeID,
		EventDate:       tp.EventDate,
	}
}
