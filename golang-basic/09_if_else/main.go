package main

import "fmt"

func main() {

	score := 72

	if score >= 90 {
		fmt.Println("A+")
	} else if score >= 75 {
		fmt.Println("A")
	} else if score >= 60 {
		fmt.Println("B")
	} else if score >= 45 {
		fmt.Println("C")
	} else if score >= 35 {
		fmt.Println("D")
	} else {
		fmt.Println("F")
	}

}
