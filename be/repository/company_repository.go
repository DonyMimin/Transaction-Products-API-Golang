package repository

import (
	"context"

	"github.com/DonyMimin/Go-crud-echo/config"
	"github.com/DonyMimin/Go-crud-echo/models"
	"gorm.io/gorm/clause"
)

// FindAllCompany is a func to find all list company data
func FindAllCompany(ctx context.Context) ([]models.Company, error) {
	var company []models.Company
	err := config.DBGorm.WithContext(ctx).
		Raw("SELECT * FROM company").
		Scan(&company).Error

	if err != nil {
		return nil, err
	}
	return company, nil
}

// UpsertCompany is a func to update/insert company data
func UpsertCompany(ctx context.Context, updateFields []string, params models.Company) (err error) {
	err = config.DBGorm.Table("company").
		Clauses(clause.OnConflict{Columns: []clause.Column{{Name: "company_id"}}, DoUpdates: clause.AssignmentColumns(updateFields)}).
		Create(&params).WithContext(ctx).Error
	return
}

// FindCompany is a func to find company data
func FindCompany(ctx context.Context, id int) (data models.Company, err error) {
	query := config.DBGorm.Table("company").
		WithContext(ctx).
		Select(
			"company_id",
			"company_name",
		)

	if id != 0 {
		query.Where("company_id = ?", id)
	}

	err = query.Find(&data).Error
	return
}

// DeleteCompany is a func to delete company data
func DeleteCompany(ctx context.Context, id int) (err error) {
	data := &models.Company{}
	query := config.DBGorm.Table("company").
		WithContext(ctx)

	if id != 0 {
		err = query.Where("company_id = ?", id).Delete(data).Error
	}

	return
}
