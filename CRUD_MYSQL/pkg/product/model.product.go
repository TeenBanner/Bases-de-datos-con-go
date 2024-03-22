package product

import "time"

type Model struct {
	ID           uint
	Name         string
	Observations string
	Price        int
	Created_at   time.Time
	Updated_at   time.Time
}

type Models []*Model

type Storage interface {
	Migrate() error
	// Create(*Model) error
	// Update(*Model) error
	// GetAll() (Models, error)
	// GetById(uint) (*Model, error)
	// Delete(id uint) error
}

type Service struct {
	storage Storage
}

func NewService(s Storage) *Service {
	return &Service{s}
}

func (s *Service) Migrate() error {
	return s.storage.Migrate()
}
