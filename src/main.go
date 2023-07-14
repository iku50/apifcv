package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/iku50/apifcv/src/api"
)

func main() {
	http.HandleFunc("/", handler)
	api.Init()

	err := http.ListenAndServe(":12321", nil)
	if err != nil {
		log.Fatal("Server error:", err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	user := r.URL.Query().Get("user")
	prompt := r.URL.Query().Get("prompt")
	response := api.GptGet(user, prompt)
	fmt.Fprint(w, response)
}
