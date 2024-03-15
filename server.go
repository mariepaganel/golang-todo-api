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

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
