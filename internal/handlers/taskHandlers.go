package handlers

import (
	"Task/internal/taskService"
	"Task/internal/web/tasks"
	"context"
)

type Handler struct {
	Service *taskService.TaskService
}

func (h *Handler) DeleteTasksId(ctx context.Context, request tasks.DeleteTasksIdRequestObject) (tasks.DeleteTasksIdResponseObject, error) {

	err := h.Service.DeleteTaskByID(uint(request.Id))
	if err != nil {
		return nil, err
	}

	deleteTask := tasks.DeleteTasksId204Response{}

	return &deleteTask, nil
}

func (h *Handler) PutTasksId(ctx context.Context, request tasks.PutTasksIdRequestObject) (tasks.PutTasksIdResponseObject, error) {

	updateTask := taskService.Task{
		Task:   *request.Body.Task,
		IsDone: *request.Body.IsDone,
	}

	updateTask, err := h.Service.UpdateTaskByID(uint(request.Id), updateTask)
	if err != nil {
		return nil, err
	}

	response := tasks.PutTasksId200JSONResponse{
		Id:     &updateTask.ID,
		Task:   &updateTask.Task,
		IsDone: &updateTask.IsDone,
	}
	return response, nil
}

func (h *Handler) GetTasks(_ context.Context, _ tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	allTasks, err := h.Service.GetAllTasks()
	if err != nil {
		return nil, err
	}

	responce := tasks.GetTasks200JSONResponse{}

	for _, task := range allTasks {
		responce = append(responce, tasks.Task{
			Id:     &task.ID,
			Task:   &task.Task,
			IsDone: &task.IsDone})
	}
	return responce, nil
}

func (h *Handler) PostTasks(_ context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {

	createNewTask := taskService.Task{
		Task:   *request.Body.Task,
		IsDone: *request.Body.IsDone,
	}

	newTask, err := h.Service.CreateTask(createNewTask)
	if err != nil {
		return nil, err
	}

	response := tasks.PostTasks201JSONResponse{
		Id:     &newTask.ID,
		Task:   &newTask.Task,
		IsDone: &newTask.IsDone,
	}

	return response, nil
}

// Нужна для создания структуры Handler на этапе инициализации приложения

func NewHandler(service *taskService.TaskService) *Handler {
	return &Handler{
		Service: service,
	}
}

//func (h *Handler) GetTasksHandler(w http.ResponseWriter, r *http.Request) {
//	tasks, err := h.Service.GetAllTasks()
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//	}
//	w.Header().Set("Content-Type", "application/json")
//	json.NewEncoder(w).Encode(tasks)
//}
//
//func (h *Handler) PostTaskHandler(w http.ResponseWriter, r *http.Request) {
//	var task taskService.Task
//	err := json.NewDecoder(r.Body).Decode(&task)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusBadRequest)
//		return
//	}
//
//	createdTask, err := h.Service.CreateTask(task)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//
//	w.Header().Set("Content-Type", "application/json")
//	json.NewEncoder(w).Encode(createdTask)
//}
//
//func (h *Handler) PutTaskHandler(w http.ResponseWriter, r *http.Request) {
//	var task taskService.Task
//	id, err := strconv.Atoi(mux.Vars(r)["id"])
//
//	if err != nil {
//		http.Error(w, "Invalid ID", http.StatusBadRequest)
//		return
//	}
//
//	err2 := json.NewDecoder(r.Body).Decode(&task)
//	if err2 != nil {
//		http.Error(w, err2.Error(), http.StatusBadRequest)
//		return
//	}
//
//	updateTask, err3 := h.Service.UpdateTaskByID(uint(id), task)
//	if err3 != nil {
//		http.Error(w, err3.Error(), http.StatusInternalServerError)
//		return
//	}
//
//	w.Header().Set("Content-Type", "application/json")
//	json.NewEncoder(w).Encode(updateTask)
//}
//
//func (h *Handler) DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
//	id, err := strconv.Atoi(mux.Vars(r)["id"])
//	if err != nil {
//		http.Error(w, "Invalid ID", http.StatusBadRequest)
//		return
//	}
//
//	err2 := h.Service.DeleteTaskByID(uint(id))
//	if err2 != nil {
//		http.Error(w, err2.Error(), http.StatusInternalServerError)
//		return
//	}
//	w.WriteHeader(http.StatusNoContent)
//}
