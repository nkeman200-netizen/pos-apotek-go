package entity

import "time"

type ProductBatch struct {
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	BatchNumber string `gorm:"type:varchar(100);not null;"`
	ExpiredDate time.Time `gorm:"type:date;not null"`
	PurchasePrice int64 `gorm:"type:bigint;not null"`
	Stock int `gorm:"type:int;not null"`
	ProductID uint `gorm:"not null"`
	Product Product `gorm:"foreignKey:ProductID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	CreatedAt   time.Time `gorm:"autoCreateTime;"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}	