package main

import (
	"crud-service/internal/database"
	"crud-service/internal/handlers"
	"crud-service/internal/models"

	"github.com/labstack/echo/v4"
)

func main() {
	database.InitDB()
	database.DB.AutoMigrate(&models.Task{}) // Миграция модели Task - создание таблицы Tasks

	e := echo.New()

	e.POST("/api/v1/createtask", handlers.CreateTask)
	e.GET("/api/v1/gettask", handlers.GetTask)
	e.DELETE("/api/v1/deletetask/:id", handlers.DeleteTask)
	e.PATCH("/api/v1/patchtask/:id", handlers.PatchTask)

	e.Start(":8080")
}
