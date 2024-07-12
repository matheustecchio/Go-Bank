package main

import (
	"log"
)

// balance is a global variable that holds the current balance.
var balance float64

func main() {
	if err := createFile("balance.txt", "0"); err != nil {
		log.Fatal(err)
	}

	balance, err := getFloatFromFile("balance.txt")
	if err != nil {
		log.Fatal(err)
	}

	for {
		if err := displayMenu(); err != nil {
			log.Fatal(err)
		}

		if err := writeFloatToFile("balance.txt", balance); err != nil {
			log.Fatal(err)
		}
	}
}
