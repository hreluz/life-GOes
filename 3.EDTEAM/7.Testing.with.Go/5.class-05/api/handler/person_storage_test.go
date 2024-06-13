package handler

import (
	"errors"

	"github.com/hreluz/go-testing-class-5/model"
)

type StorageMockError struct{}

func (sme *StorageMockError) Create(person *model.Person) error {
	return errors.New("mock error")
}

func (sme *StorageMockError) Update(ID int, person *model.Person) error {
	return errors.New("mock error")
}

func (sme *StorageMockError) Delete(ID int) error {
	return errors.New("mock error")
}

func (sme *StorageMockError) GetByID(ID int) (model.Person, error) {
	return model.Person{}, errors.New("mock error")
}

func (sme *StorageMockError) GetAll() (model.Persons, error) {
	return model.Persons{}, errors.New("mock error")
}

type StorageMockOk struct{}

func (sme *StorageMockOk) Create(person *model.Person) error {
	return nil
}

func (sme *StorageMockOk) Update(ID int, person *model.Person) error {
	return nil
}

func (sme *StorageMockOk) Delete(ID int) error {
	return nil
}

func (sme *StorageMockOk) GetByID(ID int) (model.Person, error) {
	return model.Person{}, nil
}

func (sme *StorageMockOk) GetAll() (model.Persons, error) {
	return model.Persons{}, nil
}
