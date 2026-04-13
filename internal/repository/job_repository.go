package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/l10-bhushan/mini_worker_queue/internal/model"
)

// Interface that defines all the methods on repository layer
type JobRepository interface {
	CreateJob(ctx context.Context, job model.Job) (model.Job, error)
}

// In memory db we will replace this with postgres or mongoDb in future
type InMemoryJobDb struct {
	db *pgxpool.Pool
}

// Constructor method to return instance of Job db
func NewJobRepository(db *pgxpool.Pool) *InMemoryJobDb {
	return &InMemoryJobDb{
		db: db,
	}
}

// Repo method to GetAllJobs
func (db *InMemoryJobDb) GetAllJobs(ctx context.Context) []model.Job {
	// TODO: Add query to fetch all jobs
	return []model.Job{}
}

// Repo method to create new Jobs
func (db *InMemoryJobDb) CreateJob(ctx context.Context, job model.Job) model.Job {
	// TODO: Add query to create jobs
	return job
}
