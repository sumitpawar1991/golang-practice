package main

import (
	"fmt"
	greet "go-modules/internal"
)

func main() {
	msg1 := greet.Hello("abc")

	fmt.Println(msg1)
}
