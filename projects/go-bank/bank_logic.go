package main

import (
	"fmt"
	"log"
	"os"
)

// displayMenu is a function that provides a simple command-line interface for a banking application.
// It displays a menu with options to check balance, deposit money, withdraw money, and exit the application.
// The function reads the user's choice and calls the corresponding function.
// If an error occurs while reading the user's choice, it returns the error.
func displayMenu() error {
	fmt.Println("\nWelcome to the Go Bank!")
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit")
	fmt.Println("3. Withdraw")
	fmt.Println("0. Exit")
	var choice uint8
	fmt.Print("Enter with your option: ")

	if _, err := fmt.Scan(&choice); err != nil {
		return err
	}

	switch choice {
	case 1:
		checkBalance()
	case 2:
		depositMoney()
	case 3:
		withdrawMoney()
	case 0:
		os.Exit(0)
	}

	return nil
}

// checkBalance is a function that prints the current balance of the bank account.
func checkBalance() {
	fmt.Printf("Your balance is: %.2f\n", balance)
}

// depositMoney is a function that allows the user to deposit money into their bank account.
// It reads the amount to be deposited from the user and updates the balance accordingly.
// If an error occurs while reading the amount, it logs a fatal error message.
func depositMoney() {
	fmt.Println("Insert how much do you want to deposit: ")
	var amount float64
	if _, err := fmt.Scan(&amount); err != nil {
		log.Fatal("\nError while depositing...")
	}

	balance = balance + amount

	fmt.Printf("\nYou deposited %2.f to your bank account!\n", amount)
}

// withdrawMoney is a function that allows the user to withdraw money from their bank account.
// It reads the amount to be withdrawn from the user and updates the balance accordingly.
// If the user tries to withdraw more money than their balance, it logs a fatal error message.
// If an error occurs while reading the amount, it logs a fatal error message.
func withdrawMoney() {
	fmt.Println("Insert how much do you want to withdraw: ")
	var amount float64
	if _, err := fmt.Scan(&amount); err != nil {
		log.Fatal("\nError while withdrawing...")
	}

	if amount > balance {
		log.Fatal("\nYou don't have enough money to complete this operation!")
	}

	balance = balance - amount

	fmt.Printf("\nYou withdrew %2.f to your bank account!\n", amount)
}
