package main

import (
	"apotek-pos-go/internal/config"
	"apotek-pos-go/internal/modules/inventory/handler"
	"apotek-pos-go/internal/modules/inventory/repository"
	"apotek-pos-go/internal/modules/inventory/service"
	userHndl "apotek-pos-go/internal/modules/users/handler"
	userRepo "apotek-pos-go/internal/modules/users/repository"
	userSrv "apotek-pos-go/internal/modules/users/service"
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

	userRepo:=userRepo.NewUserRepo(config.DB);
	userSrv:=userSrv.NewUserService(userRepo)
	userHndl:=userHndl.NewUserHandler(userSrv)
	
	api:=r.Group("/api")

	api.POST("/users/register",userHndl.Register)
	api.POST("/users/login", userHndl.Login)

	privateRoutes:=api.Group("/")
	privateRoutes.Use(midleware.JwtAuthMidleware())
	{
		admin := privateRoutes.Group("/")
		admin.Use(midleware.AdminMidleware()) 
		{ 
			admin.POST("/categories",categoryHndl.Create)
			admin.PUT("/categories/:id",categoryHndl.Update)
			admin.DELETE("/categories/:id",categoryHndl.Delete)
			
			admin.POST("/units",unitHndl.Create)
			admin.PUT("/units/:id",unitHndl.Update)
			admin.DELETE("/units/:id",unitHndl.Delete)

			admin.POST("/products",productHndl.Create)
			admin.PUT("/products/:id",productHndl.Update)
			admin.DELETE("/products/:id",productHndl.Delete)

			admin.POST("/product-batches",pbHndl.Create)
			admin.PUT("/product-batches/:id",pbHndl.Update)
			
			admin.DELETE("/product-batches/:id",pbHndl.Delete) 
		}
		

		owner := privateRoutes.Group("/")
		owner.Use(midleware.OwnerMidleware())
		{
			owner.PUT("/users/:id",userHndl.UpdateUser)
			owner.DELETE("/users/:id",userHndl.Delete)
			owner.GET("/users",userHndl.GetAllUser)
		}
		
		
	}

	r.Run(":8080")
}