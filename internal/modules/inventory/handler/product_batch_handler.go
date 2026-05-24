package handler

import (
	"apotek-pos-go/internal/modules/inventory/entity"
	"apotek-pos-go/internal/modules/inventory/service"
	"apotek-pos-go/internal/pkg/response"
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
		c.JSON(http.StatusBadRequest, response.Response[any]{
			Code: 400,
			Message: "Format data salah: " + err.Error(),
			Data: nil,
		})
		return 
	}
	if err:=h.srv.CreateProductBatch(&pb);err!=nil {
		c.JSON(http.StatusInternalServerError, response.Response[any]{
			Code: 500,
			Message: "Server sedang bermasalah: "+err.Error(),
			Data: nil,
		})
		return
	}

	c.JSON(http.StatusCreated,response.Response[entity.ProductBatch]{
		Code: 201,
		Message: "Berhasil menambahkan product batch baru",
		Data: pb,
	})
}

func(h *productBatchHandler) Update(c *gin.Context){
	var pb entity.ProductBatch
	if err:=c.ShouldBindJSON(&pb); err!=nil{
		c.JSON(http.StatusBadRequest, response.Response[any]{
			Code: 400,
			Message: "Format data salah: " + err.Error(),
			Data: nil,
		})
		return 
	}

	idstr:=c.Param("id")
	idInt,_:=strconv.Atoi(idstr)
	pb.ID=uint(idInt)

	if err:=h.srv.UpdateProductBatch(&pb);err!=nil {
		c.JSON(http.StatusInternalServerError, response.Response[any]{
			Code: 500,
			Message: "Server sedang bermasalah: "+err.Error(),
			Data: nil,
		})
		return
	}

	c.JSON(http.StatusOK,response.Response[entity.ProductBatch]{
		Code: 200,
		Message: "Berhasil memperbarui data product batch",
		Data: pb,
	})
}

func(h *productBatchHandler) Delete(c *gin.Context){
	idS:=c.Param("id")
	idI,err:=strconv.Atoi(idS)
	if err!=nil {
		c.JSON(http.StatusBadRequest,response.Response[any]{
			Code: 400,
			Message: "Parameter id harus berupa angka",
			Data: nil,
		})
		return 
	}
	if err:=h.srv.DeleteProductBatch(uint(idI));err!=nil{
		c.JSON(http.StatusInternalServerError,response.Response[any]{
			Code: 500,
			Message: "Server sedang bermasalah: "+err.Error(),
			Data: nil,
		})
		return
	}

	c.JSON(http.StatusOK,response.Response[any]{
		Code: 204,
		Message: "Berhasil menghapus product batch",
		Data: nil,
	})
}

func(h *productBatchHandler) GetAllProductBatch(c *gin.Context){
	var pb []entity.ProductBatch
	pb,err:=h.srv.GetAllProductBatch()
	if err!=nil {
		c.JSON(http.StatusInternalServerError,response.Response[any]{
			Code: 500,
			Message: "Server sedang bermasalah: "+err.Error(),
			Data: nil,
		})
		return
	}

	c.JSON(http.StatusOK,response.Response[[]entity.ProductBatch]{
		Code: 200,
		Message: "Berhasil mengambil semua data product batch",
		Data: pb,
	})
}

func(h *productBatchHandler) GetProductBatchById(c *gin.Context){
	var pb entity.ProductBatch
	idS:=c.Param("id")
	idI,_:=strconv.Atoi(idS)
	pb,err:=h.srv.GetProductBatchById(uint(idI))
	if err!=nil {
		c.JSON(http.StatusInternalServerError,response.Response[any]{
			Code: 500,
			Message: "Server sedang bermasalah: "+err.Error(),
			Data: nil,
		})
		return
	}

	c.JSON(http.StatusOK,response.Response[entity.ProductBatch]{
		Code: 200,
		Message: "Berhasil mengambil data product batch",
		Data: pb,
	})
}