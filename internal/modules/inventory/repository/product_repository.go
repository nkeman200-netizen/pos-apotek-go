package repository

import (
	"apotek-pos-go/internal/modules/inventory/entity"

	"gorm.io/gorm"
)

type ProductRepository interface {
	Create(product *entity.Product)error
	Update(product *entity.Product)error
	Delete(id uint) error
	FindAll(filter entity.ProductFilter) ([]entity.Product,error)
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

func(r *productRepositoryimpl) FindAll(filter entity.ProductFilter)([]entity.Product,error){
	var product []entity.Product
	query:=r.db.Model(&entity.Product{}).
		Preload("Unit").
		Preload("Category").
		Where("is_active=?",true)

	if filter.CategoryId!="" {
		query=query.Where("category_id=?",filter.CategoryId)
	}
	if filter.Name!="" {
		query=query.Where("name like ?","%"+filter.Name+"%")
	}
	if filter.SKU!="" {
		query=query.Where("sku=?",filter.SKU)
	}
	if filter.UnitId!="" {
		query=query.Where("unit_id=?",filter.UnitId)
	}

	offset := (filter.Page - 1) * filter.Limit
	err:=query.Limit(filter.Limit).Offset(offset).Find(&product).Error
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
