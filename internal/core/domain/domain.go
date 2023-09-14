package domain

type Info struct {
	Version     string `json:"version"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type PaginatedResponse struct {
	Offset    int64       `json:"offset"`
	Limit     int64       `json:"limit"`
	Total     int64       `json:"total"`
	Documents interface{} `json:"documents"`
}
