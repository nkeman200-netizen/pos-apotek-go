package entity

import "time"

type Unit struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	Name      string    `gorm:"type:varchar(50);not null"`
	ShortName string    `gorm:"type:varchar(20);not null"`
	CreatedAt time.Time `gorm:"autoCreateTime;"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}