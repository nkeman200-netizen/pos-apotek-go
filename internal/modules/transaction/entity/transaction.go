package entity

import (
	"apotek-pos-go/internal/modules/inventory/entity"
	"apotek-pos-go/internal/modules/users/entitty"
	"time"
)

type Transaction struct {
	Id               uint                 `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	InvoiceNumber    string               `gorm:"type:varchar(255);unique" json:"invoice_number"`
	CustomerId       *uint                 `gorm:"" json:"customer_id"`
	Customer         Customer             `gorm:"foreignKey:CustomerId;constraint:OnUpdate:cascade;OnDelete:restrict" json:"customer,omitempty"`
	UserId           uint                 `gorm:"not null" json:"user_id"`
	User             entitty.User         `gorm:"foreignKey:UserId;constraint:OnUpdate:cascade;OnDelete:restrict" json:"user,omitempty"`
	TotalPrice       int64                `gorm:"not null;type:bigint" json:"total_price"`
	PaymentMethod    string               `gorm:"type:varchar(50);not null" json:"payment_method"`
	PaymentReference string               `gorm:"type:varchar(100)" json:"payment_reference"`
	Pembayaran       int64                `gorm:"type:bigint;not null" json:"pembayaran"`
	Kembalian        int64                `gorm:"type:bigint;not null" json:"kembalian"`
	Status           string               `gorm:"type:varchar(100);not null;default:'complete'" json:"status"`
	VoidReason       string               `gorm:"type:varchar(100)" json:"void_reason"`
	Details          []TransactionDetails `gorm:"foreignKey:TransactionId" json:"details,omitempty"`
	CreatedAt        time.Time            `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt        time.Time            `gorm:"autoUpdateTime" json:"updated_at"`
}

type TransactionDetails struct {
	Id             uint                `gorm:"primaryKey;autoIncrement" json:"id"`
	TransactionId  uint                `gorm:"not null" json:"transaction_id"`
	ProductId      uint                `gorm:"not null" json:"product_id"`
	Product        entity.Product      `gorm:"foreignKey:ProductId;constraint:OnUpdate:cascade;OnDelete:restrict" json:"product,omitempty"`
	ProductBatchId uint                `gorm:"not null" json:"product_batch_id"`
	ProductBatch   entity.ProductBatch `gorm:"foreignKey:ProductBatchId;constraint:OnUpdate:cascade;OnDelete:restrict" json:"product_batch,omitempty"`
	Quantity       int                 `gorm:"not null" json:"quantity"`
	UnitPrice      int64               `gorm:"Type:bigint;not null" json:"unit_price"`
	SubTotal       int64               `gorm:"not null;type:bigint" json:"sub_total"`
}
