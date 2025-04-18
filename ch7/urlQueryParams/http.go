package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func getUsers(url string) ([]User, error) {
	fullURL := url + fmt.Sprintf("?sort=experience")
	res, err := http.Get(fullURL)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var users []User
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&users)
	if err != nil {
		return nil, err
	}

	return users, nil
}
