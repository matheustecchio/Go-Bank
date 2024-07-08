package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

// If a file doesn't exist, it creates the file with the designated name and format, and insert the value 0.
func createFile(fileName string) {
	_, err := os.ReadFile(fileName)
	if err != nil {
		os.Create(fileName)
		os.WriteFile(fileName, []byte("0"), 0644)
	}
}

// Synchronizes the file data to a variable in our code.
func syncData(fileName string) float64 {
	fileByte, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatalf("File '%s' doesn't exist or cannot be read: %v", fileName, err)
	}

	byteToString := string(fileByte)
	dataToBeSaved, err := strconv.ParseFloat(byteToString, 64)
	if err != nil {
		log.Fatalf("Failed to convert balance data from file '%s' to float64: %v", fileName, err)
	}

	return dataToBeSaved
}

// Save data to a file.
func saveDataToFile(fileName string, targetFloat float64) {
	floatToString := fmt.Sprint(targetFloat)
	os.WriteFile(fileName, []byte(floatToString), 0644)
}
