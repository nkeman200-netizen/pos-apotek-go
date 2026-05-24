package handler

import (
	"apotek-pos-go/internal/modules/users/entitty"
	"apotek-pos-go/internal/modules/users/service"
	"apotek-pos-go/internal/pkg/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	srv service.UserService
}

func NewUserHandler(srv service.UserService) *UserHandler{
	return &UserHandler{srv: srv}
}

func (h *UserHandler) Register(c *gin.Context){
	var user entitty.User
	if err:=c.ShouldBindJSON(&user);err!=nil {
		c.JSON(http.StatusBadRequest,response.Response[any]{
			Code: 400,
			Message: "Format data salah: "+err.Error(),
			Data: nil,
		})
		return
	}

	if err:=h.srv.RegisterUser(&user);err!=nil {
		c.JSON(http.StatusInternalServerError, response.Response[any]{
			Code: 500,
			Message: "Server sedang bermasalah: "+err.Error(),
			Data: nil,
		})
		return
	}

	c.JSON(http.StatusCreated,response.Response[entitty.User]{
		Code: 201,
		Message: "Berhasil membuat akun",
		Data: user,
	})
}

func (h *UserHandler) UpdateUser(c *gin.Context){
	var user entitty.User
	idS:=c.Param("id")
	idI,err:=strconv.Atoi(idS)
	if err!=nil {
		c.JSON(http.StatusBadRequest,response.Response[any]{
			Code: 400,
			Message: "Masukan hanya angka user id",
			Data: nil,
		})	
		return
	}
	
	if err:=c.ShouldBindJSON(&user);err!=nil {
		c.JSON(http.StatusBadRequest,response.Response[any]{
			Code: 400,
			Message: "Format data salah: "+err.Error(),
			Data: nil,
		})
		return
	}

	user.Id=uint(idI)

	if err:=h.srv.UpdateUser(&user);err!=nil {
		c.JSON(http.StatusInternalServerError, response.Response[any]{
			Code: 500,
			Message: "Server sedang bermasalah: "+err.Error(),
			Data: nil,
		})
		return
	}

	c.JSON(http.StatusOK,response.Response[entitty.User]{
		Code: 200,
		Message: "Berhasil memperbarui akun dengan id "+idS,
		Data: user,
	})
}

func (h *UserHandler) GetAllUser(c *gin.Context){
	var users []entitty.User
	users,err:=h.srv.GetAllUser()
	if err!=nil {
		c.JSON(http.StatusInternalServerError,response.Response[any]{
			Code: 500,
			Message: "Terjadi kesalahan pada server: "+err.Error(),
			Data: nil,
		})
		return 
	}

	c.JSON(http.StatusOK,response.Response[[]entitty.User]{
		Code: 200,
		Message: "Berhasil mengambil semua data user",
		Data: users,
	})
}

func(h *UserHandler) Delete(c *gin.Context){
	idS:=c.Param("id")
	idI,err:=strconv.Atoi(idS)
	if err!=nil {
		c.JSON(http.StatusBadRequest,response.Response[any]{
			Code: 400,
			Message: "Masukan hanya angka user id",
			Data: nil,
		})	
		return
	}
	
	if err:=h.srv.Delete(uint(idI));err!=nil {
		c.JSON(http.StatusInternalServerError,response.Response[any]{
			Code: 500,
			Message: "Terdapat masalah pada server: "+err.Error(),
			Data: nil,
		})
		return 
	}

	c.JSON(http.StatusOK,response.Response[any]{
		Code: 204,
		Message: "Berhasil menghapus user id "+idS,
		Data: nil,
	})

}

