package main

import (
	"Task/internal/database"
	"Task/internal/handlers"
	"Task/internal/taskService"
	"Task/internal/userService"
	"Task/internal/web/tasks"
	"Task/internal/web/users"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

func main() {

	database.InitDB()

	// Миграции для задач
	if err := database.DB.AutoMigrate(&taskService.Task{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	// Миграции для пользователей
	if err := database.DB.AutoMigrate(&userService.User{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	taskRepo := taskService.NewTaskRepository(database.DB)
	taskService := taskService.NewService(taskRepo)
	taskHandler := handlers.NewTaskHandler(taskService)

	userRepo := userService.NewUserRepository(database.DB)
	userService := userService.NewService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	taskStrictHandler := tasks.NewStrictHandler(taskHandler, nil)
	tasks.RegisterHandlers(e, taskStrictHandler)

	userStrictHandler := users.NewStrictHandler(userHandler, nil)
	users.RegisterHandlers(e, userStrictHandler)

	if err := e.Start(":8082"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}
}
