package main

import (
	"fmt"
	"strings"
)

func main() {
	firstName := "ABC"
	lastName := "PQR"

	fullName := firstName + " " + lastName

	fmt.Println("fullname", fullName)

	fmt.Println(strings.ToUpper(fullName))
}
