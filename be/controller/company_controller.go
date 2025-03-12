package controller

import (
	"net/http"
	"strconv"

	"github.com/DonyMimin/Go-crud-echo/models"
	"github.com/DonyMimin/Go-crud-echo/services"
	"github.com/labstack/echo"
)

func GetListCompany(c echo.Context) error {
	var (
		data []models.Company
		ctx  = c.Request().Context()
		err  error
	)

	data, err = services.GetListCompany(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, data)
	}
	return c.JSON(http.StatusOK, data)
}

func UpsertCompany(c echo.Context) error {
	var (
		params models.Company
		ctx    = c.Request().Context()
		err    error
	)

	if err := c.Bind(&params); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request data",
		})
	}

	err = services.UpsertCompany(ctx, params)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "error")
	}
	return c.JSON(http.StatusOK, "success")
}

func DeleteCompany(c echo.Context) error {
	var (
		params models.DeleteCompany
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

	params.CompanyID = id

	err = services.DeleteCompany(ctx, params.CompanyID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "error")
	}
	return c.JSON(http.StatusOK, "success")
}
