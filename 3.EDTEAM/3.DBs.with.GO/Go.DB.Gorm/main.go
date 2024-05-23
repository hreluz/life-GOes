package main

import (
	"fmt"

	"github.com/hreluz/go-db-gorm/model"
	"github.com/hreluz/go-db-gorm/storage"
)

func main() {
	driver := storage.Postgres
	storage.New(driver)

	migrateTables()
	// createProducts()
	readProducts()
}

func readProducts() {
	products := make([]model.Product, 0)
	storage.DB().Find(&products)

	for _, product := range products {
		fmt.Printf("%d - %s\n", product.ID, product.Name)
	}

}

func migrateTables() {
	storage.DB().AutoMigrate(
		&model.Product{},
		&model.InvoiceHeader{},
		&model.InvoiceItem{},
	)
}

func createProducts() {
	product1 := model.Product{
		Name:  "Go Course",
		Price: 88,
	}

	product2 := model.Product{
		Name:  "Node Course",
		Price: 98,
	}

	obs := "Testing with go"
	product3 := model.Product{
		Name:         "Testing Go",
		Price:        33,
		Observations: &obs,
	}

	storage.DB().Create(&product1)
	storage.DB().Create(&product2)
	storage.DB().Create(&product3)
}
