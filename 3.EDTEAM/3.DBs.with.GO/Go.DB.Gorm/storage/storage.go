package storage

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	once sync.Once
)

// Drivers
type Driver string

const (
	MySQL    Driver = "MYSQL"
	Postgres Driver = "POSTGRES"
)

type MENU_OPTIONS string

const (
	CREATE_PRODUCT     MENU_OPTIONS = "CREATE PRODUCT"
	UPDATE_PRODUCT     MENU_OPTIONS = "UPDATE PRODUCT"
	DELETE_PRODUCT     MENU_OPTIONS = "DELETE PRODUCT"
	SHOW_BY_ID_PRODUCT MENU_OPTIONS = "SHOW PRODUCT BY ID"
	SHOW_ALL_PRODUCT   MENU_OPTIONS = "SHOW ALL PRODUCTS"
	CREATE_INVOICE     MENU_OPTIONS = "CREATE INVOICE"
	SHOW_INVOICES      MENU_OPTIONS = "SHOW INVOICES"
	MIGRATE_ALL        MENU_OPTIONS = "DO MIGRATIONS"
)

// New create the conection with db
func New(d Driver) {
	switch d {
	case MySQL:
		newMySQLDB()
	case Postgres:
		newPostgresDB()
	default:
		log.Fatalf("Connection not implemented")
	}
}

func newPostgresDB() {
	once.Do(func() {
		var err error
		userName := "pguser"
		password := "pgpasswd"
		database := "pgdb"

		dsn := fmt.Sprintf("postgres://%s:%s@localhost/%s?sslmode=disable", userName, password, database)
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

		if err != nil {
			log.Fatalf("can't open db: %v", err)
		}

		fmt.Print("Connected to postgres\n")
	})
}

func newMySQLDB() {

	once.Do(func() {
		var err error
		userName := "mysqluser"
		password := "secret"
		database := "mysqldb"
		ip := "localhost"

		dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true", userName, password, ip, database)
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

		if err != nil {
			log.Fatalf("can't open db: %v", err)
		}

		fmt.Print("Connected to mysql\n")

	})
}

// Return only one instance of db
func DB() *gorm.DB {
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
