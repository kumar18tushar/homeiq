package main

import (
	"fmt"
	"net/http"
)

// This is the entry point Vercel looks for
func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, `{"message": "Hello from net/http on Vercel!"}`)
}
