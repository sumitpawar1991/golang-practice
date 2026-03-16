package main

import "fmt"

func main() {

	//change true false values and check the result
	isLogged := true

	isAdmin := false

	hasSubcription := true

	canOpenDashboard := isLogged && hasSubcription

	canDeletePost := isAdmin || (isLogged && hasSubcription)

	fmt.Println(canDeletePost, canOpenDashboard)

	age := 35
	isAdult := age > 18

	fmt.Println(isAdult)
}
