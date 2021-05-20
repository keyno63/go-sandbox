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
	//mux.HandleFunc("/json", jsonHandler)
	mux.HandleFunc("/cookie", getCookie)
	http.ListenAndServe("127.0.0.1:3000", mux)

	value := "value"
	var xxx *string
	xxx = &value
	fmt.Printf("value: %s", *xxx)

	//router := httprouter.New()
	////router.GET("/", baseHandler)
	//router.GET("/sample/:path", jsonHandler)
	//router.GET("/value/sample", getHeader)
	//http.ListenAndServe("127.0.0.1:3000", router)
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

func getCookie(w http.ResponseWriter, q *http.Request) {
	name := "sample"
	value, err := parseAndGetCookieValue(q, "sample")
	if err != nil {
		fmt.Printf("err, err=[%#v]", err)
		fmt.Fprintf(w, "Hello getCookie, err=[%#v]", err)
		return
	}
	fmt.Fprintf(w, "Hello, key=[%s], value=[%s]", name, value)
}

// router handler
func jsonHandler(w http.ResponseWriter, q *http.Request, p httprouter.Params) {
	fmt.Fprintf(w, "Hello, %q", p.ByName("user"))
}

// private function
func parseAndGetCookieValue(q *http.Request, cookieName string) (string, error) {
	cookie, err := q.Cookie(cookieName)
	if err != nil {
		return "", err
	}

	return cookie.Value, nil
}
