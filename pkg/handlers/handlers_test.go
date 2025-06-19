package handlers_test

import (
	"testing"

	"github.com/Mukam21/io-bound-task-api_Golang/pkg/entity"
	"github.com/Mukam21/io-bound-task-api_Golang/pkg/handlers"
	"github.com/stretchr/testify/assert"
)

func TestInMemoryRepo_CRUD(t *testing.T) {
	repo := handlers.NewInMemoryTaskRepo(2)

	task := &entity.Task{ID: "task1", Status: entity.StatusPending}

	err := repo.Create(task)
	assert.NoError(t, err)

	got, err := repo.GetByID("task1")
	assert.NoError(t, err)
	assert.Equal(t, task.ID, got.ID)

	task.Status = entity.StatusCompleted
	err = repo.Update(task)
	assert.NoError(t, err)

	all, err := repo.GetAll()
	assert.NoError(t, err)
	assert.Len(t, all, 1)

	err = repo.Delete("task1")
	assert.NoError(t, err)

	_, err = repo.GetByID("task1")
	assert.ErrorIs(t, err, handlers.ErrTaskNotFound)
}

func TestInMemoryRepo_Limit(t *testing.T) {
	repo := handlers.NewInMemoryTaskRepo(1)

	t1 := &entity.Task{ID: "1"}
	t2 := &entity.Task{ID: "2"}

	_ = repo.Create(t1)
	err := repo.Create(t2)

	assert.ErrorIs(t, err, handlers.ErrStorageFull)
}
