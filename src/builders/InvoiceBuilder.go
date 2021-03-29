package builders

import (
	"time"

	"github.com/ctabares06/go-builder/src/objects"
)

type InvoiceBuilder interface {
	setProducts(products []*objects.Product)
	setDates(startDate time.Time, dueDate time.Time)
	setTaxes(taxes float32)
	getInvoice() electronicInvoice
}
