package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type CatFaceResponse struct {
	Fact   string `json:"fact"`
	Length int    `json:"length"`
}

func writeJson(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(status)

	_ = json.NewEncoder(w).Encode(data)
}

func fetchCatFact() (CatFaceResponse, error) {

	url := "https://catfact.ninja/fact"

	res, err := http.Get(url)

	if err != nil {
		return CatFaceResponse{}, err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return CatFaceResponse{}, fmt.Errorf("external api failed: %s", res.Status)
	}

	bodyBytes, err := io.ReadAll(res.Body)

	if err != nil {
		return CatFaceResponse{}, err
	}

	var data CatFaceResponse

	if err := json.Unmarshal(bodyBytes, &data); err != nil {
		return CatFaceResponse{}, err
	}

	return data, nil
}

func externalHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJson(w, http.StatusMethodNotAllowed, map[string]any{
			"ok":    false,
			"error": "only get method is allowed",
		})
		return
	}

	data, err := fetchCatFact()

	if err != nil {
		writeJson(w, http.StatusBadGateway, map[string]any{
			"ok":    false,
			"error": "failed to fetch data",
		})
	}

	writeJson(w, http.StatusOK, map[string]any{

		"ok":        true,
		"timeStamp": time.Now().UTC(),
		"external": map[string]any{
			"source": "catfact.ninja",
			"fact":   data.Fact,
			"length": data.Length,
		},
	})

}

func main() {

	http.HandleFunc("/external", externalHandler)
	err := http.ListenAndServe(":8080", nil)

	fmt.Println(err)
}
