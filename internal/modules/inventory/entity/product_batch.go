package entity

import "time"

type ProductBatch struct {
	ID          uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	BatchNumber string `gorm:"type:varchar(100);not null;" json:"batch_number"`
	ExpiredDate time.Time `gorm:"type:date;not null" json:"expired_date"`
	PurchasePrice int64 `gorm:"type:bigint;not null" json:"purchase_price"`
	Stock int `gorm:"type:int;not null" json:"stock"`
	ProductID uint `gorm:"not null" json:"product_id"`
	Product Product `gorm:"foreignKey:ProductID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT" json:"product"`
	CreatedAt   time.Time `gorm:"autoCreateTime;" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}	

type ProductBatchFilter struct{
	
}