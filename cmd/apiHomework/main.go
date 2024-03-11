package main

import (
	"fmt"
	"net/http"

	"github.com/Enotisi/go-rest-api-homework/internal/handlers"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	r.Get("/tasks", handlers.GetTasks)
	r.Post("/tasks", handlers.AddTask)
	r.Get("/tasks/{id}", handlers.GetTaskById)
	r.Delete("/tasks/{id}", handlers.DeleteTaskById)

	if err := http.ListenAndServe(":8080", r); err != nil {
		fmt.Printf("Ошибка при запуске сервера: %s", err.Error())
		return
	}
}
