package handler

import (
	"apotek-pos-go/internal/modules/purchases/entity"
	"apotek-pos-go/internal/modules/purchases/service"
	"apotek-pos-go/internal/pkg/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SupplierHandler struct {
	srv service.SupplierSrv
}

func NewSupplierHandler(srv service.SupplierSrv) *SupplierHandler{
	return &SupplierHandler{srv: srv}
}

func (h *SupplierHandler) Create(c *gin.Context){
	var Supplier entity.Supplier
	if err:=c.ShouldBindJSON(&Supplier);err!=nil {
		c.JSON(http.StatusBadRequest,response.Response[any]{
			Code: 400,
			Message: "Format data salah: " + err.Error(),
			Data: nil,
		})
		return 
	}
	if err:=h.srv.Create(&Supplier);err!=nil {
		c.JSON(http.StatusInternalServerError,response.Response[any]{
			Code: 500,
			Message: "Server sedang bermasalah: "+err.Error(),
			Data: nil,
		})
		return 
	}
	c.JSON(http.StatusCreated,response.Response[entity.Supplier]{
		Code: 201,
		Message: "Berhasil membuat Supplier baru",
		Data: Supplier,
	})
}

func (h *SupplierHandler) Update(c *gin.Context){
	var Supplier entity.Supplier
	idStr:=c.Param("id")
	idInt,err:=strconv.Atoi(idStr)
	if err!=nil {
		c.JSON(http.StatusBadRequest,response.Response[any]{
			Code: 400,
			Message: "Parameter id harus berupa angka",
			Data: nil,
		})
		return 
	}
	if err:=c.ShouldBindJSON(&Supplier);err!=nil {
		c.JSON(http.StatusBadRequest,response.Response[any]{
			Code: 400,
			Message: "Format data salah: " + err.Error(),
			Data: nil,
		})
		return 
	}
	Supplier.Id=uint(idInt)
	if err:=h.srv.Update(&Supplier);err!=nil {
		c.JSON(http.StatusInternalServerError,response.Response[any]{
			Code: 500,
			Message: "Server sedang bermasalah: "+err.Error(),
			Data: nil,
		})
		return 
	}

	c.JSON(http.StatusOK,response.Response[entity.Supplier]{
		Code: 200,
		Message: "Berhasil memperbarui data Supplier",
		Data: Supplier,
	})
}

func (h *SupplierHandler) Delete(c *gin.Context){
	idStr:=c.Param("id")
	idInt,err:=strconv.Atoi(idStr)
	if err!=nil {
		c.JSON(http.StatusBadRequest,response.Response[any]{
			Code: 400,
			Message: "Parameter id harus berupa angka",
			Data: nil,
		})
		return 
	}
	if err:=h.srv.Delete(uint(idInt));err!=nil {
		c.JSON(http.StatusInternalServerError,response.Response[any]{
			Code: 500,
			Message: "Server sedang bermasalah: "+err.Error(),
			Data: nil,
		})
		return 
	}

	c.JSON(http.StatusOK,response.Response[any]{
		Code: 204,
		Message: "Berhasil menghapus data Supplier",
		Data: nil,
	})
}

func ( h *SupplierHandler) GetAll(c *gin.Context){
	var Supplier []entity.Supplier
	Supplier,err:=h.srv.GetAll();
	if err!=nil {
		c.JSON(http.StatusInternalServerError,response.Response[any]{
			Code: 500,
			Message: "Server sedang bermasalah: "+err.Error(),
			Data: nil,
		})
		return 
	}
	c.JSON(http.StatusOK,response.Response[[]entity.Supplier]{
		Code: 200,
		Message: "Berhasil mengambil semua data Supplier",
		Data: Supplier,
	})
}

func (h *SupplierHandler) GetSupplierById(c *gin.Context){
	var Supplier entity.Supplier
	idStr:=c.Param("id")
	idInt,err:=strconv.Atoi(idStr)
	if err!=nil {
		c.JSON(http.StatusBadRequest,response.Response[any]{
			Code: 400,
			Message: "Parameter id harus berupa angka",
			Data: nil,
		})
		return 
	}
	Supplier,err= h.srv.GetById(uint(idInt))
	if err!=nil{
		c.JSON(http.StatusInternalServerError,response.Response[any]{
			Code: 500,
			Message: "Server sedang bermasalah: "+err.Error(),
			Data: nil,
		})
		return 
	}	
	c.JSON(http.StatusOK,response.Response[entity.Supplier]{
		Code: 200,
		Message: "Berhasil mengambil data Supplier",
		Data: Supplier,
	})
}