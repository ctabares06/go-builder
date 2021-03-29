package builders

import (
	"time"

	"github.com/ctabares06/go-builder/src/objects"
)

type importBuilder struct {
	startDate   time.Time
	dueDate     time.Time
	companyCod  string
	products    []*objects.Product
	taxes       float32
	importTaxes float32
}

func reset() *importBuilder {
	return &importBuilder{}
}

func (i *importBuilder) setCompany(companyCode string) {
	i.companyCod = companyCode
}

func (i *importBuilder) setDates(startDate time.Time, dueDate time.Time) {
	i.startDate = startDate
	i.dueDate = dueDate
}

func (i *importBuilder) setProducts(products []*objects.Product) {
	i.products = products
}

func (i *importBuilder) setTaxes(taxes float32) {
	i.taxes = taxes
}

func (i *importBuilder) setImportTaxes(importTaxes float32) {
	i.importTaxes = importTaxes
}

func (i *importBuilder) getInvoice() electronicInvoice {

	return electronicInvoice{
		startDate:   i.startDate,
		dueDate:     i.dueDate,
		companyCod:  i.companyCod,
		products:    i.products,
		taxes:       i.taxes,
		importTaxes: i.importTaxes,
	}
}
