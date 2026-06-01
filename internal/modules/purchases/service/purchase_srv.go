package service

import (
	"apotek-pos-go/internal/modules/purchases/entity"
	"apotek-pos-go/internal/modules/purchases/repo"
	"apotek-pos-go/internal/modules/users/repository"
	"errors"
)

type PurchaseSrv interface {
	Create(p *entity.Purchase) error
	Void(p *entity.Purchase) error
	GetAll() ([]entity.Purchase, error)
	GetById(id uint) (entity.Purchase, error)
}

type purchaseSrvImp struct {
	repo repo.PurchaseRepo
	userRepo repository.UserRepo
	supRepo repo.SupplierRepo
}

func NewPurchaseRepo(repo repo.PurchaseRepo,supRepo repo.SupplierRepo,userRepo repository.UserRepo) PurchaseSrv {
	return &purchaseSrvImp{repo: repo, supRepo: supRepo,userRepo: userRepo}
}

// Create implements [PurchaseSrv].
func (s *purchaseSrvImp) Create(p *entity.Purchase) error {
	if err:=s.userRepo.FindById(p.UserId);err!=nil {
		return errors.New("Id Username tidak ditemukan")
	}
	if _,err:=s.supRepo.GetById(p.SupplierId);err!=nil {
		return errors.New("Id supplier tidak ditemukan")
	}
	if p.PurchaseDate.IsZero()  {
		return errors.New("Tanggal purchase tidak boleh kosong")
	}
	if p.TotalCost<=0 {
		return errors.New("Total cost harus lebih dari nol")
	}
	return s.repo.Create(p)
}

// Update implements [PurchaseSrv].
func (s *purchaseSrvImp) Void(p *entity.Purchase) error {
	if p.Status=="void" {
		return errors.New("Nota pembelian ini sudah berstatus void")
	}
	if p.VoidReason=="" {
		return errors.New("Void reason tidak boleh kosong")
	}
	return s.repo.Void(p)
}

// GetAll implements [PurchaseSrv].
func (p *purchaseSrvImp) GetAll() ([]entity.Purchase, error) {
	return p.repo.GetAll()
}

// GetById implements [PurchaseSrv].
func (p *purchaseSrvImp) GetById(id uint) (entity.Purchase, error) {
	return p.repo.GetById(id)
}
