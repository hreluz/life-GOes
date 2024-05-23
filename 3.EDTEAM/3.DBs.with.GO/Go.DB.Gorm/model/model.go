package model

import "gorm.io/gorm"

// import "time"

// Model of product
type Product struct {
	gorm.Model
	// ID           uint
	Name         string  `gorm:"type:varchar(100); not null;"`
	Observations *string `gorm:"type:varchar(200);"`
	Price        int     `gorm:"not null;"`
	// CreatedAt    time.Time
	// UpdatedAt    time.Time
	InvoiceItems []InvoiceItem
}

// Model of invoicehaeder
type InvoiceHeader struct {
	gorm.Model
	// ID           uint
	Client       string `gorm:"type:varchar(100); not null;"`
	InvoiceItems []InvoiceItem
	// CreatedAt time.Time
	// UpdatedAt time.Time
}

// Model of invoiceitem
type InvoiceItem struct {
	gorm.Model
	// ID              uint
	InvoiceHeaderID uint
	ProductID       uint
	// CreatedAt       time.Time
	// UpdatedAt       time.Time
}
