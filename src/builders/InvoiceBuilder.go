package builders

import (
	"github.com/ctabares06/go-builder/src/objects"
)

type InvoiceBuilder interface {
	SetCompany()
	SetProducts()
	SetDates()
	SetTaxes()
	GetInvoice() objects.ElectronicInvoice
}
