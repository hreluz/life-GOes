package storage

import (
	"fmt"

	"github.com/hreluz/echo-framework/model"
)

type Memory struct {
	currentID int
	Persons   map[int]model.Person
}

func NewMemory() Memory {
	persons := make(map[int]model.Person)

	return Memory{
		currentID: 0,
		Persons:   persons,
	}
}

// Create a person
func (m *Memory) Create(person *model.Person) error {
	if person == nil {
		return model.ErrPersonCanNotBeNil
	}

	m.currentID++
	m.Persons[m.currentID] = *person

	return nil
}

// Update a person in the memory slice
func (m *Memory) Update(ID int, person *model.Person) error {
	if person == nil {
		return model.ErrPersonCanNotBeNil
	}

	if _, ok := m.Persons[ID]; !ok {
		return fmt.Errorf("ID: %d: %w", ID, model.ErrIDPersonDoesNotExist)
	}

	m.Persons[ID] = *person

	return nil
}

// Delete a person from memory
func (m *Memory) Delete(ID int) error {
	if _, ok := m.Persons[ID]; !ok {
		return fmt.Errorf("ID: %d: %w", ID, model.ErrIDPersonDoesNotExist)
	}

	delete(m.Persons, ID)

	return nil
}

// GetById returns a person by ID
func (m *Memory) GetById(ID int) (model.Person, error) {
	person, ok := m.Persons[ID]

	if !ok {
		return person, fmt.Errorf("ID: %d: %w", ID, model.ErrIDPersonDoesNotExist)
	}

	return person, nil
}

// GetAll returns all the persons in memory
func (m *Memory) GetAll() (model.Persons, error) {
	var result model.Persons

	for _, v := range m.Persons {
		result = append(result, v)
	}

	return result, nil
}
