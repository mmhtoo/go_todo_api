package mappers

import (
	"github.com/mmhtoo/go-todo-api/entities"
	"github.com/mmhtoo/go-todo-api/internal/database"
)

func MapFromDBTodoToEntityTodo(dbTodo *database.Todo) entities.Todo {
	return entities.Todo{
		Id:        int(dbTodo.ID),
		Title:     dbTodo.Title,
		Status:    dbTodo.Status,
		CreatedAt: dbTodo.CreatedAt,
		UpdatedAt: dbTodo.UpdatedAt,
	}
}

func MapFromDBTodoListToEntityTodoList(dbTodos *[]database.Todo) []entities.Todo {
	todos := []entities.Todo{}
	for _, todo := range *dbTodos {
		todos = append(todos, MapFromDBTodoToEntityTodo(&todo))
	}
	return todos
}
