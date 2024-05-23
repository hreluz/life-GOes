package main

import (
	"fmt"

	"github.com/hreluz/go-db-gorm/model"
	"github.com/hreluz/go-db-gorm/storage"
)

func main() {
	driver := storage.MySQL
	storage.New(driver)

	migrateTables()
	// createProducts()
	// readProducts()
	// readProduct()
	// updateProduct()
	// softDeleteProduct()
	forceDeleteProduct()
}

func forceDeleteProduct() {
	product := model.Product{}
	product.ID = 3

	storage.DB().Unscoped().Delete(&product)
}

func softDeleteProduct() {
	product := model.Product{}
	product.ID = 3

	storage.DB().Delete(&product)

	readProducts()
}

func updateProduct() {
	product := model.Product{}
	product.ID = 3

	storage.DB().Model(&product).Updates(
		model.Product{Name: "CSS course", Price: 777},
	)

	readProducts()
}

func readProduct() {
	product := model.Product{}
	storage.DB().First(&product, 2)
	fmt.Println(product)
}

func readProducts() {
	products := make([]model.Product, 0)
	storage.DB().Find(&products)

	for _, product := range products {
		fmt.Printf("%d - %s - %d\n", product.ID, product.Name, product.Price)
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
