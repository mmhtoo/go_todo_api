package todo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

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

func HandleDeleteTodo(w http.ResponseWriter, r *http.Request) {

}

func HandleUpdateTodo(w http.ResponseWriter, r *http.Request) {

}

func HandleGetTodoById(w http.ResponseWriter, r *http.Request) {

}

func HandleGetTodoList(w http.ResponseWriter, r *http.Request) {

}

func HandleGetTodoByStatus(w http.ResponseWriter, r *http.Request) {

}
