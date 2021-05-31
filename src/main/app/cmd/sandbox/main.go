package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Http Handler の設定
	mux := http.NewServeMux()
	handler := NewHandler()
	mux.Handle("/api/go-app/handle", handler)

	// server の起動設定
	server := http.Server{
		Addr:    ":8180",
		Handler: mux,
	}

	// server の起動
	_ = server.ListenAndServe()
}

type Handler interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

type handler struct {
}

func NewHandler() Handler {
	return handler{}
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "sample handle")
}
