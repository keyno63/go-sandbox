package main

import (
	"encoding/json"
	"fmt"
	//"html"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Person struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Birthday int    `json:"birthday"`
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", baseHandler)

	router := httprouter.New()
	router.GET("/:path", jsonHandler)
	//router.GET("/sample/:path", jsonHandler)
	//http.ListenAndServe("127.0.0.1:3000", mux)

	http.ListenAndServe("127.0.0.1:3000", router)
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

func jsonHandler(w http.ResponseWriter, q *http.Request, p httprouter.Params) {
	fmt.Fprintf(w, "Hello, %q", p.ByName("user"))
}
