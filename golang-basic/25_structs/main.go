package main

import "fmt"

// struct groups related fields into on type
type User struct {
	ID    int
	Name  string
	Email string
	Age   int
}

func main() {
	user1 := User{
		ID:    1,
		Name:  "Abc",
		Email: "abc@gmail.com",
		Age:   100,
	}

	fmt.Println(user1, user1.ID, user1.Name)

	user1.Age = 30

	fmt.Println(user1, user1.ID, user1.Name)

	user2 := User{
		Name:  "pqr",
		Email: "pqr22gmail.com",
	}

	fmt.Println("User 2", user2)

}
