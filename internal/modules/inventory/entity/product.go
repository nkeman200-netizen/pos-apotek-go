package entity

import "time"

type Product struct {
	ID           uint      `gorm:"primaryKey;autoIncrement"`
	Name         string    `gorm:"type:varchar(100);not null"`
	SKU          string    `gorm:"type:varchar(50);not null;unique"`
	SellingPrice int64     `gorm:"type:bigint;not null"`
	CategoryID   uint      `gorm:"not null"`
	Category     Category  `gorm:"foreignKey:CategoryID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	UnitID       uint      `gorm:"not null"`
	Unit         Unit      `gorm:"foreignKey:UnitID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	MinStock     int       `gorm:"type:int;not null;default:0"`
	IsActive     bool      `gorm:"type:tinyint(1);not null;default:true"`
	TotalStock int `gorm:"-"`
	CreatedAt    time.Time `gorm:"autoCreateTime;"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`

}