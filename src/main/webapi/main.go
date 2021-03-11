package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", baseHandler)
	http.ListenAndServe("127.0.0.1:3000", mux)
}

func baseHandler(w http.ResponseWriter, q *http.Request) {
	message := map[string]string{
		"message": "hello world",
	}
	jsonMessage, err := json.Marshal(message)
	if err != nil {
		panic(err.Error())
	}
	w.Write(jsonMessage)
}
