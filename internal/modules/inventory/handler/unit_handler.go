package handler

import (
	"apotek-pos-go/internal/modules/inventory/entity"
	"apotek-pos-go/internal/modules/inventory/service"
	"apotek-pos-go/internal/pkg/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UnitHandler struct {
	srv service.UnitService
}

func NewUnitHandler(srv service.UnitService) *UnitHandler{
	return &UnitHandler{srv: srv}
}

func (h *UnitHandler) Create(c *gin.Context){
	var unit entity.Unit
	if err:=c.ShouldBindJSON(&unit);err!=nil {
		c.JSON(http.StatusBadRequest,response.Response[any]{
			Code: 400,
			Message: "Format data salah: " + err.Error(),
			Data: nil,
		})
		return 
	}
	if err:=h.srv.CreateUnit(&unit);err!=nil {
		c.JSON(http.StatusInternalServerError,response.Response[any]{
			Code: 500,
			Message: "Server sedang bermasalah: "+err.Error(),
			Data: nil,
		})
		return 
	}
	c.JSON(http.StatusCreated,response.Response[entity.Unit]{
		Code: 201,
		Message: "Berhasil membuat unit baru",
		Data: unit,
	})
}

func (h *UnitHandler) Update(c *gin.Context){
	var unit entity.Unit
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
	if err:=c.ShouldBindJSON(&unit);err!=nil {
		c.JSON(http.StatusBadRequest,response.Response[any]{
			Code: 400,
			Message: "Format data salah: " + err.Error(),
			Data: nil,
		})
		return 
	}
	unit.ID=uint(idInt)
	if err:=h.srv.UpdateUnit(&unit);err!=nil {
		c.JSON(http.StatusInternalServerError,response.Response[any]{
			Code: 500,
			Message: "Server sedang bermasalah: "+err.Error(),
			Data: nil,
		})
		return 
	}

	c.JSON(http.StatusOK,response.Response[entity.Unit]{
		Code: 200,
		Message: "Berhasil memperbarui data unit",
		Data: unit,
	})
}

func (h *UnitHandler) Delete(c *gin.Context){
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
	if err:=h.srv.DeleteUnit(uint(idInt));err!=nil {
		c.JSON(http.StatusInternalServerError,response.Response[any]{
			Code: 500,
			Message: "Server sedang bermasalah: "+err.Error(),
			Data: nil,
		})
		return 
	}

	c.JSON(http.StatusOK,response.Response[any]{
		Code: 204,
		Message: "Berhasil menghapus data unit",
		Data: nil,
	})
}

func ( h *UnitHandler) GetAll(c *gin.Context){
	var unit []entity.Unit
	unit,err:=h.srv.GetAllUnit();
	if err!=nil {
		c.JSON(http.StatusInternalServerError,response.Response[any]{
			Code: 500,
			Message: "Server sedang bermasalah: "+err.Error(),
			Data: nil,
		})
		return 
	}
	c.JSON(http.StatusOK,response.Response[[]entity.Unit]{
		Code: 200,
		Message: "Berhasil mengambil semua data unit",
		Data: unit,
	})
}

func (h *UnitHandler) GetUnitById(c *gin.Context){
	var unit entity.Unit
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
	unit,err= h.srv.GetUnitById(uint(idInt))
	if err!=nil{
		c.JSON(http.StatusInternalServerError,response.Response[any]{
			Code: 500,
			Message: "Server sedang bermasalah: "+err.Error(),
			Data: nil,
		})
		return 
	}	
	c.JSON(http.StatusOK,response.Response[entity.Unit]{
		Code: 200,
		Message: "Berhasil mengambil data unit",
		Data: unit,
	})
}