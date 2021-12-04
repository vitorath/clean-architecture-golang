package entity

import "errors"

type Transation struct {
	ID           string
	AccountID    string
	Amount       float64
	Status       string
	ErrorMessage string
}

func NewTransaction() *Transation {
	return &Transation{}
}

func (t *Transation) IsValid() error {
	if t.Amount > 1000 {
		return errors.New("you don't have limit for this transaction")
	}

	return nil
}
