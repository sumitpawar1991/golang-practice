package main

import "fmt"

func main() {

	fmt.Println(divide(10, 5))

	q, r := divide(20, 5)

	fmt.Println("values", q, r)
}

func divide(a int, b int) (p int, q int) {

	p = a / b
	q = a * b

	return
}
