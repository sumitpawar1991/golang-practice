package main

import "fmt"

func main() {

	fmt.Println(add(1, 2))
	s, p := SumAndProduct(10, 20)

	fmt.Println(s, p)
	//fmt.Printlne(add(1, "2")) // give error type is int not string
}

func add(a int, b int) int {
	return a + b
}

func SumAndProduct(a int, b int) (int, int) {
	sum := a + b
	product := a * b

	return sum, product
}
