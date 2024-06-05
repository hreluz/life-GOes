package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func createPerson(url, token string, person *Person) GeneralResponse {
	data := bytes.NewBuffer([]byte{})
	err := json.NewEncoder(data).Encode(person)

	if err != nil {
		log.Fatalf("Error processing person for creating %v", err)
	}

	resp := httpClient(http.MethodPost, url, token, data)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalf("Error reading body for creating person")
	}

	if resp.StatusCode != http.StatusCreated {
		log.Fatalf("201 was expected, it got %d, response %s", resp.StatusCode, string(body))
	}

	dataResponse := GeneralResponse{}
	err = json.NewDecoder(bytes.NewReader(body)).Decode(&dataResponse)

	if err != nil {
		log.Fatalf("Error processing body's login: %v", dataResponse)
	}

	return dataResponse
}
