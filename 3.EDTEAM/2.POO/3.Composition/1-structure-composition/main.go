package main

import (
	"fmt"

	"github.com/hreluz/composition/pkg/invoice"
	"github.com/hreluz/composition/pkg/invoice/customer"
	"github.com/hreluz/composition/pkg/invoiceitem"
)

func main() {
	c := customer.New("Batman", "123 Arkham", "555555")
	items := invoiceitem.NewItems(
		invoiceitem.New(1, "Go Course", 12.99),
		invoiceitem.New(2, "POO Go", 15.11),
		invoiceitem.New(3, "Testing with Go", 19.22),
	)

	i := invoice.New(
		"Peru",
		"Lima",
		c,
		items,
	)

	i.SetTotal()
	fmt.Printf("%+v", i)

}
