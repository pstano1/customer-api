package api

import (
	"github.com/pstano1/customer-api/internal/pkg"
	"github.com/pstano1/customer-api/service"
)

type API struct {
	repository service.IRepository
}

func New(repository service.IRepository) service.IAPI {
	return &API{
		repository: repository,
	}
}

func (a *API) CreateCustomer(request pkg.CustomerCreate) (pkg.Customer, error) {
	return a.repository.CreateCustomer(request)
}

func (a *API) ExchangeTagForId(tag string) (*string, error) {
	customer, err := a.repository.GetCustomer(nil, &tag)
	if err != nil {
		return nil, err
	}
	return &customer.Id, nil
}

func (a *API) ValidateId(id string) (bool, error) {
	_, err := a.repository.GetCustomer(&id, nil)
	if err != nil {
		return false, err
	}
	return true, nil
}
