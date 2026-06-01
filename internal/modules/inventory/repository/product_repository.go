package repository

import (
	"apotek-pos-go/internal/modules/inventory/entity"

	"gorm.io/gorm"
)

type ProductRepository interface {
	Create(product *entity.Product)error
	Update(product *entity.Product)error
	Delete(id uint) error
	FindAll() ([]entity.Product,error)
	FindById(id uint) (entity.Product,error)
	FindBySKU(sku string)(entity.Product,error)
}

type productRepositoryimpl struct{
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository{
	return &productRepositoryimpl{db: db}
}

func (r *productRepositoryimpl) Create(product *entity.Product) error{
	return r.db.Create(product).Error
}
func (r *productRepositoryimpl) Update(product *entity.Product) error{
	return r.db.Model(&entity.Product{}).Where("id=?",product.ID).Updates(product).Error
}

func (r *productRepositoryimpl) Delete(id uint) error{
	return r.db.Delete(&entity.Product{},id).Error
}

func(r *productRepositoryimpl) FindAll()([]entity.Product,error){
	var product []entity.Product
	err:=r.db.Preload("Unit").Preload("Category").Find(&product).Error
	return product,err
}

func(r *productRepositoryimpl) FindById(id uint)(entity.Product,error){
	var product entity.Product
	queryStock:="(select COALESCE(SUM(stock),0) from product_batches where product_batches.product_id=products.id)"
	err:=r.db.Select("products.*, "+queryStock+" as total_stock").Preload("Unit").Preload("Category").First(&product,id).Error
	return product,err
}
func(r *productRepositoryimpl) FindBySKU(sku string)(entity.Product,error){
	var product entity.Product
	queryStock:="(select COALESCE(SUM(stock),0) from product_batches where product_batches.product_id=products.id)"
	err:=r.db.Select("products.*, "+queryStock+" as total_stock").Preload("Unit").Preload("Category").Where("SKU = ?",sku).First(&product).Error
	return product,err
}
