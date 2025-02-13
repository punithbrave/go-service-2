// main.go
package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type Response struct {
	Message string `json:"message"`
	Pod     string `json:"pod"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	podName := os.Getenv("HOSTNAME")
	response := Response{
		Message: "Hello from Argo CD deployed service!",
		Pod:     podName,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/", handler)
	log.Println("Server starting on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
