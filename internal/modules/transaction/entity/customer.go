package entity

import "time"

type Customer struct {
	Id        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Nama      string    `gorm:"type:varchar(200);not null" json:"nama"`
	Phone     string    `gorm:"type:varchar(50);" json:"phone"`
	Address   string    `gorm:"type:text" json:"address"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
