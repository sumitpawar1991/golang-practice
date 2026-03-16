package main

import "fmt"

func main() {
	var city string

	city = "London"

	var c = "abc" //inferred to string

	sub := 5000

	sub = sub + 1

	fmt.Println("sub", sub)

	likes, comments := 100, 30

	fmt.Println("likes , commens , city , c", likes, comments, city, ccd)
}
