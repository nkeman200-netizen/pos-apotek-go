package config

import (
	"apotek-pos-go/internal/modules/inventory/entity"
	"apotek-pos-go/internal/modules/users/entitty"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB(){
	dsn := "root:@tcp(127.0.0.1:3306)/apotek_go_db?charset=utf8mb4&parseTime=True&loc=Local"

	database, err := gorm.Open(mysql.Open(dsn),&gorm.Config{}) //berikan lokasi gorm config{} bukan malah cetakannya doang
	//go bisa menyimpan dua value
	if err != nil { //nil == null
		log.Fatal("Eror konek database: ", err)
	}

	DB=database
	log.Print("Database berhasil dikoneksikan")

	err=DB.AutoMigrate(
		&entity.Category{},
		&entity.Product{},
		&entity.Unit{},
		&entity.ProductBatch{},
		&entitty.User{},
	)
	if err!=nil {
		log.Fatal("Gagal migrate: ", err)
	}
	log.Print("Migration berhasil!")
}