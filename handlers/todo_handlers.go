package handlers

import (
	"github.com/labstack/echo/v4"
	"marie_paganel/todo/models"
	"net/http"
)

var todos []models.Todo

func CreateTodo(c echo.Context) error {
	var todo models.Todo // create a new item

	// parse request for data
	err := c.Bind(&todo)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	// Generate unique ID for the new item
	todo.ID = len(todos) + 1

	// Add the new item to the in-memory store
	todos = append(todos, todo)

	return c.JSON(http.StatusOK, todo)
}

func GetTodos(c echo.Context) error {
	return c.JSON(http.StatusOK, todos)
}
