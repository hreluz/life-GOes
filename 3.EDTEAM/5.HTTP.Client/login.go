package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func loginClient(url, email, password string) LoginResponse {
	login := Login{
		Email:    email,
		Password: password,
	}

	data := bytes.NewBuffer([]byte{})
	err := json.NewEncoder(data).Encode(&login)

	if err != nil {
		log.Fatalf("Error on login client fn: %v", err)
	}

	resp := httpClient(http.MethodPost, url, data)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalf("Error reading login's body: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("It was expected 200 code, it got %d, response :%s", resp.StatusCode, string(body))
	}

	dataResponse := LoginResponse{}
	err = json.NewDecoder(bytes.NewReader(body)).Decode(&dataResponse)

	if err != nil {
		log.Fatalf("Error uncoding the body in login: %v", err)
	}

	return dataResponse
}
