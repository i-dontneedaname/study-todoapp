package tasks_service

import (
	"context"
	"fmt"

	"github.com/i-dontneedaname/study-todoapp/internal/core/domain"
)

func (s *TaskService) GetTask(ctx context.Context, id int) (domain.Task, error) {
	taskDomain, err := s.tasksRepository.GetTask(ctx, id)
	if err != nil {
		return domain.Task{}, fmt.Errorf("get task from repository: %w", err)
	}

	return taskDomain, nil
}
