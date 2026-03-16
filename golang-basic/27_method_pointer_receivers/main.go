package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {
	u := User{Name: "abc", Age: 10}

	fmt.Println(u.Age)

	u.Birthday()

	fmt.Println("After birtday", u.Age)
}

func (u *User) Birthday() {
	u.Age++
}
