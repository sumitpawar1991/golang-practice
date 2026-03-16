package main

import "fmt"

func main() {

	//in golang array is fixed given array value is fixed
	var marks [3]int

	marks[0] = 70
	marks[1] = 90
	marks[2] = 80

	fmt.Println("marks", marks)

	//array literal

	res := [5]int{1, 2, 3, 4, 5}

	fmt.Println(len(res))
}
