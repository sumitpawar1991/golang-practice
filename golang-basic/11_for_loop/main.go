package main

import "fmt"

func main() {

	for i := 0; i <= 10; i++ {

		fmt.Println(i)
	}

	N := 10

	sum := 0

	for i := 1; i <= N; i++ {
		sum = sum + 2
	}

	fmt.Println("sum value ", sum)
}
