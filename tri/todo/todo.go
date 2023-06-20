package todo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

/*
Named types
  - can be any know type (struct, string, int, slice, a new type we declared, etc)
  - methods can be declared on it
  - not an alias - Explicit type
*/

var TodoFilePath string = "/tmp/tri_todolist.json"

type Item struct {
	Text     string
	Priority int
	position int
}

func (i *Item) PrettyPriotity() string {
	if i.Priority == 0 {
		return " "
	}

	return fmt.Sprintf("(%d)", i.Priority)
}

func (i *Item) PrettyPosition() string {
	return fmt.Sprintf("%d.", i.position)
}

/*
- ByPri implements sorte.Interface for []Item based on the Priority & position field
*/
type ByPri []Item

func (s ByPri) Len() int {
	return len(s)
}

func (s ByPri) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByPri) Less(i int, j int) bool {
	if s[i].Priority == s[j].Priority {
		return s[i].position < s[j].position
	}

	return s[i].Priority < s[j].Priority
}

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

	for i := range items {
		items[i].position = i + 1
	}

	return items, nil
}

func (i *Item) SetPriority(priorityLevel int) {
	switch priorityLevel {
	case 1:
		i.Priority = 1
	case 3:
		i.Priority = 3
	default:
		i.Priority = 2
	}
}
