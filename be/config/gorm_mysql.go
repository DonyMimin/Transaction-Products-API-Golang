package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DBGorm *gorm.DB

func ConnectGorm() error {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:@tcp(localhost:3306)/go_restapi_gin"
	// var db *gorm.DB
	var err error
	DBGorm, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return err
}
