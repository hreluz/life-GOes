package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/hreluz/go-db/pkg/invoice"
	"github.com/hreluz/go-db/pkg/invoiceheader"
	"github.com/hreluz/go-db/pkg/invoiceitem"
	"github.com/hreluz/go-db/pkg/product"
	"github.com/hreluz/go-db/storage"
)

func main() {
	// postgresDB()
	mysqlDB()
}

func mysqlDB() {
	storage.NewMySQLDB()
	storage.NewMySQLDB()
}

func postgresDB() {
	storage.NewPostgresDB()
	// migrate()
	// createProduct()
	// getProducts()
	// getProductById()
	// updateProduct()
	// deleteProduct()

	createStorageHeaderAndItem()
}

func createStorageHeaderAndItem() {
	storage.NewPostgresDB()

	storageHeader := storage.NewPsqlInvoiceHeader(storage.Pool())
	storageItems := storage.NewPsqlInvoiceItem(storage.Pool())

	storageInvoice := storage.NewPsqlInvoice(
		storage.Pool(),
		storageHeader,
		storageItems,
	)

	m := &invoice.Model{
		Header: &invoiceheader.Model{
			Client: "James Bond 2",
		},
		Items: invoiceitem.Models{
			&invoiceitem.Model{ProductID: 1},
			&invoiceitem.Model{ProductID: 2},
		},
	}

	serviceInvoice := invoice.NewService(storageInvoice)

	if err := serviceInvoice.Create(m); err != nil {
		log.Fatalf("invoice.Create: %v", err)
	}

}

func deleteProduct() {
	storageProduct := storage.NewPsqlProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)

	err := serviceProduct.Delete(4)
	if err != nil {
		log.Fatalf("product.Delete: %v", err)
	}

	getProducts()
}

func updateProduct() {
	storageProduct := storage.NewPsqlProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)

	m := &product.Model{
		ID:           222,
		Name:         "Currency Go",
		Observations: "This is the go course",
		Price:        33,
	}

	err := serviceProduct.Update(m)

	if err != nil {
		log.Fatalf("Product updated: %v", err)

	}

	getProducts()
}

func getProductById() {
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
