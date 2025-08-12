package dto

type Pagination[T any] struct {
	Contents     T   `json:"contents"`
	Page         int `json:"page"`
	Size         int `json:"size"`
	TotalPage    int `json:"total_page"`
	TotalElement int `json:"total_element"`
}
