package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	Init()

	err := http.ListenAndServe(":12322", nil)
	if err != nil {
		log.Fatal("Server error:", err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	user := r.URL.Query().Get("user")
	prompt := r.URL.Query().Get("prompt")
	response := GptGet(user, prompt)
	fmt.Fprint(w, response)
}
