package handler

import (
	"apotek-pos-go/internal/modules/inventory/entity"
	"apotek-pos-go/internal/modules/inventory/service"
	"apotek-pos-go/internal/pkg/response"
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
		c.JSON(http.StatusBadRequest,response.Response[any]{
			Code: 400,
			Message: "Format data salah: " + err.Error(),
			Data: nil,
		})
		return 
	}
	err:=h.srv.CreateProduct(&product)
	if err!=nil {
		c.JSON(http.StatusInternalServerError,response.Response[any]{
			Code: 500,
			Message: "Server sedang bermasalah: "+err.Error(),
			Data: nil,
		})
		return 
	}

	c.JSON(http.StatusCreated,response.Response[entity.Product]{
		Code: 201,
		Message: "Berhasil membuat product baru",
		Data: product,
	})
}

func (h *productHandler) Update(c *gin.Context){
	var product entity.Product
	if err:=c.ShouldBindJSON(&product);err!=nil {
		c.JSON(http.StatusBadRequest,response.Response[any]{
			Code: 400,
			Message: "Format data salah: " + err.Error(),
			Data: nil,
		})
		return
	}
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
	product.ID=uint(idInt)
	if err:=h.srv.UpdateProduct(&product);err!=nil {
		c.JSON(http.StatusInternalServerError,response.Response[any]{
			Code: 500,
			Message: "Server sedang bermasalah: "+err.Error(),
			Data: nil,
		})
		return 
	}

	c.JSON(http.StatusOK,response.Response[entity.Product]{
		Code: 200,
		Message: "Berhasil memperbarui data product",
		Data: product,
	})
}

func (h *productHandler) Delete(c *gin.Context){
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

	if err:=h.srv.DeleteProduct(uint(idInt));err!=nil {
		c.JSON(http.StatusInternalServerError,response.Response[any]{
			Code: 500,
			Message: "Server sedang bermasalah: "+err.Error(),
			Data: nil,
		})
		return 
	}
	c.JSON(http.StatusOK,response.Response[any]{
		Code: 204,
		Message: "Berhasil menghapus data product",
		Data: nil,
	})
}

func (h *productHandler) GetAllProduct(c *gin.Context){
	var product []entity.Product
	var filter entity.ProductFilter

	filter.Name=c.Query("name")
	filter.SKU=c.Query("sku")
	filter.UnitId=c.Query("unit_id")
	filter.CategoryId=c.Query("category_id")
	limInt,err:=strconv.Atoi(c.DefaultQuery("limit","10"))
	filter.Limit=limInt
	pageInt,err:=strconv.Atoi(c.DefaultQuery("page","1"))
	filter.Page=pageInt

	product,err=h.srv.GetAllProduct(filter)
	if err!=nil {
		c.JSON(http.StatusInternalServerError,response.Response[any]{
			Code: 500,
			Message: "Server sedang bermasalah: "+err.Error(),
			Data: nil,
		})
		return 
	}
	c.JSON(http.StatusOK,response.Response[[]entity.Product]{
		Code: 200,
		Message: "Berhasil mengambil semua data product",
		Data: product,
	})
}

func (h *productHandler) GetProductById(c *gin.Context){
	var product entity.Product
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
	product,err=h.srv.GetProductById(uint(idInt))
	if err!=nil {
		c.JSON(http.StatusInternalServerError,response.Response[any]{
			Code: 500,
			Message: "Server sedang bermasalah: "+err.Error(),
			Data: nil,
		})
		return 
	}
	c.JSON(http.StatusOK,response.Response[entity.Product]{
		Code: 200,
		Message: "Berhasil mengambil data product",
		Data: product,
	})
}