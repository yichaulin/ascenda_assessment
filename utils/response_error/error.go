package response_error

import "ascenda_assessment/logger"

type ResponseError struct {
	Msg string `json:"msg"`
}

func New(err error) ResponseError {
	return ResponseError{
		Msg: err.Error(),
	}
}

func NewInternalServerError(err error) ResponseError {
	logger.Error(err)
	return ResponseError{
		Msg: "Internal Server Error",
	}
}
