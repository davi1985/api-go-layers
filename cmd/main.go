package main

import (
	"go-api/controller"
	"go-api/db"
	"go-api/repository"
	"go-api/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	dbConnection, err := db.ConnectDB()

	if err != nil {
		panic(err)
	}

	ProductRepository := repository.NewProductRepository(dbConnection)
	ProductUsecase := usecase.NewProductUsecase(ProductRepository)
	ProductController := controller.NewProductController(ProductUsecase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "pong"})
	})

	server.GET("/products", ProductController.GetProducts)
	server.GET("/products/:id", ProductController.GetProductById)
	server.POST("/products", ProductController.CreateProduct)

	server.Run(":3001")
}
