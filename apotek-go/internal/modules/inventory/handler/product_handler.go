package handler

import (
	"apotek-pos-go/internal/modules/inventory/entity"
	"apotek-pos-go/internal/modules/inventory/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type productHandler struct {
	srv service.ProductService
}

func NewProductHandler(srv service.ProductService) *productHandler{
	return &productHandler{srv: srv}
}

func (h *productHandler) Create(c *gin.Context){
	var product entity.Product
	if err:=c.ShouldBindJSON(&product);err!=nil {
		c.JSON(http.StatusBadRequest,gin.H{"error":"Format request salah"})
		return 
	}
	err:=h.srv.CreateProduct(&product)
	if err!=nil {
		c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})	
		return 
	}

	c.JSON(http.StatusOK,gin.H{
		"message": "Berhasil membuat product baru",
		"data": product,
	})
}

func (h *productHandler) Update(c *gin.Context){
	var product entity.Product
	if err:=c.ShouldBindJSON(&product);err!=nil {
		c.JSON(http.StatusBadRequest,gin.H{"error":"Format request salah"})
		return
	}
	idStr:=c.Param("id")
	idInt,err:=strconv.Atoi(idStr)
	if err!=nil {
		c.JSON(http.StatusBadRequest,gin.H{"error":"Parameter id harus berupa angka"})
		return
	}
	product.ID=uint(idInt)
	if err:=h.srv.UpdateProduct(&product);err!=nil {
		c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})	
		return 
	}

	c.JSON(http.StatusOK,gin.H{
		"message": "Berhasil memperbarui data product",
		"data": product,
	})
}

func (h *productHandler) Delete(c *gin.Context){
	idStr:=c.Param("id")
	idInt,err:=strconv.Atoi(idStr)
	if err!=nil {
		c.JSON(http.StatusBadRequest,gin.H{"error":"Parameter id harus berupa angka"})
		return
	}

	if err:=h.srv.DeleteProduct(uint(idInt));err!=nil {
		c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})	
		return 
	}
	c.JSON(http.StatusOK,gin.H{
		"message": "Berhasil menghapus data product",
	})
}

func (h *productHandler) GetAllProduct(c *gin.Context){
	var product []entity.Product
	product,err:=h.srv.GetAllProduct()
	if err!=nil {
		c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})	
		return 
	}
	c.JSON(http.StatusOK,gin.H{
		"message": "Berhasil mengambil seluruh data product",
		"data": product,
	})
}

func (h *productHandler) GetProductById(c *gin.Context){
	var product entity.Product
	idStr:=c.Param("id")
	idInt,err:=strconv.Atoi(idStr)
	if err!=nil {
		c.JSON(http.StatusBadRequest,gin.H{"error":"Parameter id harus berupa angka"})
		return
	}
	product,err=h.srv.GetProductById(uint(idInt))
	if err!=nil {
		c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})	
		return 
	}
	c.JSON(http.StatusOK,gin.H{
		"message": "Berhasil mengambil data product",
		"data": product,
	})
}