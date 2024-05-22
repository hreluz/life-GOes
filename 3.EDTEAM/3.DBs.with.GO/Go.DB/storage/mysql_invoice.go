package storage

import (
	"database/sql"
	"fmt"

	"github.com/hreluz/go-db/pkg/invoice"
	"github.com/hreluz/go-db/pkg/invoiceheader"
	"github.com/hreluz/go-db/pkg/invoiceitem"
)

// MySQLInvoice used for work with mysql - invoice
type MySQLInvoice struct {
	db            *sql.DB
	storageHeader invoiceheader.Storage
	storageItems  invoiceitem.Storage
}

// NewMySQLInvoice return a new pointer of MySQLInvoice
func NewMySQLInvoice(db *sql.DB, h invoiceheader.Storage, i invoiceitem.Storage) *MySQLInvoice {
	return &MySQLInvoice{
		db:            db,
		storageHeader: h,
		storageItems:  i,
	}
}

// Create implement the interface invoice.Storage
func (p *MySQLInvoice) Create(m *invoice.Model) error {
	tx, err := p.db.Begin()

	if err != nil {
		return err
	}

	if err := p.storageHeader.CreateTx(tx, m.Header); err != nil {
		tx.Rollback()
		return fmt.Errorf("Header: %w", err)
	}

	fmt.Printf("Invoice created with id: %d \n", m.Header.ID)

	if err := p.storageItems.CreateTX(tx, m.Header.ID, m.Items); err != nil {
		tx.Rollback()
		return fmt.Errorf("Items: %w", err)
	}

	fmt.Printf("Invoice items created: %d \n", len(m.Items))

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("Commit: %w", err)
	}

	return nil
}
