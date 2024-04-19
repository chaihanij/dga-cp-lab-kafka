package utils

import (
	"dga-cp-lab-kafka/app/errors"
	"encoding/json"
	"github.com/gin-gonic/gin"
	errs "github.com/pkg/errors"
	"net/http"
	"time"
)

const (
	StatusFail    string = "fail"
	StatusSuccess string = "success"
)

const (
	MessageOk string = "OK"
)

type BaseSuccessResponse struct {
	Code      string      `json:"code,omitempty" example:"10000"`
	Status    string      `json:"status" example:"success"`
	Message   string      `json:"message" example:"OK"`
	Timestamp time.Time   `json:"timestamp" example:"2021-08-10T15:00:00Z"`
	Data      interface{} `json:"data,omitempty"`
}

func (r BaseSuccessResponse) ToJSON() ([]byte, error) {
	return json.Marshal(r)
}

func NewSuccessResponse(data interface{}) BaseSuccessResponse {
	r := BaseSuccessResponse{
		Status:    StatusSuccess,
		Message:   MessageOk,
		Timestamp: time.Now(),
		Data:      data,
	}
	return r
}

type ErrorResponse struct {
	Code      string      `json:"code" example:"10000"`
	Status    string      `json:"status" example:"fail"`
	Message   string      `json:"message" example:"error message will be show here"`
	Timestamp time.Time   `json:"timestamp" example:"2021-08-10T15:00:00Z"`
	Data      interface{} `json:"data,omitempty"`
}

func (r ErrorResponse) ToJSON() ([]byte, error) {
	return json.Marshal(r)
}

func NewErrorResponse(message string) ErrorResponse {
	errorResponse := ErrorResponse{}
	errorResponse.Status = StatusFail
	errorResponse.Message = message
	return errorResponse
}

// JSONErrorResponse response error json
func JSONErrorResponse(c *gin.Context, err error, data ...interface{}) {
	statusCode := http.StatusInternalServerError
	message := ""
	code := ""

	switch err.(type) {
	case errors.ParameterError:
		statusCode = http.StatusBadRequest
		message = err.Error()
		code = errors.ErrCodeBadRequest

		var tempErr errors.ParameterError
		if errs.As(err, &tempErr) {
			code = tempErr.Code
		}
	case errors.UnprocessableEntity:
		statusCode = http.StatusBadRequest
		message = err.Error()
		code = errors.ErrCodeBadRequest

		var tempErr errors.UnprocessableEntity
		if errs.As(err, &tempErr) {
			code = tempErr.Code
		}
	case errors.InternalError:
		statusCode = http.StatusInternalServerError
		message = err.Error()
		code = errors.ErrCodeInternalServerError

		var tempErr errors.InternalError
		if errs.As(err, &tempErr) {
			code = tempErr.Code
		}
	case errors.RecordNotFoundError:
		statusCode = http.StatusNotFound
		message = err.Error()
		code = errors.ErrCodeNotFound

		var tempErr errors.RecordNotFoundError
		if errs.As(err, &tempErr) {
			code = tempErr.Code
		}
	case errors.Unauthorized:
		statusCode = http.StatusUnauthorized
		message = err.Error()
		code = errors.ErrCodeUnauthorized

		var tempErr errors.Unauthorized
		if errs.As(err, &tempErr) {
			code = tempErr.Code
		}
	case errors.Forbidden:
		statusCode = http.StatusForbidden
		message = err.Error()
		code = errors.ErrCodeForbidden

		var tempErr errors.Forbidden
		if errs.As(err, &tempErr) {
			code = tempErr.Code
		}
	default:
		message = err.Error()
		code = errors.ErrCodeInternalServerError
	}

	errorResponse := ErrorResponse{
		Code:      code,
		Status:    StatusFail,
		Message:   message,
		Timestamp: time.Now(),
		Data:      data,
	}

	c.AbortWithStatusJSON(statusCode, errorResponse)
}

func JSONSuccessResponse(c *gin.Context, data interface{}) {
	successResponse := NewSuccessResponse(data)
	c.JSON(http.StatusOK, successResponse)
}
