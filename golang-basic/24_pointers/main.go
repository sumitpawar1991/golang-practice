package main

import "fmt"

func main() {
	//store the memory address of any val

	//&x -> add of x
	//*x -> de ref ( go to that address and read / write ) change value
	score := 10
	fmt.Println("before: ", score)

	addScore(&score)

	fmt.Println("After: ", score)

}

func addScore(score *int) {
	*score = *score + 5
}
