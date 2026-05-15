package repository

import (
	"apotek-pos-go/internal/modules/inventory/entity"

	"gorm.io/gorm"
)

type ProductBatchRepository interface {
	Create(pb *entity.ProductBatch) error
	Update(pb *entity.ProductBatch) error
	Delete(id uint) error
	FindAll() ([]entity.ProductBatch,error)
	FindById(id uint) (entity.ProductBatch,error)
}

type productBatchRepositoryImpl struct{
	db *gorm.DB
}

func NewProductBatchRepository(db *gorm.DB) ProductBatchRepository{
	return &productBatchRepositoryImpl{db: db}
}

func (r *productBatchRepositoryImpl) Create(pb *entity.ProductBatch) error{
	return r.db.Create(pb).Error
}

func (r *productBatchRepositoryImpl) Update(pb *entity.ProductBatch) error{
	return r.db.Save(pb).Error //golang otomatis query update jika di dalam entity terdapat data id
}

func(r *productBatchRepositoryImpl) Delete(id uint)error{
	return r.db.Delete(&entity.ProductBatch{},id).Error
}

func(r *productBatchRepositoryImpl) FindAll() ([]entity.ProductBatch,error){
	var pb []entity.ProductBatch
	err:=r.db.Preload("Product").Find(&pb).Error
	return pb,err
}

func(r *productBatchRepositoryImpl) FindById(id uint) (entity.ProductBatch,error){
	var pb entity.ProductBatch
	err:=r.db.Preload("Product").First(&pb,id).Error
	return pb,err
}