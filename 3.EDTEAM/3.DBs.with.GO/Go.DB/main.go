package main

import (
	"github.com/hreluz/go-db/storage"
)

func main() {
	storage.NewPostgresDB()
	storage.NewPostgresDB()
	storage.NewPostgresDB()
}
