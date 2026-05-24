package entitty

import "time"

type User struct {
	Id        uint   `gorm:"autoIncrement;primaryKey" json:"id"`
	Username  string `gorm:"type:varchar(50);unique;not null" json:"username"`
	Password  string `gorm:"type:varchar(255);not null" json:"-"`
	Role      string `gorm:"type:varchar(50);not null" json:"role"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}