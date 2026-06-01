package service

import (
	"apotek-pos-go/internal/modules/transaction/entity"
	"apotek-pos-go/internal/modules/transaction/repository"
	"errors"
)

type CustomerSrv interface {
	Create(customer *entity.Customer) error
	GetAll() ([]entity.Customer,error)
	GetUnitById(id uint) (entity.Customer,error)
	Delete(id uint) error
	Update(customer *entity.Customer) error
}

type CustomerSrvImpl struct{
	repo repository.CustomerRepo
}

func NewCustomerSrv(repo repository.CustomerRepo) CustomerSrv{
	return  &CustomerSrvImpl{repo: repo} 
}

func (r *CustomerSrvImpl) Create(customer *entity.Customer) error{
	if customer.Nama=="" {
		return errors.New("Nama tidak boleh kosong")
	}

	err:=r.repo.Create(customer)
	return err;
}

func (r *CustomerSrvImpl) Update(customer *entity.Customer) error{
	if customer.Nama=="" {
		return errors.New("Nama tidak boleh kosong")
	}


	return r.repo.Update(customer)
}

func (r *CustomerSrvImpl) GetAll() ([]entity.Customer,error){
	return r.repo.FindAll()
}

func (r *CustomerSrvImpl) GetUnitById(id uint) (entity.Customer,error){
	return r.repo.FindById(id)
}

func (r *CustomerSrvImpl) Delete(id uint) error{
	return r.repo.Delete(id)
}
