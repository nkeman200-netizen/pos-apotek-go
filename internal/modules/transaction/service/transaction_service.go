package service

import (
	invRepo "apotek-pos-go/internal/modules/inventory/repository"
	"apotek-pos-go/internal/modules/transaction/entity"
	"apotek-pos-go/internal/modules/transaction/repository"
	userRep "apotek-pos-go/internal/modules/users/repository"
	"errors"
)

type TransactionService interface {
	Create(t *entity.Transaction) error
	Void(t *entity.Transaction) error
	FindAll() ([]entity.Transaction, error)
	FindById(id uint) (entity.Transaction, error)
}

type transactionServiceImpl struct {
	repo repository.TransactionRepository
	cRepo repository.CustomerRepo
	uRepo userRep.UserRepo
	batchRepo invRepo.ProductBatchRepository
}

func NewTransactionServiceImpl(
		repo repository.TransactionRepository,
		cRepo repository.CustomerRepo,
		uRepo userRep.UserRepo,
		batchRepo invRepo.ProductBatchRepository,
	) TransactionService {
	return &transactionServiceImpl{
		repo: repo,
		cRepo: cRepo,
		uRepo: uRepo,
		batchRepo: batchRepo,
	}
}

// Create implements [TransactionService].
func (srv *transactionServiceImpl) Create(t *entity.Transaction) error {
	if t.CustomerId!=nil {
		if _,err:=srv.cRepo.FindById(*t.CustomerId);err!=nil {
			return errors.New("Customer id tidak ditemukan")
		}
	}
	if err:=srv.uRepo.FindById(t.UserId);err!=nil {
		return errors.New("User id tidak ditemukan")
	}
	if t.Pembayaran<t.TotalPrice {
		return errors.New("Total pembayaran kurang dari total harga")
	}
	if t.Kembalian<0{
		return errors.New("Kembalian tidak boleh kurang dari nol")
	}
	if t.Pembayaran<0{
		return errors.New("Pembayaran tidak boleh kurang dari nol")
	}
	if t.TotalPrice<0{
		return errors.New("Total harga tidak boleh kurang dari nol")
	}
	if t.PaymentMethod=="" {
		return errors.New("Payment method tidak boleh kosong")
	}
	
	if err:=srv.repo.Create(t);err!=nil {
		return err
	}
	
	return nil
}

// Void implements [TransactionService].
func (srv *transactionServiceImpl) Void(t *entity.Transaction) error {
	if t.Status=="void" {
		return errors.New("Nota pembelian ini sudah berstatus void")
	}
	if t.VoidReason=="" {
		return errors.New("Void reason tidak boleh kosong")
	}
	if err:=srv.repo.Void(t);err!=nil {
		return err
	}
	return nil
}
// FindAll implements [TransactionService].
func (srv *transactionServiceImpl) FindAll() ([]entity.Transaction, error) {
	return srv.repo.FindAll()
}

// FindById implements [TransactionService].
func (srv *transactionServiceImpl) FindById(id uint) (entity.Transaction, error) {
	return srv.repo.FindById(id)
}



