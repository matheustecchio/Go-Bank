package main

import (
	"log"

	"github.com/matheustecchio/learn-go/fileops"
)

// balance is a global variable that holds the current balance.
var balance float64

func main() {
	if err := fileops.CreateFile("balance.txt", "0"); err != nil {
		log.Fatal(err)
	}

	balance, err := fileops.GetFloatFromFile("balance.txt")
	if err != nil {
		log.Fatal(err)
	}

	for {
		if err := displayMenu(); err != nil {
			log.Fatal(err)
		}

		if err := fileops.WriteFloatToFile("balance.txt", balance); err != nil {
			log.Fatal(err)
		}
	}
}
