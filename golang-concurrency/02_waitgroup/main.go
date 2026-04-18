package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// waitgroup -> how many rotines you have

	var wg sync.WaitGroup

	wg.Add(3)

	go func() {
		defer wg.Done()

		fmt.Println("task 1")
		time.Sleep(250 * time.Millisecond)

		fmt.Println("task 1 is now done")

	}()

	go func() {
		defer wg.Done()

		fmt.Println("task 2")
		time.Sleep(250 * time.Millisecond)

		fmt.Println("task 2 is now done")

	}()

	go func() {
		defer wg.Done()

		fmt.Println("task 3")
		time.Sleep(250 * time.Millisecond)

		fmt.Println("task 3 is now done")

	}()

	fmt.Println("main:waiting for taks for finish")

	wg.Wait()

	fmt.Println("main:all task are finish now")

}
