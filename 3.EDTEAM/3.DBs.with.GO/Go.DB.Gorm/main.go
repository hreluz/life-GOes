package main

import (
	"github.com/hreluz/go-db-gorm/model"
	"github.com/hreluz/go-db-gorm/storage"
)

func main() {
	driver := storage.Postgres
	storage.New(driver)

	// Migrate
	storage.DB().AutoMigrate(
		&model.Product{},
		&model.InvoiceHeader{},
		&model.InvoiceItem{},
	)

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
