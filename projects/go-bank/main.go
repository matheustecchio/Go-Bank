package main

var balance float64

func main() {
	createFile("balance.txt")
	balance = syncData("balance.txt")

	for {
		displayMenu()
		saveDataToFile("balance.txt", balance)
	}

}
