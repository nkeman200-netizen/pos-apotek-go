package handler

import (
	invRepo "apotek-pos-go/internal/modules/inventory/repository"
	"apotek-pos-go/internal/modules/transaction/entity"
	"apotek-pos-go/internal/modules/transaction/service"
	"apotek-pos-go/internal/pkg/response"
	"crypto/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type TransactionRequest struct {
	CustomerId       uint                       `json:"customer_id"`
	PaymentMethod    string                     `json:"payment_method"`
	Pembayaran       int64                      `json:"pembayaran"`
	Details          []TransactionDetailRequest `json:"details"`
}

type TransactionDetailRequest struct {
	ProductId      uint `json:"product_id"`
	Quantity       int  `json:"quantity"`
}

type VoidRequest struct{
	VoidReason string `json:"void_reason"`
}

type TransactionHndl struct {
	srv service.TransactionService
	batchRepo invRepo.ProductBatchRepository
}

func NewTransactionHndl(
		srv service.TransactionService,
		batchRepo invRepo.ProductBatchRepository,
	) *TransactionHndl{
		return &TransactionHndl{
			srv: srv,
			batchRepo: batchRepo,
		}
	}

func (h *TransactionHndl) Create(c *gin.Context){
	var req TransactionRequest
	var t entity.Transaction
	err:=c.ShouldBindJSON(&req)
	if err!=nil {
		c.JSON(http.StatusBadRequest,response.Response[any]{
			Code: 400,
			Message: "Format request salah",
			Data: nil,
		})
		return 
	}

	idUser,_:=c.Get("user_id")
	t.UserId=uint(idUser.(float64))
	
	t.InvoiceNumber="INV-"+strconv.FormatInt(time.Now().Unix(),10)
	t.CustomerId=req.CustomerId
	t.PaymentMethod=req.PaymentMethod
	switch req.PaymentMethod {
		case "cash":
			t.PaymentReference=""
		case "qris":
			t.PaymentReference="QRIS-"+rand.Text()
	}
	t.Pembayaran=req.Pembayaran

	var totalPrice int64
	for i := range req.Details{
		batch,_:=h.batchRepo.FindByProductId(req.Details[i].ProductId)
		qtyDibutuhkan:=req.Details[i].Quantity
		for j :=range batch{
			if qtyDibutuhkan<=batch[j].Stock {
				subTotal:=batch[j].Product.SellingPrice*int64(qtyDibutuhkan)
				t.Details=append(t.Details, entity.TransactionDetails{
					ProductId: req.Details[i].ProductId,
					ProductBatchId: batch[j].ID,
					Quantity: qtyDibutuhkan,
					UnitPrice:batch[j].Product.SellingPrice,
					SubTotal:subTotal,
				})
				
				qtyDibutuhkan=0
				totalPrice+=subTotal
				break
			}else if qtyDibutuhkan>batch[j].Stock{
				stokAda:=batch[j].Stock
				subTotal:=batch[j].Product.SellingPrice*int64(stokAda)
				t.Details=append(t.Details, entity.TransactionDetails{
					ProductId: req.Details[i].ProductId,
					ProductBatchId: batch[j].ID,
					Quantity: stokAda,
					UnitPrice:batch[j].Product.SellingPrice,
					SubTotal:subTotal,
				})
				qtyDibutuhkan-=stokAda
				totalPrice+=subTotal
				continue
			}
		}
	}

	t.TotalPrice=totalPrice
	t.Kembalian=req.Pembayaran-totalPrice
	t.Status="complete"

	if err:=h.srv.Create(&t);err!=nil {
		c.JSON(http.StatusInternalServerError,response.Response[any]{
			Code: 500,
			Message: "Bermasalah di server: "+err.Error(),
			Data: nil,
		})
		return 
	}

	c.JSON(http.StatusOK,response.Response[entity.Transaction]{
		Code: 201,
		Message: "Berhasil membuat transaksi baru",
		Data: t,
	})
}

func (h *TransactionHndl) Void(c *gin.Context){
	sId:=c.Param("id")
	IntId,err:=strconv.Atoi(sId)
	if err!=nil {
		c.JSON(http.StatusBadGateway,response.Response[any]{
			Code: 400,
			Message: "Parameter id harus diisi id transaksi: "+err.Error(),
			Data: nil,
		})
		return 
	}
	t,err:=h.srv.FindById(uint(IntId))
	if err!=nil {
		c.JSON(http.StatusInternalServerError,response.Response[any]{
			Code: 500,
			Message: "Data transaksi tidak ditemukan: "+err.Error(),
			Data: nil,
		})
		return 
	}
	
	var voidReq VoidRequest
	err=c.ShouldBindJSON(&voidReq)
	if err!=nil {
		c.JSON(http.StatusBadRequest,response.Response[any]{
			Code: 400,
			Message: "Format data salah: "+err.Error(),
			Data: nil,
		})
		return 
	}

	t.Status="void"
	t.VoidReason=voidReq.VoidReason
	if err:=h.srv.Void(&t);err!=nil {
		c.JSON(http.StatusInternalServerError,response.Response[any]{
			Code: 500,
			Message: "Gagal melakukan void transaksi: "+err.Error(),
			Data: nil,
		})
		return 
	}
	c.JSON(http.StatusOK,response.Response[entity.Transaction]{
		Code: 200,
		Message: "Berhasil melakukan void transaksi",
		Data: t,
	})
}

func (h *TransactionHndl) GetAllTransaction(c *gin.Context){
	t,err:=h.srv.FindAll()
	if err!=nil {
		c.JSON(http.StatusInternalServerError,response.Response[any]{
			Code: 500,
			Message: "Server sedang bermasalah: "+err.Error(),
			Data: nil,
		})
		return
	}
	c.JSON(http.StatusOK,response.Response[[]entity.Transaction]{
		Code: 200,
		Message: "Berhasil mengambil semua data transaksi",
		Data: t,
	})
}

func (h *TransactionHndl) GetTransactionById(c *gin.Context){
	idS:=c.Param("id")
	idI,_:=strconv.Atoi(idS)
	t,err:=h.srv.FindById(uint(idI))
	if err!=nil {
		c.JSON(http.StatusInternalServerError,response.Response[any]{
			Code: 500,
			Message: "Server sedang bermasalah: "+err.Error(),
			Data: nil,
		})
		return
	}
	c.JSON(http.StatusOK,response.Response[entity.Transaction]{
		Code: 200,
		Message: "Berhasil mengambil data transaksi id "+idS,
		Data: t,
	})
}