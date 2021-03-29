package builders

type InvoiceBuilder interface {
	Reset()
	SetProducts()
	SetDates()
	SetTaxes()
}
