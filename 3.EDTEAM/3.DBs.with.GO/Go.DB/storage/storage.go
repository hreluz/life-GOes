package storage

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/hreluz/go-db/pkg/invoice"
	"github.com/hreluz/go-db/pkg/invoiceheader"
	"github.com/hreluz/go-db/pkg/invoiceitem"
	"github.com/hreluz/go-db/pkg/product"
	_ "github.com/lib/pq"
)

var (
	db   *sql.DB
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
	}
}

func DAOInvoice(driver Driver, ih invoiceheader.Storage, ii invoiceitem.Storage) (invoice.Storage, error) {
	switch driver {
	case Postgres:
		return newPsqlInvoice(db, ih, ii), nil
	case MySQL:
		return newMySQLInvoice(db, ih, ii), nil
	default:
		return nil, fmt.Errorf("Driver not implemented")
	}
}

func DAOInvoiceHeader(driver Driver) (invoiceheader.Storage, error) {
	switch driver {
	case Postgres:
		return newPsqlInvoiceHeader(db), nil
	case MySQL:
		return newMySQLInvoiceHeader(db), nil
	default:
		return nil, fmt.Errorf("Driver not implemented")
	}
}

func DAOInvoiceItem(driver Driver) (invoiceitem.Storage, error) {
	switch driver {
	case Postgres:
		return newPsqlInvoiceItem(db), nil
	case MySQL:
		return newMySQLInvoiceItem(db), nil
	default:
		return nil, fmt.Errorf("Driver not implemented")
	}
}

func DAOProduct(driver Driver) (product.Storage, error) {
	switch driver {
	case Postgres:
		return newPsqlProduct(db), nil
	case MySQL:
		return newMySQLProduct(db), nil
	default:
		return nil, fmt.Errorf("Driver not implemented")
	}
}

func newPostgresDB() {
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

func newMySQLDB() {

	once.Do(func() {
		var err error
		userName := "mysqluser"
		password := "secret"
		database := "mysqldb"
		ip := "localhost"

		dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true", userName, password, ip, database)
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
