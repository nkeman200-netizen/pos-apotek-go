package repo

import (
	"apotek-pos-go/internal/modules/purchases/entity"

	"gorm.io/gorm"
)

type SupplierRepo interface {
	Create(s *entity.Supplier) error
	Update(s *entity.Supplier) error
	Delete(id uint) error
	GetAll() ([]entity.Supplier, error)
	GetById(id uint) (entity.Supplier, error)
}

type supplierRepoImpl struct {
	db *gorm.DB
}

func NewSupplierRepo(db *gorm.DB) SupplierRepo {
	return &supplierRepoImpl{db: db}
}

// Create implements [SupplierRepo].
func (sr *supplierRepoImpl) Create(s *entity.Supplier) error {
	return sr.db.Create(s).Error
}

// Update implements [SupplierRepo].
func (sr *supplierRepoImpl) Update(s *entity.Supplier) error {
	return sr.db.Model(&entity.Supplier{}).Where("id=?",s.Id).Updates(s).Error
}

// Delete implements [SupplierRepo].
func (s *supplierRepoImpl) Delete(id uint) error {
	return s.db.Delete(&entity.Supplier{},id).Error
}

// GetAll implements [SupplierRepo].
func (s *supplierRepoImpl) GetAll() ([]entity.Supplier, error) {
	var supplier []entity.Supplier
	err:=s.db.Find(&supplier).Error
	return supplier,err
}

// GetById implements [SupplierRepo].
func (s *supplierRepoImpl) GetById(id uint) (entity.Supplier, error) {
	var supplier entity.Supplier
	err:=s.db.First(&supplier,id).Error
	return supplier,err
}

