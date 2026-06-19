package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from fshare!"))
	})

	http.ListenAndServe(":8000", nil)
}
