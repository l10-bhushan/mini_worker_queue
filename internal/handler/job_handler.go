package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/l10-bhushan/mini_worker_queue/dto"
	"github.com/l10-bhushan/mini_worker_queue/internal/service"
)

// struct to hold the dependency injection for Job service
type JobHandler struct {
	service *service.JobService
}

func (handler *JobHandler) handleError(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application.json")
	switch {
	case errors.Is(err, service.ErrorJobNotFound):
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Job not found",
		})
	case errors.Is(err, service.ErrorInternalServerError):
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Internal server error",
		})
	case errors.Is(err, service.ErrorBadRequest):
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "please check the request shcema",
		})
	default:
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Internal server error, please contact support",
		})
	}
}

// constructor that will return us an instance of JobHandler
func NewJobHandler(service *service.JobService) *JobHandler {
	return &JobHandler{
		service: service,
	}
}

// Handler for getting all jobs
func (handler *JobHandler) GetAllJobs(w http.ResponseWriter, r *http.Request) {
	data, err := handler.service.GetAllJobs(r.Context())
	if err != nil {
		handler.handleError(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(dto.JobCreationSuccess{
		Status: true,
		Data:   data,
	})
}

// Handler for creating a job
func (handler *JobHandler) CreateJob(w http.ResponseWriter, r *http.Request) {
	var request dto.JobCreationRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		handler.handleError(w, service.ErrorBadRequest)
		return
	}
	result, err := handler.service.CreateJob(r.Context(), request.Type, request.Description)
	if err != nil {
		handler.handleError(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(dto.JobCreationSuccess{
		Status: true,
		Data:   result,
	})
}
