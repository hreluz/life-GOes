package main

import (
	"github.com/hreluz/go-db-gorm/model"
	"github.com/hreluz/go-db-gorm/storage"
)

func main() {
	driver := storage.MySQL
	storage.New(driver)

	storage.DB().AutoMigrate(
		&model.Product{},
		&model.InvoiceHeader{},
		&model.InvoiceItem{},
	)
}
