package storage

import (
	"database/sql"
	"fmt"

	invoiceheader "github.com/TeenBanner/db-go/pkg/invoiceHeader"
)

const (
	psqlMigrateInvoiceHeader = `CREATE TABLE IF NOT EXISTS invoice_headers(
		id SERIAL NOT NULL,
		client VARCHAR(100) NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT now(),
		updated_at TIMESTAMP,
		CONSTRAINT invoice_headers_id_pk PRIMARY KEY (id)
	)`
	psqlCreateInvoiceHeader = `
		INSERT INTO invoice_headers(client) VALUES($1) RETURNING id, created_at
	`
)

// Psql used for work with postgres -Invoice_Headers
type PsqlInvoiceHeader struct {
	db *sql.DB
}

// NewpsqlInvoiceHader returns a new pointer of psqlInvoiceHeader
func NewpsqlInvoiceHeader(db *sql.DB) *PsqlInvoiceHeader {
	return &PsqlInvoiceHeader{db}
}

// Migrate Implement the interface InvoiceHeader.Storage
func (p *PsqlInvoiceHeader) Migrate() error {
	stmt, err := p.db.Prepare(psqlMigrateInvoiceHeader)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	fmt.Println("Migracion de InvoiceHeader ejecutada correctamente")

	return nil
}

// CreateTx implements invoiceHeader.Storage, it creates the invoice Header of invoice
func (p *PsqlInvoiceHeader) CreateTx(tx *sql.Tx, h *invoiceheader.Model) error {
	// prepare Sql Sentence
	stmt, err := tx.Prepare(psqlCreateInvoiceHeader)
	if err != nil {
		return err
	}
	defer stmt.Close()

	fmt.Println("Transaccion de invoiceHeader Realizada")
	// hace la consulta y le pasamos el campo cliente del modelo de InvoiceHeader recibido para pasarlo como el valor del marcador de poscicion de la sentencia sql
	// para despues mappear los campos del modelo de InvoiceHeader con los datos que la DB nos esta regresando
	return stmt.QueryRow(h.Client).Scan(&h.ID, &h.CreatedAt) // esta funcion devuelve un nil por lo que podemos retornarla al final de la funcio
}
