package types

import (
	"net/http"

	"github.com/mmhtoo/go-todo-api/internal/database"
)

type ApiConfig struct {
	DB *database.Queries
}

type RouteHandler = func(w http.ResponseWriter, r *http.Request)
