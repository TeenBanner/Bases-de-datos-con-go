package product

import (
	"errors"
	"fmt"
	"time"
)

var (
	ErrIDNotFound = errors.New("El producto no contiene ese id")
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
	Update(*Model) error
	GetAll() (Models, error)
	GetByID(uint) (*Model, error)
	Delete(id uint) error
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

// GetAll Service it's used for get all products
func (s *Service) GetAll() (Models, error) {
	return s.storage.GetAll()
}

// GetByID service it's used for get an element by it id
func (s *Service) GetByID(ID uint) (*Model, error) {
	return s.storage.GetByID(ID)
}

// Update service implement the Storage interfaces updating a product
func (s *Service) Update(m *Model) error {
	if m.ID == 0 {
		return ErrIDNotFound
	}
	m.UpdatedAt = time.Now()
	return s.storage.Update(m)
}

// Delete it's used for delete a product, it implements the storage interface
func (s *Service) Delete(Id uint) error {
	return s.storage.Delete(Id)
}
