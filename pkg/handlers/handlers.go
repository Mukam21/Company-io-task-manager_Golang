package handlers

import (
	"errors"
	"sync"

	"github.com/Mukam21/io-bound-task-api_Golang/pkg/entity"
)

var (
	ErrTaskNotFound = errors.New("task not found")
	ErrStorageFull  = errors.New("storage limit reached")
)

type TaskHandlers interface {
	Create(task *entity.Task) error
	GetByID(id string) (*entity.Task, error)
	GetAll() ([]*entity.Task, error)
	Delete(id string) error
	Update(task *entity.Task) error
}

type inMemoryTaskHandlers struct {
	sync.RWMutex
	tasks    map[string]*entity.Task
	maxTasks int
}

func NewInMemoryTaskRepo(maxTasks int) TaskHandlers {
	return &inMemoryTaskHandlers{
		tasks:    make(map[string]*entity.Task),
		maxTasks: maxTasks,
	}
}

func (r *inMemoryTaskHandlers) Create(task *entity.Task) error {
	r.Lock()
	defer r.Unlock()

	if len(r.tasks) >= r.maxTasks {
		return ErrStorageFull
	}

	r.tasks[task.ID] = task
	return nil
}

func (r *inMemoryTaskHandlers) GetByID(id string) (*entity.Task, error) {
	r.RLock()
	defer r.RUnlock()

	task, exists := r.tasks[id]
	if !exists {
		return nil, ErrTaskNotFound
	}
	return task, nil
}

func (r *inMemoryTaskHandlers) GetAll() ([]*entity.Task, error) {
	r.RLock()
	defer r.RUnlock()

	tasks := make([]*entity.Task, 0, len(r.tasks))
	for _, task := range r.tasks {
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (r *inMemoryTaskHandlers) Delete(id string) error {
	r.Lock()
	defer r.Unlock()

	if _, exists := r.tasks[id]; !exists {
		return ErrTaskNotFound
	}
	delete(r.tasks, id)
	return nil
}

func (r *inMemoryTaskHandlers) Update(task *entity.Task) error {
	r.Lock()
	defer r.Unlock()

	if _, exists := r.tasks[task.ID]; !exists {
		return ErrTaskNotFound
	}
	r.tasks[task.ID] = task
	return nil
}
