package invoice

import (
	invoiceheader "github.com/TeenBanner/db-go/pkg/invoiceHeader"
	invoiceitem "github.com/TeenBanner/db-go/pkg/invoiceItem"
)

type Model struct {
	Header *invoiceheader.Model
	Items  invoiceitem.Models
}

// storage interface that must implement a db storage
type Storage interface {
	Create(*Model) error
}

// Services of invoice
type Service struct {
	storage Storage
}

// contruct a new service struct
func NewService(s Storage) *Service {
	return &Service{s}
}

// Create Implement Create method of Storage.Create
func (s *Service) Create(m *Model) error {
	return s.storage.Create(m)
}
