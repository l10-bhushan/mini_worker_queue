package model

import "time"

// Created a job struct to describe job data
type Job struct {
	Id           string     `json:"id"`
	Type         string     `json:"type"`
	Description  string     `json:"description"`
	Status       string     `json:"status"`
	Created_at   time.Time  `json:"created_at"`
	Completed_at *time.Time `json:"completed_at,omitempty"`
}
