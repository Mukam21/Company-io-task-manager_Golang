package entity_test

import (
	"testing"
	"time"

	"github.com/Mukam21/io-bound-task-api_Golang/pkg/entity"
	"github.com/stretchr/testify/assert"
)

func TestTask_Validate(t *testing.T) {
	task := &entity.Task{
		ID:     "t1",
		Status: entity.StatusCompleted,
	}

	err := task.Validate()
	assert.Error(t, err)

	now := time.Now()
	task.CompletedAt = &now
	err = task.Validate()
	assert.NoError(t, err)
}
