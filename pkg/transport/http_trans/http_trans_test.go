package http_trans_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Mukam21/io-bound-task-api_Golang/pkg/handlers"
	"github.com/Mukam21/io-bound-task-api_Golang/pkg/router"
	"github.com/Mukam21/io-bound-task-api_Golang/pkg/service"
	"github.com/stretchr/testify/assert"
)

func TestTaskHandlers_Flow(t *testing.T) {
	repo := handlers.NewInMemoryTaskRepo(10)
	svc := service.NewTaskService(repo, 1)
	handler := router.NewRouter(svc)

	req := httptest.NewRequest(http.MethodPost, "/tasks", nil)
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var created map[string]interface{}
	json.NewDecoder(w.Body).Decode(&created)
	id := created["id"].(string)

	req = httptest.NewRequest(http.MethodGet, "/tasks/"+id, nil)
	w = httptest.NewRecorder()
	handler.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	req = httptest.NewRequest(http.MethodDelete, "/tasks/"+id, nil)
	w = httptest.NewRecorder()
	handler.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNoContent, w.Code)

	req = httptest.NewRequest(http.MethodGet, "/tasks/"+id, nil)
	w = httptest.NewRecorder()
	handler.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
}
