package controller

import (
	"net/http"
	"strconv"

	"github.com/DonyMimin/Go-crud-echo/models"
	"github.com/DonyMimin/Go-crud-echo/services"
	"github.com/labstack/echo"
)

func GetListProduct(c echo.Context) error {
	var (
		data []models.Product
		ctx  = c.Request().Context()
		err  error
	)

	data, err = services.GetListProduct(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, data)
	}
	return c.JSON(http.StatusOK, data)
}

func GetProduct(c echo.Context) error {
	var (
		params models.GetProduct
		data   models.Product
		ctx    = c.Request().Context()
		err    error
	)

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid product ID",
		})
	}

	params.ProductID = id

	data, err = services.GetProduct(ctx, params.ProductID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, data)
	}
	return c.JSON(http.StatusOK, data)
}

func UpsertProduct(c echo.Context) error {
	var (
		params models.Product
		ctx    = c.Request().Context()
		err    error
	)

	if err := c.Bind(&params); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request data",
		})
	}

	err = services.UpsertProduct(ctx, params)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "error")
	}
	return c.JSON(http.StatusOK, "success")
}

func DeleteProduct(c echo.Context) error {
	var (
		params models.DeleteProduct
		ctx    = c.Request().Context()
		err    error
	)

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid product ID",
		})
	}

	params.ProductID = id

	err = services.DeleteProduct(ctx, params.ProductID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "error")
	}
	return c.JSON(http.StatusOK, "success")
}
