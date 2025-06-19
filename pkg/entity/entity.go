package entity

import (
	"errors"
	"time"
)

type TaskStatus string

const (
	StatusPending    TaskStatus = "pending (в ожидании)"
	StatusProcessing TaskStatus = "processing (в процессе)"
	StatusCompleted  TaskStatus = "completed (завершено)"
	StatusFailed     TaskStatus = "failed (ошибка)"
)

type Task struct {
	ID          string     `json:"id"`
	Status      TaskStatus `json:"status"`
	CreatedAt   time.Time  `json:"created_at"`
	StartedAt   *time.Time `json:"started_at,omitempty"`
	CompletedAt *time.Time `json:"completed_at,omitempty"`
	Duration    *float64   `json:"duration_seconds,omitempty"`
	Result      *string    `json:"result,omitempty"`
	Error       *string    `json:"error,omitempty"`
}

func (t *Task) Validate() error {
	if t.Status == StatusCompleted && t.CompletedAt == nil {
		return errors.New("completed task must have completion time")
	}
	return nil
}
