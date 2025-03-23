package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func createUser(url, apiKey string, data User) (User, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return User{}, fmt.Errorf("%w", err)
	}

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return User{}, err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("X-API-Key", apiKey)

	client := &http.Client{}

	res, err := client.Do(request)
	if err != nil {
		return User{}, nil
	}

	defer res.Body.Close()

	var user User

	decoder := json.NewDecoder(res.Body)

	if err := decoder.Decode(&user); err != nil {
		return User{}, err
	}

	return user, nil
}
