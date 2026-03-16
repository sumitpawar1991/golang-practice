package main

import "fmt"

func main() {
	fmt.Println(sumAll(1, 2, 3, 4, 5, 6, 7, 8, 9))

	values := []int{10, 20, 30, 40}
	fmt.Println(sumAll(values...))

	res := func(n int) int {
		return n * 2
	}

	fmt.Println(res(2))

	//IIFE style
	res1 := func(a int, b int) int {
		return a + b
	}(5, 10)

	fmt.Println(res1)
}

func sumAll(nums ...int) int {

	total := 0

	for _, currentValue := range nums {
		total = currentValue + total
	}

	return total
}
