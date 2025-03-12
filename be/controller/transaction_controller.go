package controller

import (
	"net/http"
	"strconv"

	"github.com/DonyMimin/Go-crud-echo/models"
	"github.com/DonyMimin/Go-crud-echo/services"
	"github.com/DonyMimin/Go-crud-echo/utils"
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

// ExportTransactionPDF generates a PDF report of transactions
func ExportTransactionPDF(c echo.Context) error {
	ctx := c.Request().Context()

	// Get data from service
	data, err := services.GetListTransaction(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve data"})
	}

	// Convert data to [][]string for the PDF
	var pdfData [][]string
	for _, transaction := range data {
		pdfData = append(pdfData, []string{
			strconv.Itoa(transaction.TransactionID),
			transaction.ProductName,
			strconv.Itoa(transaction.Quantity),
			strconv.Itoa(transaction.TotalPrice),
			transaction.CreatedAt,
		})
	}

	// Define headers
	headers := []string{"Transaction ID", "Product Name", "Quantity", "Total Price", "Date"}

	// Generate PDF
	outputFile := "transactions.pdf"
	err = utils.GeneratePDF("Transaction List", headers, pdfData, outputFile)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to generate PDF"})
	}

	return c.File(outputFile)
}
