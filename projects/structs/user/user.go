package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	FirstName string
	LastName  string
	BirthDate string
	CreatedAt time.Time
}

// Prints the user's first name, last name, and birth date.
func (u *User) OutputUserDetails() {
	fmt.Println(u.FirstName, u.LastName, u.BirthDate)
}

// Sets the user's first name and last name to an empty string.
func (u *User) ClearUserName() {
	u.FirstName = ""
	u.LastName = ""
}

// Creates a new User instance with the provided first name, last name, and birth date.
// It returns a pointer to the new User instance and an error if any of the fields are empty.
func New(firstName, lastName, birthDate string) (*User, error) {

	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("all fields must be filled")
	}

	return &User{
		FirstName: firstName,
		LastName:  lastName,
		BirthDate: birthDate,
		CreatedAt: time.Now(),
	}, nil

}

type Admin struct {
	Email    string
	Password string
	User     User
}

func NewAdmin(email string, password string) *Admin {
	return &Admin{
		Email:    email,
		Password: password,
		User: User{
			FirstName: "ADMIN",
			LastName:  "ADMIN",
			BirthDate: "01/01/01",
			CreatedAt: time.Now(),
		},
	}
}
