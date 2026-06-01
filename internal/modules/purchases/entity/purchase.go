package entity

import (
	"apotek-pos-go/internal/modules/inventory/entity"
	"apotek-pos-go/internal/modules/users/entitty"
	"time"
)

type Purchase struct {
	Id         uint     `gorm:"autoIncrement;primaryKey" json:"id"`
	SupplierId uint     `json:"supplier_id"`
	Supplier   Supplier `gorm:"foreignKey:SupplierId;constraint:OnUpdate:cascade;OnDelete:set null" json:"supplier"`
	UserId     uint     `gorm:"not null" json:"user_id"`
	User       entitty.User `gorm:"foreignKey:UserId;constraint:OnUpdate:cascade;OnDelete:restrict" json:"user"`
	TotalCost uint64 `gorm:"type:bigint;not null" json:"total_cost"`
	PurchaseDate time.Time `gorm:"type:date" json:"purchase_date"`
	Status           string               `gorm:"type:varchar(100);not null;default:'complete'" json:"status"`
	VoidReason       string               `gorm:"type:varchar(100)" json:"void_reason"`
	Details []PurchaseDetails `gorm:"foreignKey:PurchaseId" json:"details"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type PurchaseDetails struct{
	Id uint `gorm:"autoIncrement;primaryKey" json:"id"`
	ProductId uint `gorm:"not null" json:"product_id"`
	Product entity.Product `gorm:"foreignKey:ProductId;constraint:OnUpdate:cascade;OnDelete:restrict" json:"product"`
	PurchaseId uint `gorm:"not null" json:"purchase_id"`
	Purchase Purchase `gorm:"foreignKey:PurchaseId;constraint:OnUpdate:cascade;OnDelete:restrict" json:"-"`
	ProductBatchId uint `gorm:"not null" json:"product_batch_id"`
	ProductBatch entity.ProductBatch `gorm:"foreignKey:ProductBatchId;constraint:OnUpdate:cascade;OnDelete:restrict" json:"product_batch"`
	Quantity int `gorm:"not null" json:"quantity"`
	PurchasePrice int64 `gorm:"type:bigint;not null" json:"purchase_price"`
	SubTotal int64 `gorm:"type:bigint;not null" json:"sub_total"`
	Diskon *int `gorm:"" json:"diskon"`
}