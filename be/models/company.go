package models

type Company struct {
	CompanyID   int    `gorm:"primaryKey" json:"company_id"`
	CompanyName string `gorm:"type:varchar(200)" json:"company_name"`
}

type DeleteCompany struct {
	CompanyID int `gorm:"primaryKey" json:"company_id" param:"company_id"`
}
