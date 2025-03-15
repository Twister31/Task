package main

import (
	"Task/internal/database"
	"Task/internal/handlers"
	"Task/internal/taskService"
	"Task/internal/web/tasks"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

func main() {

	database.InitDB()

	database.DB.AutoMigrate(&taskService.Task{})

	repo := taskService.NewTaskRepository(database.DB)
	service := taskService.NewService(repo)
	handler := handlers.NewHandler(service)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	strictHandler := tasks.NewStrictHandler(handler, nil)
	tasks.RegisterHandlers(e, strictHandler)
	if err := e.Start(":8082"); err != nil {
		log.Fatal("failed to start with err: %v", err)
	}
}
