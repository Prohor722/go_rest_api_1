package student

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/Prohor722/go_rest_api_1/internal/types"
)

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		var student types.Student
		json.NewDecoder(r.Body).Decode(&student)
		
		slog.Info("creating a student")
		w.Write([]byte("welcome to students api..."))
	}
}