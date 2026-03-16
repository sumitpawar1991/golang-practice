package main

import "fmt"

func main() {
	day := 3

	switch day {

	case 1:
		fmt.Println("This is first day of week")
		break

	case 2:
		fmt.Println("This is second day of week")
		break

	case 3:
		fmt.Println("This is third day of week")
		break

	default:
		fmt.Println("4 day or more in week")
	}
}
