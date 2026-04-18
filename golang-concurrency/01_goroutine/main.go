package main

import (
	"fmt"
	"time"
)

func main() {
	//concepts
	//doing multiple tasks in a way that they overlap in time
	//overlap -> while task A (network db), task b can run

	//overlapping ->concurrency
	//executing multiple tasks at the exact same time

	//backend ->multiple service calls ->combine all result

	start := time.Now()

	go func() {
		time.Sleep(300 * time.Millisecond)

		fmt.Println("go routine A: finish simulated API at", time.Since(start))
	}()

	go func() {
		time.Sleep(150 * time.Millisecond)

		fmt.Println("go routine B: finish simulated API at", time.Since(start))
	}()

	fmt.Println("main:started two go routine at", time.Since(start))

	fmt.Println("main: doing step 1", time.Since(start))
	time.Sleep(100 * time.Millisecond)

	fmt.Println("main: doing step 2", time.Since(start))
	time.Sleep(100 * time.Millisecond)

	fmt.Println("main: doing step 3", time.Since(start))
	time.Sleep(100 * time.Millisecond)

	time.Sleep(500 * time.Millisecond)

	fmt.Println("main:existing at", time.Since(start))
}
