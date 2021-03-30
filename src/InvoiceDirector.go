package src

import (
	"github.com/ctabares06/go-builder/src/builders"
	"github.com/ctabares06/go-builder/src/objects"
)

type InvoiceDirector struct {
	builder builders.InvoiceBuilder
}

func (i *InvoiceDirector) SetBuilder(builder builders.InvoiceBuilder) {
	i.builder = builder
}

func (i *InvoiceDirector) BuildInvoice() objects.ElectronicInvoice {
	i.builder.SetCompany()
	i.builder.SetDates()
	i.builder.SetProducts()
	i.builder.SetTaxes()
	return i.builder.GetInvoice()
}
