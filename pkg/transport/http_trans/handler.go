package http_trans

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Mukam21/io-bound-task-api_Golang/pkg/handlers"
	"github.com/Mukam21/io-bound-task-api_Golang/pkg/service"
)

type TaskHandler struct {
	service service.TaskService
}

func NewTaskHandler(service service.TaskService) *TaskHandler {
	return &TaskHandler{service: service}
}

func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	task, err := h.service.CreateTask(r.Context())
	if err != nil {
		if errors.Is(err, handlers.ErrStorageFull) {
			respondWithError(w, http.StatusTooManyRequests, err)
		} else {
			respondWithError(w, http.StatusInternalServerError, err)
		}
		return
	}
	respondWithJSON(w, http.StatusCreated, task)
}

func (h *TaskHandler) GetTask(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		respondWithError(w, http.StatusBadRequest, errors.New("task ID is required"))
		return
	}

	task, err := h.service.GetTask(r.Context(), id)
	if err != nil {
		if errors.Is(err, handlers.ErrTaskNotFound) {
			respondWithError(w, http.StatusNotFound, err)
		} else {
			respondWithError(w, http.StatusInternalServerError, err)
		}
		return
	}

	respondWithJSON(w, http.StatusOK, task)
}

func (h *TaskHandler) GetAllTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := h.service.GetAllTasks(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err)
		return
	}
	respondWithJSON(w, http.StatusOK, tasks)
}

func (h *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		respondWithError(w, http.StatusBadRequest, errors.New("task ID is required"))
		return
	}

	err := h.service.DeleteTask(r.Context(), id)
	if err != nil {
		if errors.Is(err, handlers.ErrTaskNotFound) {
			respondWithError(w, http.StatusNotFound, err)
		} else {
			respondWithError(w, http.StatusInternalServerError, err)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(payload)
}

func respondWithError(w http.ResponseWriter, code int, err error) {
	respondWithJSON(w, code, map[string]string{"error": err.Error()})
}
