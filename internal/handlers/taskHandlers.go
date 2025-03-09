package handlers

import (
	"Task/internal/taskService"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type Handler struct {
	Service *taskService.TaskService
}

// Нужна для создания структуры Handler на этапе инициализации приложения

func NewHandler(service *taskService.TaskService) *Handler {
	return &Handler{
		Service: service,
	}
}

func (h *Handler) GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	tasks, err := h.Service.GetAllTasks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func (h *Handler) PostTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task taskService.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdTask, err := h.Service.CreateTask(task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdTask)
}

func (h *Handler) PutTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task taskService.Task
	id, err := strconv.Atoi(mux.Vars(r)["id"])

	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	err2 := json.NewDecoder(r.Body).Decode(&task)
	if err2 != nil {
		http.Error(w, err2.Error(), http.StatusBadRequest)
		return
	}

	updateTask, err3 := h.Service.UpdateTaskByID(uint(id), task)
	if err3 != nil {
		http.Error(w, err3.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updateTask)
}

func (h *Handler) DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	err2 := h.Service.DeleteTaskByID(uint(id))
	if err2 != nil {
		http.Error(w, err2.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
