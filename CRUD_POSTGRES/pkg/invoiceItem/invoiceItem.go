package invoiceitem

import (
	"time"
)

// Model of invoice item
type Model struct {
	ID              uint
	invoiceHeaderID uint
	ProductID       uint
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
