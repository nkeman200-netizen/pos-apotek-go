package main

import (
	"apotek-pos-go/internal/config"
	"apotek-pos-go/internal/modules/inventory/handler"
	"apotek-pos-go/internal/modules/inventory/repository"
	"apotek-pos-go/internal/modules/inventory/service"
	userHndler "apotek-pos-go/internal/modules/users/handler"
	userRepository "apotek-pos-go/internal/modules/users/repository"
	userSrvice "apotek-pos-go/internal/modules/users/service"
	transactionHndler "apotek-pos-go/internal/modules/transaction/handler"
	transactionRepository "apotek-pos-go/internal/modules/transaction/repository"
	transactionSrvice "apotek-pos-go/internal/modules/transaction/service"
	purchasesHndler "apotek-pos-go/internal/modules/purchases/handler"
	purchasesRepository "apotek-pos-go/internal/modules/purchases/repo"
	purchasesSrvice "apotek-pos-go/internal/modules/purchases/service"
	"apotek-pos-go/internal/pkg/midleware"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config.ConnectDB()
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "Amana",
			"pesan":  "Server apotek go siap temput",
		})
	})

	categoryRepo := repository.NewCategoryRepository(config.DB)
	categorySrv := service.NewCategoryService(categoryRepo)
	categoryHndl := handler.NewCategoryHandler(categorySrv)

	unitRepo := repository.NewUnitRepository(config.DB)
	unitSrv := service.NewUnitService(unitRepo)
	unitHndl := handler.NewUnitHandler(unitSrv)

	productRepo := repository.NewProductRepository(config.DB)
	productSrv := service.NewProductService(productRepo, categoryRepo, unitRepo)
	productHndl := handler.NewProductHandler(productSrv)

	pbRepo := repository.NewProductBatchRepository(config.DB)
	pbSrv := service.NewProductBatchService(pbRepo, productRepo)
	pbHndl := handler.NewProductBatchHandler(pbSrv)

	userRepo := userRepository.NewUserRepo(config.DB)
	userSrv := userSrvice.NewUserService(userRepo)
	userHndl := userHndler.NewUserHandler(userSrv)

	transRepo := transactionRepository.NewTransactionRepository(config.DB)
	customerRepo := transactionRepository.NewCustomerRepo(config.DB)
	transSrv := transactionSrvice.NewTransactionServiceImpl(transRepo,customerRepo,userRepo,pbRepo)
	transHndl := transactionHndler.NewTransactionHndl(transSrv,pbSrv)
	
	customerSrv:=transactionSrvice.NewCustomerSrv(customerRepo)
	customerHndler:=transactionHndler.NewCustomerHndler(customerSrv)

	supplierRepo := purchasesRepository.NewSupplierRepo(config.DB)
	supplierSrv:=purchasesSrvice.NewSupplierSrv(supplierRepo)
	supplierHandler:=purchasesHndler.NewSupplierHandler(supplierSrv)

	purchaseRepo := purchasesRepository.NewPurchaseRepo(config.DB)
	purchaseSrv := purchasesSrvice.NewPurchaseRepo(purchaseRepo,supplierRepo,userRepo)
	purchaseHndl := purchasesHndler.NewPurchaseHandler(purchaseSrv)

	api := r.Group("/api")

	api.POST("/users/register", userHndl.Register)
	api.POST("/users/login", userHndl.Login)

	privateRoutes := api.Group("/")
	privateRoutes.Use(midleware.JwtAuthMiddleware())
	{
		admKsr := privateRoutes.Group("/")
		admKsr.Use(midleware.AdminKasirMiddleware())
		{
			admKsr.POST("/transactions",transHndl.Create)
			admKsr.PUT("/transactions/:id",transHndl.Void)
			admKsr.POST("/customers",customerHndler.Create)
			admKsr.PUT("/customers/:id",customerHndler.Update)
			admKsr.DELETE("/customers/:id",customerHndler.Delete)
			
		}

		admOwnr:=privateRoutes.Group("/")
		admOwnr.Use(midleware.OwnerAdminMiddleware())
		{
			admOwnr.GET("/purchases",purchaseHndl.GetAll)
			admOwnr.GET("/purchases/:id",purchaseHndl.GetById)

			admOwnr.GET("/units", unitHndl.GetAll)
			admOwnr.GET("/units/:id", unitHndl.GetUnitById)

			admOwnr.GET("/categories", categoryHndl.GetAll)
			admOwnr.GET("/categories/:id", categoryHndl.GetById)

			admOwnr.GET("/supplier/",supplierHandler.GetAll)
			admOwnr.GET("/supplier/:id",supplierHandler.GetSupplierById)

		}

		admin := privateRoutes.Group("/")
		admin.Use(midleware.AdminMiddleware())
		{
			admin.POST("/categories", categoryHndl.Create)
			admin.PUT("/categories/:id", categoryHndl.Update)
			admin.DELETE("/categories/:id", categoryHndl.Delete)

			admin.POST("/units", unitHndl.Create)
			admin.PUT("/units/:id", unitHndl.Update)
			admin.DELETE("/units/:id", unitHndl.Delete)

			admin.POST("/products", productHndl.Create)
			admin.PUT("/products/:id", productHndl.Update)
			admin.DELETE("/products/:id", productHndl.Delete)

			admin.POST("/product-batches", pbHndl.Create)
			admin.PUT("/product-batches/:id", pbHndl.Update)
			admin.DELETE("/product-batches/:id", pbHndl.Delete)

			admin.POST("/purchases", purchaseHndl.Create)
			admin.PUT("/purchases/:id", purchaseHndl.Void)

			admin.POST("/supplier",supplierHandler.Create)
			admin.PUT("/supplier/:id",supplierHandler.Update)
			admin.DELETE("/supplier/:id",supplierHandler.Delete)
		}

		owner := privateRoutes.Group("/")
		owner.Use(midleware.OwnerMiddleware())
		{
			owner.PUT("/users/:id", userHndl.UpdateUser)
			owner.DELETE("/users/:id", userHndl.Delete)
			owner.GET("/users", userHndl.GetAllUser)
		}

		privateRoutes.GET("/products", productHndl.GetAllProduct)
		privateRoutes.GET("/products/:id", productHndl.GetProductById)
		
		privateRoutes.GET("/product-batches", pbHndl.GetAllProductBatch)
		privateRoutes.GET("/product-batches/:id", pbHndl.GetProductBatchById)

		privateRoutes.GET("/transactions",transHndl.GetAllTransaction)
		privateRoutes.GET("/transactions/:id",transHndl.GetTransactionById)

		privateRoutes.GET("/customers",customerHndler.GetAll)
		privateRoutes.GET("/customers/:id",customerHndler.GetById)

	}

	r.Run(":8080")
}
