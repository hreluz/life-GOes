package handler

import (
	"encoding/json"
	"net/http"

	"github.com/hreluz/go-api-crud/authorization"
	"github.com/hreluz/go-api-crud/model"
)

type login struct {
	storage Storage
}

func newLogin(s Storage) login {
	return login{s}
}

func (l *login) login(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		response := newResponse(Error, "Method not supported", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	data := model.Login{}
	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		resp := newResponse(Error, "Structure not valid", nil)
		responseJSON(w, http.StatusBadRequest, resp)
		return
	}

	if !isLoginValid(&data) {
		resp := newResponse(Error, "User or password not valid", nil)
		responseJSON(w, http.StatusBadRequest, resp)
		return
	}

	token, err := authorization.GenerateToken(&data)
	if err != nil {
		resp := newResponse(Error, "Token could not be generated", nil)
		responseJSON(w, http.StatusBadRequest, resp)
		return
	}

	dataToken := map[string]string{"token": token}
	resp := newResponse(Message, "OK", dataToken)
	responseJSON(w, http.StatusOK, resp)
}

func isLoginValid(data *model.Login) bool {
	return data.Email == "batman@gmail.com" && data.Password == "123456"
}
