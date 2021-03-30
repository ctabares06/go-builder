package main

import (
	"fmt"
	"time"

	"github.com/ctabares06/go-builder/src"
	"github.com/ctabares06/go-builder/src/builders"
	"github.com/ctabares06/go-builder/src/objects"
)

func main() {
	fmt.Println("project initialized")

	singleProduct := objects.Product{
		Name:     "producto random",
		Price:    2000.32,
		Quantity: 10,
		Iva:      19,
		ValIva:   200.2,
		Impo:     0.0,
	}

	products := []*objects.Product{
		&singleProduct,
	}

	importBuilder := &builders.ImportBuilder{
		CompanyCode: "g12b124612622dsff",
		StartDate:   time.Time{},
		DueDate:     time.Time{},
		Products:    products,
		Taxes:       5000.2,
		ImportTaxes: 4000.4,
	}

	var director src.InvoiceDirector

	director.SetBuilder(importBuilder)

	importInvoice := director.BuildInvoice()

	importInvoice.SendInvoice()

}
