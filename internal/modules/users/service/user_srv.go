package service

import (
	"apotek-pos-go/internal/modules/users/entitty"
	"apotek-pos-go/internal/modules/users/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	RegisterUser(user *entitty.User) error
	UpdateUser(user *entitty.User) error
	GetAllUser() ([]entitty.User,error)
	Delete(id uint) error
}

type userServiceImpl struct{
	repo repository.UserRepo
}

func NewUserService(repo repository.UserRepo) UserService{
	return &userServiceImpl{repo: repo}
}

func (s *userServiceImpl) RegisterUser(user *entitty.User) error{
	if err:= s.repo.FindByUsername(user.Username);err!=nil{
		return err
	}

	hashedPass,err:=bcrypt.GenerateFromPassword([]byte(user.Password),bcrypt.DefaultCost)
	if err!=nil {
		return err
	}
	user.Password=string(hashedPass)
	return s.repo.Create(user)
}

func(s *userServiceImpl) UpdateUser(user *entitty.User) error{
	if err:=s.repo.FindById(user.Id); err!=nil{
		return err
	}
	hashedPass,err:=bcrypt.GenerateFromPassword([]byte(user.Password),bcrypt.DefaultCost)
	if err!=nil {
		return err
	}
	user.Password=string(hashedPass)
	return s.repo.Update(user)
}


func(s *userServiceImpl) GetAllUser() ([]entitty.User,error){
	return s.repo.FindAll()
}

func(s *userServiceImpl) Delete(id uint) error{
	return s.repo.Delete(id)
}
