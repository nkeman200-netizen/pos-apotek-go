package dto

import (
	"apotek-pos-go/internal/modules/inventory/entity"
)

type productResponse struct {
	Id         uint   `json:"id"`
	Name       string `json:"name"`
	SKU        string `json:"sku"`
	Price      int64  `json:"price"`
	TotalStock int    `json:"total_stock"`
	Category   string `json:"category"`
	Unit       string `json:"unit"`
}

func FormatProduct(p entity.Product) productResponse{
	return productResponse{
		Id: p.ID,
		Name: p.Name,
		SKU: p.SKU,
		Price: p.SellingPrice,
		TotalStock: p.TotalStock,
		Category: p.Category.Name,
		Unit: p.Unit.Name,
	}
}

func FormatProducts(p []entity.Product) []productResponse{
	var formatProducts []productResponse
	for _,product:=range p{
		formatedProduct:=FormatProduct(product)
		formatProducts=append(formatProducts, formatedProduct)
	}
	return formatProducts
}