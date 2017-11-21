package main

import (
	"math/rand"
	"fmt"
	"net/http"
	"encoding/json"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	j, err := json.Marshal(rand.Intn(1001))

	if err == nil {
		fmt.Fprintf(w, "Número randomico %v!", j)
	}

}
