package src

import "github.com/ctabares06/go-builder/src/builders"

type InvoiceDirector struct {
	builder *builders.InvoiceBuilder
}

func (i *InvoiceDirector) SetBuilder(builder *builders.InvoiceBuilder) {
	i.builder = builder
}

func (i *InvoiceDirector) BuildInvoice() {

}
