package model

import "time"

type JobStatus string

var (
	JobStatusPending    JobStatus = "pending"
	JobStatusCompleted  JobStatus = "completed"
	JobStatusFailed     JobStatus = "failed"
	JobStatusProcessing JobStatus = "processing"
)

// Created a job struct to describe job data
type Job struct {
	Id           string     `json:"id"`
	Type         string     `json:"type"`
	Description  string     `json:"description"`
	Status       JobStatus  `json:"status"`
	Created_at   time.Time  `json:"created_at"`
	Completed_at *time.Time `json:"completed_at,omitempty"`
}
