package storage

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/lib/pq"
)

var (
	db   *sql.DB
	once sync.Once
)

func NewPostgresDB() {
	once.Do(func() {
		var err error
		userName := "pguser"
		password := "pgpasswd"
		database := "pgdb"

		dataSourceName := fmt.Sprintf("postgres://%s:%s@localhost/%s?sslmode=disable", userName, password, database)
		driverName := "postgres"
		db, err = sql.Open(driverName, dataSourceName)

		if err != nil {
			log.Fatalf("can't open db: %v", err)
		}

		if err := db.Ping(); err != nil {
			log.Fatalf("can't do ping: %v", err)
		}

		fmt.Print("Connected to postgres")

	})
}

// Return only one instance of db
func Pool() *sql.DB {
	return db
}
