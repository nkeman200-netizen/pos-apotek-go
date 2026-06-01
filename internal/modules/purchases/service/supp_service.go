package service

import (
	"apotek-pos-go/internal/modules/purchases/entity"
	"apotek-pos-go/internal/modules/purchases/repo"
	"errors"
)

type SupplierSrv interface {
	Create(s *entity.Supplier) error
	Update(s *entity.Supplier) error
	Delete(id uint) error
	GetAll() ([]entity.Supplier, error)
	GetById(id uint) (entity.Supplier, error)
}

type supplierSrvImp struct {
	repo repo.SupplierRepo
}

func NewSupplierSrv(repo repo.SupplierRepo) SupplierSrv {
	return &supplierSrvImp{repo: repo}
}

// Create implements [SupplierSrv].
func (srv *supplierSrvImp) Create(s *entity.Supplier) error {
	if s.Nama=="" {
		return errors.New("Nama supplier tidak boleh kosong")
	}
	return srv.repo.Create(s)
}

// Update implements [SupplierSrv].
func (srv *supplierSrvImp) Update(s *entity.Supplier) error {
	if s.Nama=="" {
		return errors.New("Nama supplier tidak boleh kosong")
	}
	return srv.repo.Update(s)
}

// Delete implements [SupplierSrv].
func (s *supplierSrvImp) Delete(id uint) error {
	return s.repo.Delete(id)
}

// GetAll implements [SupplierSrv].
func (s *supplierSrvImp) GetAll() ([]entity.Supplier, error) {
	return s.repo.GetAll()
}

// GetById implements [SupplierSrv].
func (s *supplierSrvImp) GetById(id uint) (entity.Supplier, error) {
	return s.repo.GetById(id)
}

