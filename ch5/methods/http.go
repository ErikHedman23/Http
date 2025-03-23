package main

import (
	"encoding/json"
	"net/http"
)

func getUsers(url string) ([]User, error) {
	//Another way to write this is to use the two-step approach:
	//var client *http.Client
	//client = new(http.Client)
	//What you are doing here is creating a new instance of the http.Client struct
	//And taking the memory address of that instance,
	//and assigning that address (a pointer), to the client variable.
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	var user []User
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&user); err != nil {
		return nil, err
	}

	return user, nil
}
