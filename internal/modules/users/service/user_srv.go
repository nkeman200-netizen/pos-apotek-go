package service

import (
	"apotek-pos-go/internal/modules/users/entitty"
	"apotek-pos-go/internal/modules/users/repository"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	RegisterUser(user *entitty.User) error
	UpdateUser(user *entitty.User) error
	GetAllUser() ([]entitty.User,error)
	Delete(id uint) error
	Login(username string, password string) (string,error)
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

var JwtSecretKey = []byte("ApotekPOS_SuperSecret_2026")
func(s *userServiceImpl) Login(username string, password string) (string,error){
	var user entitty.User
	user,err:=s.repo.GetUserByUsername(username)
	if err!=nil {
		return "",errors.New("Username tidak ditemukan")
	}

	err=bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(password))
	if err!=nil {
		return "",errors.New("Password salah")
	}

	claim:=jwt.MapClaims{
		"id":user.Id,
		"username":user.Username,
		"role":user.Role,
		"exp":time.Now().Add(time.Hour*24).Unix(),
	}


	token:=jwt.NewWithClaims(jwt.SigningMethodHS256,claim)
	tokenStrg,err:=token.SignedString(JwtSecretKey)
	if err!=nil {
		return "",errors.New("Gagal mencetak token")
	}
	return tokenStrg,nil
}