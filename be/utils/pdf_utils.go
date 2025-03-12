package utils

import "github.com/jung-kurt/gofpdf"

// GeneratePDF is a reusable function to create a PDF report
func GeneratePDF(title string, headers []string, data [][]string, outputFilename string) error {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetFont("Arial", "", 12)
	pdf.AddPage()
	pageWidth, _ := pdf.GetPageSize()
	marginLeft := 10.0
	marginRight := 10.0
	tableWidth := pageWidth - marginLeft - marginRight
	colWidth := tableWidth / float64(len(headers)) // Auto column width

	// Title
	pdf.SetFont("Arial", "B", 14)
	pdf.CellFormat(190, 10, title, "", 1, "C", false, 0, "")
	pdf.Ln(5)

	// Table Header
	pdf.SetFont("Arial", "B", 12)
	for _, header := range headers {
		pdf.CellFormat(colWidth, 7, header, "1", 0, "C", false, 0, "")
	}
	pdf.Ln(-1)

	/// Table Content
	pdf.SetFont("Arial", "", 10)
	for _, row := range data {
		x := pdf.GetX() // Simpan posisi X awal
		y := pdf.GetY() // Simpan posisi Y awal

		for _, col := range row {
			pdf.Rect(x, y, colWidth, 7, "D") // Buat border sel
			pdf.CellFormat(colWidth, 7, col, "", 0, "L", false, 0, "")
			x += colWidth   // Geser ke kolom berikutnya
			pdf.SetXY(x, y) // Set posisi X dan Y untuk kolom berikutnya
		}

		pdf.Ln(7) // Pindah ke baris berikutnya
	}

	// Output PDF
	return pdf.OutputFileAndClose(outputFilename)
}
