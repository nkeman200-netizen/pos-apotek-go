package handler

import (
	"apotek-pos-go/internal/modules/inventory/entity"
	"apotek-pos-go/internal/modules/inventory/service"
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
		c.JSON(http.StatusBadRequest,gin.H{"error":"Request salah format"})
		return 
	}
	if err:=h.srv.CreateUnit(&unit);err!=nil {
		c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
		return 
	}
	c.JSON(http.StatusOK,gin.H{
		"message": "Berhasil membuat unit baru",
		"data":unit,
	})
}

func (h *UnitHandler) Update(c *gin.Context){
	var unit entity.Unit
	idStr:=c.Param("id")
	idInt,err:=strconv.Atoi(idStr)
	if err!=nil {
		c.JSON(http.StatusBadRequest,gin.H{"error":"Parameter id harus berupa angka"})
		return 
	}
	if err:=c.ShouldBindJSON(&unit);err!=nil {
		c.JSON(http.StatusBadRequest,gin.H{"error":"Format data salah"})
		return 
	}
	unit.ID=uint(idInt)
	if err:=h.srv.UpdateUnit(&unit);err!=nil {
		c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
		return 
	}

	c.JSON(http.StatusOK,gin.H{
		"message":"Berhasil update data unit",
		"data":unit,
	})
}

func (h *UnitHandler) Delete(c *gin.Context){
	idStr:=c.Param("id")
	idInt,err:=strconv.Atoi(idStr)
	if err!=nil {
		c.JSON(http.StatusBadRequest,gin.H{"error":"Param id harus berupa angka"})
		return 
	}
	if err:=h.srv.DeleteUnit(uint(idInt));err!=nil {
		c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
		return 
	}

	c.JSON(http.StatusOK,gin.H{
		"message":"Berhasil menghapus data unit",
	})
}

func ( h *UnitHandler) GetAll(c *gin.Context){
	var unit []entity.Unit
	unit,err:=h.srv.GetAllUnit();
	if err!=nil {
		c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
		return 
	}
	c.JSON(http.StatusOK,gin.H{
		"message":"Berhasil mengambil semua data unit",
		"data":unit,
	})
}

func (h *UnitHandler) GetUnitById(c *gin.Context){
	var unit entity.Unit
	idStr:=c.Param("id")
	idInt,err:=strconv.Atoi(idStr)
	if err!=nil {
		c.JSON(http.StatusBadRequest,gin.H{"error":"Parameter id harus berbentuk angka"})
		return 
	}
	unit,err= h.srv.GetUnitById(uint(idInt))
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
		return 
	}	
	c.JSON(http.StatusOK,gin.H{
		"message":"Berhasil mengambil data unit",
		"data":unit,
	})
}