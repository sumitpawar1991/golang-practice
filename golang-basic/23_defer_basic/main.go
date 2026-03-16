package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Case 1 : success")

	if err := doWork(true); err != nil {
		fmt.Println("error", err)
	}

	fmt.Println("Case 2 : falied")

	if err := doWork(false); err != nil {
		fmt.Println("error", err)
	}
}

func doWork(success bool) error {

	//resouce related
	//start message -> resource acquried
	//clearnup message -> recource released

	fmt.Println("Start : resource acquried")

	defer fmt.Println("cleanup: resource released")

	if !success {
		return errors.New("someting went wrong. i am returning error")
	}

	fmt.Println("work:doing something")
	fmt.Println("work is done")

	return nil
}
