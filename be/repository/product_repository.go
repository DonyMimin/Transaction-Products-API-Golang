package repository

import (
	"context"

	"github.com/DonyMimin/Go-crud-echo/config"
	"github.com/DonyMimin/Go-crud-echo/models"
	"gorm.io/gorm/clause"
)

// FindAllProducts is a func to find all list products data
func FindAllProducts(ctx context.Context) ([]models.Product, error) {
	var products []models.Product
	err := config.DBGorm.WithContext(ctx).
		Raw("SELECT * FROM products").
		Scan(&products).Error

	if err != nil {
		return nil, err
	}
	return products, nil
}

// FindProduct is a func to find product data
func FindProduct(ctx context.Context, id int) (data models.Product, err error) {
	query := config.DBGorm.Table("products").
		WithContext(ctx).
		Select(
			"product_id",
			"nama_product",
			"price",
			"stock",
		)

	if id != 0 {
		query.Where("product_id = ?", id)
	}

	err = query.Find(&data).Error
	return
}

// UpsertProduct is a func to update or insert product data
func UpsertProduct(ctx context.Context, updateFields []string, params models.Product) (err error) {
	err = config.DBGorm.Table("products").
		Clauses(clause.OnConflict{Columns: []clause.Column{{Name: "product_id"}}, DoUpdates: clause.AssignmentColumns(updateFields)}).
		Create(&params).WithContext(ctx).Error
	return
}

// DeleteProduct is a func to delete product
func DeleteProduct(ctx context.Context, id int) (err error) {
	data := &models.Product{}
	query := config.DBGorm.Table("products").
		WithContext(ctx)

	if id != 0 {
		err = query.Where("product_id = ?", id).Delete(data).Error
	}

	return
}
