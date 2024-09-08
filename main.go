package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/mmhtoo/go-todo-api/handlers/todo"
	"github.com/mmhtoo/go-todo-api/helpers"
	"github.com/mmhtoo/go-todo-api/internal/database"
	"github.com/mmhtoo/go-todo-api/types"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	// load env file
	godotenv.Load(".env")
	// create new router
	router := chi.NewRouter()
	// config cors
	configCors(router)

	// get port from env
	port, _ := helpers.LoadFromEnv("PORT", false, "3000")
	// get db url from env
	dbURL, _ := helpers.LoadFromEnv("DB_URL", true, "")
	// open db connection
	dbConn, dbConnError := sql.Open("postgres", dbURL)
	if dbConnError != nil {
		log.Fatal("Failed to open database connnection at ", dbURL, "\n ", dbConnError)
	}
	// make queries instance
	queries := database.New(dbConn)
	apiConfig := types.ApiConfig{
		DB: queries,
	}

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		helpers.NewDataResponse(w, 200, "Success", struct{}{})
	})
	router.Post("/todos", todo.HandleCreateTodo(&apiConfig))

	server := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	fmt.Printf("Server is listening on port %s \n", port)
	serverError := server.ListenAndServe()

	if serverError != nil {
		fmt.Printf("Error: %s", serverError)
	}

}

func configCors(router *chi.Mux) {
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://*", "https://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Origin", "Content-Type", "Accept"},
		ExposedHeaders:   []string{"Content-Length", "Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
}
