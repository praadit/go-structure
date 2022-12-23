package services

// import (
// 	"bytes"
// 	"errors"
// 	"fmt"
// 	"log"
// 	"os"
// 	"strings"

// 	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
// 	"gitlab.com/m8851/pmo-echo-api/handlers/dto"
// 	"gitlab.com/m8851/pmo-echo-api/pdfs"
// 	"gitlab.com/m8851/pmo-echo-api/pdfsdata"
// )

// type PdfGeneratorInterface interface {
// 	SavePdf(group string, identifier string, filename string) (filePath string, err error)
// 	CreatePageFromHtml(html string, pageName string)
// 	CreatePageFromFile(path string, pageName string)
// 	NewPageBuffer()
// 	AppendPageBuffer(html string)
// 	ClosePageBuffer()
// }

// type PdfGenerator struct {
// 	pdfg    *wkhtmltopdf.PDFGenerator
// 	bufPage *bytes.Buffer
// }

// func GenerateNewPdf(orientation string) PdfGeneratorInterface {
// 	pdfg, err := wkhtmltopdf.NewPDFGenerator()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	pdfg.Dpi.Set(300)
// 	pdfg.PageSize.Set(wkhtmltopdf.PageSizeLetter)
// 	pdfg.Orientation.Set(orientation)
// 	pdfg.Grayscale.Set(false)
// 	pdfg.MarginBottom.Set(0)
// 	pdfg.MarginLeft.Set(5)
// 	pdfg.MarginRight.Set(5)
// 	pdfg.MarginTop.Set(10)

// 	return &PdfGenerator{
// 		pdfg: pdfg,
// 	}
// }
// func GenerateNewPdfSptjm(orientation string) PdfGeneratorInterface {
// 	pdfg, err := wkhtmltopdf.NewPDFGenerator()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	pdfg.Dpi.Set(300)
// 	pdfg.PageSize.Set(wkhtmltopdf.PageSizeLetter)
// 	pdfg.Orientation.Set(orientation)
// 	pdfg.Grayscale.Set(false)
// 	pdfg.MarginBottom.Set(0)
// 	pdfg.MarginLeft.Set(5)
// 	pdfg.MarginRight.Set(5)
// 	pdfg.MarginTop.Set(5)

// 	return &PdfGenerator{
// 		pdfg: pdfg,
// 	}
// }
// func (service *PdfGenerator) SavePdf(group string, identifier string, filename string) (filepath string, err error) {
// 	err = service.pdfg.Create()
// 	if err != nil {
// 		return
// 	}

// 	if group == "" {
// 		err = errors.New("Group name should be filled!")
// 		return
// 	}
// 	if identifier == "" {
// 		err = errors.New("Identifier should be filled by unique value!")
// 		return
// 	}
// 	if filename == "" {
// 		err = errors.New("Filename should be filled")
// 		return
// 	}

// 	err = os.MkdirAll(fmt.Sprintf("./tmp/%s/%s", group, identifier), os.ModePerm)
// 	if err != nil {
// 		return
// 	}

// 	filepath = fmt.Sprintf("./tmp/%s/%s/%s", group, identifier, filename)

// 	err = service.pdfg.WriteFile(filepath)
// 	if err != nil {
// 		return
// 	}

// 	if service.bufPage != nil {
// 		service.bufPage.Truncate(0)
// 	}

// 	return
// }

// func (service *PdfGenerator) NewPageBuffer() {
// 	service.bufPage = new(bytes.Buffer)
// }
// func (service *PdfGenerator) AppendPageBuffer(html string) {
// 	service.bufPage.Write([]byte(html))
// 	service.bufPage.WriteString(`<P style="page-break-before: always">`)
// }
// func (service *PdfGenerator) ClosePageBuffer() {
// 	service.pdfg.AddPage(wkhtmltopdf.NewPageReader(service.bufPage))
// 	// service.bufPage.Truncate(0)
// }

// func (service *PdfGenerator) CreatePageFromHtml(html string, pageName string) {
// 	page := wkhtmltopdf.NewPageReader(strings.NewReader(html))

// 	opt := wkhtmltopdf.NewPageOptions()
// 	opt.EnableLocalFileAccess.Set(true)

// 	page.PageOptions = opt
// 	page.FooterRight.Set(pageName)
// 	page.FooterFontSize.Set(10)

// 	service.appendPdfPage(page)
// }
// func (service *PdfGenerator) CreatePageFromFile(path string, pageName string) {
// 	page := wkhtmltopdf.NewPage(path)
// 	opt := wkhtmltopdf.NewPageOptions()
// 	opt.EnableLocalFileAccess.Set(true)

// 	page.PageOptions = opt
// 	page.FooterRight.Set(pageName)
// 	page.FooterFontSize.Set(10)

// 	service.appendPdfPage(page)
// }

// func (service *PdfGenerator) appendPdfPage(page wkhtmltopdf.PageProvider) {
// 	service.pdfg.AddPage(page)
// }

// func ExampleNewPDFGenerator() {
// 	newPdf := GenerateNewPdf(wkhtmltopdf.OrientationLandscape)

// 	htmlFiller := pdfs.HtmlFiller{}

// 	propString := htmlFiller.GenerateProposalA(dto.TemplateRubrikProposalA{
// 		AsalPt: "Test markites",
// 	})
// 	newPdf.CreatePageFromHtml(propString, "Penilaian Proposal B")

// 	newPdf.SavePdf("testing", "pdf", "pdf-html.pdf")

// 	fmt.Println("Done")
// }

// func ExampleImageToPdf() {
// 	newPdf := GenerateNewPdf(wkhtmltopdf.OrientationLandscape)

// 	htmlFiller := pdfs.HtmlFiller{}
// 	imgData := pdfsdata.ImageData{
// 		Image: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQjzFV1rJVagbTWb9uFUC5TB4JLuVoYJMpqxA&usqp=CAU",
// 	}
// 	imgString := htmlFiller.FillImage(imgData)
// 	newPdf.CreatePageFromHtml(imgString, "Lampiran KTP")

// 	newPdf.SavePdf("testing", "pdf", "pdf-image.pdf")

// 	fmt.Println("Done")
// }
