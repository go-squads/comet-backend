package domain

import (
	"net/http"
)

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func SuccessResponse() Response {
	return Response{Status: http.StatusCreated, Message: http.StatusText(http.StatusCreated)}
}

func FailedResponse(err error) Response {
	return Response{Status: http.StatusBadRequest, Message: err.Error()}
}
