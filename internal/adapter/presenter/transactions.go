package presenter

import "silver-clean-code/internal/entity"

type Transactions struct {
	TotalItems int            `json:"total_items"`
	Items      []*Transaction `json:"items"`
}

func NewTransactionsPresenter(entities []*entity.Transaction) *Transactions {
	transactions := &Transactions{}
	if entities != nil {
		return transactions.Parse(entities)
	}
	return transactions
}

func (tps *Transactions) Parse(entities []*entity.Transaction) *Transactions {
	for _, ent := range entities {
		tps.Items = append(tps.Items, NewTransactionPresenter(ent))
	}
	tps.TotalItems = len(tps.Items)
	return tps
}
