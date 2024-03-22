package storage

import (
	"database/sql"
	"fmt"

	invoiceitem "github.com/TeenBanner/db-go/pkg/invoiceItem"
)

const (
	psqlMigrateInvoiceItem = `CREATE TABLE IF NOT EXISTS invoice_items(
		id SERIAL NOT NULL,
		invoice_header_id INT NOT NULL,
		product_id INT NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT now(),
		updated_at TIMESTAMP,

		CONSTRAINT invoice_items_id_pk PRIMARY KEY (id),

		CONSTRAINT invoice_items_invoice_header_id_fk FOREIGN KEY (invoice_header_id)
		REFERENCES invoice_headers (id) ON UPDATE RESTRICT ON DELETE RESTRICT,

		CONSTRAINT invoice_items_product_id_fk FOREIGN KEY (product_id)
		REFERENCES products (id) ON UPDATE RESTRICT ON DELETE RESTRICT
	);`

	psqlCreateInvoiceItem = `
		INSERT INTO invoice_items(invoice_header_id, product_id) VALUES($1, $2) RETURNING id, created_at
		`
)

// Psql used for work with postgres -Invoice_items
type PsqlInvoiceItem struct {
	db *sql.DB
}

// NewpsqlInvoiceHader returns a new pointer of psqlInvoiceItemeee
func NewpsqlInvoiceItem(db *sql.DB) *PsqlInvoiceItem {
	return &PsqlInvoiceItem{db}
}

// Migrate Implement the interface InvoiceItem.Storage
func (p *PsqlInvoiceItem) Migrate() error {
	stmt, err := p.db.Prepare(psqlMigrateInvoiceItem)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	fmt.Println("Migracion de InvoiceItem ejecutada correctamente")

	return nil
}

// CreateTx creates a new transaction of invoiceItems
func (p *PsqlInvoiceItem) CreateTx(tx *sql.Tx, headerID uint, ItemSlice invoiceitem.Models) error {
	// Prepara la sentencia sql
	stmt, err := tx.Prepare(psqlCreateInvoiceItem)
	if err != nil {
		return err
	}
	defer stmt.Close()
	// itera por cada uno de los items para crear el InvoiceItems de Invoice
	for _, item := range ItemSlice {
		err = stmt.QueryRow(headerID, item.ProductID).Scan( // Hace la consutla y mappea los datos en los campos del InvoiceItems
			&item.ID,
			&item.CreatedAt,
		)
		if err != nil {
			return err
		}

	}

	fmt.Println("transaccion de invoice Item Realizada")

	return nil
}
