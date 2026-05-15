package service

import (
	"apotek-pos-go/internal/modules/inventory/entity"
	"apotek-pos-go/internal/modules/inventory/repository"
	"errors"
)

type CategoryService interface {
	CreateCategory(category *entity.Category) error
	GetAllCategory() ([]entity.Category,error)
	GetCategoryById(id uint) (entity.Category,error)
	DeleteCategory(id uint) error
	UpdateCategory(category *entity.Category) error
}

type CategoryServiceImpl struct{
	repo repository.CategoryRepository
}

func NewCategoryService(repo repository.CategoryRepository) CategoryService{
	return  &CategoryServiceImpl{repo: repo} 
}

func (r *CategoryServiceImpl) CreateCategory(category *entity.Category) error{
	if category.Name=="" {
		return errors.New("GAGAL: Nama category tidak boleh kosong")
	}
	if len(category.Name)>50 {
		return errors.New("GAGAL: Nama category tidak boleh lebih dari 50 huruf")
	}

	err:=r.repo.Create(category)
	return err;
}

func (r *CategoryServiceImpl) GetAllCategory() ([]entity.Category,error){
	return r.repo.FindAll()
}

func (r *CategoryServiceImpl) GetCategoryById(id uint) (entity.Category,error){
	return r.repo.FindById(id)
}

func (r *CategoryServiceImpl) DeleteCategory(id uint) error{
	return r.repo.Delete(id)
}

func (r *CategoryServiceImpl) UpdateCategory(category *entity.Category) error{
	if  category.Name=="" {
		return errors.New("GAGAL: Nama category tidak boleh kosong")
	}
	if len(category.Name)>50 {
		return errors.New("GAGAL: Nama category tidak boleh lebih dari 50 huruf")
	}

	return r.repo.Update(category)
}