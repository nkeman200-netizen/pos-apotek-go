package repository

import (
	"apotek-pos-go/internal/modules/inventory/entity"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	Create(category *entity.Category) error //keahlian create, apa yg dibutuhkan entity categori, dia mereturn apa = error
	Update(category *entity.Category) error //update butuh data categoru apa yang diapdate, ngapa ga butuh id ya? kalo di laravel kan butuh 
	Delete(id uint) error //delete butuh id, terus return eror
	FindAll() ([]entity.Category, error) //mereturn kumpulan kategori dan error
	FindById(id uint) (entity.Category,error) // find by id nyarinya butuh id unsign int, terus mereturn 1 entity categoti dan eror

}

type categoryRepositoryImpl struct{ 
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository{
	return &categoryRepositoryImpl{db:db}
}

func (r *categoryRepositoryImpl) Create(category *entity.Category) error {
	err := r.db.Create(category).Error
	return err
}

func (r *categoryRepositoryImpl) FindAll() ([]entity.Category, error){
	var category []entity.Category
	err := r.db.Find(&category).Error
	return  category,err
}

func ( r *categoryRepositoryImpl) FindById(id uint) (entity.Category,error){
	var categori entity.Category 
	err:=r.db.First(&categori,id).Error //hasilnya temuan taro di alamat categori, id categori yang mana yang id di parameter ke dua
	return  categori,err
}

func (r *categoryRepositoryImpl) Delete(id uint) error{ //unsigne integer
	err :=r.db.Delete(&entity.Category{},id).Error
	return err
}

func (r *categoryRepositoryImpl) Update(category *entity.Category) error{
	err:=r.db.Model(&entity.Category{}).Where("id=?",category.Id).Updates(category).Error
	return  err
}