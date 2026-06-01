package handler

import (
	"apotek-pos-go/internal/modules/transaction/entity"
	"apotek-pos-go/internal/modules/transaction/service"
	"apotek-pos-go/internal/pkg/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CustomerHndler struct {
	srv service.CustomerSrv
}

func NewCustomerHndler(srv service.CustomerSrv) *CustomerHndler{
	return &CustomerHndler{srv: srv}
}

func (h *CustomerHndler) Create(c *gin.Context){
	var customer entity.Customer
	if err:=c.ShouldBindJSON(&customer); err!=nil {
		c.JSON(http.StatusBadRequest, response.Response[any]{
			Code: 400,
			Message: "Format data salah: " + err.Error(),
			Data: nil,
		})
		return
	}

	if err:=h.srv.Create(&customer); err!=nil {
		c.JSON(http.StatusInternalServerError, response.Response[any]{
			Code: 500,
			Message: "Gagal membuat customer baru",
			Data: nil,
		})
		return 
	}

	c.JSON(http.StatusCreated,response.Response[entity.Customer]{
		Code: 201,
		Message: "Berhasil menambahkan customer",
		Data: customer,
	})
}

func (h *CustomerHndler) GetAll(c *gin.Context) {
	var customer []entity.Customer
	customer,err:=h.srv.GetAll()
	if err!=nil{
		c.JSON(http.StatusInternalServerError,response.Response[any]{
			Code: 500,
			Message: "Server sedang bermasalah: "+err.Error(),
			Data: nil,
		})
		return
	}

	c.JSON(http.StatusOK,response.Response[[]entity.Customer]{
		Code: 200,
		Message: "Berhasil mengambil semua data customer",
		Data: customer,
	})
}

func (h *CustomerHndler) GetById(c *gin.Context){
	idStr:=c.Param("id")
	idInt,err:=strconv.Atoi(idStr)
	if err!=nil {
		c.JSON(http.StatusBadRequest,response.Response[any]{
			Code: 400,
			Message: "Masukan hanya angka customer id",
			Data: nil,
		})	
		return
	}

	customer, err:= h.srv.GetUnitById(uint(idInt))
	if err!=nil {
		c.JSON(http.StatusNotFound,response.Response[any]{
			Code: 404,
			Message: "customer tidak ditemukan",
			Data: nil,
		})
		return
	}

	c.JSON(http.StatusOK,response.Response[entity.Customer]{
		Code: 200,
		Message: "get customer by id berhasil",
		Data: customer,
	})
}

func (h *CustomerHndler) Update(c *gin.Context){
	var customer entity.Customer
	idStr:=c.Param("id")
	idInt,err:=strconv.Atoi(idStr)
	if err!=nil {
		c.JSON(http.StatusBadRequest,response.Response[any]{
			Code: 400,
			Message: "Masukan id customer",
			Data: nil,
		})
		return
	}
	if err:=c.ShouldBindJSON(&customer);err!=nil {
		c.JSON(http.StatusBadRequest,response.Response[any]{
			Code: 400,
			Message: "Format data salah: " + err.Error(),
			Data: nil,
		})
		return
	}
	customer.Id=uint(idInt)

	err= h.srv.Update(&customer)
	if err!=nil {
		c.JSON(http.StatusInternalServerError,response.Response[any]{
			Code: 500,
			Message: "Server sedang bermasalah: "+err.Error(),
			Data: nil,
		})
		return
	}
	c.JSON(http.StatusOK,response.Response[entity.Customer]{
		Code: 200,
		Message: "Berhasil memperbarui data customer",
		Data: customer,
	})
}

func (h *CustomerHndler) Delete(c *gin.Context){
	idStr:=c.Param("id")
	idInt,err:=strconv.Atoi(idStr)
	if err!=nil {
		c.JSON(http.StatusBadRequest,response.Response[any]{
			Code: 400,
			Message: "Format data salah: " + err.Error(),
			Data: nil,
		})
		return
	}
	err=h.srv.Delete(uint(idInt))
	if err!=nil {
		c.JSON(http.StatusInternalServerError,response.Response[any]{
			Code: 500,
			Message: "Server sedang bermasalah: "+err.Error(),
			Data: nil,
		})
		return
	}
	c.JSON(http.StatusOK,response.Response[any]{
		Code: 204,
		Message: "Data customer berhasil dihapus",
		Data: nil,
	})
}
