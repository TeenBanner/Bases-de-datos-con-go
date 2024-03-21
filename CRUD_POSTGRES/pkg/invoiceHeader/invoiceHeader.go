package invoiceheader

import (
	"database/sql"
	"time"
)

// Model of invoiceHeader
type Model struct {
	ID        uint
	Client    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Storage interface {
	Migrate() error
	CreateTx(*sql.Tx, *Model) error
}

// service of invoiceHeader
type Service struct {
	storage Storage
}

// Service Return a pointer of Service
func NewService(s Storage) *Service {
	return &Service{s}
}

// Migrate is  used to create Table invoiceHeader on a database migration
func (s *Service) Migrate() error {
	return s.storage.Migrate()
}
