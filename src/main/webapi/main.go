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
	mux.HandleFunc("/func", func(w http.ResponseWriter, q *http.Request) {
		orgHeaderValue := q.Header.Get("org-header")
		fmt.Fprintf(w, "func, headerValue: %s", orgHeaderValue)
	})
	mux.HandleFunc("/getJson", getJsonHandle)
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
	type SampleJsonResponse struct {
		Message string `json:"message"`
	}
	value := "hello world"
	//message := map[string]string{
	//	"message": "hello world",
	//}
	jsonMessage, err := json.Marshal(SampleJsonResponse{value})
	if err != nil {
		//panic(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
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

func getJsonHandle(w http.ResponseWriter, q *http.Request) {

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
