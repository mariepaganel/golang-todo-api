package main

import (
	"marie_paganel/todo/handlers"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.POST("/todo", handlers.CreateTodo)
	e.GET("/todos", handlers.GetTodos)
	e.GET("/todo/:id", handlers.ReadTodo)
	e.PUT("/todo/:id", handlers.UpdateTodo)
	e.DELETE("/todo/:id", handlers.DeleteTodo)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "My Todo App")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
