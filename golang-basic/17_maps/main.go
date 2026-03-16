package main

import "fmt"

func main() {

	//map[keyType]valueType

	ages := map[string]int{
		"abc": 10,
		"pqr": 35,
		"xyz": 40,
	}

	fmt.Println(ages["abc"])
	fmt.Println(len(ages))

	scores := make(map[string]int)

	scores["math"] = 90

	fmt.Println(scores, scores["math"])

	users := map[string]string{
		"u1": "s1",
		"u2": "s2",
		"u3": "s3",
		"u4": "s4",
	}

	fmt.Println(users)

	delete(users, "u4")
	delete(users, "u33")

	fmt.Println(users)
}
