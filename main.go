package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

var task string
var count int

func postTask(w http.ResponseWriter, r *http.Request) {
	json.NewDecoder(r.Body).Decode(&task)

}

func getTask(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s", task)
}

func GetMethod(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		fmt.Fprintf(w, "это не гет запрос!")
		return
	}

	fmt.Fprintf(w, "число = %s", strconv.Itoa(count))
}

func PostMethod(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		fmt.Fprintf(w, "это поле не пост запрос!")
		return
	}

	fmt.Fprintf(w, "увеличение значения числа на 1")
	count++
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}

func main() {

	router := mux.NewRouter()
	//router.HandleFunc("/api/hello", HelloHandler).Methods("GET")
	router.HandleFunc("/api/hello", getTask).Methods("GET")
	router.HandleFunc("/api/task", postTask).Methods("POST")
	//router.HandleFunc("/api/get", GetMethod).Methods("GET")
	//router.HandleFunc("/api/post", PostMethod).Methods("POST")

	//http.HandleFunc("/Get", GetMethod)
	//http.HandleFunc("/Post", PostMethod)

	http.ListenAndServe("localhost:8082", router)

}
