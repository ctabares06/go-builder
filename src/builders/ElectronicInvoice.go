package builders

import (
	"fmt"
	"time"

	"github.com/ctabares06/go-builder/src/objects"
)

type electronicInvoice struct {
	startDate       time.Time
	dueDate         time.Time
	companyCod      string
	products        []*objects.Product
	taxes           float32
	importTaxes     float32
	exportTaxes     float32
	transportTaxes  float32
	navigationTaxes float32
}

func (e *electronicInvoice) sendInvoice() {
	fmt.Printf("%v", e)
}
