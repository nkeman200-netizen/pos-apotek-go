package repository

import (
	"apotek-pos-go/internal/modules/transaction/entity"

	"gorm.io/gorm"
)

type CustomerRepo interface {
	Create(c *entity.Customer) error
	Update(c *entity.Customer) error
	FindAll() ([]entity.Customer, error)
	FindById(id uint) (entity.Customer, error)
}

type customerRepoImp struct {
	db *gorm.DB
}

// Create implements [CustomerRepo].
func (r *customerRepoImp) Create(c *entity.Customer) error {
	return r.db.Create(c).Error
}

// FindAll implements [CustomerRepo].
func (r *customerRepoImp) FindAll() ([]entity.Customer, error) {
	var cus []entity.Customer
	err:=r.db.Find(&cus).Error
	return cus,err
}

// FindById implements [CustomerRepo].
func (r *customerRepoImp) FindById(id uint) (entity.Customer, error) {
	var cus entity.Customer
	err:=r.db.Find(&cus,id).Error
	return cus,err
}

// Update implements [CustomerRepo].
func (r *customerRepoImp) Update(c *entity.Customer) error {
	return r.db.Save(&c).Error
	
}

func NewCustomerRepo(db *gorm.DB) CustomerRepo {
	return &customerRepoImp{db: db}
}
