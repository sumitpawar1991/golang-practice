package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only get method is allowed", http.StatusMethodNotAllowed)
		return
	}

	_, _ = w.Write([]byte("Hello from go net/http server"))
}
func main() {

	http.HandleFunc("/hello", helloHandler)

	fmt.Println("try to going to 8080 poart")

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println(err)
	}

}
