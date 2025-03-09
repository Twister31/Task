package main

import (
	"Task/internal/database"
	"Task/internal/handlers"
	"Task/internal/taskService"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {

	database.InitDB()

	database.DB.AutoMigrate(&taskService.Task{})

	repo := taskService.NewTaskRepository(database.DB)
	service := taskService.NewService(repo)
	handler := handlers.NewHandler(service)

	router := mux.NewRouter()

	router.HandleFunc("/api/get", handler.GetTasksHandler).Methods("GET")
	router.HandleFunc("/api/post", handler.PostTaskHandler).Methods("POST")
	router.HandleFunc("/api/put/{id}", handler.PutTaskHandler).Methods("PUT")
	router.HandleFunc("/api/delete/{id}", handler.DeleteTaskHandler).Methods("DELETE")

	http.ListenAndServe("localhost:8082", router)

}
