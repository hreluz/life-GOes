package storage

import (
	"database/sql"
	"fmt"
	"github.com/hreluz/go-db/pkg/invoiceitem"
)

const (
	mySQLMigrateInvoiceItem = `CREATE TABLE IF NOT EXISTS invoice_items(
		id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
		invoice_header_id INT NOT NULL,
		product_id INT NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT now(),
		updated_at TIMESTAMP,
		CONSTRAINT invoice_items_invoice_header_id_fk FOREIGN KEY(invoice_header_id) REFERENCES invoice_headers (id) ON UPDATE RESTRICT ON DELETE RESTRICT,
		CONSTRAINT invoice_items_product_id_fk FOREIGN KEY(product_id) REFERENCES products (id) ON UPDATE RESTRICT ON DELETE RESTRICT
	) ENGINE = InnoDB`
	mySQLCreateInvoiceItem = `INSERT INTO invoice_items(invoice_header_id, product_id) VALUES(?, ?)`
)

type mySQLInvoiceItem struct {
	db *sql.DB
}

// newMySQLInvoiceItem return a new pointer of mySQLInvoiceItem
func newMySQLInvoiceItem(db *sql.DB) *mySQLInvoiceItem {
	return &mySQLInvoiceItem{db}
}

// Migrate implement the interface invoiceItem.Storage
func (p *mySQLInvoiceItem) Migrate() error {
	stmt, err := p.db.Prepare(mySQLMigrateInvoiceItem)

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
func (p *mySQLInvoiceItem) CreateTX(tx *sql.Tx, headerID uint, ms invoiceitem.Models) error {
	stmt, err := tx.Prepare(mySQLCreateInvoiceItem)

	if err != nil {
		return err
	}

	defer stmt.Close()

	for _, item := range ms {
		result, err := stmt.Exec(
			headerID,
			item.ProductID,
		)

		if err != nil {
			return err
		}

		id, err := result.LastInsertId()

		if err != nil {
			return err
		}

		item.ID = uint(id)
	}

	return nil
}
