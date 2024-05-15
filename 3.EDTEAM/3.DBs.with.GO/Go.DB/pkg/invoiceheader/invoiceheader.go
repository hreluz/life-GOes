package invoiceheader

import "time"

// Model of invoicehaeder
type Model struct {
	ID        uint
	Client    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Storage interface {
	Migrate() error
}

// Service of invoiceheader
type Service struct {
	storage Storage
}

// NewService return a pointer of Service
func NewService(s Storage) *Service {
	return &Service{s}
}

// Migrate is used for migrate invoiceheader
func (s *Service) Migrate() error {
	return s.storage.Migrate()
}
