package service

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/l10-bhushan/mini_worker_queue/internal/model"
	"github.com/l10-bhushan/mini_worker_queue/internal/repository"
)

var (
	ErrorJobNotFound         = errors.New("job not found")
	ErrorBadRequest          = errors.New("please check request")
	ErrorInternalServerError = errors.New("internal server error")
	ErrorJobCreationFailed   = errors.New("job creation failed")
)

// struct to hold the dependency injection of repository layer
type JobService struct {
	repo *repository.PostgresDb
}

// constructor function that returns us an instance of Jobservice struct
func NewJobService(repo *repository.PostgresDb) *JobService {
	return &JobService{
		repo: repo,
	}
}

// Service for getting all jobs
func (service *JobService) GetAllJobs(ctx context.Context) []model.Job {
	data := service.repo.GetAllJobs(ctx)
	return data
}

// Service for creating a job
func (service *JobService) CreateJob(ctx context.Context, typ string, description string) (model.Job, error) {
	job := model.Job{
		Id:           uuid.New().String(),
		Type:         typ,
		Description:  description,
		Status:       "processing",
		Created_at:   time.Now().String(),
		Completed_at: "",
	}
	result, err := service.repo.CreateJob(ctx, job)
	if err != nil {
		return model.Job{}, err
	}
	return result, nil
}
