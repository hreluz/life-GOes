package storage

import (
	"database/sql"
	"fmt"

	"github.com/hreluz/go-db/pkg/invoiceitem"
)

const (
	psqlMigrateInvoiceItem = `CREATE TABLE IF NOT EXISTS invoice_items(
		id SERIAL NOT NULL,
		invoice_header_id INT NOT NULL,
		product_id INT NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT now(),
		updated_at TIMESTAMP,
		CONSTRAINT invoice_items_id_pk PRIMARY KEY(id),
		CONSTRAINT invoice_items_invoice_header_id_fk FOREIGN KEY(invoice_header_id) REFERENCES invoice_headers (id) ON UPDATE RESTRICT ON DELETE RESTRICT,
		CONSTRAINT invoice_items_product_id_fk FOREIGN KEY(product_id) REFERENCES products (id) ON UPDATE RESTRICT ON DELETE RESTRICT
	)`
	psqlCreateInvoiceItem = `INSERT INTO invoice_items(invoice_header_id, product_id) VALUES($1, $2) RETURNING id, created_at`
)

type psqlInvoiceItem struct {
	db *sql.DB
}

// newPsqlInvoiceItem return a new pointer of psqlInvoiceItem
func newPsqlInvoiceItem(db *sql.DB) *psqlInvoiceItem {
	return &psqlInvoiceItem{db}
}

// Migrate implement the interface invoiceItem.Storage
func (p *psqlInvoiceItem) Migrate() error {
	stmt, err := p.db.Prepare(psqlMigrateInvoiceItem)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec()

	if err != nil {
		return err
	}

	fmt.Println("Migration invoiceItem executed correctly")
	return nil
}

// CreateTx implement the interface invoiceitem.Storage
func (p *psqlInvoiceItem) CreateTX(tx *sql.Tx, headerID uint, ms invoiceitem.Models) error {
	stmt, err := tx.Prepare(psqlCreateInvoiceItem)

	if err != nil {
		return err
	}

	defer stmt.Close()

	for _, item := range ms {
		err = stmt.QueryRow(headerID, item.ProductID).Scan(
			&item.ID,
			&item.CreatedAt,
		)

		if err != nil {
			return err
		}
	}

	return nil
}
