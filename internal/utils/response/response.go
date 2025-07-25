package response

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Response struct {
	Status string `json:"status"`
	Message string `json:"message"`
}
const (
	StatusOK = "OK"
	StatusError = "Error"
)
func GeneralOK(msg string) Response {
	return Response {
		Status: StatusOK,
		Message: msg,
	}
}

func GeneralError(err error) Response {
	return Response {
		Status: StatusError,
		Message: err.Error(),
	}
}

func WriteJson(w http.ResponseWriter, status int, data interface{})error {
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}

func ValidationError(errs validator.ValidationErrors) Response {
	var errMsg[]string
	for _, err := range errs {
		switch err.ActualTag() {
			case "required":
			errMsg = append(errMsg, fmt.Sprintf("Field %s is required field", err.Field()))
			default:
			errMsg = append(errMsg, fmt.Sprintf("Field %s is invalid", err.Field()))
		}
	}
	return Response{
		Status: StatusError,
		Message: strings.Join(errMsg, ", "),
	}
}