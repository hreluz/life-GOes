package handler

import "github.com/hreluz/go-api-crud/model"

type Storage interface {
	Create(person *model.Person) error
	Update(ID int, person *model.Person) error
	Delete(ID int) error
	GetById(ID int) (model.Person, error)
	GetAll() (model.Persons, error)
}
