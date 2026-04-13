package handler

import (
	"net/http"

	"github.com/l10-bhushan/mini_worker_queue/internal/model"
	"github.com/l10-bhushan/mini_worker_queue/internal/service"
)

type JobHandler struct {
	service *service.JobService
}

func NewJobHandler(service *service.JobService) *JobHandler {
	return &JobHandler{
		service: service,
	}
}

func (handler *JobHandler) CreateJob(w http.ResponseWriter, r *http.Request) {
	var job model.Job
	handler.service.CreateJob(r.Context(), job)
	w.WriteHeader(http.StatusOK)
}

func (handler *JobHandler) GetAllJobs(w http.ResponseWriter, r *http.Request) {
	handler.service.GetAllJobs(r.Context())
	w.WriteHeader(http.StatusOK)
}
