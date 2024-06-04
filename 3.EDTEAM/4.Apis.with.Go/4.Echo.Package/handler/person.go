package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/hreluz/echo-framework/model"
	"github.com/labstack/echo/v4"
)

type person struct {
	storage Storage
}

func newPerson(storage Storage) person {
	return person{storage}
}

func (p *person) create(c echo.Context) error {
	data := model.Person{}
	err := c.Bind(&data)

	if err != nil {
		response := newResponse(Error, "The person sent is not valid", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	err = p.storage.Create(&data)

	if err != nil {
		response := newResponse(Error, "There was a problem creating the person", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := newResponse(Message, "Person was created successfully", nil)
	return c.JSON(http.StatusCreated, response)
}

func (p *person) update(c echo.Context) error {
	ID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		response := newResponse(Error, "The id must be an integer", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	data := model.Person{}
	err = c.Bind(&data)

	if err != nil {
		response := newResponse(Error, "The person sent is not valid", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	err = p.storage.Update(ID, &data)

	if err != nil {
		response := newResponse(Error, "There was a problem updating", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := newResponse(Message, "Person was updated correctly", nil)
	return c.JSON(http.StatusOK, response)
}

func (p *person) delete(c echo.Context) error {

	ID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		response := newResponse(Error, "The id must be an integer", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	err = p.storage.Delete(ID)

	if errors.Is(err, model.ErrIDPersonDoesNotExist) {
		response := newResponse(Error, "The id of the person does not exist", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	if err != nil {
		response := newResponse(Error, "There was an error deleting the person", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := newResponse(Message, "It was deleted successfuly", nil)
	return c.JSON(http.StatusOK, response)

}

func (p *person) getById(c echo.Context) error {
	ID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		response := newResponse(Error, "The id must be an integer", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	person, err := p.storage.GetById(ID)

	if err != nil {
		response := newResponse(Error, "Person not found", nil)
		return c.JSON(http.StatusNotFound, response)
	}

	response := newResponse(Message, "Ok", person)
	return c.JSON(http.StatusOK, response)
}

func (p *person) getAll(c echo.Context) error {
	data, err := p.storage.GetAll()

	if err != nil {
		response := newResponse(Error, "There was a problem when getting people", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := newResponse(Message, "Ok", data)
	return c.JSON(http.StatusOK, response)
}
