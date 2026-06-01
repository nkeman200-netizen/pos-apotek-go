package service

import (
	"apotek-pos-go/internal/modules/inventory/entity"
	"apotek-pos-go/internal/modules/inventory/repository"
	"errors"
	"time"
)

type ProductBatchService interface {
	CreateProductBatch(pb *entity.ProductBatch) error
	UpdateProductBatch(pb *entity.ProductBatch) error
	DeleteProductBatch(id uint) error
	GetAllProductBatch() ([]entity.ProductBatch,error)
	GetProductBatchById(id uint) (entity.ProductBatch,error)
	GetProductBatchByProductId(id uint) ([]entity.ProductBatch,error)
}

type productBatchServiceImpl struct{
	repo repository.ProductBatchRepository
	productRepo repository.ProductRepository
}

func NewProductBatchService(
	repo repository.ProductBatchRepository,
	productRepo repository.ProductRepository,
) ProductBatchService{
	return &productBatchServiceImpl{repo: repo, productRepo:productRepo}
}

func (s *productBatchServiceImpl) CreateProductBatch(pb *entity.ProductBatch) error{
	if pb.BatchNumber=="" {
		return errors.New("GAGAL: Batch number tidak boleh kosong")
	}
	if pb.ExpiredDate.Before(time.Now()) {
		return errors.New("GAGAL: Tanggal ED telah terlampui")
	}
	if pb.PurchasePrice<0 {
		return errors.New("GAGAL: Harga beli tidak boleh kurang dari nol")
	}
	if pb.Stock<0 {
		return errors.New("GAGAL: Stok tidak boleh kurang dari nol")
	}

	if pb.ProductID==0 {
		return errors.New("GAGAL: Product id harus diisi")
	}
	if _,err:=s.productRepo.FindById(pb.ProductID);err==nil	 { //artinya gada eror, ya berarti berhasil ditemukan
		return errors.New("GAGAL: Id product tidak ditemukan di database")
	}

	return s.repo.Create(pb)
}

func (s *productBatchServiceImpl) UpdateProductBatch(pb *entity.ProductBatch) error{
	if pb.BatchNumber=="" {
		return errors.New("GAGAL: Batch number tidak boleh kosong")
	}
	if pb.ExpiredDate.Before(time.Now()) {
		return errors.New("GAGAL: Tanggal ED telah terlampui")
	}
	if pb.PurchasePrice<0 {
		return errors.New("GAGAL: Harga beli tidak boleh kurang dari nol")
	}
	if pb.Stock<0 {
		return errors.New("GAGAL: Stok tidak boleh kurang dari nol")
	}

	if pb.ProductID==0 {
		return errors.New("GAGAL: Product id harus diisi")
	}
	if _,err:=s.productRepo.FindById(pb.ProductID);err!=nil	 {
		return errors.New("GAGAL: Id product tidak ditemukan di database")
	}

	return s.repo.Update(pb)
}

func (s *productBatchServiceImpl) DeleteProductBatch(id uint) error{
	return s.repo.Delete(id)
}

func(s *productBatchServiceImpl) GetAllProductBatch() ([]entity.ProductBatch,error){
	var pb []entity.ProductBatch
	pb,err:=s.repo.FindAll()
	return pb,err
}

func(s *productBatchServiceImpl) GetProductBatchByProductId(id uint) ([]entity.ProductBatch,error){
	var pb []entity.ProductBatch
	pb,err:=s.repo.FindByProductId(id)
	return pb,err
}

func(s *productBatchServiceImpl) GetProductBatchById(id uint) (entity.ProductBatch,error){
	var pb entity.ProductBatch
	pb,err:=s.repo.FindById(id)
	return pb,err
}
