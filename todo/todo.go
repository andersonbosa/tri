package todo

import (
	"encoding/json"
	"io/ioutil"
)

/*
Named types
  - can be any know type (struct, string, int, slice, a new type we declared, etc)
  - methods can be declared on it
  - not an alias - Explicit type
*/
type Item struct {
	/* EXPORTED: */
	Text string
}

var TodoFilePath string = "/tmp/tri_todolist.json"

func SaveItems(filename string, items []Item) error {
	b, err := json.Marshal(items)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, b, 0644)
	if err != nil {
		return err
	}

	return nil
}

func ReadItems(filename string) ([]Item, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return []Item{}, nil
	}

	var items []Item
	if err := json.Unmarshal(b, &items); err != nil {
		return []Item{}, err
	}

	return items, nil
}
