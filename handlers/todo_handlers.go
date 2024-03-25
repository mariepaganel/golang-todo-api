package handlers

import (
	"github.com/labstack/echo/v4"
	"marie_paganel/todo/models"
	"net/http"
	"strconv"
	"sync"
)

var (
	todos = map[int]*models.Todo{}
	seq   = 1
	lock  = sync.Mutex{}
)

func CreateTodo(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()
	todo := &models.Todo{
		ID: seq,
	} // create a new item

	// parse request for data
	err := c.Bind(todo)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	// Add the new item to the in-memory store
	todos[todo.ID] = todo
	seq++

	return c.JSON(http.StatusOK, todo)
}

func GetTodos(c echo.Context) error {
	return c.JSON(http.StatusOK, todos)
}

func ReadTodo(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid todo ID"})
	}
	return c.JSON(http.StatusOK, todos[id])
}

func UpdateTodo(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()
	todo := new(models.Todo)
	if err := c.Bind(todo); err != nil {
		return err
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid todo ID"})
	}
	todos[id].Title = todo.Title
	todos[id].Description = todo.Description
	todos[id].Completed = todo.Completed
	return c.JSON(http.StatusOK, todos[id])
}

func DeleteTodo(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid todo ID"})
	}
	delete(todos, id)
	return c.NoContent(http.StatusNoContent)
}
