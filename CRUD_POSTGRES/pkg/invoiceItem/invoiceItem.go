package invoiceitem

import (
	"database/sql"
	"time"
)

// Model of invoice item
type Model struct {
	ID              uint
	InvoiceHeaderID uint
	ProductID       uint
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

// slice of models
type Models []*Model

type Storage interface {
	Migrate() error
	CreateTx(*sql.Tx, uint, Models) error
}

// service of invoice item
type Service struct {
	storage Storage
}

// Service Return a pointer of Service
func NewService(s Storage) *Service {
	return &Service{s}
}

// Migrate is used to run the Create Invoice_items migration
func (s *Service) Migrate() error {
	return s.storage.Migrate()
}
