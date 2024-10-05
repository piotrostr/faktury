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
	pdf.AddPage()

	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Faktura")

	pdf.SetFont("Arial", "", 12)
	pdf.Ln(10)
	pdf.Cell(0, 6, fmt.Sprintf("Sprzedawca: %s", g.config.CompanyName))
	pdf.Ln(6)
	pdf.Cell(0, 6, fmt.Sprintf("NIP: %s", g.config.NIP))
	pdf.Ln(6)
	pdf.Cell(0, 6, fmt.Sprintf("REGON: %s", g.config.REGON))
	pdf.Ln(6)
	pdf.Cell(0, 6, fmt.Sprintf("Email: %s", g.config.Email))
	pdf.Ln(6)
	pdf.Cell(0, 6, fmt.Sprintf("Telefon: %s", g.config.Phone))

	pdf.Ln(10)
	pdf.Cell(0, 6, fmt.Sprintf("Nabywca: %s", project.Client.Name))
	pdf.Ln(6)
	pdf.Cell(0, 6, fmt.Sprintf("Adres: %s", project.Client.Address))
	pdf.Ln(6)
	pdf.Cell(0, 6, fmt.Sprintf("NIP: %s", project.Client.NIP))

	pdf.Ln(10)
	pdf.Cell(0, 6, fmt.Sprintf("Tytuł projektu: %s", project.Title))
	pdf.Ln(6)
	pdf.Cell(0, 6, fmt.Sprintf("Koszt: %.2f PLN", project.Cost))
	pdf.Ln(6)
	pdf.Cell(0, 6, fmt.Sprintf("Deliverable: %s", project.Deliverable))

	pdf.Ln(10)
	pdf.Cell(0, 6, "Usługa kwalifikowana do IP Box")

	err := pdf.OutputFileAndClose("invoice.pdf")
	if err != nil {
		panic(err)
	}

	fmt.Println("Invoice generated: invoice.pdf")
}
