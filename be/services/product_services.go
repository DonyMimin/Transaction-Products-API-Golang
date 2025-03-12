package services

import (
	"context"

	"github.com/DonyMimin/Go-crud-echo/models"
	"github.com/DonyMimin/Go-crud-echo/repository"
)

func GetListProduct(ctx context.Context) (data []models.Product, err error) {
	// Get list data product
	data, err = repository.FindAllProducts(ctx)
	if err != nil {
		return
	}

	return
}

func GetProduct(ctx context.Context, id int) (data models.Product, err error) {
	// Get data product
	data, err = repository.FindProduct(ctx, id)
	if err != nil {
		return
	}

	return
}

func UpsertProduct(ctx context.Context, params models.Product) (err error) {
	// Update and Insert product
	err = repository.UpsertProduct(ctx, []string{"product_id", "nama_product", "price", "stock"}, models.Product{
		ProductID:   params.ProductID,
		NamaProduct: params.NamaProduct,
		Price:       params.Price,
		Stock:       params.Stock,
	})
	if err != nil {
		return
	}

	return
}

func DeleteProduct(ctx context.Context, id int) (err error) {
	// Delete product
	err = repository.DeleteProduct(ctx, id)
	if err != nil {
		return
	}

	return
}
