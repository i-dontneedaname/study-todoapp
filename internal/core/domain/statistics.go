package domain

import "time"

type Statistics struct {
	TasksCreated           int
	TasksCompleted         int
	TasksCompletedRate     *float64
	TasksAVGCompletionTime *time.Duration
}

func NewStatistics(
	tasksCreated int,
	tasksCompleted int,
	tasksCompletedRate *float64,
	tasksAVGCompletionTime *time.Duration,
) Statistics {
	return Statistics{
		TasksCreated:           tasksCreated,
		TasksCompleted:         tasksCompleted,
		TasksCompletedRate:     tasksCompletedRate,
		TasksAVGCompletionTime: tasksAVGCompletionTime,
	}
}
