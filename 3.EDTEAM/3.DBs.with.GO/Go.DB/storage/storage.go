package storage

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
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

		fmt.Print("Connected to postgres\n")

	})
}

func NewMySQLDB() {

	once.Do(func() {
		var err error
		userName := "mysqluser"
		password := "secret"
		database := "mysqldb"
		ip := "localhost"

		dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", userName, password, ip, database)
		driverName := "mysql"
		db, err = sql.Open(driverName, dataSourceName)

		if err != nil {
			log.Fatalf("can't open db: %v", err)
		}

		if err := db.Ping(); err != nil {
			log.Fatalf("can't do ping: %v", err)
		}

		fmt.Print("Connected to mysql\n")

	})
}

// Return only one instance of db
func Pool() *sql.DB {
	return db
}

// if it is not empty, it will set valid as true, and it will not set a null value in the db
func stringToNull(s string) sql.NullString {
	null := sql.NullString{String: s}

	if null.String != "" {
		null.Valid = true
	}

	return null
}

func timeToNull(t time.Time) sql.NullTime {
	null := sql.NullTime{Time: t}

	if !null.Time.IsZero() {
		null.Valid = true
	}

	return null
}
