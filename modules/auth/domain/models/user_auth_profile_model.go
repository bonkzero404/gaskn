package models

type UserAuthProfileResponse struct {
	ID       string `json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	IsActive bool   `json:"is_active"`
}