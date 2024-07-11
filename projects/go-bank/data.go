package main

import (
	"fmt"
	"os"
	"strconv"
)

// createFile checks if a file exists. If it doesn't, it creates the file with the designated name and format,
// and inserts the `initialValue` into it.
//
// `file`: The name of the file to be created.
// `initialValue`: The initial value to be written into the file.
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

// getFileData reads a file and returns its content as a float64 value.
//
// `file`: The name of the file to be read.
//
// Returns the content of the file as a float64 value and an error if any operation fails.
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

// saveDataToFile writes a float64 value to a file.
//
// `file`: The name of the file to be written to.
// `targetFloat`: The float64 value to be written into the file.
func saveDataToFile(file string, targetFloat float64) error {
	floatToString := fmt.Sprint(targetFloat)
	if err := os.WriteFile(file, []byte(floatToString), 0644); err != nil {
		return err
	}

	return nil
}
