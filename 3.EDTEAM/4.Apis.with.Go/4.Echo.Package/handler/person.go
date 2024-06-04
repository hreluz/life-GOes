package handler

import (
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

// func (p *person) delete(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodDelete {
// 		response := newResponse(Error, "Method not supported", nil)
// 		responseJSON(w, http.StatusBadRequest, response)
// 		return
// 	}

// 	ID, err := strconv.Atoi(r.URL.Query().Get("id"))

// 	if err != nil {
// 		response := newResponse(Error, "The id must be an integer", nil)
// 		responseJSON(w, http.StatusBadRequest, response)
// 		return
// 	}

// 	err = p.storage.Delete(ID)

// 	if errors.Is(err, model.ErrIDPersonDoesNotExist) {
// 		response := newResponse(Error, "The id of the person does not exist", nil)
// 		responseJSON(w, http.StatusBadRequest, response)
// 		return
// 	}

// 	if err != nil {
// 		response := newResponse(Error, "There was an error deleting the person", nil)
// 		responseJSON(w, http.StatusInternalServerError, response)
// 		return
// 	}

// 	response := newResponse(Message, "It was deleted successfuly", nil)
// 	responseJSON(w, http.StatusOK, response)

// }

// func (p *person) getById(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodGet {
// 		response := newResponse(Error, "Method not supported", nil)
// 		responseJSON(w, http.StatusBadRequest, response)
// 		return
// 	}

// 	ID, err := strconv.Atoi(r.URL.Query().Get("id"))

// 	if err != nil {
// 		response := newResponse(Error, "The id must be an integer", nil)
// 		responseJSON(w, http.StatusBadRequest, response)
// 		return
// 	}

// 	person, err := p.storage.GetById(ID)

// 	if err != nil {
// 		response := newResponse(Error, "Person not found", nil)
// 		responseJSON(w, http.StatusNotFound, response)
// 		return
// 	}

// 	response := newResponse(Message, "Ok", person)
// 	responseJSON(w, http.StatusOK, response)
// }

func (p *person) getAll(c echo.Context) error {
	data, err := p.storage.GetAll()

	if err != nil {
		response := newResponse(Error, "There was a problem when getting people", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := newResponse(Message, "Ok", data)
	return c.JSON(http.StatusOK, response)
}
