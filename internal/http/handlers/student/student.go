package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/Prohor722/go_rest_api_1/internal/storage"
	"github.com/Prohor722/go_rest_api_1/internal/types"
	"github.com/Prohor722/go_rest_api_1/internal/utils/response"
	"github.com/go-playground/validator/v10"
)

func New(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		slog.Info("creating a student")
		var student types.Student
		err := json.NewDecoder(r.Body).Decode(&student)

		if errors.Is(err, io.EOF){
			// response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("empty body")))
			return
		}

		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		//Request Validation
		reqErr := validator.New().Struct(student)

		if reqErr != nil {
			validateErrs := reqErr.(validator.ValidationErrors)
			response.WriteJson(w, http.StatusBadRequest, response.ValidationError(validateErrs))
			return
		}

		w.Write([]byte("welcome to students api..."))

		response.WriteJson(w, http.StatusCreated, map[string]string{"success": "OK"})
	}
}
