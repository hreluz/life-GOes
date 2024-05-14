package invoiceheader

import "time"

// Model of invoicehaeder
type Model struct {
	ID        uint
	Client    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
