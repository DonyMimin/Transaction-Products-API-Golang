package controller

import (
	"net/http"

	"github.com/DonyMimin/Go-crud-echo/models"
	"github.com/DonyMimin/Go-crud-echo/services"
	"github.com/labstack/echo"
)

func GetListTransaction(c echo.Context) error {
	var (
		data []models.Transaction
		ctx  = c.Request().Context()
		err  error
	)

	data, err = services.GetListTransaction(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, data)
	}
	return c.JSON(http.StatusOK, data)
}

func CreateTransaction(c echo.Context) error {
	var (
		params models.UpsertTransaction
		ctx    = c.Request().Context()
		err    error
	)

	if err := c.Bind(&params); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request data",
		})
	}

	err = services.UpsertTransaction(ctx, params)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "error")
	}
	return c.JSON(http.StatusOK, "success")
}
