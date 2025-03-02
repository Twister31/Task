package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func postTask(w http.ResponseWriter, r *http.Request) {
	var task = Task{}
	json.NewDecoder(r.Body).Decode(&task)
	DB.Create(&task)
}

func getTask(w http.ResponseWriter, r *http.Request) {
	var task []Task
	DB.Find(&task)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func main() {

	InitDB()

	DB.AutoMigrate(&Task{})

	router := mux.NewRouter()

	router.HandleFunc("/api/hello", getTask).Methods("GET")
	router.HandleFunc("/api/task", postTask).Methods("POST")

	http.ListenAndServe("localhost:8082", router)

}
