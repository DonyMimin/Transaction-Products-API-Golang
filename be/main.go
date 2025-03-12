package main

import (
	"github.com/DonyMimin/Go-crud-echo/config"
	"github.com/DonyMimin/Go-crud-echo/controller"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// config.ConnectDB()
	config.ConnectGorm()

	// Products
	e.GET("/api/products", controller.GetListProduct)
	e.GET("/api/product/:id", controller.GetProduct)
	e.POST("/api/product", controller.UpsertProduct)
	e.DELETE("/api/product/:id", controller.DeleteProduct)

	// Company
	e.GET("/api/company", controller.GetListCompany)
	e.POST("/api/company", controller.UpsertCompany)
	e.DELETE("/api/company/:id", controller.DeleteCompany)

	// Transaction
	e.GET("/api/transaction", controller.GetListTransaction)
	e.POST("/api/transaction", controller.CreateTransaction)

	e.Logger.Fatal(e.Start(":8080"))
}
