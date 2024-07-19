package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"text"`
}

func New(text string) (Todo, error) {
	// Validate input
	if text == "" {
		return Todo{}, errors.New("text is required")
	}

	return Todo{
		Text: text,
	}, nil
}

func (todo Todo) Display() {
	fmt.Printf("Todo: %s\n", todo.Text)
}

func (todo Todo) Save() error {
	fileName := "todo.json"

	json, err := json.Marshal(todo)
	if err != nil {
		return err
	}

	if err := os.WriteFile(fileName, json, 0644); err != nil {
		return err
	}

	fmt.Print("Saved successfully!\n")

	return nil
}
