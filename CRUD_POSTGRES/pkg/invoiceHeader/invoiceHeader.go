package invoiceheader

import "time"

// Model of invoiceHeader
type Model struct {
	ID        uint
	Client    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
