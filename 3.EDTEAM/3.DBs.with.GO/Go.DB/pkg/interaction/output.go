package interaction

import (
	"fmt"

	"github.com/common-nighthawk/go-figure"
	"github.com/hreluz/go-db/storage"
)

func PrintGreeting() {
	asciiFigure := figure.NewFigure("MySQL and PostgresDB app", "", true)
	asciiFigure.Print()
}

func ShowAvailableDatabases() {
	fmt.Println("Choose database to interact with: ")
	fmt.Println("-------------------------------")
	fmt.Println("(1) MySQL")
	fmt.Println("(2) PostgresDB")
	fmt.Println("(3) EXIT")
}

func ShowAvailableActions() {
	fmt.Println("Choose an action: ")
	fmt.Println("-------------------------------")
	fmt.Printf("(1) %s\n", storage.CREATE_PRODUCT)
	fmt.Printf("(2) %s\n", storage.UPDATE_PRODUCT)
	fmt.Printf("(3) %s\n", storage.DELETE_PRODUCT)
	fmt.Printf("(4) %s\n", storage.SHOW_ALL_PRODUCT)
	fmt.Printf("(5) %s\n", storage.SHOW_BY_ID_PRODUCT)
	fmt.Printf("(6) %s\n", storage.CREATE_INVOICE)
	fmt.Printf("(7) %s\n", storage.MIGRATE_ALL)
}
