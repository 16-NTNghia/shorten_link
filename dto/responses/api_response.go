package responses

import (
	"net/http"
)

type ApiResponse[T any] struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	Data       *T     `json:"data"`
}

func SuccessResponse[T any](data T) ApiResponse[T] {
	return ApiResponse[T]{
		StatusCode: http.StatusOK,
		Message:    "Success",
		Data:       &data,
	}
}

func ErrorResponse[T any](e error) ApiResponse[T] {
	return ApiResponse[T]{
		StatusCode: http.StatusBadRequest,
		Message:    e.Error(),
	}
}
