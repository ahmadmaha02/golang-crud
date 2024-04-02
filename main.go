package main

import (
	"github.com/gin-gonic/gin"
	productcontroller "github.com/golang-crud/controllers"
	"github.com/golang-crud/database"
	"github.com/golang-crud/repositories"
)

func main() {
	r := gin.Default()
	database.ConnectDatabase()

	productRepo := repositories.NewProductRepository(database.DB)
	productController := productcontroller.NewProductController(productRepo)

	r.GET("/api/products", productController.Index)
	r.GET("/api/product/:id", productController.Show)
	r.POST("/api/product", productController.Create)
	r.PUT("/api/product/:id", productController.Update)
	r.DELETE("/api/product/:id", productController.Delete)

	r.Run(":8080")
}
