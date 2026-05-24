package repository

import (
	"apotek-pos-go/internal/modules/users/entitty"
	"errors"

	"gorm.io/gorm"
)

type UserRepo interface {
	Create(user *entitty.User) error
	FindByUsername(username string) error
	Update(u *entitty.User) error
	Delete(id uint) error
	FindAll()([]entitty.User,error)
	FindById(id uint) error
	GetUserByUsername(username string) (entitty.User,error)
}

type userRepoImpl struct{
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo{
	return &userRepoImpl{db: db}
}

func (r *userRepoImpl) Create(user *entitty.User) error{
	return r.db.Save(user).Error
}

func(r *userRepoImpl) FindAll()([]entitty.User,error){
	var users []entitty.User
	err:=r.db.Find(&users).Error
	return users,err
}

func(r *userRepoImpl) Delete(id uint) error{
	return r.db.Delete(&entitty.User{},id).Error
}

func(r *userRepoImpl) FindByUsername(username string) error{
	var user entitty.User
	if err:=r.db.Where("username = ?",username).First(&user).Error; err==nil {
		return errors.New("Username telah digunakan")
	}
	return nil
}

func(r *userRepoImpl) FindById(id uint) error{
	return r.db.First(&entitty.User{},id).Error
}

func(r *userRepoImpl) GetUserByUsername(username string) (entitty.User,error){
	var user entitty.User
	err:=r.db.Where("username=?",username).First(&user).Error
	return user,err
}

func(r *userRepoImpl) Update(u *entitty.User) error{
	return r.db.Updates(u).Error
}

