package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/hreluz/echo-framework/authorization"
	"github.com/hreluz/echo-framework/model"
)

type login struct {
	storage Storage
}

func newLogin(s Storage) login {
	return login{s}
}

func (l *login) login(c echo.Context) error {
	data := model.Login{}
	err := c.Bind(&data)

	if err != nil {
		resp := newResponse(Error, "Structure not valid", nil)
		return c.JSON(http.StatusBadRequest, resp)
	}

	if !isLoginValid(&data) {
		resp := newResponse(Error, "User or password not valid", nil)
		return c.JSON(http.StatusBadRequest, resp)
	}

	token, err := authorization.GenerateToken(&data)
	if err != nil {
		resp := newResponse(Error, "Token could not be generated", nil)
		return c.JSON(http.StatusInternalServerError, resp)
	}

	dataToken := map[string]string{"token": token}
	resp := newResponse(Message, "OK", dataToken)
	return c.JSON(http.StatusOK, resp)
}

func isLoginValid(data *model.Login) bool {
	return data.Email == "batman@gmail.com" && data.Password == "123456"
}
