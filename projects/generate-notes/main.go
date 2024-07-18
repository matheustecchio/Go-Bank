package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"example.com/note"
)

func getUserInput(promptText string) string {
	fmt.Print(promptText) // Prompt user for input
	var input string

	// Read user input and remove trailing newline character using TrimSuffix function from strings package.
	input, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		return ""
	}

	input = strings.TrimSuffix(input, "\n")

	return input
}

func getNoteData() (string, string) {
	title := getUserInput("Enter the title of your note: ")

	description := getUserInput("Enter the description of your note: ")

	return title, description
}

func main() {
	title, content := getNoteData()

	userNote, err := note.New(title, content)
	if err != nil {
		log.Fatal(err)
	}

	userNote.DisplayNote()
	err = userNote.Save()
	if err != nil {
		log.Fatal(err)
	}

}
