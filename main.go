package main

import (
	"net/http"
)

func main() {
	println("listening on http://localhost:3000...")
	http.ListenAndServe(":3000", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		index("World").Render(r.Context(), w)
	}))
}
