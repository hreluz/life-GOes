package storage

import (
	"database/sql"
	"fmt"

	"github.com/hreluz/go-db/pkg/product"
)

type scanner interface {
	Scan(dest ...interface{}) error
}

const (
	psqlMigrateProduct = `CREATE TABLE IF NOT EXISTS products(
		id SERIAL NOT NULL,
		name VARCHAR(25) NOT NULL,
		observations VARCHAR(100),
		price INT NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT now(),
		updated_at TIMESTAMP,
		CONSTRAINT products_id_pk PRIMARY KEY(id)
	)`

	psqlCreateProduct  = `INSERT INTO products(name, observations, price, created_at) VALUES ($1, $2, $3, $4) RETURNING id`
	psqlGetAllProduct  = `SELECT id, name, observations, price, created_at, updated_at FROM products`
	psqlGetProductById = psqlGetAllProduct + " WHERE id = $1"
)

type PsqlProduct struct {
	db *sql.DB
}

// NewPsqlProduct return a new pointer of PsqlProduct
func NewPsqlProduct(db *sql.DB) *PsqlProduct {
	return &PsqlProduct{db}
}

// Migrate implement the interface product.Storage
func (p *PsqlProduct) Migrate() error {
	stmt, err := p.db.Prepare(psqlMigrateProduct)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec()

	if err != nil {
		return err
	}

	fmt.Println("Migration Product executed correctly")
	return nil

}

// Create implement the interface product.Storage
func (p *PsqlProduct) Create(m *product.Model) error {
	stmt, err := p.db.Prepare(psqlCreateProduct)

	if err != nil {
		return err
	}

	defer stmt.Close()

	stmt.QueryRow(
		m.Name,
		stringToNull(m.Observations),
		m.Price,
		m.CreatedAt,
	).Scan(&m.ID)

	fmt.Println("Product created correctly")
	return nil
}

// Get All implement the interface product.Storage
func (p *PsqlProduct) GetAll() (product.Models, error) {
	stmt, err := p.db.Prepare(psqlGetAllProduct)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	ms := make(product.Models, 0)

	for rows.Next() {
		m, err := scanRowProduct(rows)

		if err != nil {
			return nil, err
		}

		ms = append(ms, m)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return ms, nil
}

// GetByID implement the interface product.Storage
func (p *PsqlProduct) GetByID(id uint) (*product.Model, error) {
	stmt, err := p.db.Prepare(psqlGetProductById)

	if err != nil {
		return &product.Model{}, err
	}

	defer stmt.Close()

	return scanRowProduct(stmt.QueryRow(id))
}

func scanRowProduct(s scanner) (*product.Model, error) {
	m := &product.Model{}
	observationNull := sql.NullString{}
	updatedAtNull := sql.NullTime{}

	err := s.Scan(
		&m.ID,
		&m.Name,
		&observationNull,
		&m.Price,
		&m.CreatedAt,
		&updatedAtNull,
	)

	if err != nil {
		return nil, err
	}

	m.Observations = observationNull.String
	m.UpdatedAt = updatedAtNull.Time

	return m, nil
}
