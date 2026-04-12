package model

// Created a job struct to describe job data
type Job struct {
	Id           string `json:"id"`
	Type         string `json:"type"`
	Status       string `json:"status"`
	Created_at   string `json:"created_at"`
	Completed_at string `json:"completed_at"`
}
