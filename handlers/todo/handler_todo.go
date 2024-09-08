package todo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	"github.com/mmhtoo/go-todo-api/helpers"
	"github.com/mmhtoo/go-todo-api/internal/database"
	"github.com/mmhtoo/go-todo-api/mappers"
	"github.com/mmhtoo/go-todo-api/types"
)

type createTodoDto struct {
	Title  string `json:"title"`
	Status string `json:"status"`
}

func HandleCreateTodo(apiConfig *types.ApiConfig) types.RouteHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		todoDto := createTodoDto{}
		error := decoder.Decode(&todoDto)
		if error != nil {
			helpers.NewErrorResponse(
				w,
				http.StatusBadRequest,
				"Invalid request body to process!",
				struct{}{},
			)
			return
		}
		savedTodo, saveTodoError := apiConfig.DB.Save(
			r.Context(), database.SaveParams{
				Title:     todoDto.Title,
				Status:    todoDto.Status,
				CreatedAt: time.Now().UTC(),
				UpdatedAt: time.Now().UTC(),
			},
		)
		if saveTodoError != nil {
			fmt.Println(saveTodoError)
			helpers.NewErrorResponse(
				w,
				http.StatusInternalServerError,
				"Failed to save todo!",
				saveTodoError,
			)
			return
		}
		helpers.NewDataResponse(
			w,
			http.StatusCreated,
			"Success!",
			mappers.MapFromDBTodoToEntityTodo(&savedTodo),
		)
	}
}

func HandleDeleteTodo(apiConfig *types.ApiConfig) types.RouteHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		todoId := chi.URLParam(r, "todoId")
		parsedTodoId, error := strconv.Atoi(todoId)
		if todoId == "" || error != nil {
			helpers.NewErrorResponse(
				w,
				http.StatusBadRequest,
				"Invalid todo id!",
				struct{}{},
			)
			return
		}
		deleteError := apiConfig.DB.DeleteById(r.Context(), int32(parsedTodoId))
		if deleteError != nil {
			helpers.NewErrorResponse(
				w,
				http.StatusInternalServerError,
				"Failed to delete!",
				struct{}{},
			)
		}
		w.WriteHeader(http.StatusOK)
	}
}

type updateTodoDto struct {
	createTodoDto
}

func HandleUpdateTodoById(apiConfig *types.ApiConfig) types.RouteHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		todoId := chi.URLParam(r, "todoId")
		parsedTodoId, error := strconv.Atoi(todoId)
		if todoId == "" || error != nil {
			helpers.NewErrorResponse(
				w,
				http.StatusBadRequest,
				"Invalid todo id!",
				struct{}{},
			)
			return
		}
		decoder := json.NewDecoder(r.Body)
		payload := updateTodoDto{}
		decodeErr := decoder.Decode(&payload)
		if decodeErr != nil {
			helpers.NewErrorResponse(
				w,
				http.StatusBadRequest,
				"Invalid payload to handle!",
				struct{}{},
			)
			return
		}
		updatedTodo, updateError := apiConfig.DB.UpdateById(
			r.Context(),
			database.UpdateByIdParams{
				ID:     int32(parsedTodoId),
				Title:  payload.Title,
				Status: payload.Status,
			},
		)
		if updateError != nil {
			helpers.NewErrorResponse(
				w,
				http.StatusInternalServerError,
				"Failed to update!",
				struct{}{},
			)
			return
		}
		helpers.NewDataResponse(
			w,
			http.StatusOK,
			"Success!",
			mappers.MapFromDBTodoToEntityTodo(&updatedTodo),
		)
		return
	}
}

func HandleGetTodoById(apiConfig *types.ApiConfig) types.RouteHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		todoId := chi.URLParam(r, "todoId")
		parsedTodoId, error := strconv.Atoi(todoId)
		if todoId == "" || error != nil {
			helpers.NewErrorResponse(
				w,
				http.StatusBadRequest,
				"Invalid todo id!",
				struct{}{},
			)
			return
		}
		retrivedTodo, error := apiConfig.DB.FindById(r.Context(), int32(parsedTodoId))
		if error != nil {
			helpers.NewErrorResponse(
				w,
				http.StatusNotFound,
				"Content not found!",
				struct{}{},
			)
			return
		}
		helpers.NewDataResponse(
			w, http.StatusOK,
			"Success!",
			mappers.MapFromDBTodoToEntityTodo(&retrivedTodo),
		)
		return
	}
}

func HandleGetTodoList(apiConfig *types.ApiConfig) types.RouteHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		selectedTodos, error := apiConfig.DB.FindAll(r.Context())
		if error != nil {
			helpers.NewErrorResponse(
				w,
				http.StatusInternalServerError,
				"Failed to retrieve!",
				struct{}{},
			)
		}
		helpers.NewDataResponse(
			w,
			http.StatusOK,
			"Success!",
			mappers.MapFromDBTodoListToEntityTodoList(&selectedTodos),
		)
		return
	}
}
