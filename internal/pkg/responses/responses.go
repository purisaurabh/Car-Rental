package specs

import (
	"encoding/json"
	errs "errors"
	"fmt"
	"net/http"

	"go.uber.org/zap"
)

// errorResponse is a response that is returned when an error is encountered
type errorResponse struct {
	ErrorCode    int    `json:"error_code"`
	ErrorMessage string `json:"error_message"`
}

// successResponse is a response that is returned when an success is encountered
type successResponse struct {
	Data interface{} `json:"data"`
}

type MessageResponseWithID struct {
	ID      int64  `json:"id"`
	Message string `json:"message"`
}

// SuccessResponse function returns a response that is returned when an success is encountered
func SuccessResponse(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	payload := successResponse{
		Data: data,
	}

	out, err := json.Marshal(payload)
	if err != nil {
		zap.S().Error("cannot marshal success response payload")
		writeServerErrorResponse(w)
		return
	}

	_, err = w.Write(out)
	if err != nil {
		zap.S().Error("cannot write json success response")
		writeServerErrorResponse(w)
		return
	}
}

// ErrorResponse function returns a response that is returned when an failure is encountered
func ErrorResponse(w http.ResponseWriter, httpStatus int, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)

	payload := errorResponse{
		ErrorCode:    httpStatus,
		ErrorMessage: err.Error(),
	}

	out, err := json.Marshal(payload)
	if err != nil {
		zap.S().Error("error occurred while marshaling response payload")
		writeServerErrorResponse(w)
		return
	}

	_, err = w.Write(out)
	if err != nil {
		zap.S().Error("error occurred while writing response")
		writeServerErrorResponse(w)
		return
	}
}

func HandleError(w http.ResponseWriter, statusCode int, message string, err error) {
	zap.S().Errorw(message, "error", err)
	ErrorResponse(w, statusCode, errs.New(message))
}

// writeServerErrorResponse writes the error response to help with ErrorResponse
func writeServerErrorResponse(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	_, err := w.Write([]byte(fmt.Sprintf("{\"message\":%s}", "internal server error")))
	if err != nil {
		zap.S().Error("error occurred while writing response")
	}
}
