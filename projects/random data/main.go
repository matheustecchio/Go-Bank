package main

import (
	"fmt"

	"github.com/Pallinder/go-randomdata" // use `go get github.com/Pallinder/go-randomdata` to install the package
)

func main() {
	fmt.Println(randomdata.SillyName())
	fmt.Println(randomdata.Adjective())
	fmt.Println(randomdata.City())

}
