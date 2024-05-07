package service

import "github.com/pstano1/customer-api/internal/pkg"

type IRepository interface {
	CreateCustomer(pkg.CustomerCreate) (pkg.Customer, error)
	GetCustomer(id, tag *string) (pkg.Customer, error)
}

type IAPI interface {
	CreateCustomer(customer pkg.CustomerCreate) (pkg.Customer, error)
	ExchangeTagForId(tag string) (*string, error)
	ValidateId(id string) (bool, error)
}
