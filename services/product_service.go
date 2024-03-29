package services

import (
	"errors"
	"go_project/models"
)

type ProductService interface {
	ListProducts() []models.Product
	GetProduct(id int) (models.Product, error)
	CreateProduct(newProduct models.Product) (models.Product, error)
	UpdateProduct(id int, updatedProduct models.Product) (models.Product, error)
	DeleteProduct(id int) error
}

type ProductServiceImpl struct {
	products []models.Product
}

func NewProductService() ProductService {
	return &ProductServiceImpl{
		products: []models.Product{},
	}
}

func (ps *ProductServiceImpl) ListProducts() []models.Product {
	return ps.products
}

func (ps *ProductServiceImpl) GetProduct(id int) (models.Product, error) {
	for _, product := range ps.products {
		if product.ID == id {
			return product, nil
		}
	}
	return models.Product{}, errors.New("Product not found")
}

func (ps *ProductServiceImpl) CreateProduct(newProduct models.Product) (models.Product, error) {
	newProduct.ID = len(ps.products) + 1

	ps.products = append(ps.products, newProduct)

	return newProduct, nil
}

func (ps *ProductServiceImpl) UpdateProduct(id int, updatedProduct models.Product) (models.Product, error) {
	for i, product := range ps.products {
		if product.ID == id {
			updatedProduct.ID = id
			ps.products[i] = updatedProduct
			return updatedProduct, nil
		}
	}
	return models.Product{}, errors.New("Product not found")
}

func (ps *ProductServiceImpl) DeleteProduct(id int) error {
	for i, product := range ps.products {
		if product.ID == id {
			ps.products = append(ps.products[:i], ps.products[i+1:]...)
			return nil
		}
	}
	return errors.New("Product not found")
}
