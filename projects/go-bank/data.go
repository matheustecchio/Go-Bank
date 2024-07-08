package main

import (
	"fmt"
	"os"
	"strconv"
)

// If a file doesn't exist, it creates the file with the designated name and format, and insert `initialValue` in it.
func createFile(file string, initialValue string) error {
	if _, err := os.ReadFile(file); err != nil {
		if _, err := os.Create(file); err != nil {
			return err
		}
		if err := os.WriteFile(file, []byte(initialValue), 0644); err != nil {
			return err
		}
	}

	return nil
}

// Return a value from the file. Used to save the value in the file to a variable in the code.
func getFileData(file string) (float64, error) {
	fileByte, err := os.ReadFile(file)
	if err != nil {
		return 0, err
	}

	byteToString := string(fileByte)
	dataToBeSaved, err := strconv.ParseFloat(byteToString, 64)
	if err != nil {
		return 0, err
	}

	return dataToBeSaved, nil
}

// Save data to a file.
func saveDataToFile(file string, targetFloat float64) error {
	floatToString := fmt.Sprint(targetFloat)
	if err := os.WriteFile(file, []byte(floatToString), 0644); err != nil {
		return err
	}

	return nil
}
