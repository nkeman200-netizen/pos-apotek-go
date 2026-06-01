package repository

import (
	inventoryEntity "apotek-pos-go/internal/modules/inventory/entity"
	"apotek-pos-go/internal/modules/transaction/entity"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	Create(t *entity.Transaction) error
	Void(t *entity.Transaction) error
	FindAll() ([]entity.Transaction,error)
	FindById(id uint) (entity.Transaction,error)
}

type transactionRepositoryImpl struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepositoryImpl{db: db}
}

func (r *transactionRepositoryImpl) Create(t *entity.Transaction) error {
	
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(t).Error; err != nil {
			return err
		}

		for i := range t.Details {
			err:=tx.Model(&inventoryEntity.ProductBatch{}).
			Where("id =?",t.Details[i].ProductBatchId).
			Update("stock",gorm.Expr("stock -?",t.Details[i].Quantity)).Error

			if err!=nil {
				return err
			}
		}
		return nil
	})
}

func(r *transactionRepositoryImpl) Void(t *entity.Transaction) error{
	return r.db.Transaction(func(tx *gorm.DB) error {
		t.Status="void"
		if err:=tx.Save(t).Error;err!=nil { //status dan pesn void diatur di service
			return err
		}
		for i := range t.Details {
			err:=tx.Model(&inventoryEntity.ProductBatch{}).
				Where("id=?",t.Details[i].ProductBatchId).
				Update("stock",gorm.Expr("stock +?",t.Details[i].Quantity)).Error
			if err!=nil {
				return err
			}
		}
		return nil
	})
}

func(r *transactionRepositoryImpl) FindAll() ([]entity.Transaction,error){
	var trns []entity.Transaction
	err:=r.db.Preload("Details").
	Preload("Details.Product").
	Preload("User").
	Preload("Customer").
	Find(&trns).Error
	return trns,err
}

func(r *transactionRepositoryImpl) FindById(id uint) (entity.Transaction,error){
	var trns entity.Transaction
	err:=r.db.Preload("Details").
	Preload("Details.Product").
	Preload("User").
	Preload("Customer").
	First(&trns,id).Error
	return trns,err
}
