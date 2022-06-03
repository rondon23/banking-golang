package domain

import "banking/errs"

type Customer struct {
	Id          string
	Name        string
	City        string
	ZipCode     string
	DateofBirth string
	Status      string
}

type CustomerRepository interface {
	FindAll() ([]Customer, *errs.AppError)
	ById(string) (*Customer, *errs.AppError)
}
