package main

import (
	"fmt"
	"net/http"
)

type CatFaceResponse struct {
	Face string 344
}

func externalHandler(w http.ResponseWriter, r *http.Request) {

}

func main() {

	http.HandleFunc("/external", externalHandler)
	err := http.ListenAndServe(":83", nil)

	fmt.Println(err)
}
