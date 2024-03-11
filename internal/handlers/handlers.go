package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/Enotisi/go-rest-api-homework/internal/taskbase"
	"github.com/go-chi/chi/v5"
)

// Метод для получения всех задач
func GetTasks(w http.ResponseWriter, r *http.Request) {

	taskList := taskbase.List()

	resp, err := json.Marshal(taskList)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write(resp)
}

// Метод для добавления новой задачи
func AddTask(w http.ResponseWriter, r *http.Request) {

	var task taskbase.Task

	data, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err = json.Unmarshal(data, &task); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	taskbase.Add(task)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

// Метод для получения задачи по ID
func GetTaskById(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")

	task, err := taskbase.ById(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp, err := json.Marshal(task)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write(resp)

}

// Метод для удаления задачи по ID
func DeleteTaskById(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")

	err := taskbase.Delete(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
