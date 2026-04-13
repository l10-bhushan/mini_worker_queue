package service

import (
	"context"

	"github.com/l10-bhushan/mini_worker_queue/internal/model"
	"github.com/l10-bhushan/mini_worker_queue/internal/repository"
)

type JobService struct {
	repo *repository.InMemoryJobDb
}

func NewJobService(repo *repository.InMemoryJobDb) *JobService {
	return &JobService{
		repo: repo,
	}
}

func (service *JobService) CreateJob(ctx context.Context, job model.Job) {
	_, err := service.repo.CreateJob(ctx, job)
	if err != nil {
		return
	}
}

func (service *JobService) GetAllJobs(ctx context.Context) {
	err := service.repo.GetAllJobs(ctx)
	if err != nil {
		return
	}
}
