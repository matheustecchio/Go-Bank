package main

import "fmt"

func getAdultAge(age *int) int {
	return *age - 18
}

func main() {

	age := 32
	adultAge := getAdultAge(&age)

	agePointer := &age
	*agePointer = 20

	fmt.Printf("My age is %d years old.\n", age)
	fmt.Printf("Age as an adult is %d years old.\n", adultAge)
}
