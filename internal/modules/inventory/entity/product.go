package entity

import "time"

type Product struct {
	ID           uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Name         string    `gorm:"type:varchar(100);not null" json:"name"`
	SKU          string    `gorm:"type:varchar(50);not null;unique" json:"sku"`
	SellingPrice int64     `gorm:"type:bigint;not null" json:"selling_price"`
	CategoryID   uint      `gorm:"not null" json:"category_id"`
	Category     Category  `gorm:"foreignKey:CategoryID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;" json:"category"`
	UnitID       uint      `gorm:"not null" json:"unit_id"`
	Unit         Unit      `gorm:"foreignKey:UnitID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT" json:"unit"`
	MinStock     int       `gorm:"type:int;not null;default:0" json:"min_stok"`
	IsActive     bool      `gorm:"type:tinyint(1);not null;default:true" json:"is_active"`
	TotalStock int `gorm:"-" json:"total_stok"`
	CreatedAt    time.Time `gorm:"autoCreateTime;" json:"created_at"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime" json:"updated_at"`

}