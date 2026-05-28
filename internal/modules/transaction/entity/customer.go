package entity

import "time"

type Customer struct {
	Id        uint      `gorm:"primaryKey;autoIncrement"`
	Nama      string    `gorm:"type:varchar(200);not null"`
	Phone     string    `gorm:"type:varchar(50);"`
	Address   string    `gorm:"type:text"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
