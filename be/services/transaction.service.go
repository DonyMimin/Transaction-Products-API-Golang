package services

import (
	"context"
	"time"

	"github.com/DonyMimin/Go-crud-echo/models"
	"github.com/DonyMimin/Go-crud-echo/repository"
	"github.com/DonyMimin/Go-crud-echo/utils"
)

func GetListTransaction(ctx context.Context) (data []models.Transaction, err error) {
	// Get list data transaction
	data, err = repository.FindAllTransaction(ctx)
	if err != nil {
		return
	}

	return
}

func UpsertTransaction(ctx context.Context, params models.UpsertTransaction) (err error) {
	// Update and Insert transaction
	// Find product data first
	productDetail := models.Product{}
	productDetail, err = repository.FindProduct(ctx, params.ProductID)
	if err != nil {
		return
	}

	if productDetail.ProductID == 0 {
		return
	}

	// Find company data
	companyDetail := models.Company{}
	companyDetail, err = repository.FindCompany(ctx, params.CompanyID)
	if err != nil {
		return
	}

	if companyDetail.CompanyID == 0 {
		return
	}

	// Find transaction data
	transactionDetail := []models.Transaction{}
	transactionDetail, err = repository.FindAllTransaction(ctx)
	if err != nil {
		return
	}

	// Check if transaction is already done with the same product and company
	for _, buyedItemCompany := range transactionDetail {
		if buyedItemCompany.CompanyName == companyDetail.CompanyName && buyedItemCompany.ProductName == productDetail.NamaProduct {
			// Update the remaining stock and update quantity before + now
			err = repository.UpsertProduct(ctx, []string{"product_id", "stock"}, models.Product{
				ProductID: params.ProductID,
				Stock:     productDetail.Stock - params.Quantity,
			})
			if err != nil {
				return
			}

			params.TransactionID = buyedItemCompany.TransactionID
			params.Quantity = params.Quantity + buyedItemCompany.Quantity
		}
	}

	// Get current local time
	currentTime, err := utils.GetCurrentLocalTime()
	if err != nil {
		return
	}

	// Update or Insert the Transaction
	totalPrice := params.Quantity * productDetail.Price
	err = repository.UpsertTransaction(ctx, []string{"transaction_id", "product_name", "company_name", "created_at", "quantity", "total_price"}, models.Transaction{
		TransactionID: params.TransactionID,
		ProductName:   productDetail.NamaProduct,
		CompanyName:   companyDetail.CompanyName,
		CreatedAt:     currentTime.Format(time.DateTime),
		Quantity:      params.Quantity,
		TotalPrice:    totalPrice,
	})
	if err != nil {
		return
	}

	return
}
