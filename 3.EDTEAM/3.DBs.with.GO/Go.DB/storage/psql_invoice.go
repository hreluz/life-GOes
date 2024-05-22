package storage

import (
	"database/sql"
	"fmt"

	"github.com/hreluz/go-db/pkg/invoice"
	"github.com/hreluz/go-db/pkg/invoiceheader"
	"github.com/hreluz/go-db/pkg/invoiceitem"
)

// psqlInvoice used for work with postgres - invoice
type psqlInvoice struct {
	db            *sql.DB
	storageHeader invoiceheader.Storage
	storageItems  invoiceitem.Storage
}

// newPsqlInvoice return a new pointer of psqlInvoice n
func newPsqlInvoice(db *sql.DB, h invoiceheader.Storage, i invoiceitem.Storage) *psqlInvoice {
	return &psqlInvoice{
		db:            db,
		storageHeader: h,
		storageItems:  i,
	}
}

// Create implement the interface invoice.Storage
func (p *psqlInvoice) Create(m *invoice.Model) error {
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

	return tx.Commit()
}
