package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	handler := http.FileServer(http.Dir(wd))

	http.Handle("GET /", handler)

	http.ListenAndServe(":8000", nil)
}
