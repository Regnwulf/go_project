// main.go
package main

import (
	"go_project/controllers" // Importando o pacote controllers
	"go_project/services"    // Importando o pacote services

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Importando o pacote do service
	productService := services.NewProductService()

	// Importando o pacote do handler
	productHandler := &controllers.ProductHandler{ProductService: productService}

	// CRUD routes for products
	r.GET("/products", productHandler.ListProducts)
	r.GET("/products/:id", productHandler.GetProduct)
	r.POST("/products", productHandler.CreateProduct)
	r.PUT("/products/:id", productHandler.UpdateProduct)
	r.DELETE("/products/:id", productHandler.DeleteProduct)

	r.Run(":8080")
}
