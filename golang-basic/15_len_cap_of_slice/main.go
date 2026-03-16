package main

import "fmt"

func main() {

	//type len capactity
	scores := make([]int, 0, 5) //if hit cap then cap gets double

	fmt.Println(scores, len(scores), cap(scores))

	scores = append(scores, 10)
	scores = append(scores, 20)
	scores = append(scores, 30)
	scores = append(scores, 40)
	scores = append(scores, 50)

	fmt.Println("before cap", scores, len(scores), cap(scores))

	scores = append(scores, 60)

	fmt.Println("before cap exceed ", scores, len(scores), cap(scores))

	todos := []string{"task 1", "tasl 2", "tasl 3", "task 4"}

	more := []string{"task 5"}

	todos = append(todos, more...)

	fmt.Println(todos)
}
