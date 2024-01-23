// controllers/product_handler.go
package controllers

import (
	"fmt"
	"go_project/models"
	"go_project/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	ProductService services.ProductService
}

func (ph *ProductHandler) ListProducts(c *gin.Context) {
	c.JSON(200, ph.ProductService.ListProducts())
}

func (ph *ProductHandler) GetProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}

	product, err := ph.ProductService.GetProduct(id)
	if err != nil {
		c.JSON(404, gin.H{"error": fmt.Sprintf("Product with ID %d not found", id)})
		return
	}

	c.JSON(200, product)
}

func (ph *ProductHandler) CreateProduct(c *gin.Context) {
	var newProduct models.Product
	if err := c.ShouldBindJSON(&newProduct); err != nil {
		c.JSON(400, gin.H{"error": "Invalid data"})
		return
	}

	product, err := ph.ProductService.CreateProduct(newProduct)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error creating the product"})
		return
	}

	c.JSON(201, product)
}

func (ph *ProductHandler) UpdateProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}

	var updatedProduct models.Product
	if err := c.ShouldBindJSON(&updatedProduct); err != nil {
		c.JSON(400, gin.H{"error": "Invalid data"})
		return
	}

	product, err := ph.ProductService.UpdateProduct(id, updatedProduct)
	if err != nil {
		c.JSON(404, gin.H{"error": fmt.Sprintf("Product with ID %d not found", id)})
		return
	}

	c.JSON(200, product)
}

func (ph *ProductHandler) DeleteProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}

	if err := ph.ProductService.DeleteProduct(id); err != nil {
		c.JSON(404, gin.H{"error": fmt.Sprintf("Product with ID %d not found", id)})
		return
	}

	c.JSON(200, gin.H{"message": fmt.Sprintf("Product with ID %d deleted", id)})
}
