package main

import (
	"apotek-pos-go/internal/config"
	"apotek-pos-go/internal/modules/inventory/handler"
	"apotek-pos-go/internal/modules/inventory/repository"
	"apotek-pos-go/internal/modules/inventory/service"
	"net/http"

	"github.com/gin-gonic/gin"
)
func main() {
	config.ConnectDB()
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "Amana",
			"pesan": "Server apotek go siap temput",
		})
	})

	categoryRepo:=repository.NewCategoryRepository(config.DB)
	categorySrv:=service.NewCategoryService(categoryRepo)
	categoryHndl:=handler.NewCategoryHandler(categorySrv)

	unitRepo:=repository.NewUnitRepository(config.DB)
	unitSrv:=service.NewUnitService(unitRepo)
	unitHndl:=handler.NewUnitHandler(unitSrv)

	productRepo:=repository.NewProductRepository(config.DB)
	productSrv:=service.NewProductService(productRepo,categoryRepo,unitRepo)
	productHndl:=handler.NewProductHandler(productSrv)

	pbRepo:=repository.NewProductBatchRepository(config.DB)
	pbSrv:=service.NewProductBatchService(pbRepo,productRepo)
	pbHndl:=handler.NewProductBatchHandler(pbSrv)
	
	api:=r.Group("/api")
	{
		api.POST("/categories",categoryHndl.Create)
		api.GET("/categories",categoryHndl.GetAll)
		api.GET("/categories/:id",categoryHndl.GetById)
		api.PUT("/categories/:id",categoryHndl.Update)
		api.DELETE("/categories/:id",categoryHndl.Delete)

		api.POST("/units",unitHndl.Create)
		api.PUT("/units/:id",unitHndl.Update)
		api.DELETE("/units/:id",unitHndl.Delete)
		api.GET("/units",unitHndl.GetAll)
		api.GET("/units/:id",unitHndl.GetUnitById)

		api.POST("/products",productHndl.Create)
		api.PUT("/products/:id",productHndl.Update)
		api.DELETE("/products/:id",productHndl.Delete)
		api.GET("/products",productHndl.GetAllProduct)
		api.GET("/products/:id",productHndl.GetProductById)

		api.POST("/product-batches",pbHndl.Create)
		api.PUT("/product-batches/:id",pbHndl.Update)
		api.DELETE("/product-batches/:id",pbHndl.Delete)
		api.GET("/product-batches",pbHndl.GetAllProductBatch)
		api.GET("/product-batches/:id",pbHndl.GetProductBatchById)
		
		
	}

	r.Run(":8080")
}