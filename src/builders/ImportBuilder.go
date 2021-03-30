package builders

import (
	"time"

	"github.com/ctabares06/go-builder/src/objects"
)

type ImportBuilder struct {
	invoice            objects.ElectronicInvoice
	CompanyCode        string
	StartDate, DueDate time.Time
	Products           []*objects.Product
	Taxes, ImportTaxes float32
}

func Reset(companyCode string, startDate, dueDate time.Time, products []*objects.Product, taxes, importTaxes float32) *ImportBuilder {
	return &ImportBuilder{
		CompanyCode: companyCode,
		StartDate:   startDate,
		DueDate:     dueDate,
		Products:    products,
		Taxes:       taxes,
		ImportTaxes: importTaxes,
	}
}

func (i *ImportBuilder) SetCompany() {
	i.invoice.CompanyCod = i.CompanyCode
}

func (i *ImportBuilder) SetDates() {
	i.invoice.StartDate = i.StartDate
	i.invoice.DueDate = i.DueDate
}

func (i *ImportBuilder) SetProducts() {
	i.invoice.Products = i.Products
}

func (i *ImportBuilder) SetTaxes() {
	i.invoice.Taxes = i.Taxes
	i.invoice.ImportTaxes = i.ImportTaxes

}

func (i *ImportBuilder) GetInvoice() objects.ElectronicInvoice {
	return i.invoice
}
