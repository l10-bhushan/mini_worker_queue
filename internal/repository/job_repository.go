package repository

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/l10-bhushan/mini_worker_queue/internal/model"
)

var (
	DbErrorJobCreation  = errors.New("Job creation failed")
	DbErrorFetchingJobs = errors.New("failed while fetching jobs")
	DbErrorParsingRow   = errors.New("failed while parsing database row")
)

// Interface that defines all the methods on repository layer
type JobRepository interface {
	GetAllJobs(ctx context.Context) ([]model.Job, error)
	CreateJob(ctx context.Context, job model.Job) (model.Job, error)
}

// In memory db we will replace this with postgres or mongoDb in future
type PostgresDb struct {
	db *pgxpool.Pool
}

// Constructor method to return instance of Job db
func NewJobRepository(db *pgxpool.Pool) *PostgresDb {
	return &PostgresDb{
		db: db,
	}
}

// Repo method to GetAllJobs
func (repo *PostgresDb) GetAllJobs(ctx context.Context) ([]model.Job, error) {
	// TODO: Add query to fetch all jobs
	var jobs []model.Job
	query := `SELECT * FROM jobs`
	rows, err := repo.db.Query(ctx, query)
	if err != nil {
		return []model.Job{}, DbErrorFetchingJobs
	}
	defer rows.Close()

	for rows.Next() {
		var j model.Job
		err := rows.Scan(&j.Id, &j.Type, &j.Description, &j.Status, &j.Created_at, &j.Completed_at)
		if err != nil {
			return []model.Job{}, DbErrorParsingRow
		}
		jobs = append(jobs, j)
	}

	if err = rows.Err(); err != nil {
		return []model.Job{}, DbErrorParsingRow
	}

	return jobs, nil
}

// Repo method to create new Jobs
func (repo *PostgresDb) CreateJob(ctx context.Context, job model.Job) (model.Job, error) {
	// TODO: Add query to create jobs
	query := `INSERT INTO jobs (id , type , description, status, created_at, completed_at) VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := repo.db.Exec(ctx, query, job.Id, job.Type, job.Description, job.Status, job.Created_at, job.Completed_at)
	if err != nil {
		return model.Job{}, DbErrorJobCreation
	}
	return job, nil
}

func (repo *PostgresDb) UpdateStatus(ctx context.Context, id string, status model.JobStatus) (model.Job, error) {
	return model.Job{}, nil
}
