package controller

import (
	"net/http"
	"strconv"

	"github.com/DonyMimin/Go-crud-echo/models"
	"github.com/DonyMimin/Go-crud-echo/services"
	"github.com/DonyMimin/Go-crud-echo/utils"
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

// ExportProductPDF generates a PDF report of products
func ExportProductPDF(c echo.Context) error {
	ctx := c.Request().Context()

	// Get data from service
	data, err := services.GetListProduct(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve data"})
	}

	// Convert data to [][]string for the PDF
	var pdfData [][]string
	for _, product := range data {
		pdfData = append(pdfData, []string{
			strconv.Itoa(product.ProductID),
			product.NamaProduct,
			strconv.Itoa(product.Price),
			strconv.Itoa(product.Stock),
		})
	}

	// Define headers
	headers := []string{"ID", "Product Name", "Price", "Stock"}

	// Generate PDF
	outputFile := "products.pdf"
	err = utils.GeneratePDF("Product List", headers, pdfData, outputFile)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to generate PDF"})
	}

	return c.File(outputFile)
}
