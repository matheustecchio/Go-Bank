package main

import (
	"log"
)

var balance float64

func main() {
	if err := createFile("balance.txt", "0"); err != nil {
		log.Fatal(err)
	}

	balance, err := getFileData("balance.txt")
	if err != nil {
		log.Fatal(err)
	}

	for {
		if err := displayMenu(); err != nil {
			log.Fatal(err)
		}

		if err := saveDataToFile("balance.txt", balance); err != nil {
			log.Fatal(err)
		}
	}

}
