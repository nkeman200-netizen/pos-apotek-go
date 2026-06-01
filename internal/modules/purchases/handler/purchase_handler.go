package handler

import (
	"apotek-pos-go/internal/modules/purchases/entity"
	"apotek-pos-go/internal/modules/purchases/service"
	"apotek-pos-go/internal/pkg/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PurchaseHandler struct {
	srv service.PurchaseSrv
}

type VoidRequest struct{
	VoidReason string `json:"void_reason"`
}

func NewPurchaseHandler(srv service.PurchaseSrv) *PurchaseHandler{
	return &PurchaseHandler{srv: srv}
}

func(h *PurchaseHandler) Create(c *gin.Context){
	var purchase entity.Purchase
	err:=c.ShouldBindJSON(&purchase)
	if err!=nil {
		c.JSON(http.StatusBadRequest,response.Response[any]{
			Code: 400,
			Message: "Format request salah: "+err.Error(),
			Data: nil,
		})
		return 
	}
	userId,_:=c.Get("user_id")
	userIdUint:=uint(userId.(float64))
	purchase.UserId=userIdUint

	if err:=h.srv.Create(&purchase);err!=nil {
		c.JSON(http.StatusInternalServerError,response.Response[any]{
			Code: 500,
			Message: "Masalah pada server: "+err.Error(),
			Data: nil,
		})
		return 
	}
	c.JSON(http.StatusCreated,response.Response[entity.Purchase]{
		Code: 201,
		Message: "Berhasil membuat purchase baru",
		Data: purchase,
	})
}

func(h *PurchaseHandler) Void(c *gin.Context){
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
	var req VoidRequest
	err=c.ShouldBindJSON(&req)
	if err!=nil {
		c.JSON(http.StatusBadRequest,response.Response[any]{
			Code: 400,
			Message: "Format request salah: "+err.Error(),
			Data: nil,
		})
		return 
	}
	purchase,err:=h.srv.GetById(uint(idInt))
	purchase.VoidReason=req.VoidReason
	if err:=h.srv.Void(&purchase);err!=nil {
		c.JSON(http.StatusInternalServerError,response.Response[any]{
			Code: 500,
			Message: "Gagal melakukan void purchase: "+err.Error(),
			Data: nil,
		})
		return 
	}
	c.JSON(http.StatusOK,response.Response[entity.Purchase]{
		Code: 200,
		Message: "Berhasiil melakukan void purchase",
		Data: purchase,
	})
}


func ( h *PurchaseHandler) GetAll(c *gin.Context){
	var Purchase []entity.Purchase
	Purchase,err:=h.srv.GetAll();
	if err!=nil {
		c.JSON(http.StatusInternalServerError,response.Response[any]{
			Code: 500,
			Message: "Server sedang bermasalah: "+err.Error(),
			Data: nil,
		})
		return 
	}
	c.JSON(http.StatusOK,response.Response[[]entity.Purchase]{
		Code: 200,
		Message: "Berhasil mengambil semua data Purchase",
		Data: Purchase,
	})
}

func (h *PurchaseHandler) GetById(c *gin.Context){
	var Purchase entity.Purchase
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
	Purchase,err= h.srv.GetById(uint(idInt))
	if err!=nil{
		c.JSON(http.StatusInternalServerError,response.Response[any]{
			Code: 500,
			Message: "Server sedang bermasalah: "+err.Error(),
			Data: nil,
		})
		return 
	}	
	c.JSON(http.StatusOK,response.Response[entity.Purchase]{
		Code: 200,
		Message: "Berhasil mengambil data Purchase",
		Data: Purchase,
	})
}