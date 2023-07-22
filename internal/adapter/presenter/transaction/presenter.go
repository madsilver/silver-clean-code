package transaction

import "silver-clean-code/internal/entity"

type Transaction struct {
	TransactionID   uint64  `json:"transaction_id"`
	AccountID       uint64  `json:"account_id"`
	Amount          float64 `json:"amount"`
	OperationTypeID int     `json:"operation_type_id"`
	EventDate       string  `json:"eventDate"`
}

type Transactions struct {
	TotalItems int           `json:"total_items"`
	Items      []Transaction `json:"items"`
}

func ToPresenter(ent *entity.Transaction) Transaction {
	return Transaction{
		TransactionID:   ent.TransactionID,
		AccountID:       ent.AccountID,
		Amount:          ent.Amount,
		OperationTypeID: ent.OperationTypeID,
		EventDate:       ent.EventDate,
	}
}

func ToPresenters(entities []*entity.Transaction) Transactions {
	var presenters []Transaction
	for _, ent := range entities {
		presenters = append(presenters, ToPresenter(ent))
	}
	return Transactions{
		Items:      presenters,
		TotalItems: len(presenters),
	}
}

func ToEntity(pre *Transaction) *entity.Transaction {
	return &entity.Transaction{
		TransactionID:   pre.TransactionID,
		AccountID:       pre.AccountID,
		Amount:          pre.Amount,
		OperationTypeID: pre.OperationTypeID,
		EventDate:       pre.EventDate,
	}
}
