package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sort"
)

func getResources(url string) ([]map[string]any, error) {
	var resources []map[string]any

	res, err := http.Get(url)
	if err != nil {
		return resources, err
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return resources, err
	}

	if err = json.Unmarshal(data, &resources); err != nil {
		return resources, err
	}

	return resources, nil
}

func logResources(resources []map[string]any) {
	var formattedStrings []string

	for _, mapItem := range resources {
		for key, value := range mapItem {
			store := fmt.Sprintf(
				"Key: %s - Value: %v",
				fmt.Sprint(key),
				value,
			) // can also use fmt.Sprint(key) instead.
			formattedStrings = append(formattedStrings, store)
		}

	}

	sort.Strings(formattedStrings)

	for _, str := range formattedStrings {
		fmt.Println(str)
	}
}
