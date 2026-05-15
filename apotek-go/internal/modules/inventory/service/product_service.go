package service

import (
	"apotek-pos-go/internal/modules/inventory/entity"
	"apotek-pos-go/internal/modules/inventory/repository"
	"errors"
)

type ProductService interface{
	CreateProduct(product *entity.Product) error
	UpdateProduct(product *entity.Product) error
	DeleteProduct(id uint) error
	GetAllProduct() ([]entity.Product,error)
	GetProductById(id uint) (entity.Product,error)
}

type productServiceImpl struct {
	repo repository.ProductRepository
	categoryRepo repository.CategoryRepository
	unitRepo repository.UnitRepository
}

func NewProductService(
	repo repository.ProductRepository,
	categoryRepo repository.CategoryRepository,
	unitRepo repository.UnitRepository,
) ProductService{
	return &productServiceImpl{
		repo: repo,categoryRepo: categoryRepo,unitRepo: unitRepo,
	}	
}	

func (s *productServiceImpl) CreateProduct(p *entity.Product) error{
	//nama dan harga
	if p.Name=="" {
		return errors.New("Nama tidak boleh kosong")
	}
	if  p.SellingPrice<0 {
		return errors.New("Harga tidak boleh negatif")
	}
	//sku
	if p.SKU=="" {
		return errors.New("GAGAL: SKU tidak boleh kosong")
	}
	if _,err:=s.repo.FindBySKU(p.SKU); err==nil{
		return errors.New("GAGAL: SKU sudah terdaftar, SKU bersifat unik")
	}
	//min stok
	if p.MinStock<0 {
		return errors.New("GAGAL: Min stock tidak boleh kurang dari 0")
	}
	//categori dan unit
	if _,err:=s.categoryRepo.FindById(p.CategoryID);err!=nil {
		return errors.New("Category id tidak ditemukan")
	}
	if _,err:=s.unitRepo.FindById(p.UnitID);err!=nil {
		return errors.New("Unit id tidak ditemukan")	
	}
	return s.repo.Create(p)
}

func (s *productServiceImpl) UpdateProduct(p *entity.Product) error{
	if _,err:=s.repo.FindById(p.ID);err!=nil {
		return errors.New("GAGAL: Product yang anda update tidak ditemukan")
	}
	
	if p.Name=="" {
		return errors.New("Nama tidak boleh kosong")
	}
	if  p.SellingPrice<0 {
		return errors.New("Harga tidak boleh negatif")
	}

	//sku
	if p.SKU=="" {
		return errors.New("GAGAL: SKU tidak boleh kosong")
	}
	if _,err:=s.repo.FindBySKU(p.SKU); err==nil{
		return errors.New("GAGAL: SKU sudah terdaftar, SKU bersifat unik")
	}
	//min stok
	if p.MinStock<0 {
		return errors.New("GAGAL: Min stock tidak boleh kurang dari 0")
	}

	if _,err:=s.categoryRepo.FindById(p.CategoryID);err!=nil {
		return errors.New("Category id tidak ditemukan")
	}
	if _,err:=s.unitRepo.FindById(p.UnitID);err!=nil {
		return errors.New("Unit id tidak ditemukan")	
	}
	return s.repo.Update(p)
}

func (s *productServiceImpl) DeleteProduct(id uint) error{
	if _,err:=s.repo.FindById(id);err!=nil {
		return errors.New("GAGAL: ID product tidak ditemukan")
	}
	return s.repo.Delete(id)
}

func (s *productServiceImpl) GetAllProduct() ([]entity.Product,error){
	var product []entity.Product
	product,err:=s.repo.FindAll()
	return product,err
}

func (s *productServiceImpl) GetProductById(id uint) (entity.Product,error){
	var product entity.Product
	product,err:=s.repo.FindById(id)
	return product,err
}