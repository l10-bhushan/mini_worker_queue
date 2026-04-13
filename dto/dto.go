package dto

type JobCreationRequest struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}

type JobCreationSuccess struct {
	Status bool `json:"status"`
	Data   any  `json:"data"`
}
