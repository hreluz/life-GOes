package handler

import "github.com/hreluz/echo-framework/model"

type Storage interface {
	Create(person *model.Person) error
	Update(ID int, person *model.Person) error
	Delete(ID int) error
	GetById(ID int) (model.Person, error)
	GetAll() (model.Persons, error)
}
