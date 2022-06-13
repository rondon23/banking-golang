package domain

import (
	"banking/dto"
	"banking/errs"
)

type Account struct {
	AccountId   string
	CustomerId  string
	OpeningDate string
	AccountType string
	Amount      float64
	Status      string
}

func (a Account) ToNewAccountResponseDto() dto.NewAccountResponse {
	return dto.NewAccountResponse{a.AccountId}
}

type AccountRepository interface {
	Save(Account) (*Account, *errs.AppError)
	SaveTransaction(transaction Transaction) (*Transaction, *errs.AppError)
	FindById(accountId string) (*Account, *errs.AppError)
}

func (a Account) CanWithDrawal(amount float64) bool {
	if a.Amount < amount {
		return false
	}
	return true
}
