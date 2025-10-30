package httpimg_test

import (
	gofpdf "github.com/akkaraponph/ghostpdf"
	"github.com/akkaraponph/ghostpdf/contrib/httpimg"
	"github.com/akkaraponph/ghostpdf/internal/example"
)

func ExampleRegister() {
	pdf := gofpdf.New("L", "mm", "A4", "")
	pdf.SetFont("Helvetica", "", 12)
	pdf.SetFillColor(200, 200, 220)
	pdf.AddPage()

	url := "https://github.com/akkaraponph/ghostpdf/raw/main/image/logo_gofpdf.jpg?raw=true"
	httpimg.Register(pdf, url, "")
	pdf.Image(url, 15, 15, 267, 0, false, "", 0, "")
	fileStr := example.Filename("contrib_httpimg_Register")
	err := pdf.OutputFileAndClose(fileStr)
	example.Summary(err, fileStr)
	// Output:
	// Successfully generated ../../pdf/contrib_httpimg_Register.pdf
}
