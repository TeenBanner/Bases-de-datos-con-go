package product

import (
	"fmt"
	"time"
)

// model of product
type Model struct {
	ID          uint
	Name        string
	Observation string
	Price       int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// Format Data Recived From DB
func (m *Model) String() string {
	return fmt.Sprintf("%02d | %-20s | %-20s | %5d | %10s | %10s", m.ID, m.Name, m.Observation, m.Price, m.CreatedAt.Format("2006-01-02"), m.UpdatedAt.Format("2006-01-02"))
}

type Models []*Model

type Storage interface {
	Migrate() error
	Create(*Model) error
	// Update(*Model) error
	GetAll() (Models, error)
	// GetByID(uint) (*Model, error)
	// Delete() error
}

// servicio of product
type Service struct {
	storage Storage
}

// Service Return a pointer of Service
func NewService(s Storage) *Service {
	return &Service{s}
}

// Migrate is  used to run the database migration
func (s *Service) Migrate() error {
	return s.storage.Migrate()
}

// Create it's use for insert a product
func (s *Service) Create(m *Model) error {
	// inicializamos la fecha de creacion del producto
	m.CreatedAt = time.Now()
	// Retornamos el posible error al ejecutar el metodo Create de la interfaz Storage
	return s.storage.Create(m)
}

func (s *Service) GetAll() (Models, error) {
	return s.storage.GetAll()
}
