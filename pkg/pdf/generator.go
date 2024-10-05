package pdf

import (
	"fmt"

	"github.com/jung-kurt/gofpdf"
	"github.com/piotrostr/essadev/faktury/pkg/config"
	"github.com/piotrostr/essadev/faktury/pkg/invoice"
)

type Generator struct {
	config *config.Config
}

func NewGenerator(cfg *config.Config) *Generator {
	return &Generator{config: cfg}
}

func (g *Generator) GenerateInvoice(project *invoice.Project) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddUTF8Font("SF", "", "SFMono.ttf")
	pdf.SetFont("SF", "", 12)

	pdf.AddPage()

	// Set margins
	pdf.SetMargins(20, 20, 20)
	pdf.SetAutoPageBreak(true, 20)

	// Helper function for MultiCell
	writeMultiCell := func(text string) {
		pdf.MultiCell(0, 6, text, "", "", false)
		pdf.Ln(2)
	}

	pdf.SetFont("SF", "", 16)
	writeMultiCell("Faktura")
	pdf.Ln(5)

	pdf.SetFont("SF", "", 12)

	pdf.SetFont("SF", "U", 13)
	writeMultiCell("Sprzedawca:")
	pdf.SetFont("SF", "", 12)
	writeMultiCell(g.config.CompanyName)
	writeMultiCell(fmt.Sprintf("NIP: %s", g.config.NIP))
	writeMultiCell(fmt.Sprintf("REGON: %s", g.config.REGON))
	writeMultiCell(fmt.Sprintf("Email: %s", g.config.Email))
	writeMultiCell(fmt.Sprintf("Telefon: %s", g.config.Phone))
	pdf.Ln(5)

	pdf.SetFont("SF", "U", 13)
	writeMultiCell("Nabywca:")
	pdf.SetFont("SF", "", 12)
	writeMultiCell(project.Client.Name)
	writeMultiCell(fmt.Sprintf("Adres: %s", project.Client.Address))
	writeMultiCell(fmt.Sprintf("NIP: %s", project.Client.NIP))
	pdf.Ln(5)

	pdf.SetFont("SF", "U", 13)
	writeMultiCell("Projekt:")
	pdf.SetFont("SF", "", 12)
	writeMultiCell(project.Title)
	writeMultiCell(fmt.Sprintf("Koszt: %.2f PLN", project.Cost))
	writeMultiCell(fmt.Sprintf("Deliverable: %s", project.Deliverable))
	pdf.Ln(5)

	err := pdf.OutputFileAndClose("invoice.pdf")
	if err != nil {
		panic(err)
	}

	fmt.Println("Invoice generated: invoice.pdf")
}
