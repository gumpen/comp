package main

import (
	"net/http"
	"os"

	"github.com/gumpen/comp/server/handler"
)

func main() {
	var PORT string
	if PORT = os.Getenv("PORT"); PORT == "" {
		PORT = "3001"
	}

	http.HandleFunc("/hello", handler.Hello)
	http.HandleFunc("/", handler.Hoge)

	http.ListenAndServe(":"+PORT, nil)
}
