package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func postTask(w http.ResponseWriter, r *http.Request) {
	var task = Task{}
	json.NewDecoder(r.Body).Decode(&task)
	DB.Create(&task)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func getTask(w http.ResponseWriter, r *http.Request) {
	var task []Task
	DB.Find(&task)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func putTask(w http.ResponseWriter, r *http.Request) {
	var task Task
	var id = mux.Vars(r)["id"]
	json.NewDecoder(r.Body).Decode(&task)

	ID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	DB.Model(&task).Where("id = ?", ID).Updates(task)

	DB.First(&task, ID)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	var task Task
	var id = mux.Vars(r)["id"]
	ID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	DB.Find(&task, ID)
	DB.Delete(&task)
	w.WriteHeader(http.StatusNoContent)
}

func main() {

	InitDB()

	DB.AutoMigrate(&Task{})

	router := mux.NewRouter()

	router.HandleFunc("/api/tasks", getTask).Methods("GET")
	router.HandleFunc("/api/tasks", postTask).Methods("POST")
	router.HandleFunc("/api/tasks/{id}", putTask).Methods("PUT")
	router.HandleFunc("/api/tasks/{id}", deleteTask).Methods("DELETE")

	http.ListenAndServe("localhost:8082", router)

}
