package repository

import (
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"github.com/pstano1/customer-api/internal/pkg"
	"github.com/pstano1/customer-api/service"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) service.IRepository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) CreateCustomer(request pkg.CustomerCreate) (pkg.Customer, error) {
	var customer pkg.Customer
	err := copier.Copy(&customer, request)
	if err != nil {
		return customer, err
	}
	customer.Id = uuid.New().String()

	return customer, r.db.Create(customer).Error
}

func (r *Repository) GetCustomer(id, tag *string) (pkg.Customer, error) {
	var customer pkg.Customer
	var err error
	if id != nil {
		err = r.db.Where("id = ?", id).First(&customer).Error
	} else {
		err = r.db.Where("tag = ?", tag).First(&customer).Error
	}

	return customer, err
}
