package main

import (
	"encoding/json"
	"fmt"
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

func putTask(w http.ResponseWriter, r *http.Request) {
	var task Task
	json.NewDecoder(r.Body).Decode(&task)
	DB.Model(&task).Updates(task)
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	var task Task
	json.NewDecoder(r.Body).Decode(&task)
	id := task.ID
	DB.Delete(&task)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Task with ID %d deleted", id)
}

func main() {

	InitDB()

	DB.AutoMigrate(&Task{})

	router := mux.NewRouter()

	router.HandleFunc("/api/task", getTask).Methods("GET")
	router.HandleFunc("/api/task", postTask).Methods("POST")
	router.HandleFunc("/api/task", putTask).Methods("PUT")
	router.HandleFunc("/api/task", deleteTask).Methods("DELETE")

	http.ListenAndServe("localhost:8082", router)

}
