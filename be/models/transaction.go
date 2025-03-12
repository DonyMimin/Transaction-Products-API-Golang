package models

type Transaction struct {
	TransactionID int    `json:"transaction_id"`
	CompanyName   string `json:"company_name"`
	ProductName   string `json:"product_name"`
	CreatedAt     string `json:"created_at"`
	Quantity      int    `json:"quantity"`
	TotalPrice    int    `json:"total_price"`
}

type UpsertTransaction struct {
	TransactionID int `json:"transaction_id"`
	ProductID     int `json:"product_id"`
	CompanyID     int `json:"company_id"`
	Quantity      int `json:"quantity"`
}

type GetTransaction struct {
	TransactionID int `gorm:"primaryKey" json:"transaction_id" param:"transaction_id"`
}
