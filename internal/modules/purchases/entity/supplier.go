package entity

import (
	"time"
)

type Supplier struct {
	Id        uint   `gorm:"autoIncrement;primaryKey" json:"id"`
	Nama      string `gorm:"type:varchar(100);not null" json:"nama"`
	Address   string `gorm:"type:text" json:"address"`
	Phone     string `gorm:"type:varchar(20)" json:"phone"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}