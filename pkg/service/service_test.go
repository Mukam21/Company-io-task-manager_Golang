package service

import (
	"context"
	"testing"
	"time"

	"github.com/Mukam21/io-bound-task-api_Golang/pkg/entity"
	"github.com/Mukam21/io-bound-task-api_Golang/pkg/handlers"
	"github.com/stretchr/testify/assert"
)

func TestTaskService_CreateAndGetTask(t *testing.T) {
	repo := handlers.NewInMemoryTaskRepo(10)
	svc := NewTaskService(repo, 1)

	task, err := svc.CreateTask(context.Background())
	assert.NoError(t, err)
	assert.NotNil(t, task)

	fetchedTask, err := svc.GetTask(context.Background(), task.ID)
	assert.NoError(t, err)
	assert.Equal(t, task.ID, fetchedTask.ID)
	assert.Equal(t, entity.StatusPending, fetchedTask.Status)
}

func TestTaskService_DeleteTask(t *testing.T) {
	repo := handlers.NewInMemoryTaskRepo(10)
	svc := NewTaskService(repo, 1)

	task, err := svc.CreateTask(context.Background())
	assert.NoError(t, err)

	err = svc.DeleteTask(context.Background(), task.ID)
	assert.NoError(t, err)

	_, err = svc.GetTask(context.Background(), task.ID)
	assert.ErrorIs(t, err, handlers.ErrTaskNotFound)
}

func TestTaskService_ProcessTask(t *testing.T) {
	repo := handlers.NewInMemoryTaskRepo(10)
	svc := TestOnlyNewTaskService(repo, 0)

	go func() {
		for {
			taskID := <-svc.TaskChan()
			task, _ := repo.GetByID(taskID)
			task.Status = entity.StatusProcessing
			now := time.Now()
			task.StartedAt = &now
			time.Sleep(100 * time.Millisecond)
			task.Status = entity.StatusCompleted
			completed := time.Now()
			task.CompletedAt = &completed
			duration := completed.Sub(now).Seconds()
			task.Duration = &duration
			result := "done"
			task.Result = &result
			repo.Update(task)
		}
	}()

	task, err := svc.CreateTask(context.Background())
	assert.NoError(t, err)

	time.Sleep(200 * time.Millisecond)

	fetchedTask, err := svc.GetTask(context.Background(), task.ID)
	assert.NoError(t, err)
	assert.Equal(t, entity.StatusCompleted, fetchedTask.Status)
	assert.NotNil(t, fetchedTask.CompletedAt)
	assert.NotNil(t, fetchedTask.Result)
}
