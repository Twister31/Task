package handlers

import (
	"Task/internal/taskService"
	"Task/internal/web/tasks"
	"context"
)

type TaskHandler struct {
	Service *taskService.TaskService
}

func (h *TaskHandler) DeleteTasksId(_ context.Context, request tasks.DeleteTasksIdRequestObject) (tasks.DeleteTasksIdResponseObject, error) {

	err := h.Service.DeleteTaskByID(uint(request.Id))
	if err != nil {
		return nil, err
	}

	deleteTask := tasks.DeleteTasksId204Response{}

	return &deleteTask, nil
}

func (h *TaskHandler) PutTasksId(_ context.Context, request tasks.PutTasksIdRequestObject) (tasks.PutTasksIdResponseObject, error) {

	updateTask := taskService.Task{
		Task:   *request.Body.Task,
		IsDone: *request.Body.IsDone,
	}

	updateTaskN, err := h.Service.UpdateTaskByID(uint(request.Id), updateTask)
	if err != nil {
		return nil, err
	}

	response := tasks.PutTasksId200JSONResponse{
		Id:     &updateTaskN.ID,
		Task:   &updateTaskN.Task,
		IsDone: &updateTaskN.IsDone,
	}
	return response, nil
}

func (h *TaskHandler) GetTasks(_ context.Context, _ tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
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

func (h *TaskHandler) PostTasks(_ context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {

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

// Нужна для создания структуры TaskHandler на этапе инициализации приложения

func NewTaskHandler(service *taskService.TaskService) *TaskHandler {
	return &TaskHandler{
		Service: service,
	}
}
