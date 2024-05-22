package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/hreluz/go-db/pkg/interaction"
	"github.com/hreluz/go-db/pkg/invoice"
	"github.com/hreluz/go-db/pkg/invoiceheader"
	"github.com/hreluz/go-db/pkg/invoiceitem"
	"github.com/hreluz/go-db/pkg/product"
	"github.com/hreluz/go-db/storage"
)

func main() {
	interaction.PrintGreeting()
	interaction.ShowAvailableDatabases()
	driver := interaction.GetDBChoice()
	storage.New(driver)

	// product
	productStorage, err := storage.DAOProduct(driver)

	if err != nil {
		log.Fatalf("DAO Product: %v", err)
	}

	serviceProduct := product.NewService(productStorage)

	// invoice header
	invoiceHeaderStorage, err := storage.DAOInvoiceHeader(driver)

	if err != nil {
		log.Fatalf("DAO InvoiceHeader: %v", err)
	}

	serviceInvoiceHeader := invoiceheader.NewService(invoiceHeaderStorage)

	// invoice item
	invoiceItemStorage, err := storage.DAOInvoiceItem(driver)

	if err != nil {
		log.Fatalf("DAO InvoiceItems: %v", err)
	}

	serviceInvoiceItems := invoiceitem.NewService(invoiceItemStorage)

	interaction.ShowAvailableActions()
	choice := interaction.GetActionChoice()

	switch choice {
	case storage.CREATE_PRODUCT:
		createProduct(serviceProduct)
	case storage.UPDATE_PRODUCT:
		updateProduct(serviceProduct)
	case storage.DELETE_PRODUCT:
		deleteProduct(serviceProduct)
	case storage.SHOW_ALL_PRODUCT:
		showProducts(serviceProduct)
	case storage.SHOW_BY_ID_PRODUCT:
		showProductById(serviceProduct)
	case storage.MIGRATE_ALL:
		if err := serviceProduct.Migrate(); err != nil {
			log.Fatalf("product.Migrate: %v", err)
		}

		if err := serviceInvoiceHeader.Migrate(); err != nil {
			log.Fatalf("header.Migrate: %v", err)
		}

		if err := serviceInvoiceItems.Migrate(); err != nil {
			log.Fatalf("item.Migrate: %v", err)
		}
	case storage.CREATE_INVOICE:
		invoiceStorage, err := storage.DAOInvoice(
			driver,
			invoiceHeaderStorage,
			invoiceItemStorage,
		)

		if err != nil {
			log.Fatalf("DAO Invoice: %v", err)
		}

		createInvoice(&invoiceStorage, serviceProduct)
	}
}

func createInvoice(invoiceStorage *invoice.Storage, serviceProuct *product.Service) {
	items := invoiceitem.Models{}
	clientName, _ := interaction.GetUserInput("Insert Client Name: ")
	fmt.Println("\n\nProduct Items")
	fmt.Println("-----------------------------------")
	showProducts(serviceProuct)

	for {
		fmt.Printf("Insert Product ID only: ")
		productId := interaction.GetInputNumber()
		if productId == 0 {
			continue
		}
		items = append(items, &invoiceitem.Model{ProductID: productId})
		exit, _ := interaction.GetUserInput("Enter (q) to stop inserting product ids: ")

		if exit == "q" {
			break
		}
	}

	m := &invoice.Model{
		Header: &invoiceheader.Model{
			Client: clientName,
		},
		Items: items,
	}

	serviceInvoice := invoice.NewService(*invoiceStorage)

	if err := serviceInvoice.Create(m); err != nil {
		log.Fatalf("invoice.Create: %v", err)
	}
}

func showProductById(serviceProduct *product.Service) {
	fmt.Printf("Insert Product ID: ")
	id := interaction.GetInputNumber()
	ms, err := serviceProduct.GetByID(id)
	switch {
	case errors.Is(err, sql.ErrNoRows):
		fmt.Println("There is no product with this id ")
	case err != nil:
		log.Fatalf("product.GetById: %v", err)
	default:
		fmt.Println(ms)
	}
}

func showProducts(serviceProduct *product.Service) {
	ms, err := serviceProduct.GetAll()

	if err != nil {
		log.Fatalf("product.GetAll: %v", err)
	}

	fmt.Println(ms)
}

func deleteProduct(serviceProduct *product.Service) {
	fmt.Printf("Insert Product ID: ")
	id := interaction.GetInputNumber()
	err := serviceProduct.Delete(id)

	if err != nil {
		log.Fatalf("product.Delete: %v", err)
	}
}

func updateProduct(serviceProduct *product.Service) {
	fmt.Printf("Insert Product ID: ")
	id := interaction.GetInputNumber()
	ms, err := serviceProduct.GetByID(id)
	switch {
	case errors.Is(err, sql.ErrNoRows):
		fmt.Println("There is no product with this id ")
	case err != nil:
		log.Fatalf("product.GetById: %v", err)
	default:
		fmt.Println(ms)
	}
	productName, _ := interaction.GetUserInput("Insert New Product Name: ")
	fmt.Printf("Insert New Product Price: ")
	productPrice := interaction.GetInputNumber()
	productObservation, _ := interaction.GetUserInput("Insert New Product Observation: ")

	m := &product.Model{
		ID:           id,
		Name:         productName,
		Observations: productObservation,
		Price:        int(productPrice),
	}

	if err := serviceProduct.Update(m); err != nil {
		log.Fatalf("Product updated: %v", err)

	}

	fmt.Println(m)
}

func createProduct(serviceProduct *product.Service) {
	productName, _ := interaction.GetUserInput("Insert Product Name: ")
	fmt.Printf("Insert Product Price: ")
	productPrice := interaction.GetInputNumber()
	productObservation, _ := interaction.GetUserInput("Insert Product Observation: ")

	m := &product.Model{
		Name:         productName,
		Price:        int(productPrice),
		Observations: productObservation,
	}

	if err := serviceProduct.Create(m); err != nil {
		log.Fatalf("product.Create: %v", err)
	}

	fmt.Printf("%v", m)
}
