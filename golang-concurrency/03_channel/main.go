package main

import (
	"fmt"
	"time"
)

func main() {

	/*
	* 1) run work concurrently + collect results
	* 2) pipe -> send values between goroutines
	* 3) one goroutine sends ch < value
	* 4) another receives value := (assign) ch
	* 4) make []
	 */

	type User struct {
		ID   int
		Name string
	}

	ch := make(chan User) //user values

	go func() {
		time.Sleep(200 * time.Millisecond)

		ch <- User{ID: 100, Name: "ABC"} //send : block until main recevies
	}()

	fmt.Println("main:waiting to receive user....")

	u := <-ch

	fmt.Println("main: now go user", u.ID)
}
