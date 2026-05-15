package handler

import (
	"apotek-pos-go/internal/modules/inventory/entity"
	"apotek-pos-go/internal/modules/inventory/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type productBatchHandler struct {
	srv service.ProductBatchService
}

func NewProductBatchHandler(srv service.ProductBatchService) *productBatchHandler{
	return &productBatchHandler{srv: srv}
}

func(h *productBatchHandler) Create(c *gin.Context){
	var pb entity.ProductBatch
	if err:=c.ShouldBindJSON(&pb); err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error":"Format request salah"})
		return 
	}
	if err:=h.srv.CreateProductBatch(&pb);err!=nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
		return
	}

	c.JSON(http.StatusOK,gin.H{
		"message":"Berhasil membuat product batch baru",
		"data": pb,
	})
}

func(h *productBatchHandler) Update(c *gin.Context){
	var pb entity.ProductBatch
	if err:=c.ShouldBindJSON(&pb); err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error":"Format request salah"})
		return 
	}

	idstr:=c.Param("id")
	idInt,_:=strconv.Atoi(idstr)
	pb.ID=uint(idInt)

	if err:=h.srv.UpdateProductBatch(&pb);err!=nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
		return
	}

	c.JSON(http.StatusOK,gin.H{
		"message":"Berhasil update product batch",
		"data": pb,
	})
}

func(h *productBatchHandler) Delete(c *gin.Context){
	idS:=c.Param("id")
	idI,err:=strconv.Atoi(idS)
	if err!=nil {
		c.JSON(http.StatusBadRequest,gin.H{"error":"Parameter id harus berupa angka"})
		return 
	}
	if err:=h.srv.DeleteProductBatch(uint(idI));err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
		return
	}

	c.JSON(http.StatusOK,gin.H{
		"message":"Berhasil menghapus product batch",
	})
}

func(h *productBatchHandler) GetAllProductBatch(c *gin.Context){
	var pb []entity.ProductBatch
	pb,err:=h.srv.GetAllProductBatch()
	if err!=nil {
		c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
		return
	}

	c.JSON(http.StatusOK,gin.H{
		"message":"Berhasil mengambiil semua product batch",
		"data": pb,
	})
}

func(h *productBatchHandler) GetProductBatchById(c *gin.Context){
	var pb entity.ProductBatch
	idS:=c.Param("id")
	idI,_:=strconv.Atoi(idS)
	pb,err:=h.srv.GetProductBatchById(uint(idI))
	if err!=nil {
		c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
		return
	}

	c.JSON(http.StatusOK,gin.H{
		"message":"Berhasil mengambiil data product batch",
		"data": pb,
	})
}