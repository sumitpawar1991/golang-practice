package main

import "fmt"

/*
* Problem Statement : Reverse the string [handle Unicode properly]
 */
func main() {

	var name string
	name = "this is sentence use for reversing the string"

	fmt.Println("string name is - ", name)

	fmt.Printf("text reversed is - %s\n", reverse(name))
}

func reverse(s string) string {

	uni := []rune(s) //handling unicode

	//swap elements . pointers
	/*
	*  Initialization - i, j := 0, len(uni)-1;
	*  Condition -  i < j;
	*  Update - i, j = i+1, j-1 //  Move i forward / Move j backward
	*  Swap =  uni[i], uni[j] = uni[j], uni[i]
	 */
	for i, j := 0, len(uni)-1; i < j; i, j = i+1, j-1 {
		uni[i], uni[j] = uni[j], uni[i]
	}

	return string(uni)
}
