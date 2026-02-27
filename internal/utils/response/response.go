package response

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator"
)

type Response struct{
	Status string `json:"status"`	//to small latters
	Error string `json:"error_msg"`	//custom also
}

const (
	Status = "OK"
	StatusError = "Error"
)

func WriteJson(w http.ResponseWriter, status int, data interface {}) error {
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(data)
}

func GeneralError(err error) Response {
	return Response{
		Status: StatusError,
		Error: err.Error(),
	}
}

func validationError(errs validator.ValidationErrors) Response {
	var errMsgs []string

	for _,err:=range errs{
		switch err.ActualTag() {
		case "required":
			errMsgs = append(errMsgs, fmt.Sprintf("field %s is required.", err.Field()))
		default:
			errMsgs = append(errMsgs, fmt.Sprintf("field %s is invalid.", err.Field()))
		}
	}

	
	return Response{

	}
}