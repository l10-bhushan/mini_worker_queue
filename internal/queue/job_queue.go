package queue

import "github.com/l10-bhushan/mini_worker_queue/internal/model"

type JobQueue struct {
	Jobs chan model.Job
}

func NewJobQueue(buffer int) *JobQueue {
	return &JobQueue{
		Jobs: make(chan model.Job, buffer),
	}
}
