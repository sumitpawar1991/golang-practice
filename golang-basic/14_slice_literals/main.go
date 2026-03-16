package main

import "fmt"

func main() {

	//common colletection type
	// ---- dyanamic size

	results := []string{"abc", "pqr"}
	fmt.Println(results)

	fmt.Println(results[0]) //first result

	results[1] = "qwe"
	fmt.Println(results[len(results)-1])

	var nums []int

	nums = append(nums, 10)
	nums = append(nums, 20)
	nums = append(nums, 30)
	nums = append(nums, 40)

}
