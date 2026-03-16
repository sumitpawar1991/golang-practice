package main

import "fmt"

func main() {
	points := map[string]int{
		"a": 10,
		"b": 0,
	}

	fmt.Println("a value is ", points["a"])
	fmt.Println("b value is ", points["b"])
	fmt.Println("c value is ", points["c"])

	//value ,  ok

	valB, okB := points["c"]
	fmt.Println(valB, okB)

	if val, ok := points["b"]; ok {
		fmt.Println(val)
	} else {
		fmt.Println("value key is not present")
	}

	//for range for maps

	prices := map[string]int{
		"xyz": 500,
		"qwe": 2000,
	}

	total := 0
	for item, price := range prices {
		fmt.Println(item, price)
		total = total + price
	}

	fmt.Println("total price ", total)

	for item := range prices {
		fmt.Println(item)
	}
}
