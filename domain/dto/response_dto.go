package dto

import "fmt"

type Response struct {
	Valid bool        `json:"valid"`
	Meta  Meta        `json:"meta"`
	Error interface{} `json:"errors"`
	Data  interface{} `json:"data"`
}

type Meta struct {
	Route  string `json:"route"`
	Method string `json:"method"`
	Query  string `json:"query"`
	Code   int    `json:"code"`
	Status string `json:"status"`
}

type ErrorResponse struct {
	FailedField string `json:"field"`
	Tag         string `json:"tag"`
	Message     string `json:"message"`
}

type ApiErrorResponse struct {
	StatusCode int    `json:"code"`
	Message    string `json:"message"`
}

func (r *ApiErrorResponse) Error() string {
	return fmt.Sprintf("status %d: err %v", r.StatusCode, r.Message)
}

func (r *ApiErrorResponse) GetStatusCode() int {
	return r.StatusCode
}
