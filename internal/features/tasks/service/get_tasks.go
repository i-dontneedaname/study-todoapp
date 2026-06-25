package tasks_service

import (
	"context"
	"fmt"

	"github.com/i-dontneedaname/study-todoapp/internal/core/domain"
	core_errors "github.com/i-dontneedaname/study-todoapp/internal/core/errors"
)

func (s *TaskService) GetTasks(ctx context.Context, userId *int, limit *int, offset *int) ([]domain.Task, error) {
	if limit != nil && *limit < 0 {
		return nil, fmt.Errorf("limit must be non-negative: %w", core_errors.ErrInvalidArg)
	}
	if offset != nil && *offset < 0 {
		return nil, fmt.Errorf("limit must be non-negative: %w", core_errors.ErrInvalidArg)
	}

	tasksDomains, err := s.tasksRepository.GetTasks(ctx, userId, limit, offset)
	if err != nil {
		return []domain.Task{}, fmt.Errorf("get tasks from repository: %w", err)
	}

	return tasksDomains, nil
}
