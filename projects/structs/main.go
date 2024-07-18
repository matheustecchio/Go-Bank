package main

import (
	"fmt"
	"log"

	"example.com/user"
)

var (
	UserFirstName string
	UserLastName  string
	UserBirthdate string
)

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var input string
	fmt.Scanln(&input)
	return input
}

func main() {
	UserFirstName = getUserData("Please enter your first name: ")
	UserLastName = getUserData("Please enter your last name: ")
	UserBirthdate = getUserData("Please enter your birthdate (DD/MM/YYYY): ")

	var appUser *user.User

	appUser, err := user.New(UserFirstName, UserLastName, UserBirthdate)
	if err != nil {
		log.Fatal(err)
	}

	appUser.OutputUserDetails()
	appUser.ClearUserName()

	appAdmin := user.NewAdmin("admin@example.com", "adminadmin")

	appAdmin.User.OutputUserDetails()
	appAdmin.User.ClearUserName()
}
