package repository

import (
	"context"

	"github.com/DonyMimin/Go-crud-echo/config"
	"github.com/DonyMimin/Go-crud-echo/models"
	"gorm.io/gorm/clause"
)

// FindAllTransaction is a func to find list transaction data
func FindAllTransaction(ctx context.Context) ([]models.Transaction, error) {
	var transaction []models.Transaction
	err := config.DBGorm.WithContext(ctx).
		Raw("SELECT * FROM transaction").
		Scan(&transaction).Error

	if err != nil {
		return nil, err
	}
	return transaction, nil
}

// DeleteProduct is a func to update or insert transaction data
func UpsertTransaction(ctx context.Context, updateFields []string, params models.Transaction) (err error) {
	err = config.DBGorm.Table("transaction").
		Clauses(clause.OnConflict{Columns: []clause.Column{{Name: "transaction_id"}}, DoUpdates: clause.AssignmentColumns(updateFields)}).
		Create(&params).WithContext(ctx).Error
	return
}
