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
