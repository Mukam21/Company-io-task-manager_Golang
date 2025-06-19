package service

import (
	"context"
	"math/rand"
	"time"

	"github.com/Mukam21/io-bound-task-api_Golang/pkg/entity"
	"github.com/Mukam21/io-bound-task-api_Golang/pkg/handlers"
)

type TaskService interface {
	CreateTask(ctx context.Context) (*entity.Task, error)
	GetTask(ctx context.Context, id string) (*entity.Task, error)
	GetAllTasks(ctx context.Context) ([]*entity.Task, error)
	DeleteTask(ctx context.Context, id string) error
}

type taskService struct {
	repo     handlers.TaskHandlers
	taskChan chan string
	workers  int
}

func NewTaskService(repo handlers.TaskHandlers, workers int) TaskService {
	svc := &taskService{
		repo:     repo,
		taskChan: make(chan string, 100),
		workers:  workers,
	}
	svc.startWorkers()
	return svc
}

func (s *taskService) CreateTask(ctx context.Context) (*entity.Task, error) {
	task := &entity.Task{
		ID:        generateID(),
		Status:    entity.StatusPending,
		CreatedAt: time.Now(),
	}

	if err := s.repo.Create(task); err != nil {
		return nil, err
	}

	s.taskChan <- task.ID
	return task, nil
}

func (s *taskService) GetTask(ctx context.Context, id string) (*entity.Task, error) {
	return s.repo.GetByID(id)
}

func (s *taskService) GetAllTasks(ctx context.Context) ([]*entity.Task, error) {
	return s.repo.GetAll()
}

func (s *taskService) DeleteTask(ctx context.Context, id string) error {
	return s.repo.Delete(id)
}

func (s *taskService) startWorkers() {
	for i := 0; i < s.workers; i++ {
		go s.worker()
	}
}

func (s *taskService) worker() {
	for taskID := range s.taskChan {
		s.processTask(taskID)
	}
}

func (s *taskService) processTask(taskID string) {
	task, err := s.repo.GetByID(taskID)
	if err != nil {
		return
	}

	task.Status = entity.StatusProcessing
	now := time.Now()
	task.StartedAt = &now
	s.repo.Update(task)

	time.Sleep(3*time.Minute + time.Duration(rand.Intn(120))*time.Second)

	task.Status = entity.StatusCompleted
	completedAt := time.Now()
	task.CompletedAt = &completedAt
	duration := completedAt.Sub(*task.StartedAt).Seconds()
	task.Duration = &duration
	result := "Task completed successfully"
	task.Result = &result
	s.repo.Update(task)
}

func generateID() string {
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, 16)
	for i := range b {
		b[i] = chars[rand.Intn(len(chars))]
	}
	return "task_" + string(b)
}
