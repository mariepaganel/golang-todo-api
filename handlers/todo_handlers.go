package handlers

import (
	"github.com/labstack/echo/v4"
	"marie_paganel/todo/database"
	"marie_paganel/todo/models"
	"net/http"
	"strconv"
	"sync"
)

var lock = sync.Mutex{}

func CreateTodo(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()
	todo := new(models.Todo)

	// parse request for data
	err := c.Bind(todo)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "bad request"})
	}

	// Save to the database
	if err := database.DB.Create(todo).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to create todo"})
	}

	return c.JSON(http.StatusOK, todo)
}

func GetTodos(c echo.Context) error {
	var todos []models.Todo
	if err := database.DB.Find(&todos).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to fetch todos"})
	}
	return c.JSON(http.StatusOK, todos)
}

func ReadTodo(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid todo ID"})
	}

	// look for entry with this id
	var todo models.Todo
	database.DB.Find(&todo, id)
	return c.JSON(http.StatusOK, todo)
}

func UpdateTodo(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()
	// parse information for new data
	updatedTodo := new(models.Todo)
	if err := c.Bind(updatedTodo); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "bad request"})
	}

	// get id from context
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid todo ID"})
	}

	// Retrieve the existing to do from the database
	var existingTodo models.Todo
	if err := database.DB.First(&existingTodo, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Todo not found"})
	}

	// Update the retrieved to do with the new information
	existingTodo.Title = updatedTodo.Title
	existingTodo.Description = updatedTodo.Description
	existingTodo.Completed = updatedTodo.Completed

	// Save the updated to do back to the database
	if err := database.DB.Save(&existingTodo).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update todo"})
	}

	return c.JSON(http.StatusOK, existingTodo)
}

func DeleteTodo(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid todo ID"})
	}

	database.DB.Delete(models.Todo{}, id)
	return c.NoContent(http.StatusNoContent)
}
