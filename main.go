package main

import (
	"log"
	"net/http"
	"os"

	"github.com/justinemmanuelmercado/avatargo/avatar"
)

func main() {
	http.Handle("/", http.HandlerFunc(handle))
	err := http.ListenAndServe(":2003", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func handle(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	avatar := avatar.NewAvatar("EJ", avatar.Options{
		Shape: avatar.Square,
		Size:  128,
	})

	f, err := os.Create("new.svg")
	if err != nil {
		log.Print(err)
	}

	avatar.Generate(w)
	avatar.Generate(f)
}
