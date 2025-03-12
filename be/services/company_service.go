package services

import (
	"context"

	"github.com/DonyMimin/Go-crud-echo/models"
	"github.com/DonyMimin/Go-crud-echo/repository"
)

func GetListCompany(ctx context.Context) (data []models.Company, err error) {
	// Get list data company
	data, err = repository.FindAllCompany(ctx)
	if err != nil {
		return
	}

	return
}

func UpsertCompany(ctx context.Context, params models.Company) (err error) {
	// Update and Insert company
	err = repository.UpsertCompany(ctx, []string{"company_id", "company_name"}, models.Company{
		CompanyID:   params.CompanyID,
		CompanyName: params.CompanyName,
	})
	if err != nil {
		return
	}

	return
}

func DeleteCompany(ctx context.Context, id int) (err error) {
	// Delete product
	err = repository.DeleteCompany(ctx, id)
	if err != nil {
		return
	}

	return
}
