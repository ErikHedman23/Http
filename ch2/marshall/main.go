package main

import (
	"encoding/json"
)

func marshalAll[T any](items []T) ([][]byte, error) {
	var slices [][]byte
	slices = make([][]byte, len(items))
	var err error
	for i := range slices {
		slices[i], err = json.Marshal(items[i])
		if err != nil {
			return nil, err
		}
	}

	return slices, nil
}
