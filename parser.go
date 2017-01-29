package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Event struct {
	Type string
  Branch string
  Container string
	Project string
}

type BranchMapping map[string][]string

type ProjectMapping map[string]*BranchMapping

type Mapping map[string]*ProjectMapping

func ReadEvent(file string) (*Event, error) {
	data := readFromFile(file)
	event := Event{
		Type: "service.update",
	}
	err := json.Unmarshal(data, &event)
	if err != nil {
		return nil, err
	}
	return &event, nil
}

func ReadMapping(file string) (*Mapping, error) {
	data := readFromFile(file)
	var mapping Mapping
	err := json.Unmarshal(data, &mapping)
	if err != nil {
		return nil, err
	}
	return &mapping, nil
}

func readFromFile(file string) []byte {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "Error reading version file: %s", err)
		return nil
	}

	data, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading version file: %s", err)
		return nil
	}
	return data
}
