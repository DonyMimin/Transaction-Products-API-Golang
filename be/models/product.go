package models

type Product struct {
	ProductID   int    `gorm:"primaryKey" json:"product_id"`
	NamaProduct string `gorm:"type:varchar(300)" json:"nama_product"`
	Price       int    `gorm:"type:int" json:"price"`
	Stock       int    `gorm:"type:int" json:"stock"`
}

type GetProduct struct {
	ProductID int `gorm:"primaryKey" json:"product_id" param:"product_id"`
}

type DeleteProduct struct {
	ProductID int `gorm:"primaryKey" json:"product_id" param:"product_id"`
}
