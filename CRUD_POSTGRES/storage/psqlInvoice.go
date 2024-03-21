package storage

import (
	"database/sql"
	"fmt"

	"github.com/TeenBanner/db-go/pkg/invoice"
	invoiceheader "github.com/TeenBanner/db-go/pkg/invoiceHeader"
	invoiceitem "github.com/TeenBanner/db-go/pkg/invoiceItem"
)

type PsqlInvoice struct {
	db            *sql.DB
	storageHeader invoiceheader.Storage
	storageItem   invoiceitem.Storage
}

func NewPsqlInvoice(db *sql.DB, h invoiceheader.Storage, i invoiceitem.Storage) *PsqlInvoice {
	return &PsqlInvoice{
		db:            db,
		storageHeader: h,
		storageItem:   i,
	}
}

// Create Implement the interface invoice.Storage
func (p *PsqlInvoice) Create(m *invoice.Model) error {
	//iniciamos la transaccion
	tx, err := p.db.Begin()
	if err != nil {
		return fmt.Errorf("Error al Hacer la transaccion de factura: %v", err)
	}
	// Creamos la transaccion de InvoiceHeader
	if err = p.storageHeader.CreateTx(tx, m.Header); err != nil {
		tx.Rollback()
		return err
	}
	// Creamos la transaccion de InvoiceItem
	if err := p.storageItem.CreateTx(tx, m.Header.ID, m.Items); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
