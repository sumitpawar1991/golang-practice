package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {
	u := User{Name: "abc", Age: 30}
	fmt.Println(u.Intro())
}

//receive a copy of user

func (u User) Intro() string {
	return fmt.Sprintf("Hi , i am %s", u.Name)
}
