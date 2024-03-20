package invoiceheader

import "time"

// Model of invoiceHeader
type Model struct {
	ID        uint
	Client    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Storage interface {
	Migrate() error
	// Create(*Model) error
	// Update(*Model) error
	// GetAll() (Models, error)
	// GetByID(uint) (*Model, error)
	// Delete() error
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
