package main

import "fmt"

func main() {

	views1 := 1000
	views2 := 2000

	totalViews := views1 + views2

	fmt.Println("Totel Views", totalViews)

	likes := 10
	likes++
	likes++

	avgViews := totalViews / 2

	fmt.Println("Logged All Information", totalViews, avgViews, likes)
}
