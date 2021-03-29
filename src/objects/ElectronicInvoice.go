package objects

import (
	"fmt"
	"time"
)

type ElectronicInvoice struct {
	StartDate       time.Time
	DueDate         time.Time
	CompanyCod      string
	products        []Product
	ImportTaxes     float32
	ExportTaxes     float32
	TransportTaxes  float32
	NavigationTaxes float32
}

func (e *ElectronicInvoice) SendInvoice() {
	fmt.Printf("%v", e)
}
