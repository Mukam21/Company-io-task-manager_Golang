package router

import (
	"net/http"

	"github.com/Mukam21/io-bound-task-api_Golang/pkg/service"
	"github.com/Mukam21/io-bound-task-api_Golang/pkg/transport/http_trans"
)

func NewRouter(service service.TaskService) http.Handler {
	handler := http_trans.NewTaskHandler(service)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /tasks", handler.CreateTask)
	mux.HandleFunc("GET /tasks", handler.GetAllTasks)
	mux.HandleFunc("GET /tasks/{id}", handler.GetTask)
	mux.HandleFunc("DELETE /tasks/{id}", handler.DeleteTask)

	return withMiddleware(mux)
}

func withMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
