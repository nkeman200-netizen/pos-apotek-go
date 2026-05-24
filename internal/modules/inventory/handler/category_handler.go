package handler

import (
	"apotek-pos-go/internal/modules/inventory/entity"
	"apotek-pos-go/internal/modules/inventory/service"
	"apotek-pos-go/internal/pkg/response"
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
		c.JSON(http.StatusBadRequest, response.Response[any]{
			Code: 400,
			Message: "Format data salah: " + err.Error(),
			Data: nil,
		})
		return
	}

	if err:=h.srv.CreateCategory(&category); err!=nil {
		c.JSON(http.StatusInternalServerError, response.Response[any]{
			Code: 500,
			Message: "Gagal membuat category baru",
			Data: nil,
		})
		return 
	}

	c.JSON(http.StatusCreated,response.Response[entity.Category]{
		Code: 201,
		Message: "Berhasil menambahkan category",
		Data: category,
	})
}

func (h *CategoryHandler) GetAll(c *gin.Context) {
	var category []entity.Category
	category,err:=h.srv.GetAllCategory()
	if err!=nil{
		c.JSON(http.StatusInternalServerError,response.Response[any]{
			Code: 500,
			Message: "Server sedang bermasalah: "+err.Error(),
			Data: nil,
		})
		return
	}

	c.JSON(http.StatusOK,response.Response[[]entity.Category]{
		Code: 200,
		Message: "Berhasil mengambil semua data category",
		Data: category,
	})
}

func (h *CategoryHandler) GetById(c *gin.Context){
	idStr:=c.Param("id")
	idInt,err:=strconv.Atoi(idStr)
	if err!=nil {
		c.JSON(http.StatusBadRequest,response.Response[any]{
			Code: 400,
			Message: "Masukan hanya angka category id",
			Data: nil,
		})	
		return
	}

	category, err:= h.srv.GetCategoryById(uint(idInt))
	if err!=nil {
		c.JSON(http.StatusNotFound,response.Response[any]{
			Code: 404,
			Message: "Kategori tidak ditemukan",
			Data: nil,
		})
		return
	}

	c.JSON(http.StatusOK,response.Response[entity.Category]{
		Code: 200,
		Message: "get category by id berhasil",
		Data: category,
	})
}

func (h *CategoryHandler) Update(c *gin.Context){
	var category entity.Category
	idStr:=c.Param("id")
	idInt,err:=strconv.Atoi(idStr)
	if err!=nil {
		c.JSON(http.StatusBadRequest,response.Response[any]{
			Code: 400,
			Message: "Masukan id categgory",
			Data: nil,
		})
		return
	}
	if err:=c.ShouldBindJSON(&category);err!=nil {
		c.JSON(http.StatusBadRequest,response.Response[any]{
			Code: 400,
			Message: "Format data salah: " + err.Error(),
			Data: nil,
		})
		return
	}
	category.ID=uint(idInt)

	err= h.srv.UpdateCategory(&category)
	if err!=nil {
		c.JSON(http.StatusInternalServerError,response.Response[any]{
			Code: 500,
			Message: "Server sedang bermasalah: "+err.Error(),
			Data: nil,
		})
		return
	}
	c.JSON(http.StatusOK,response.Response[entity.Category]{
		Code: 200,
		Message: "Berhasil memperbarui data category",
		Data: category,
	})
}

func (h *CategoryHandler) Delete(c *gin.Context){
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
	err=h.srv.DeleteCategory(uint(idInt))
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
		Message: "Data category berhasil dihapus",
		Data: nil,
	})
}
