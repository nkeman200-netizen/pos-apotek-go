package repo

import (
	invenEnti "apotek-pos-go/internal/modules/inventory/entity"
	"apotek-pos-go/internal/modules/purchases/entity"
	"errors"

	"gorm.io/gorm"
)

type PurchaseRepo interface {
	Create(p *entity.Purchase) error
	Void(p *entity.Purchase) error
	GetAll() ([]entity.Purchase, error)
	GetById(id uint) (entity.Purchase, error)
}

type purchaseRepoImp struct {
	db *gorm.DB
}

func NewPurchaseRepo(db *gorm.DB) PurchaseRepo {
	return &purchaseRepoImp{db: db}
}

// Create implements [PurchaseRepo].
func (pr *purchaseRepoImp) Create(p *entity.Purchase) error {
	return pr.db.Transaction(func(tx *gorm.DB) error {
		for i:=range p.Details{
			var pb invenEnti.ProductBatch
			err:=tx.Model(&invenEnti.ProductBatch{}).
			Where("batch_number=?",p.Details[i].ProductBatch.BatchNumber).
			First(&pb).Error
			if err==nil {
				pb.Stock+=p.Details[i].Quantity
				if err:= tx.Save(&pb).Error;err!=nil{
					return err
				}
				p.Details[i].ProductBatchId=pb.ID
			}else{
				if err:=tx.Create(&p.Details[i].ProductBatch).Error;err!=nil{
					return err
				}
				p.Details[i].ProductBatchId=p.Details[i].ProductBatch.ID
			}
			p.Details[i].ProductBatch=invenEnti.ProductBatch{}
		}
		
		err:=tx.Create(p).Error
		if err!=nil {
			return err
		}
		return nil
	})
}

func(pr *purchaseRepoImp) Void(p *entity.Purchase) error{
	return pr.db.Transaction(func(tx *gorm.DB) error {
		for i:=range p.Details{
			var pb invenEnti.ProductBatch
			err:=tx.Model(&invenEnti.ProductBatch{}).
				Where("batch_number=?",p.Details[i].ProductBatch.BatchNumber).
				First(&pb).Error
			if err!=nil {
				return err
			}
			if pb.Stock<p.Details[i].Quantity {
				return errors.New("Tidak bisa melakukan void: stok product "+p.Details[i].Product.Name+" sudah terpakai")
			}
			pb.Stock-=p.Details[i].Quantity
			if err:= tx.Save(&pb).Error;err!=nil{
				return err
			}
		}
		p.Status="void"
		err:=tx.Save(p).Error
		if err!=nil {
			return err
		}
		return nil
	})
}

// GetAll implements [PurchaseRepo].
func (p *purchaseRepoImp) GetAll() ([]entity.Purchase, error) {
	var purchase []entity.Purchase
	err:=p.db.Preload("Supplier").
	Preload("Details").
	Preload("Details.Product").
	Preload("Details.ProductBatch").
	Preload("Details.Product.Category").
	Preload("Details.Product.Unit").
	Find(&purchase).Error
	return purchase,err
}

// GetById implements [PurchaseRepo].
func (p *purchaseRepoImp) GetById(id uint) (entity.Purchase, error) {
	var purchase entity.Purchase
	err:=p.db.Preload("Supplier").
	Preload("User").
	Preload("Details").
	Preload("Details.Product").
	Preload("Details.ProductBatch").
	Preload("Details.Product.Category").
	Preload("Details.Product.Unit").
	First(&purchase,id).Error
	return purchase,err
}

