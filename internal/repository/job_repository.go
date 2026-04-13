package repository

import (
	"context"

	"github.com/l10-bhushan/mini_worker_queue/internal/model"
)

// Interface that defines all the methods on repository layer
type JobRepository interface {
	CreateJob(ctx context.Context, job model.Job) (model.Job, error)
}

// In memory db we will replace this with postgres or mongoDb in future
type InMemoryJobDb struct {
	data map[string]model.Job
}

// Constructor method to return instance of Job db
func NewJobRepository() *InMemoryJobDb {
	return &InMemoryJobDb{
		data: make(map[string]model.Job),
	}
}

// Repo method to GetAllJobs
func (db *InMemoryJobDb) GetAllJobs(ctx context.Context) []model.Job {
	var jobs []model.Job
	for _, job := range db.data {
		jobs = append(jobs, job)
	}
	return jobs
}

// Repo method to create new Jobs
func (db *InMemoryJobDb) CreateJob(ctx context.Context, job model.Job) model.Job {
	db.data[job.Id] = job
	return job
}
