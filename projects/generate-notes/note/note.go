package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

func New(title, description string) (Note, error) {
	// Validate input
	if title == "" || description == "" {
		return Note{}, errors.New("description is required")
	}

	return Note{
		Title:       title,
		Description: description,
		CreatedAt:   time.Now(),
	}, nil
}

func (n Note) Display() {
	fmt.Printf("Note Title: %s\nNote Description:%s\n", n.Title, n.Description)
}

func (n Note) Save() error {
	fileName := strings.ReplaceAll(n.Title, " ", "_")
	fileName = strings.ToLower(fileName)
	fileName += ".json"

	json, err := json.Marshal(n)
	if err != nil {
		return err
	}

	if err := os.WriteFile(fileName, json, 0644); err != nil {
		return err
	}

	fmt.Print("Saved successfully!\n")

	return nil
}
