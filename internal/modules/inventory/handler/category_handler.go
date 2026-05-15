package handler

import (
	"apotek-pos-go/internal/modules/inventory/entity"
	"apotek-pos-go/internal/modules/inventory/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	srv service.CategoryService
}

func NewCategoryHandler(srv service.CategoryService) *CategoryHandler{
	return &CategoryHandler{srv: srv}
}

func (h *CategoryHandler) Create(c *gin.Context){
	var category entity.Category
	if err:=c.ShouldBindJSON(&category); err!=nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format data salah"})
		return
	}

	if err:=h.srv.CreateCategory(&category); err!=nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return 
	}

	c.JSON(http.StatusCreated,gin.H{
		"message":"Kategori berhasil ditambahkan",
		"data":category,
	})
}

func (h *CategoryHandler) GetAll(c *gin.Context) {
	var category []entity.Category
	category,err:=h.srv.GetAllCategory()
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error":"Server bermasalah"})
		return
	}

	c.JSON(http.StatusOK,gin.H{
		"message": "Get all category berhasil",
		"data": category,
	})
}

func (h *CategoryHandler) GetById(c *gin.Context){
	idStr:=c.Param("id")
	idInt,err:=strconv.Atoi(idStr)
	if err!=nil {
		c.JSON(http.StatusBadRequest,gin.H{"error":"Masukan hanya angka"})	
		return
	}

	category, err:= h.srv.GetCategoryById(uint(idInt))
	if err!=nil {
		c.JSON(http.StatusNotFound,gin.H{"error":"Kategori tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK,gin.H{
		"message": "Bet category by id berhasil",
		"data":category,
	})
}

func (h *CategoryHandler) Update(c *gin.Context){
	var category entity.Category
	idStr:=c.Param("id")
	idInt,err:=strconv.Atoi(idStr)
	if err!=nil {
		c.JSON(http.StatusBadRequest,gin.H{"error":"Masukan id categgory"})
		return
	}
	if err:=c.ShouldBindJSON(&category);err!=nil {
		c.JSON(http.StatusBadRequest,gin.H{"error":"Format data salah"})
		return
	}
	category.ID=uint(idInt)

	err= h.srv.UpdateCategory(&category)
	if err!=nil {
		c.JSON(http.StatusInternalServerError,gin.H{"error":"Data category gagal terupdate"})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"message":"Data category berhasil diperbarui",
		"data":category,
	})
}

func (h *CategoryHandler) Delete(c *gin.Context){
	idStr:=c.Param("id")
	idInt,err:=strconv.Atoi(idStr)
	if err!=nil {
		c.JSON(http.StatusBadRequest,gin.H{"error":"Masukan id hanya angka"})
		return
	}
	err=h.srv.DeleteCategory(uint(idInt))
	if err!=nil {
		c.JSON(http.StatusInternalServerError,gin.H{"error":"Gagal menghapus data category"})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"message":"Data category berhasil dihapus",
	})
}
