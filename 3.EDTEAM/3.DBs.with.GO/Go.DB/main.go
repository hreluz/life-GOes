package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/hreluz/go-db/pkg/invoiceheader"
	"github.com/hreluz/go-db/pkg/invoiceitem"
	"github.com/hreluz/go-db/pkg/product"
	"github.com/hreluz/go-db/storage"
)

func main() {
	storage.NewPostgresDB()
	// migrate()
	// createProduct()
	// getProducts()

	storageProduct := storage.NewPsqlProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)

	ms, err := serviceProduct.GetByID(1)
	switch {
	case errors.Is(err, sql.ErrNoRows):
		fmt.Println("There is no product with this id ")
	case err != nil:
		log.Fatalf("product.GetById: %v", err)
	default:
		fmt.Println(ms)
	}
}

func getProducts() {
	storageProduct := storage.NewPsqlProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)

	ms, err := serviceProduct.GetAll()

	if err != nil {
		log.Fatalf("product.GetAll: %v", err)
	}

	fmt.Println(ms)
}

func createProduct() {
	storageProduct := storage.NewPsqlProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)

	m := &product.Model{
		Name:         "Go Course",
		Price:        50,
		Observations: "almost done",
	}

	if err := serviceProduct.Create(m); err != nil {
		log.Fatalf("product.Create: %v", err)
	}

	fmt.Printf("%v", m)
}

func migrate() {

	storageProduct := storage.NewPsqlProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)

	if err := serviceProduct.Migrate(); err != nil {
		log.Fatalf("product.Migrate: %v", err)
	}

	storageInvoiceHeader := storage.NewPsqlInvoiceHeader(storage.Pool())
	serviceInvoiceHeader := invoiceheader.NewService(storageInvoiceHeader)

	if err := serviceInvoiceHeader.Migrate(); err != nil {
		log.Fatalf("invoiceHeader.Migrate: %v", err)
	}

	storageInvoiceItem := storage.NewPsqlInvoiceItem(storage.Pool())
	serviceInvoiceItem := invoiceitem.NewService(storageInvoiceItem)

	if err := serviceInvoiceItem.Migrate(); err != nil {
		log.Fatalf("invoiceItem.Migrate: %v", err)
	}

}
