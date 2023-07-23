package presenter

import "silver-clean-code/internal/entity"

type Accounts struct {
	TotalItems int        `json:"total_items"`
	Items      []*Account `json:"items"`
}

func NewAccountsPresenter(entities []*entity.Account) *Accounts {
	accounts := &Accounts{}
	if entities != nil {
		return accounts.Parse(entities)
	}
	return accounts
}

func (acs *Accounts) Parse(entities []*entity.Account) *Accounts {
	for _, ent := range entities {
		acs.Items = append(acs.Items, NewAccountPresenter(ent))
	}
	acs.TotalItems = len(acs.Items)
	return acs
}
