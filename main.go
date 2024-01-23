package main

import (
	"go_project/controllers"
	"go_project/services"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	productService := services.NewProductService()

	productHandler := &controllers.ProductHandler{ProductService: productService}

	r.GET("/products", productHandler.ListProducts)
	r.GET("/products/:id", productHandler.GetProduct)
	r.POST("/products", productHandler.CreateProduct)
	r.PUT("/products/:id", productHandler.UpdateProduct)
	r.DELETE("/products/:id", productHandler.DeleteProduct)

	r.Run(":8080")
}
